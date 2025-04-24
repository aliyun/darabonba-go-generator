'use strict';

const assert = require('assert');

const path = require('path');
const fs = require('fs');

const REQUEST = 'request_';
const RESPONSE = 'response_';
const CORE = 'Base';
const DSL = require('@darabonba/parser');
const Annotation = require('@darabonba/annotation-parser');
const { Tag } = DSL.Tag;
const getBuiltin = require('./builtin');

const {
  _name, _format, _string, _type, _initValue, _avoidReserveName, _importFilter, _avoidVariableKeywords,
  _setExtendFunc, _isFilterType, _getAttr, _setValueFunc, _vid, _pointerType, _lowerFirst, _escape,
  _isBinaryOp, _upperFirst, _isIterator, _modelName, _deleteWithSuffix, _snakeCase
} = require('./helper');

function getAttr(node, attrName) {
  for (let i = 0; i < node.attrs.length; i++) {
    if (_name(node.attrs[i].attrName) === attrName) {
      return node.attrs[i].attrValue.string || node.attrs[i].attrValue.lexeme || node.attrs[i].attrValue.value;
    }
  }
}

class Visitor {
  constructor(option = {}) {
    this.baseClient = option && option.baseClient;
    this.typedef = option.go && option.go.typedef ? option.go.typedef : {};
    this.importsTypedef = {};
    const pack = option && option.package || [];
    this.uselessPack = [];
    Object.keys(this.typedef).forEach((def) => {
      if (this.typedef[def].import && this.typedef[def].type) {
        if (!pack.includes(this.typedef[def].import)) {
          pack.push(this.typedef[def].import);
          this.uselessPack.push(this.typedef[def].import);
        }
      }
    });
    this.goPackages = pack.map(packageName => {
      if (packageName.includes(' ')) {
        return `  ${packageName}`;
      }
      if (packageName !== 'io') {
        return `  "${packageName}"`;
      }
    });
    this.config = {
      outputDir: '',
      indent: '    ',
      clientPath: '/client/client.go',
      ...option
    };
    this.output = '';
    this.outputDir = option.outputDir;
    this.__module = {};
    this.exec = this.config.exec;
    this.editable = this.config.editable;
    this.noCompatible = this.config.noCompatible;
    if (this.exec) {
      this.config.clientPath = '/main/main.go';
    }

    if (!this.outputDir) {
      throw new Error('`option.outputDir` should not empty');
    }

    if (!fs.existsSync(this.outputDir)) {
      fs.mkdirSync(this.outputDir, {
        recursive: true
      });
    }
  }

  save(filepath) {
    const targetPath = path.join(this.outputDir, filepath);
    fs.mkdirSync(path.dirname(targetPath), {
      recursive: true
    });

    const modPath = path.join(this.outputDir, 'go.mod');
    if (!fs.existsSync(modPath)) {
      const modFile = this.generatorMod(this.module);
      fs.writeFileSync(modPath, modFile);
    } else {
      const original = fs.readFileSync(modPath, 'UTF-8');
      const modFile = this.reWriteMod(this.module, original);
      fs.writeFileSync(modPath, modFile);
    }
    fs.writeFileSync(targetPath, this.header + this.output);
    this.output = '';
  }


  reWriteMod(module, original) {
    let content = `require (\n`;
    const originMods = {};
    const regex = /require\s*\(\s*([\s\S]*?)\s*\)/gm;
    const match = regex.exec(original);

    if (match) {
      const requireBlock = match[1].trim();
      const packageRegex = /^\s*([^()\s]+)\s+([^\s()]+)$/gm;
      let packageMatch;

      while ((packageMatch = packageRegex.exec(requireBlock)) !== null) {
        const path = packageMatch[1];
        const version = packageMatch[2];
        originMods[path] = version;
      }
    }

    const mods = [];
    Object.keys(module.importPackages).forEach((aliasId) => {
      const importPackage = module.importPackages[aliasId];
      if(!importPackage.version || !importPackage.path) {
        return;
      }
      if(originMods[importPackage.path]) {
        originMods[importPackage.path] = importPackage.version;
        return;
      }
      mods.push('	' + importPackage.path + ' ' + importPackage.version);
    });
    Object.keys(originMods).forEach((path) => {
      mods.push('	' + path + ' ' + originMods[path]);
    });
    content += [...new Set(mods)].join('\n') + `\n)`;

    return original.replace(regex, content);
  }

  generatorMod(module) {
    var content = `module ${module.moduleName}\n`;
    content += `\n`;
    content += `require (\n`;
    const mods = [];
    Object.keys(module.importPackages).forEach((aliasId) => {
      const importPackage = module.importPackages[aliasId];
      if(!importPackage.version || !importPackage.path) {
        return;
      }
      mods.push('	' + importPackage.path + ' ' + importPackage.version);
    });
    content += [...new Set(mods)].join('\n') + `\n)\n`;
    return content;
  }

  getPkgName(filepath) {
    if(filepath.startsWith(this.outputDir)) {
      return filepath.replace(this.outputDir, `${this.module.moduleName}/`);
    } else if(filepath.startsWith('./')) {
      return filepath.replace('./', `${this.module.moduleName}/`);
    } else if(filepath.startsWith('../')) {
      return filepath.replace('../', `${this.module.moduleName}/`);
    } 
    throw Error(`${filepath} is not valid path.`);
    
  }

  emit(str, level) {
    this.output += ' '.repeat(level * 2) + str;
  }

  visit(ast, level = 0) {
    const filepath = this.exec ? 'main/main.go' : 'client/client.go';
    this.ast = ast;
    this.visitModule(ast, filepath, true, level);
  }

  overwrite(ast, filepath) {
    if(!ast.moduleBody.nodes || !ast.moduleBody.nodes.length) {
      return;
    }
    const beginNotes = DSL.note.getNotes(this.notes, 0, ast.moduleBody.nodes[0].tokenRange[0]);
    const overwirte = beginNotes.find(note => note.note.lexeme === '@overwrite');
    const targetPath = path.join(this.outputDir, filepath);
    if(overwirte && overwirte.arg.value === false && fs.existsSync(targetPath)) {
      return false;
    }
    return true;
  }

  clientNameDefine(aliasId, pointer = true) {
    const self = this;
    Object.defineProperty(this.clientName, aliasId, {
      get() {
        let prefix = pointer ? '*' : '';
        const { pkgName } = self.module.importPackages[aliasId];
        self.imports.push({
          aliasId,
          pkgName,
        });
        return prefix + _importFilter(_format(aliasId).toLowerCase()) + '.' + self.constructFunc[aliasId];
      },
      enumerable: true,
      configurable: false
    });
  }

  visitImports(imports, usedTypes, innerModule) {
    const teafilePath = fs.existsSync(path.join(this.config.pkgDir, 'Teafile')) ?
      path.join(this.config.pkgDir, 'Teafile') : path.join(this.config.pkgDir, 'Darafile');
    const teaFile = JSON.parse(fs.readFileSync(teafilePath, 'utf8'));
    const release = teaFile.releases && teaFile.releases.go || '';
    const strs = release.split(':');
    this.module = {
      importPackages: {},
      moduleName: strs[0].substring(0, strs[0].lastIndexOf('/')) || 'client'
    };
    if(!this.structName) {
      this.structName = this.config.go ? this.config.go.clientName || 'Client' : 'Client';
    }
    this.clientName = {};
    this.constructFunc = {};
    if (imports.length > 0) {
      const lockPath = path.join(this.config.pkgDir, '.libraries.json');
      const lock = fs.existsSync(lockPath) ? JSON.parse(fs.readFileSync(lockPath, 'utf8')) : {};
      for (let i = 0; i < imports.length; i++) {
        const item = imports[i];
        const aliasId = _name(item);
        const main = item.mainModule;
        const inner = item.module;
        const moduleDir = main ? this.config.libraries[main] : this.config.libraries[aliasId];
        const innerPath = item.innerPath;
        
        if (innerPath) {
          const filepath = innerPath.replace(/(\.tea)$|(\.spec)$|(\.dara)$/gi, '');
          const pkgName = this.getPkgName(filepath);
          innerModule.set(aliasId, filepath);
          this.module.importPackages[aliasId] = {
            pkgName: pkgName,
            name: aliasId,
            inner: true
          };
          this.constructFunc[aliasId] = this.getInnerClient(aliasId);
          this.clientNameDefine(aliasId);
          continue;
        }
        let targetPath = '';
        if (moduleDir.startsWith('./') || moduleDir.startsWith('../')) {
          targetPath = path.join(this.config.pkgDir, moduleDir);
        } else if (moduleDir.startsWith('/')) {
          targetPath = moduleDir;
        } else {
          targetPath = path.join(this.config.pkgDir, lock[moduleDir]);
        }
        const pkgPath = fs.existsSync(path.join(targetPath, 'Teafile')) ? path.join(targetPath, 'Teafile') : path.join(targetPath, 'Darafile');
        const importTeaFile = JSON.parse(fs.readFileSync(pkgPath));
        const goPkg = importTeaFile.releases && importTeaFile.releases.go;
        if (importTeaFile.go && importTeaFile.go.typedef) {
          this.importsTypedef[aliasId] = {};
          const moduleTypedef = importTeaFile.go && importTeaFile.go.typedef;
          Object.keys(moduleTypedef).forEach((types) => {
            if (!this.importsTypedef[aliasId][types]) {
              this.importsTypedef[aliasId][types] = {};
            }
            this.importsTypedef[aliasId][types].import = moduleTypedef[types].import;
            this.importsTypedef[aliasId][types].type = moduleTypedef[types].type;
          });
        }
        if (importTeaFile.go && importTeaFile.go.clientName) {
          if (importTeaFile.go.clientName.indexOf('*') === 0) {
            this.constructFunc[aliasId] = importTeaFile.go.clientName.substring(1);
            this.clientNameDefine(aliasId);
          } else {
            this.constructFunc[aliasId] = importTeaFile.go.clientName;
            this.clientNameDefine(aliasId, false);
          }
        } else {
          this.constructFunc[aliasId] = 'Client';
          this.clientNameDefine(aliasId);
        }
        if (!goPkg) {
          throw new Error(`The '${aliasId}' has no Go supported.`);
        }

        const [pkgName, version] = goPkg.split(':'); 
        this.module.importPackages[aliasId] = {
          path: pkgName.substring(0, pkgName.lastIndexOf('/')),
          pkgName: pkgName,
          version: version,
          name: aliasId,
          goPkg: importTeaFile.go
        };
        if(inner) {
          const mainPath = pkgName.substring(0, pkgName.lastIndexOf('/'));
          const innerPath = importTeaFile.exports[inner].replace(/(\.tea)$|(\.spec)$|(\.dara)$/gi, '');
          const pkgPath = path.join(mainPath, innerPath);
          this.module.importPackages[aliasId] = {
            pkgName: pkgPath,
            name: aliasId,
            inner: true,
          };
        }
      }
    }
  }

  getInnerClient(aliasId) {
    const moduleAst = this.ast.innerDep.get(aliasId);
    const beginNotes = DSL.note.getNotes(moduleAst.notes, 0, moduleAst.moduleBody.nodes[0].tokenRange[0]);
    const clientNote = beginNotes.find(note => note.note.lexeme === '@clientName');
    if(clientNote) {
      return _string(clientNote.arg.value);
    }
    return 'Client';
  }

  saveInnerModule(ast) {
    const keys = ast.innerModule.keys();
    let data = keys.next();
    while (!data.done) {
      const aliasId = data.value;
      const moduleAst = ast.innerDep.get(aliasId);

      this.structName = this.constructFunc[aliasId];
      const filepath = path.join(ast.innerModule.get(aliasId), 'client.go');
      this.visitModule(moduleAst, filepath, false, 0);
      data = keys.next();
    }
  }

  visitInterface(ast, apis, level) {
    // 被引用时需要使用interface实现多态
    if (this.config.go && this.config.go.interface) {
      let interfaceName = this.config.go.clientName || 'Client';
      this.emit(`type ${interfaceName}Interface interface {\n`, level);
      let NonStaticFuncs = ast.moduleBody.nodes.filter(item => {
        return item.type === 'function' && !item.isStatic;
      });
      for (let func of NonStaticFuncs) {
        const env = {
          apis,
          local: new Map(),
          returnType: func.returnType,
          hasThrow: func.isAsync || func.hasThrow,
          nestFuncParamName: new Map(),
          yieldFunc: _isIterator(func.returnType),
          nestFuncParamNameSubscript: { 'count': 0 }
        };
        this.emit(`${_format(_name(func.functionName))} `, level + 1);
        this.visitParams(func.params, level, env);
        this.visitReturnType(func, level, env, false);
      }
      this.emit(`}\n\n`, level);
    }
  }

  visitModuleClient(ast, apis, nonStaticFuncs, level) {
    if (!ast.extends && (apis.length > 0 || nonStaticFuncs.length > 0)) {
      this.emit(`type ${this.structName} struct {\n`, level);
      if(!this.noCompatible) {
        this.emit(`DisableSDKError *bool\n`, level + 1);
      }
      for (let i = 0; i < ast.moduleBody.nodes.length; i++) {
        const node = ast.moduleBody.nodes[i];
        if (node.type === 'type' && node.vid) {
          let comments = DSL.comment.getFrontComments(this.comments, node.tokenRange[0]);
          this.visitComments(comments, level + 1);
          this.emit(`${_format(_name(node.vid).substring(1))}  `, level + 1);
          this.visitPointerType(node.value, level);
          this.emit(`\n`);
        }
      }
      this.emit(`}\n\n`, level);
    } else if (ast.extends) {
      this.emit(`type ${this.structName} struct {\n`, level);
      this.emit('', level + 1);
      this.visitModuleName(ast.extends);
      this.emit(`.${this.constructFunc[_name(ast.extends)]}\n`);
      if(!this.noCompatible) {
        this.emit(`DisableSDKError *bool\n`, level + 1);
      }
      for (let i = 0; i < ast.moduleBody.nodes.length; i++) {
        const node = ast.moduleBody.nodes[i];
        if (node.type === 'type' && node.vid) {
          let comments = DSL.comment.getFrontComments(this.comments, node.tokenRange[0]);
          this.visitComments(comments, level + 1);
          this.emit(`${_avoidVariableKeywords(_format(_name(node.vid).substring(1)))}  `, level + 1);
          this.visitPointerType(node.value, level);
          this.emit(`\n`);
        }
      }
      this.emit(`}\n\n`, level);
    }
  }

  saveBuffer(name, type) {
    const filename = `${_snakeCase(name)}_${type}.go`;
    if(this.fileBuffer[filename]) {
      this.fileBuffer[filename].imports = this.imports.concat(this.fileBuffer[filename].imports);
      this.fileBuffer[filename].builtinModule = this.builtinModule.concat(this.fileBuffer[filename].builtinModule);
      this.fileBuffer[filename].output = this.output + this.fileBuffer[filename].output;
      return;
    }

    this.fileBuffer[filename] = {
      imports: this.imports,
      output: this.output,
      builtinModule: this.builtinModule,
    };
  }

  flushBuffer(packageName, filepath) {
    Object.keys(this.fileBuffer).map(filename => {
      this.output = this.fileBuffer[filename].output;
      this.imports = this.fileBuffer[filename].imports;
      this.builtinModule = this.fileBuffer[filename].builtinModule;
      filepath = path.join(path.dirname(filepath), filename);
      this.visitHeader(this.__module, packageName);
      this.save(filepath);
      this.output = '';
      this.imports = [];
      this.builtinModule = [];
    });
  }

  visitModule(ast, filepath, main, level) {
    assert.equal(ast.type, 'module');
    this.notes = ast.notes;
    this.fileBuffer = {};
    if(this.overwrite(ast, filepath) === false) {
      return;
    }

    const apis = ast.moduleBody.nodes.filter((item) => {
      return item.type === 'api';
    });

    const models = ast.moduleBody.nodes.filter((item) => {
      return item.type === 'model';
    });

    const exceptions = ast.moduleBody.nodes.filter((item) => {
      return item.type === 'exception';
    });

    const nonStaticFuncs = ast.moduleBody.nodes.filter((item) => {
      return item.type === 'function' && !item.isStatic;
    });

    const initParts = ast.moduleBody.nodes.filter((item) => {
      return item.type === 'init';
    });

    ast.innerModule = new Map();
    this.comments = ast.comments;
    this.predefined = ast.models;
    this.usedExternException = ast.usedExternException;
    this.builtin = getBuiltin(this);
    this.builtinModule = [];
    this.tryFunc = [];
    this.yieldFunc = [];
    this.imports = [];
    this.visitAnnotation(ast.annotation, level);
    this.importPackages = '';
    
    this.visitImports(ast.imports, ast.usedTypes, ast.innerModule);

    const __module = this.__module;
    const packageName = path.dirname(filepath).split('/').pop();

    // global definition
    this.modelBefore(filepath, level);

    for (let i = 0; i < models.length; i++) {
      this.visitModelInterface(models[i], level);
      this.visitModel(models[i], level);
      this.saveBuffer(_format(_name(models[i].modelName)), 'model');
      this.output = '';
      this.imports = [];
      this.builtinModule = [];
    }


    for (let i = 0; i < exceptions.length; i++) {
      this.visitExceptionInterface(exceptions[i], level);
      this.visitException(exceptions[i], level);
      this.saveBuffer(_upperFirst(_name(exceptions[i].exceptionName)), 'error');
      this.output = '';
      this.imports = [];
      this.builtinModule = [];
    }

    this.modelAfter(__module, level);

    this.flushBuffer(packageName, filepath);
    
    this.moduleBefore(ast, main, level);

    this.visitInterface(ast, apis, level);

    this.visitModuleClient(ast, apis, nonStaticFuncs, level);

    // models definition 
    if (initParts.length > 0 && (apis.length > 0 || nonStaticFuncs.length > 0)) {
      this.visitInit(initParts[0], level);
      this.apiBefore(__module, level);
    } else if (ast.extends) {
      this.emit(`func NewClient(config *${_name(ast.extends).toLowerCase()}.Config)(*${this.structName}, error) {\n`, level);
      this.emit(`client := new(${this.structName})\n`, level + 1);
      this.emit(`err := client.Init(config)\n`, level + 1);
      this.emit(`return client, err\n`, level + 1);
      this.emit(`}\n\n`, level);
    } 
    // else {
    //   this.emit(`func NewClient() error {\n`, level);
    //   this.emit(`err := dara.NewSDKError(map[string]interface{}{\n`, level + 1);
    //   this.emit(`"message": "Un-Support!"\n`, level + 2);
    //   this.emit('})\n', level + 1);
      
    //   this.emit(`return err\n`, level + 1);
    //   this.emit(`}\n\n`, level);
    // }

    for (let i = 0; i < apis.length; i++) {
      this.eachAPI(apis[i], level, ast.predefined);
    }

    if (apis.length > 0) {
      this.apiAfter(__module, level);
    }

    this.wrapBefore(__module, level);

    const funcs = ast.moduleBody.nodes.filter((item) => {
      return item.type === 'function';
    });

    for (let i = 0; i < funcs.length; i++) {
      this.eachFunc(funcs[i], level, ast.predefined, apis);
    }

    for (let i = 0; i < this.tryFunc.length; i++) {
      this.eachTryFunc(this.tryFunc[i], level);
    }

    for (let i = 0; i < this.yieldFunc.length; i++) {
      this.eachYieldFunc(this.yieldFunc[i], level);
    }
    

    if (this.exec) {
      this.emit(`\nfunc main() {\n`, 0);
      this.emit(`err := _main(dara.StringSlice(os.Args[1:]))\n`, 1);
      this.emit(`if err != nil {\n`, 1);
      this.emit('panic(err)\n', 2);
      this.emit(`}\n`, 1);
      this.emit('}\n', 0);
    }

    this.wrapAfter(filepath, level);

    this.visitHeader(this.__module, packageName, level);

    this.moduleAfter(__module, level);

    this.save(filepath);
    this.mainModule = false;
    this.saveInnerModule(ast);
  }

  visitConstructor(){
    
  }

  visitAnnotation(annotation, level) {
    if (!annotation || !annotation.value) {
      return;
    }
    let comments = DSL.comment.getFrontComments(this.comments, annotation.index);
    this.visitComments(comments, level);
    var ast = Annotation.parse(annotation.value);
    var description = ast.items.find((item) => {
      return item.type === 'description';
    });
    var summary = ast.items.find((item) => {
      return item.type === 'summary';
    });
    var _return = ast.items.find((item) => {
      return item.type === 'return';
    });
    var deprecated = ast.items.find((item) => {
      return item.type === 'deprecated';
    });
    var params = ast.items.filter((item) => {
      return item.type === 'param';
    }).map((item) => {
      return {
        name: item.name.id,
        text: item.text.text
      };
    });
    var throws = ast.items.filter((item) => {
      return item.type === 'throws';
    }).map((item) => {
      return item.text.text;
    });

    let hasNextSection = false;
    const summaryText = summary ? _escape(summary.text.text).trimEnd() : '';
    const descriptionText = description ? _escape(description.text.text).trimEnd() : '';
    const returnText = _return ? _escape(_return.text.text).trimEnd() : '';

    if (deprecated) {
      const deprecatedText = _escape(deprecated.text.text).trimEnd();
      deprecatedText.split('\n').forEach((line, index, array) => {
        if(index === 0) {
          this.emit(`// Deprecated: ${line}\n`, level);
          if (array.length > 1) {
            this.emit(`//\n`, level);
          }
        } else {
          this.emit(`// ${line}\n`, level);
          if (index < array.length - 1){
            this.emit(`//\n`, level);
          }
        }
      });
      hasNextSection = true;
    }
    if (summaryText !== '') {
      if (hasNextSection) {
        this.emit(`// \n`, level);
      }
      this.emit(`// Summary:\n`, level);
      this.emit(`// \n`, level);
      summaryText.split('\n').forEach((line, index, array) => {
        this.emit(`// ${line}\n`, level);
        if (index < array.length - 1) {
          this.emit(`// \n`, level);
        }
      });
      hasNextSection = true;
    }
    if (descriptionText !== '') {
      if (hasNextSection) {
        this.emit(`// \n`, level);
      }
      this.emit(`// Description:\n`, level);
      this.emit(`// \n`, level);
      descriptionText.split('\n').forEach((line, index, array) => {
        this.emit(`// ${line}\n`, level);
        if (index < array.length - 1) {
          this.emit(`// \n`, level);
        }
      });
      hasNextSection = true;
    }
    if (params.length > 0) {
      if (hasNextSection) {
        this.emit(`// \n`, level);
      }
      params.forEach((item, i) => {
        this.emit(`// @param ${item.name} - `, level);
        const items = item.text.trimEnd().split('\n');
        items.forEach((line, j) => {
          if (j === 0) {
            this.emit(`${line}\n`);
          } else {
            this.emit(`// ${line}\n`, level);
          }
          if (j < items.length - 1 || (j === items.length - 1 && i < params.length -1)) {
            this.emit(`// \n`, level);
          }
        });
      });
      hasNextSection = true;
    }
    if (returnText) {
      if (hasNextSection) {
        this.emit(`// \n`, level);
      }
      this.emit(`// @return `, level);
      const returns = returnText.split('\n');
      returns.forEach((line, index) => {
        if (index === 0) {
          this.emit(`${line}\n`);
        } else {
          this.emit(`// ${line}\n`, level);
        }
        if (index < returns.length - 1) {
          this.emit(`// \n`, level);
        }
      });
      hasNextSection = true;
    }
    if (throws.length > 0) {
      if (hasNextSection) {
        this.emit(`// \n`, level);
      }
      throws.forEach((item, i) => {
        this.emit(`// @throws `, level);
        const items = item.trimEnd().split('\n');
        items.forEach((line, j) => {
          if (j === 0) {
            this.emit(`${line}\n`);
          } else {
            this.emit(`// ${line}\n`, level);
          }
          if (j < items.length - 1 || (j === items.length - 1 && i < params.length -1)) {
            this.emit(`// \n`, level);
          }
        });
      });
    }
  }

  visitComments(comments, level) {
    comments.forEach(comment => {
      this.emit(`${comment.value}`, level);
      this.emit(`\n`);
    });
  }

  visitParams(ast, level, env) {
    assert.equal(ast.type, 'params');
    env.pointerParams = [];
    this.emit('(');
    for (var i = 0; i < ast.params.length; i++) {
      if (i !== 0) {
        this.emit(', ');
      }
      const node = ast.params[i];
      assert.equal(node.type, 'param');
      const name = _name(node.paramName);
      env.pointerParams.push(_avoidReserveName(name));
      this.emit(`${_avoidReserveName(name)}`);
      if (node.paramType) {
        const paramType = node.paramType;
        this.emit(` `);
        this.visitPointerType(paramType, level, env);
      }
    }

    if(env.yieldFunc) {
      if(ast.params.length > 0) {
        this.emit(', ');
      }
      this.emit('_yield chan ');
      this.visitPointerType(env.returnType, level, env);
      if(env.hasThrow) {
        this.emit(', _yieldErr chan error');
      }
    }

    this.emit(')');
  }

  visitInit(ast, level) {
    const env = {
      local: new Map(),
      funcName: 'NewClient'
    };
    assert.equal(ast.type, 'init');
    this.visitAnnotation(ast.annotation, level);
    let comments = DSL.comment.getFrontComments(this.comments, ast.tokenRange[0]);
    this.visitComments(comments, level);
    this.emit(`func NewClient`, level);
    this.visitParams(ast.params, level, env);
    this.emit(`(*${this.structName}, error) {\n`);
    this.emit(`client := new(${this.structName})\n`, level + 1);
    this.emit(`err := client.Init(`, level + 1);
    for (let i = 0; i < ast.params.params.length; i++) {
      const param = ast.params.params[i];
      this.emit(`${_name(param.paramName)}`);
      if (i < ast.params.params.length - 1) {
        this.emit(`, `);
      }
    }
    this.emit(`)\n`);
    this.emit(`return client, err\n`, level + 1);
    this.emit(`}\n\n`, level);
    this.emit(`func (client *${this.structName})Init`, level);
    this.visitParams(ast.params, level, env);
    this.emit(`(_err error) {\n`);
    if (ast.initBody && ast.initBody.stmts && ast.initBody.stmts.length > 0) {
      this.visitStmts(ast.initBody, level + 1, env);
    }
    this.emit(`return nil\n`, level + 1);
    this.emit(`}\n\n`, level);
  }

  visitFunctionNested(ast, level, env) {
    var argHasThrowFunc = new Map;
    for (let i = 0; i < ast.args.length; i++) {
      if (ast.args[i].type === 'call' && ast.args[i].hasThrow) {
        let paramName;
        if (ast.args[i].left.type === 'method_call') {
          paramName = _lowerFirst(_name(ast.args[i].left.id)) + 'Tmp';
        } else {
          paramName = _lowerFirst(_name(ast.left.propertyPath[0])) + 'Tmp';
        }

        if (env.nestFuncParamName.has(paramName)) {
          let count = env.nestFuncParamNameSubscript['count'] + 1;
          paramName = paramName + count;
        }
        env.nestFuncParamName.set(paramName);
        argHasThrowFunc.set(i, paramName);
        this.emit(`${paramName}, err := `, level);
        this.visitExpr(ast.args[i], level, env, { pointer: true });
        this.emit('\n');
        this.emit(`if err != nil {\n`, level);
        this.emit(`_err = err\n`, level + 1);
        if (env.returnType && _name(env.returnType) !== 'void') {
          this.emit(`return _result, _err\n`, level + 1);
        } else {
          this.emit(`return _err\n`, level + 1);
        }
        this.emit(`}\n`, level);
      }
    }
    return argHasThrowFunc;
  }

  visitReturnType(ast, level, env, needClose = true) {
    if(env.yieldFunc) {
      this.emit(` ${needClose ? `{\n` : `\n`}`);
      return;
    }
    if (_name(ast.returnType) !== 'void') {
      this.emit(` (_result `);
      this.visitPointerType(ast.returnType, level, env);
      if (env.hasThrow) {
        this.emit(`, _err error) ${needClose ? `{\n` : `\n`}`);
      } else {
        this.emit(`) ${needClose ? `{\n` : `\n`}`);
      }
      return;
    }
    if (env.hasThrow) {
      this.emit(` (_err error) ${needClose ? `{\n` : `\n`}`);
    } else {
      this.emit(` ${needClose ? `{\n` : `\n`}`);
    }

  }

  checkExpr(expr) {
    if(!expr) {
      return 0;
    }
    if(( expr.type === 'call' &&
      expr.hasThrow) || expr.type === 'construct') {
      return 1;
    }
    return 0;
  }

  stmtsErrCount(stmts) {
    let errNum = 0;
    for (var i = 0; i < stmts.length; i++) {
      const ast = stmts[i];
      if (ast.type === 'return') {
        errNum += this.checkExpr(ast.expr);
      } else if (ast.type === 'if') {
        errNum += this.checkExpr(ast.condition);
        errNum += this.stmtsErrCount(ast.stmts.stmts);
        if(ast.elseIfs && ast.elseIfs.length > 0) {
          ast.elseIfs.map(elseIf => {
            errNum += this.checkExpr(elseIf.condition);
            errNum += this.stmtsErrCount(elseIf.stmts.stmts);
          });
        }
        if(ast.elseStmts) {
          errNum += this.stmtsErrCount(ast.elseStmts.stmts);
        }
      } else if (ast.type === 'throw') {
        errNum += 1;
      } else if (ast.type === 'assign') {
        errNum += this.checkExpr(ast.expr);
      } else if (ast.type === 'declare') {
        errNum += this.checkExpr(ast.expr);
      } else if (ast.type === 'while') {
        errNum += this.checkExpr(ast.condition);
        errNum += this.stmtsErrCount(ast.stmts.stmts);
      } else if (ast.type === 'for') {
        errNum += this.checkExpr(ast.list);
        errNum += this.stmtsErrCount(ast.stmts.stmts);
      } else if (ast.type === 'try') {
        errNum += 1;
      } else if (ast.type === 'call' && ast.hasThrow) {
        const name = _format(_name(ast.left.id));
        if (!name.startsWith('$') || !this.builtin[name]) {
          errNum += 1;
        }
      }
    }
    return errNum;
  }

  getObjectVars(ast) {
    let vars = [];
    const fields = ast.fields.filter((field) => {
      return field.type === 'objectField';
    });
    const expandFields = ast.fields.filter((field) => {
      return field.type === 'expandField';
    });
    for(let i = 0; i < fields.length; i++) {
      vars = vars.concat(this.getExprVars(fields[i].expr));
    }
    for(let i = 0; i < expandFields.length; i++) {
      vars = vars.concat(this.getExprVars(expandFields[i].expr));
    }
    return vars;
  }

  getName(ast) {
    let id = _name(ast.id);

    if(ast.id.tag === Tag.VID){
      id = _vid(ast.id);
    }

    var expr = '';
    if (id === '__response') {
      expr += RESPONSE;
    } else if (id === '__request') {
      expr += REQUEST;
    } else {
      expr += _avoidReserveName(id);
    }

    return expr;
  }

  getModuleType(moduleType) {
    if(moduleType.type === 'basic' || moduleType.type === 'model') {
      return {
        type: moduleType.type,
        name: moduleType.name
      };
    }
    return  { 
      idType: 'module',
      name: moduleType.name 
    };
  }

  getExprVars(ast) {
    let vars = [];
    if (ast.type === 'property_access') {
      vars.push({
        name: this.getName(ast),
        type: ast.id.inferred,
      });
    } else if (ast.type === 'object') {
      vars.concat(this.getObjectVars(ast));
    } else if (ast.type === 'variable') {
      vars.push({
        name: this.getName(ast),
        type: ast.inferred,
      });
    } else if (ast.type === 'virtualVariable') {
      vars.push({
        name: 'client',
        type: this.structName,
      });
    } else if (ast.type === 'decrement') {
      vars = vars.concat(this.getExprVars(ast.expr));
    } else if (ast.type === 'increment') {
      vars = vars.concat(this.getExprVars(ast.expr));
    } else if (ast.type === 'template_string') {
      for (let i = 0; i < ast.elements.length; i++) {
        const item = ast.elements[i];
        if (item.type === 'expr') {
          const expr = item.expr;
          vars = vars.concat(this.getExprVars(expr));
        }
      }
    } else if (ast.type === 'call') {
      for (let i = 0; i < ast.args.length; i++) {
        const expr = ast.args[i];
        vars = vars.concat(this.getExprVars(expr));
      }

      if(ast.left.type === 'instance_call') {

        if (_name(ast.left.id).indexOf('@') === 0) {
          vars.push({
            name: 'client',
            type: this.structName,
          });
        } else {
          const type = this.getModuleType(ast.left.id.moduleType);
          vars.push({
            name: _name(ast.left.id),
            type: type
          });
        }
      } else if(ast.left.type === 'method_call') {
        const name = _format(_name(ast.left.id));
        if (name.startsWith('$') && this.builtin[name]) {
          return vars;
        } 
        if (!ast.isStatic) {
          vars.push({
            name: 'client',
            type: this.structName,
          });
        }
      }
    } else if (ast.type === 'group') {
      vars = vars.concat(this.getExprVars(ast.expr));
    } else if (_isBinaryOp(ast.type)) {
      vars = vars.concat(this.getExprVars(ast.left));
      vars = vars.concat(this.getExprVars(ast.right));
    } else if (ast.type === 'construct') {
      for (let i = 0; i < ast.args.length; i++) {
        vars = vars.concat(this.getExprVars(ast.args[i]));
      }
    } else if (ast.type === 'construct_model' && ast.object) {
      vars = vars.concat(this.getObjectVars(ast.object));
    } else if (ast.type === 'not') {
      vars = vars.concat(this.getExprVars(ast.expr));
    } else if (ast.type === 'array') {
      for (let i = 0; i < ast.items.length; i++) {
        vars = vars.concat(this.getExprVars(ast.items[i]));
      }
    } else if (ast.type === 'map_access') {
      vars.push({
        name: this.getName(ast),
        type: {
          type: 'map', 
          keyType: { 
            type: 'basic', 
            name: 'string' 
          }, 
          valueType: ast.inferred 
        },
      });
    } else if (ast.type === 'array_access') {
      vars.push({
        name: this.getName(ast),
        type: { type: 'array', itemType: ast.inferred },
      });
    } else if (ast.type === 'super') {
      for (let i = 0; i < ast.args.length; i++) {
        vars = vars.concat(this.getExprVars(ast.args[i]));
      }
    }
    return vars;
  }

  getStmtsVars(stmts) {
    let args = [];
    let declare = [];
    for (var i = 0; i < stmts.length; i++) {
      const ast = stmts[i];
      if (ast.type === 'return') {
        args = args.concat(this.getExprVars(ast.expr));
      } else if (ast.type === 'if') {
        args = args.concat(this.getExprVars(ast.condition));
        args = args.concat(this.getStmtsVars(ast.stmts.stmts));
        if(ast.elseIfs && ast.elseIfs.length > 0) {
          ast.elseIfs.map(elseIf => {
            args = args.concat(this.getExprVars(elseIf.condition));
            args = args.concat(this.getStmtsVars(elseIf.stmts.stmts));
          });
        }
        if(ast.elseStmts) {
          args = args.concat(this.getStmtsVars(ast.elseStmts.stmts));
        }
      } else if (ast.type === 'throw') {
        args = args.concat(this.getExprVars(ast.expr));
      } else if (ast.type === 'assign') {
        args = args.concat(this.getExprVars(ast.left));
        args = args.concat(this.getExprVars(ast.expr));
      } else if (ast.type === 'declare') {
        declare.push(_name(ast.id));
        args = args.concat(this.getExprVars(ast.expr));
      } else if (ast.type === 'while') {
        args = args.concat(this.getExprVars(ast.condition));
        args = args.concat(this.getStmtsVars(ast.stmts.stmts));
      } else if (ast.type === 'for') {
        args = args.concat(this.getExprVars(ast.list));
        args = args.concat(this.getStmtsVars(ast.stmts.stmts));
      } else if (ast.type === 'try') {
        args = args.concat(this.getStmtsVars(ast.stmts.stmts));
      } else if (ast.type === 'call') {
        args = args.concat(this.getExprVars(ast));
      }  else {
        args = args.concat(this.getExprVars(ast));
      }
    }
    return args.filter(arg => {
      if(!arg) {
        return false;
      }
      if(declare.includes(arg.name)) {
        return false;
      }
      return true;
    });
  }

  visitAPIBody(ast, level, env) {
    assert.equal(ast.type, 'apiBody');
    this.emit(`${REQUEST} ${env.runtimeBody ? '' : ':'}= dara.NewRequest()\n`, level);
    if (ast.stmts) {
      this.visitStmts(ast.stmts, level, env);
    }
  }

  visitRuntimeBefore(ast, level, env) {
    assert.equal(ast.type, 'object');
    this.emit('_runtime := dara.NewRuntimeObject(', level);
    this.visitObject(ast, level, env, 'map[string]interface{}');
    this.emit(')\n\n');
    this.emit('var retryPolicyContext *dara.RetryPolicyContext\n', level);
    this.emit('var request_ *dara.Request\n', level);
    this.emit('var response_ *dara.Response\n', level);
    this.emit('var _resultErr error\n', level);
    this.emit('retriesAttempted := int(0)\n', level);
    this.emit('retryPolicyContext = &dara.RetryPolicyContext{\n', level);
    this.emit('RetriesAttempted: retriesAttempted,\n', level + 1);
    this.emit('}\n\n', level);
    if (_name(env.returnType) && _name(env.returnType) !== 'void') {
      this.emit(`_result = ${_initValue(_name(env.returnType))}\n`, level);
    } else if (env.returnType.path) {
      this.emit(`_result = new(`, level);
      for (let i = 0; i < env.returnType.path.length; i++) {
        const path = env.returnType.path[i];
        if (i === 0) {
          this.emit(_name(path).toLowerCase());
        } else {
          this.emit(`.${_name(path)}`);
        }
      }
      this.emit(`)\n`);
    } else if (env.returnType.type === 'map') {
      this.emit(`_result = make(`, level);
      this.visitPointerType(env.returnType, level);
      this.emit(`)\n`);
    }
    this.emit(`for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {\n`, level);
    this.emit(`_resultErr = nil\n`, level + 1);
    this.emit(`_backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)\n`, level + 1);
    this.emit(`dara.Sleep(_backoffDelayTime)\n`, level + 1);
    this.emit(`\n`);
  }

  visitStmt(ast, level, env) {
    let comments = DSL.comment.getFrontComments(this.comments, ast.tokenRange[0]);
    this.visitComments(comments, level);
    if (ast.type === 'return') {
      this.visitReturn(ast, level, env);
    } else if (ast.type === 'yield') {
      this.visitYield(ast, level, env);
    } else if (ast.type === 'if') {
      this.visitIf(ast, level, env);
    } else if (ast.type === 'throw') {
      this.visitThrow(ast, level, env);
    } else if (ast.type === 'assign') {
      this.visitAssign(ast, level, env);
    } else if (ast.type === 'retry') {
      this.visitRetry(ast, level);
    } else if (ast.type === 'declare') {
      this.visitDeclare(ast, level, env);
    } else if (ast.type === 'while') {
      this.visitWhile(ast, level, env);
    } else if (ast.type === 'for') {
      this.visitFor(ast, level, env);
    } else if (ast.type === 'try') {
      this.visitTry(ast, level, env);
    } else if (ast.type === 'break') {
      this.emit(`break\n`, level);
    } else {
      if (ast.type === 'call' && ast.hasThrow) {
        if (ast.inferred && _name(ast.inferred) !== 'void') {
          this.emit(`_, _err ${env.yieldFunc ? ':' : ''}= `, level);
        } else {
          this.emit(`_err ${env.yieldFunc ? ':' : ''}= `, level);
        }
        this.visitExpr(ast, level, env, { pointer: false });
        this.emit(`\n`);
        if(env.runtimeBody){
          this.visitAPIErrCatch(level, env);
        } else if(env.try) {
          const tryStmt = env.try;
          env.try = null;
          this.visitCatch(tryStmt, level, env);
        } else {
          this.emit(`if _err != nil {\n`, level);
          if (env.returnType && _name(env.returnType) !== 'void') {
            this.emit(`return _result, _err\n`, level + 1);
          } else if(env.yieldFunc){
            this.emit(`_yieldErr <- _err\n`, level + 1);
            this.emit(`return\n`, level + 1);
          } else {
            this.emit(`return _err\n`, level + 1);
          }
          this.emit(`}\n`, level);
        }
      } else {
        this.emit(``, level);
        this.visitExpr(ast, level, env, { pointer: false });
        this.emit(`\n`);
      }
    }
  }

  getReturnType(stmts, env) {
    for (var i = 0; i < stmts.length; i++) {
      const ast = stmts[i];
      if (ast.type === 'return') {
        return env.returnType ||  { lexeme: 'void' };
      } 
    }
    return { lexeme: 'void' };
  }

  visitCatch(ast, level, env) {

    if(ast.finallyBlock) {
      env.finallyBlock = true;
      this.visitStmts(ast.finallyBlock, level, env);
      env.finallyBlock = false;
    }

    this.emit(`if _err != nil {\n`, level);
    if (ast.catchBlocks && ast.catchBlocks.length > 0) {
      ast.catchBlocks.forEach(catchBlock => {
        if (!catchBlock.id) {
          return;
        }
        
        if (!catchBlock.id.type) {
          this.emit(`if _t, ok := _err.(*dara.SDKError); ok {\n`, level + 1);
        } else {
          this.emit(`if _t, ok := _err.(`, level + 1);
          this.visitType(catchBlock.id.type);
          this.emit('); ok {\n');
        }
        if(catchBlock.catchStmts && catchBlock.catchStmts.stmts.length > 0){
          this.emit(`${_name(catchBlock.id)} := _t;\n`, level + 2);
          this.visitStmts(catchBlock.catchStmts, level + 2, env);
        }
        
        
        this.emit('}\n', level + 1);
      });
    } else if (ast.catchBlock && ast.catchBlock.stmts.length > 0) {
      this.emit(`if _t, ok := _err.(*dara.SDKError); ok {\n`, level + 1);
      this.emit(`${_name(ast.catchId)} := _t\n`, level + 2);
      this.emit(`}\n`, level + 1);
      this.visitStmts(ast.catchBlock, level + 1, env);
    }

    this.emit(`}\n`, level);
  }

  visitTry(ast, level, env) {
    assert.equal(ast.type, 'try');
    env = env || {
      local: new Map(),
    };
    const tryBlock = ast.tryBlock;
    const errCounts = this.stmtsErrCount(tryBlock.stmts);
    if(errCounts > 1) {
      const funcName = `${_lowerFirst(env.funcName)}_opTryFunc`;
      const args = this.getStmtsVars(tryBlock.stmts);
      const returnType = this.getReturnType(tryBlock.stmts, env);
      this.tryFunc.push({
        args,
        functionBody: tryBlock,
        returnType,
        name: funcName,
        pointerParams: env.pointerParams,
      });
      if (returnType !== 'void') {
        this.emit(`_result, _err  = ${funcName}`, level);
      } else {
        this.emit(`_err = ${funcName}`, level);
      }
      this.visitTryArgs(args, level, env);
      this.emit('\n');
      this.visitCatch(ast, level, env);
      env.hasReturn = true;
      return;
    } 
    env.try = ast;
    this.visitStmts(tryBlock, level, env);
  }

  visitWhile(ast, level, env) {
    assert.equal(ast.type, 'while');
    let argHasThrowFunc;
    if (ast.condition.type === 'not' && ast.condition.expr && ast.condition.expr.type === 'call') {
      argHasThrowFunc = this.visitFunctionNested(ast.condition.expr, level, env);
    } else if (ast.condition.type === 'call') {
      argHasThrowFunc = this.visitFunctionNested(ast.condition, level, env);
    }

    this.emit('for ', level);
    let setFunc;
    if(ast.condition && ast.condition.type === 'call') {
      let dealFunc = this.getVarDealFunc(ast.condition, true);
      setFunc = dealFunc && dealFunc(_name(ast.condition.inferred));
    }
    if(setFunc) {
      this.emit(`${setFunc}`);
    }
    this.visitExpr(ast.condition, level + 1, env, false, argHasThrowFunc);
    if(setFunc) {
      this.emit(')');
    }
    this.emit(' {\n');
    this.visitStmts(ast.stmts, level + 1, env);
    this.emit('}\n', level);
  }

  visitFor(ast, level, env) {
    assert.equal(ast.type, 'for');
    if(ast.list.inferred && _isIterator(ast.list.inferred)) {
      this.emit(`for ${_name(ast.id)} := range `, level);
    } else {
      this.emit(`for _, ${_name(ast.id)} := range `, level);
    }
    this.visitExpr(ast.list, level + 1, env, { pointer: true });
    this.emit(' {\n');
    this.visitStmts(ast.stmts, level + 1, env);
    this.emit('}\n', level);
  }

  visitFieldValue(ast, structName, level) {
    if (ast.type === 'fieldType') {
      if (ast.fieldType === 'array') {
        if (ast.fieldItemType.type === 'modelBody') {
          this.emit('{\n');
          this.visitModelBody(ast.fieldItemType, ast.fieldItemType.nodes, structName, level);
        }
        return;
      }
    }

    if (ast.type === 'modelBody') {
      this.emit('{\n');
      this.visitModelBody(ast, ast.nodes, structName, level);
      return;
    }

    throw new Error('unimpelemented');
  }

  visitType(ast, level) {
    if (ast.type === 'moduleModel') {
      const [ mainId, ...rest ] = ast.path;
      let moduleName = _importFilter(_name(ast.path[0]).toLowerCase());
      let modelName = rest.map(node => {
        return _upperFirst(_name(node));
      }).join('');
      const externEx = this.usedExternException.get(_name(mainId));
      if (externEx && externEx.has(modelName)) {
        modelName += 'Error';
      }
      this.emit(`*${moduleName}.${_format(modelName)}`);
    } else if (ast.type === 'moduleTypedef') {
      this.emit(`*`);
      for (let i = 1; i < ast.path.length; i++) {
        this.emit(`${this.typeRelover(ast.path[i], ast.path[0])}`);
      }
    } else if (ast.type === 'subModel') {
      this.emit(`*${_format(_name(ast.path[0]))}`);
      for (let i = 1; i < ast.path.length; i++) {
        this.emit(`${_format(_name(ast.path[i]))}`);
      }
    } else if ((ast.type === 'map' || ast.fieldType === 'map') && _name(ast.keyType)) {
      this.emit(`map[${this.getType(_name(ast.keyType), false)}]`);
      this.visitPointerType(ast.valueType, level);
    } else if (ast.fieldType === 'array' || ast.type === 'array') {
      this.emit(`[]`);
      this.visitPointerType(ast.subType || ast.itemType, level);
    } else if (ast.idType === 'module' || this.clientName[_name(ast)]) {
      this.emit(`${this.clientName[_name(ast)]}`);
    } else if (this.typeRelover(ast)) {
      this.emit(this.getType(this.typeRelover(ast), false));
    } else if (ast.fieldType && DSL.util.isBasicType(ast.fieldType)) {
      this.emit(this.getType(ast.fieldType, false));
    } else if (ast.fieldType && this.typeRelover(ast.fieldType)) {
      this.emit(this.getType(this.typeRelover(ast.fieldType), false));
    } else {
      this.emit(this.getType(ast, false));
    }
  }

  visitModuleName(aliasId) {
    const moduleName = _importFilter(_format(_name(aliasId)).toLowerCase());
    this.emit(`${moduleName}`);
    const { pkgName } = this.module.importPackages[_name(aliasId)];
    this.imports.push({
      aliasId: _name(aliasId),
      pkgName,
    });
  }

  visitPointerType(ast, level) {
    if (ast.type === 'moduleModel') {
      this.emit('*');
      this.visitModuleName(ast.path[0]);
      for (let i = 1; i < ast.path.length; i++) {
        if (i === 1) {
          this.emit(`.`);
        }
        this.emit(`${_format(_name(ast.path[i]))}`);
      }
    } else if (ast.type === 'moduleTypedef') {
      this.emit(`*`);
      for (let i = 1; i < ast.path.length; i++) {
        this.emit(`${this.typeRelover(ast.path[i], ast.path[0])}`);
      }
    } else if (ast.type === 'subModel') {
      this.emit(`*${_format(_name(ast.path[0]))}`);
      for (let i = 1; i < ast.path.length; i++) {
        this.emit(`${_format(_name(ast.path[i]))}`);
      }
    } else if ((ast.type === 'map' || ast.fieldType === 'map') && _name(ast.keyType)) {
      this.emit(`map[${this.getType(_name(ast.keyType), false)}]`);
      this.visitPointerType(ast.valueType, level);
    } else if (ast.type === 'asyncIterator' || ast.type === 'iterator') {
      this.visitPointerType(ast.valueType, level);
    }else if (ast.fieldType === 'array' || ast.type === 'array') {
      this.emit(`[]`);
      this.visitPointerType(ast.subType || ast.itemType, level);
    } else if (ast.idType === 'module' || this.clientName[_name(ast)]) {
      this.emit(`${this.clientName[_name(ast)]}`);
    } else if (ast.idType === 'builtin_model') {
      this.emit(`${this.getType(_name(ast), false)}`);
    } else if (ast.type === 'model') {
      this.emit(`*`);
      if (ast.moduleName) {
        this.emit(`${ast.moduleName.replace(/-/g, '_').toLowerCase()}.`);
      }
      let strs = _format(_name(ast)).split('.');
      strs.forEach(str => {
        this.emit(`${_modelName(_format(str))}`);
      });
    } else if (_name(ast)) {
      this.emit(this.getType(this.typeRelover(ast)));
    } else if (ast.fieldType && DSL.util.isBasicType(ast.fieldType)) {
      this.emit(this.getType(ast.fieldType));
    } else if (ast.fieldType && this.typeRelover(ast.fieldTyp)) {
      this.emit(this.getType(this.typeRelover(ast.fieldType)));
    } else {
      this.emit(this.getType(ast));
    }
  }

  visitModelField(ast, structName, level) {
    //assert.equal(ast.fieldValue.type, 'fieldType');
    this.emit(`type ${structName} struct `);
    this.visitFieldValue(ast, structName, level);
  }

  visitFieldType(node, structName, fields, structMap, level) {
    let type = '', omitemptyEnable = true;
    if (node.fieldValue.fieldType === 'array') {
      type = `type:"Repeated"`;
      if (this.config.go && this.config.go.mapAndSliceWithoutOmitempty === true) {
        omitemptyEnable = false;
      }
      if (_name(node.fieldValue.fieldItemType)) {
        this.emit(`[]${this.getType(_name(node.fieldValue.fieldItemType))} `);
      } else if (node.fieldValue.fieldItemType.type === 'map') {
        this.emit(`[]`);
        this.visitType(node.fieldValue.fieldItemType);
        this.emit(` `);
      } else if (node.fieldValue.fieldItemType.type === 'modelBody') {
        structMap.push(structName);
        this.emit(`[]*${structName} `);
        fields.push(node.fieldValue);
      } else if (node.fieldValue.fieldItemType.fieldType === 'array') {
        this.emit(`[][]`);
        this.emitModelArray(node.fieldValue.fieldItemType, structMap, fields, structName);
      }
    } else if (node.fieldValue.type === 'modelBody') {
      this.emit(`*${structName} `);
      structMap.push(structName);
      fields.push(node.fieldValue);
      type = `type:"Struct"`;
    } else {
      const fieldType = node.fieldValue.fieldType;
      if (!_name(fieldType) && (fieldType === 'map' || fieldType === 'object')) {
        if (this.config.go && this.config.go.mapAndSliceWithoutOmitempty === true) {
          omitemptyEnable = false;
        }
        this.visitPointerType(node.fieldValue, level);
        this.emit(` `);
      } else {
        this.visitPointerType(fieldType, level);
        this.emit(` `);
      }
    }
    return { type, omitemptyEnable };
  }

  visitModelBody(ast, nodes, lastName, level) {
    assert.equal(ast.type, 'modelBody');
    var fields = [];
    const structMap = [];
    let node;
    for (let i = 0; i < nodes.length; i++) {
      node = nodes[i];
      let comments = DSL.comment.getFrontComments(this.comments, node.tokenRange[0]);
      this.visitComments(comments, level);
      var fieldName = _name(node.fieldName);
      const structName = lastName + _format(fieldName);
      const description = getAttr(node, 'description');
      const example = getAttr(node, 'example');
      const checkBlank = getAttr(node, 'checkBlank');
      const nullable = getAttr(node, 'nullable');
      const sensitive = getAttr(node, 'sensitive');
      const deprecated = getAttr(node, 'deprecated');
      let hasNextSection = false;
      if (deprecated === 'true') {
        this.emit(`// Deprecated\n`, level);
        hasNextSection = true;
      }
      if (description || example || typeof checkBlank !== 'undefined' || typeof nullable !== 'undefined' || typeof sensitive !== 'undefined') {
        if (description) {
          if (hasNextSection) {
            this.emit(`// \n`, level);
          }
          const descriptions = _escape(description).split('\n');
          for (let j = 0; j < descriptions.length; j++) {
            if (descriptions[j] === '') {
              this.emit(`// \n`, level);
            }
            else {
              this.emit(`// ${descriptions[j]}\n`, level);
              if (j < descriptions.length - 1 && descriptions[j + 1] !== '') {
                this.emit(`// \n`, level);
              }
            }
          }
          hasNextSection = true;
        }
        if (typeof checkBlank !== 'undefined') {
          if (hasNextSection) {
            this.emit(`// \n`, level);
          }
          this.emit('// check if is blank:\n', level);
          this.emit(`// ${checkBlank}\n`, level);
          hasNextSection = true;
        }
        if (typeof nullable !== 'undefined') {
          if (hasNextSection) {
            this.emit(`// \n`, level);
          }
          this.emit('// if can be null:\n', level);
          this.emit(`// ${nullable}\n`, level);
          hasNextSection = true;
        }
        if (typeof sensitive !== 'undefined') {
          if (hasNextSection) {
            this.emit(`// \n`, level);
          }
          this.emit('// if sensitive:\n', level);
          this.emit(`// ${sensitive}\n`, level);
          hasNextSection = true;
        }
        if (example) {
          if (hasNextSection) {
            this.emit(`// \n`, level);
          }
          const examples = _escape(example).split('\n');
          this.emit('// example:\n', level);
          this.emit(`// \n`, level);
          for (let j = 0; j < examples.length; j++) {
            if (examples[j] === '') {
              this.emit(`// \n`, level);
            } else {
              this.emit(`// ${examples[j]}\n`, level);
              if (j < examples.length - 1 && examples[j + 1] !== '') {
                this.emit(`// \n`, level);
              }
            }
          }
        }
      }
      this.emit(`${_format(fieldName)} `, level);
      let { type, omitemptyEnable } = this.visitFieldType(node, structName, fields, structMap, level);
      var realName = _getAttr(node, 'name');
      if (!realName) {
        realName = fieldName;
      }
      var tag = `json:"${realName}${omitemptyEnable ? ',omitempty' : ''}" xml:"${realName}${omitemptyEnable ? ',omitempty' : ''}"`;
      const anno = this.parseAnnotation(node, {
        'signed': 'string', 'encode': 'string'
        , 'pattern': 'string', 'maxLength': 'value', 'minLength': 'value',
        'maximum': 'value', 'minimum': 'value'
      });
      if (node.required) {
        tag = tag + ` require:"true"`;
      }
      if (anno !== '') {
        tag = tag + anno;
      }
      if (type !== '') {
        tag = tag + ` ${type}`;
      }
      this.emit(`\`${tag}\``);
      this.emit(`\n`);
    }
    if (node) {
      //find the last node's back comment
      let comments = DSL.comment.getBetweenComments(this.comments, node.tokenRange[0], ast.tokenRange[1]);
      this.visitComments(comments, level);
    }

    if (ast.nodes.length === 0) {
      //empty block's comment
      let comments = DSL.comment.getBetweenComments(this.comments, ast.tokenRange[0], ast.tokenRange[1]);
      this.visitComments(comments, level);
    }
    this.emit(`}\n`);
    this.emit(`\n`);
    this.eachGetFunc(nodes, lastName, 'model');
    this.eachSetFunc(nodes, lastName);
    this.visitValidate(nodes, lastName);
    for (let i = 0; i < fields.length; i++) {
      this.visitModelField(fields[i], structMap[i], level);
    }
  }

  emitModelArray(node, structMap, fields, structName) {
    if (node.fieldItemType.fieldType === 'array') {
      this.emit(`[]`);
      this.emitModelArray(node, structMap, fields, structName);
    } else if (node.fieldItemType.type === 'modelBody') {
      structMap.push(structName);
      this.emit(`*${structName} `);
      fields.push(node);
    } else {
      this.visitPointerType(node.fieldItemType);
      this.emit(` `);
    }
  }

  checkAnnotation(node, attrName) {
    for (let i = 0; i < node.attrs.length; i++) {
      if (attrName === _name(node.attrs[i].attrName)) {
        return true;
      }
    }
    return false;
  }

  parseAnnotation(node, annos) {
    var tag = '';
    for (let i = 0; i < node.attrs.length; i++) {
      const attrValueType = annos[_name(node.attrs[i].attrName)];
      if (attrValueType) {
        var attrName = _name(node.attrs[i].attrName);
        attrName = attrName.split('-').join('');
        tag = tag + ` ${attrName}:"${node.attrs[i].attrValue[attrValueType]}"`;
      }
    }
    return tag;
  }

  emitFuncArray(node, structName) {
    if (node.fieldItemType.fieldType === 'array') {
      this.emit(`[]`);
      this.emitFuncArray(node, structName);
    } else if (node.fieldItemType.type === 'modelBody') {
      this.emit(`*${structName}`);
    } else {
      this.visitPointerType(node.fieldItemType);
    }
  }

  eachGetFunc(nodes, structName, type = 'model', level = 0) {
    if(type === 'model') {
      this.emit(`func (s ${structName}) String() string {\n`, level);
      this.emit(`return dara.Prettify(s)\n`, level + 1);
      this.emit(`}\n`, level);
      this.emit(`\n`, level);
      this.emit(`func (s ${structName}) GoString() string {\n`, level);
      this.emit(`return s.String()\n`, level + 1);
      this.emit(`}\n`, level);
      this.emit(`\n`, level);
    } else { 
      this.builtinModule.push({
        path: 'fmt'
      });
      this.emit(`func (err ${structName}Error) Error() string {\n`, level);
      this.emit('if err.Message == nil {\n', level + 1);
      this.emit(`str := fmt.Sprintf("${structName}Error:\\n   Name: %s\\n   Code: %s\\n",\n`, level + 2);
      this.emit('dara.StringValue(err.Name), dara.StringValue(err.Code))\n', level + 3);
      this.emit('err.Message = dara.String(str)\n', level + 2);
      this.emit(`}\n`, level + 1);
      this.emit(`return dara.StringValue(err.Message)\n`, level + 1);
      this.emit(`}\n\n`, level);
    }


    // for (let i = 0; i < ast.extendFileds.length; i++) {
    //   const node = ast.extendFileds[i];
    //   this.visitGetFunc(node, structName, level);
    // }
    
    for (let i = 0; i < nodes.length; i++) {
      const node = nodes[i];
      this.visitGetFunc(node, structName, type, level);
    }
  }

  visitGetFunc(ast, structName, type = 'model',level = 0){
    const fieldName = _format(_name(ast.fieldName));
    this.emit(`func (s *${structName}${type === 'model' ? '' : 'Error'}) Get${fieldName}() `, level);
    structName = structName + _format(fieldName);
    this.visitFieldType(ast, structName, [], [], level);
    this.emit(' {\n');
    this.emit(`return s.${fieldName}\n`, level + 1);
    this.emit(`}\n`, level);
    this.emit(`\n`, level);
  }

  eachSetFunc(nodes, structName, level = 0) {
    // for (let i = 0; i < ast.extendFileds.length; i++) {
    //   const node = ast.extendFileds[i];
    //   this.visitSetFunc(node, structName, level);
    // }

    for (let i = 0; i < nodes.length; i++) {
      const node = nodes[i];
      this.visitSetFunc(node, structName, level);
    }
  }

  visitValidate(nodes, structName, level = 0) {
    this.emit(`func (s *${structName}) Validate() error {\n`, level);
    this.emit(`return dara.Validate(s)\n`, level + 1);
    this.emit(`}\n`, level);
    this.emit(`\n`, level);
  }


  visitSetFunc(ast, structName, level = 0){
    const fieldName = _format(_name(ast.fieldName));
    const fileldtype = structName + _format(fieldName);
    const itemName = structName + _format(fieldName);
    if (ast.fieldValue.fieldType === 'array') {
      if (_name(ast.fieldValue.fieldItemType)) {
        this.emit(`func (s *${structName}) Set${fieldName}(v []${this.getType(_name(ast.fieldValue.fieldItemType))}) *${structName} {\n`, level);
        this.emit(`s.${fieldName} = v\n`, level + 1);
        this.emit(`return s\n`, level + 1);
        this.emit(`}\n`, level);
        this.emit(`\n`, level);
      } else if (ast.fieldValue.fieldItemType.type === 'map') {
        this.emit(`func (s *${structName}) Set${fieldName}(v []`, level);
        this.visitType(ast.fieldValue.fieldItemType);
        this.emit(`) *${structName} {\n`);
        this.emit(`s.${fieldName} = v\n`, level + 1);
        this.emit(`return s\n`, level + 1);
        this.emit(`}\n`, level);
        this.emit(`\n`, level);
      } else if (ast.fieldValue.fieldItemType.type === 'modelBody') {
        this.emit(`func (s *${structName}) Set${fieldName}(v []${this.getType(fileldtype)}) *${structName} {\n`, level);
        this.emit(`s.${fieldName} = v\n`, level + 1);
        this.emit(`return s\n`, level + 1);
        this.emit(`}\n`, level);
        this.emit(`\n`, level);
      } else if (ast.fieldValue.fieldItemType.fieldType === 'array') {
        this.emit(`func (s *${structName}) Set${fieldName}(v [][]`, level);
        this.emitFuncArray(ast.fieldValue.fieldItemType, itemName);
        this.emit(`) *${structName} {\n`);
        this.emit(`s.${fieldName} = v\n`, level + 1);
        this.emit(`return s\n`, level + 1);
        this.emit(`}\n`, level);
        this.emit(`\n`, level);
      }
    } else if (ast.fieldValue.type === 'modelBody') {
      this.emit(`func (s *${structName}) Set${fieldName}(v *${fileldtype}) *${structName} {\n`, level);
      this.emit(`s.${fieldName} = v\n`, level + 1);
      this.emit(`return s\n`, level + 1);
      this.emit(`}\n`, level);
      this.emit(`\n`, level);
    } else if (_name(ast.fieldValue.fieldType) && ast.fieldValue.fieldType.idType === 'module') {
      const fieldType = ast.fieldValue.fieldType;
      this.emit(`func (s *${structName}) Set${fieldName}(v ${this.clientName[_name(fieldType)]}) *${structName} {\n`, level);
      this.emit(`s.${fieldName} = v\n`, level + 1);
      this.emit(`return s\n`, level + 1);
      this.emit(`}\n`, level);
      this.emit(`\n`, level);
    } else if (ast.fieldValue.fieldType.type === 'moduleModel' || ast.fieldValue.fieldType.type === 'moduleTypedef' || ast.fieldValue.fieldType.type === 'subModel') {
      this.emit(`func (s *${structName}) Set${fieldName}(v `, level);
      this.visitType(ast.fieldValue.fieldType, level);
      this.emit(`) *${structName} {\n`, level);
      this.emit(`s.${fieldName} = v\n`, level + 1);
      this.emit(`return s\n`, level + 1);
      this.emit(`}\n`, level);
      this.emit(`\n`, level);
    } else {
      var fieldType = '';
      if (!_name(ast.fieldValue.fieldType)) {
        fieldType = ast.fieldValue.fieldType;
      } else {
        fieldType = _name(ast.fieldValue.fieldType);
      }

      this.emit(`func (s *${structName}) Set${fieldName}(v `, level);
      this.visitType(ast.fieldValue, level, {});
      this.emit(`) *${structName} {\n`, level);
      if (!DSL.util.isBasicType(fieldType) || _isFilterType(fieldType)) {
        this.emit(`s.${fieldName} = v\n`, level + 1);
      } else {
        this.emit(`s.${fieldName} = &v\n`, level + 1);
      }
      this.emit(`return s\n`, level + 1);
      this.emit(`}\n`, level);
      this.emit(`\n`, level);
    }
  }

  visitExtendOn(extendOn, level, type = 'model') {
    if(!extendOn) {
      type === 'exception' ? this.emit(`dara.${CORE}Error\n`, level) : this.emit(`dara.Model\n`, level);
      return;
    }

    switch(_name(extendOn)) {
    case '$Error': 
      this.emit(`dara.${CORE}Error\n`, level);
      return;
    case '$ResponseError': 
      this.emit(`dara.ResponseError\n`, level);
      return;
    case '$Model': 
      this.emit(`dara.Model\n`, level);
      return;
    }
    
    if (extendOn.type === 'moduleModel') {
      const [moduleId, ...rest] = extendOn.path;
      this.emit(`${_importFilter(_name(moduleId).toLowerCase())}.`, level);
      this.emit('i');
      this.emit(rest.map((item) => {
        return _format(_name(item));
      }).join(''));
    } else if (extendOn.type === 'subModel') {
      this.emit(extendOn.path.map((item) => {
        return _format(_name(item));
      }).join(''), level);
      const [moduleId, ...rest] = extendOn.path;
      this.emit(`i${_format(_name(moduleId))}`);
      this.emit(rest.map((item) => {
        return _format(_name(item));
      }).join(''));
    } else {
      this.emit('i', level);
      if (extendOn.moduleName) {
        this.emit(`${_importFilter(_name(extendOn.moduleName).toLowerCase())}.`);
      }
      this.emit(`${_modelName(_format(_name(extendOn)))}`);
    }
    if(type === 'exception') {
      this.emit('Error');
    }
    this.emit('\n');
  }

  dealExtendFileds(ast) {
    const fileds = [];
    for (let i = 0; i < ast.nodes.length; i++) {
      const node = ast.nodes[i];
      const fieldName = _name(node.fieldName);
      fileds.push(fieldName);
    }
    const extendFileds = [];
    for (let i = 0; i < ast.extendFileds.length; i++) { 
      const node = ast.extendFileds[i];
      node.extend = true;
      const fieldName = _name(node.fieldName);
      if(fileds.includes(fieldName)) {
        continue;
      }
      extendFileds.push(node);

    }
    return extendFileds.concat(ast.nodes);
  }

  visitExceptionBody(ast, nodes, lastName, level) {
    assert.equal(ast.type, 'exceptionBody');
    const fields = [];
    const structMap = [];
    let node;
    for (let i = 0; i < nodes.length; i++) {
      node = nodes[i];
      if(!node.extend) {
        let comments = DSL.comment.getFrontComments(this.comments, node.tokenRange[0]);
        this.visitComments(comments, level);
      }
      if(i !== nodes.length - 1) {
        node.extend = false;
      }
      var fieldName = _name(node.fieldName);
      let tag = '';
      const structName = lastName + _format(fieldName);
      this.emit(`${_format(fieldName)} `, level);
      const { type } = this.visitFieldType(node, structName, fields, structMap, level);
      var realName = _getAttr(node, 'name');
      if (!realName) {
        realName = fieldName;
      }
      if (node.required) {
        tag = tag + ` require:"true"`;
      }
      if (type !== '') {
        tag = tag + ` ${type}`;
      }
      this.emit(`\`${tag}\``);
      this.emit(`\n`);
    }
    if (node && !node.extend) {
      //find the last node's back comment
      let comments = DSL.comment.getBetweenComments(this.comments, node.tokenRange[0], ast.tokenRange[1]);
      this.visitComments(comments, level);
    }

    if (ast.nodes.length === 0) {
      //empty block's comment
      let comments = DSL.comment.getBetweenComments(this.comments, ast.tokenRange[0], ast.tokenRange[1]);
      this.visitComments(comments, level);
    }
    this.emit(`}\n`);
    this.emit(`\n`);
    this.eachGetFunc(nodes, lastName, 'exception');
    for (let i = 0; i < fields.length; i++) {
      this.visitModelField(fields[i], structMap[i], level);
    }
  }

  visitException(ast, level) {
    assert.equal(ast.type, 'exception');
    const exceptionName = _upperFirst(_name(ast.exceptionName));
    this.visitAnnotation(ast.annotation, level);
    let comments = DSL.comment.getFrontComments(this.comments, ast.tokenRange[0]);
    this.visitComments(comments, level);
    this.emit(`type ${exceptionName}Error struct {\n`, level);
    // this.visitExtendOn(ast.extendOn, level + 1, 'exception');
    const nodes = this.dealExtendFileds(ast.exceptionBody);
    this.visitExceptionBody(ast.exceptionBody, nodes, exceptionName, level + 1);
  }

  visitExceptionInterface(ast, level) {
    assert.equal(ast.type, 'exception');
    const exceptionName = _upperFirst(_name(ast.exceptionName));
    this.emit(`type i${exceptionName}Error interface {\n`, level);
    // this.visitExtendOn(ast.extendOn, level + 1, 'exception');
    const nodes = this.dealExtendFileds(ast.exceptionBody);
    this.visitExceptionInterfceBody(ast.exceptionBody, nodes, exceptionName, level + 1);
    this.emit('}\n\n', level);
  }

  visitExceptionInterfceBody(ast, nodes, structName, level) {
    assert.equal(ast.type, 'exceptionBody');
    this.emit(`Error() string\n`, level);
    for (let i = 0; i < nodes.length; i++) {
      const node = nodes[i];
      const fieldName = _format(_name(node.fieldName));
      this.emit(`Get${fieldName}() `, level);
      this.visitFieldType(node, structName + fieldName, [], [], level);
      this.emit('\n');
    }
  }

  visitModelInterface(ast, level) {
    assert.equal(ast.type, 'model');
    const modelName = _format(_name(ast.modelName));
    this.emit(`type i${modelName} interface {\n`, level);
    this.visitExtendOn(ast.extendOn, level + 1, 'model');
    this.visitModelInterfaceBody(ast.modelBody, modelName, level + 1);
    this.emit('}\n\n', level);
  }

  visitModelInterfaceBody(ast, structName, level) {
    assert.equal(ast.type, 'modelBody');
    this.emit(`String() string\n`, level);
    this.emit(`GoString() string\n`, level);
    for (let i = 0; i < ast.nodes.length; i++) {
      const node = ast.nodes[i];
      const fieldName = _format(_name(node.fieldName));
      const fileldtype = structName + _format(fieldName);
      const itemName = structName + _format(fieldName);
      if (node.fieldValue.fieldType === 'array') {
        if (_name(node.fieldValue.fieldItemType)) {
          this.emit(`Set${fieldName}(v []${this.getType(_name(node.fieldValue.fieldItemType))}) *${structName}\n`, level);
        } else if (node.fieldValue.fieldItemType.type === 'map') {
          this.emit(`Set${fieldName}(v []`, level);
          this.visitType(node.fieldValue.fieldItemType);
          this.emit(`) *${structName}\n`);
        } else if (node.fieldValue.fieldItemType.type === 'modelBody') {
          this.emit(`Set${fieldName}(v []${this.getType(fileldtype)}) *${structName}\n`, level);
        } else if (node.fieldValue.fieldItemType.fieldType === 'array') {
          this.emit(`Set${fieldName}(v [][]`, level);
          this.emitFuncArray(node.fieldValue.fieldItemType, itemName);
          this.emit(`) *${structName}\n`);
        }
      } else if (node.fieldValue.type === 'modelBody') {
        this.emit(`Set${fieldName}(v *${fileldtype}) *${structName}\n`, level);
      } else if (_name(node.fieldValue.fieldType) && node.fieldValue.fieldType.idType === 'module') {
        const fieldType = node.fieldValue.fieldType;
        this.emit(`Set${fieldName}(v ${this.clientName[_name(fieldType)]}) *${structName}\n`, level);
      } else if (node.fieldValue.fieldType.type === 'moduleModel' || node.fieldValue.fieldType.type === 'moduleTypedef' || node.fieldValue.fieldType.type === 'subModel') {
        this.emit(`Set${fieldName}(v `, level);
        this.visitType(node.fieldValue.fieldType, level);
        this.emit(`) *${structName}\n`);
      } else {
        this.emit(`Set${fieldName}(v `, level);
        this.visitType(node.fieldValue, level, {});
        this.emit(`) *${structName}\n`);
      }
      this.emit(`Get${fieldName}() `, level);
      this.visitFieldType(node, itemName, [], [], level);
      this.emit('\n');
    }
  }

  visitModel(ast, level) {
    assert.equal(ast.type, 'model');
    const modelName = _format(_name(ast.modelName));
    this.visitAnnotation(ast.annotation, level);
    let comments = DSL.comment.getFrontComments(this.comments, ast.tokenRange[0]);
    this.visitComments(comments, level);
    this.emit(`type ${modelName} struct {\n`, level);
    // this.visitExtendOn(ast.extendOn, level + 1, 'model');
    const nodes = this.dealExtendFileds(ast.modelBody);
    this.visitModelBody(ast.modelBody, nodes, modelName, level + 1);
  }

  getModelFileds(ast) {
    const fileds = {};
    for (let i = 0; i < ast.modelBody.nodes.length; i++) {
      const node =ast.modelBody.nodes[i];
      const fieldName = _name(node.fieldName);
      fileds[fieldName] = node;
    }
    if(_name(ast.extendOn) === '$Error') {
      this.emit(`${_modelName(_format(_name(ast.extendOn)))}\n`);
      return;
    }
  }

  visitObjectFieldValue(ast, level, env, expected) {
    let setFunc;
    if(ast && ast.type === 'call') {
      let dealFunc = this.getVarDealFunc(ast, expected);
      setFunc = dealFunc && dealFunc(_name(ast.inferred));
    }
    if(setFunc) {
      this.emit(`${setFunc}`);
    }
    this.visitExpr(ast, level, env, expected);
    if(setFunc) {
      this.emit(')');
    }
  }

  visitObjectField(ast, level, env, expected) {
    assert.equal(ast.type, 'objectField');
    let comments = DSL.comment.getFrontComments(this.comments, ast.tokenRange[0]);
    this.visitComments(comments, level);
    var key = _name(ast.fieldName) || _string(ast.fieldName);
    this.emit(`"${key}": `, level);
    this.visitObjectFieldValue(ast.expr, level, env, expected);
    this.emit(`,\n`);
  }

  visitObject(ast, level, env, expected) {
    assert.equal(ast.type, 'object');
    if (ast.fields.length === 0) {
      if (ast.inferred && ast.inferred.type === 'map' &&
        !(_name(ast.inferred.keyType) === 'string' && _name(ast.inferred.valueType) === 'any')) {
        this.emit(`make(`);
        this.visitType(ast.inferred, level, env);
        this.emit(`)`);
      } else if (ast.inferred && ast.inferred.type === 'model') {
        this.emit('{}');
      } else if (ast.inferred && ast.inferred.type === 'array') {
        this.emit(`make(`);
        this.visitPointerType(expected.type ? expected : ast.inferred, level, env);
        this.emit(`, 0)`);
      } else {
        this.emit('map[string]interface{}{');
        let comments = DSL.comment.getBetweenComments(this.comments, ast.tokenRange[0], ast.tokenRange[1]);
        if (comments.length > 0) {
          this.emit('\n');
          this.visitComments(comments, level + 1);
          this.emit('', level);
        }
        this.emit(`}`);
      }
    } else {
      var hasExpandField = false;
      for (var i = 0; i < ast.fields.length; i++) {
        const field = ast.fields[i];
        if (!field) {
          continue;
        }
        if (field.type === 'expandField') {
          hasExpandField = true;
          break;
        }
      }

      var fieldType = false;
      if (!hasExpandField) {
        if (expected === 'map[string]interface{}') {
          this.emit('map[string]interface{}{\n');
        } else if (ast.inferred && ast.inferred.type === 'map') {
          expected = ast.inferred.valueType;
          if (!_isFilterType(_name(expected))) {
            fieldType = true;
          }
          this.visitType(ast.inferred, level, env);
          this.emit(`{\n`);
        } else if (ast.inferred && ast.inferred.type === 'model') {
          this.emit('{\n');
        } else if (ast.inferred && ast.inferred.type === 'array') {
          if (ast.inferred.itemType !== 'any') {
            fieldType = true;
          }
          this.visitPointerType(expected.type ? expected : ast.inferred, level, env);
          this.emit(`{\n`);
        } else {
          this.emit('map[string]interface{}{\n');
        }

        for (i = 0; i < ast.fields.length; i++) {
          this.visitObjectField(ast.fields[i], level + 1, env, expected.type ? { ...expected, pointer: fieldType } : { pointer: fieldType });
        }
        let comments = DSL.comment.getBetweenComments(this.comments, ast.fields[i - 1].tokenRange[0], ast.tokenRange[1]);
        this.visitComments(comments, level + 1);
        this.emit('}', level);
        return;
      }

      const fields = ast.fields.filter((field) => {
        return field.type === 'objectField';
      });

      const expandFields = ast.fields.filter((field) => {
        return field.type === 'expandField';
      });

      var isMerge = false;
      if (expandFields.length > 0) {
        isMerge = true;
        if (ast.inferred && ast.inferred.type === 'map' &&
          ast.inferred.valueType.name === 'string') {
          this.emit('dara.Merge(');
        } else {
          this.emit('dara.ToMap(');
        }
      }
      if (fields.length > 0) {
        var isFieldPointer = false;
        if (!_isFilterType(_name(ast.inferred.valueType))) {
          isFieldPointer = true;
        }
        this.visitType(ast.inferred, level, env);
        this.emit(`{\n`);
        for (i = 0; i < fields.length; i++) {
          this.visitObjectField(fields[i], level + 1, env, { pointer: isFieldPointer });
        }
        if (isMerge) {
          //find the last item's back comment
          let comments = DSL.comment.getBetweenComments(this.comments, ast.fields[i - 1].tokenRange[0], ast.tokenRange[1]);
          this.visitComments(comments, level + 1);
          this.emit('}, ', level);
        } else {
          //find the last item's back comment
          let comments = DSL.comment.getBetweenComments(this.comments, ast.fields[i - 1].tokenRange[0], ast.tokenRange[1]);
          this.visitComments(comments, level + 1);
          this.emit('}', level);
        }
      }

      for (let i = 0; i < expandFields.length; i++) {
        this.visitExpr(expandFields[i].expr, level + 1, env);
        if (expandFields.length > 1 && i < expandFields.length - 1) {
          this.emit(',\n');
          this.emit(``, level + 1);
        }
      }
      if (isMerge) {
        this.emit(')');
      }
    }
  }

  visitArgs(args, level, env, expected, argHasThrowFunc) {
    this.emit('(');
    for (let i = 0; i < args.length; i++) {
      const expr = args[i];
      if(expr.yieldArg) {
        this.visitYieldArgs(expr.ast, level, env);
      } else if (expr.needCast) {
        this.emit('dara.ToMap(');
        if (argHasThrowFunc && argHasThrowFunc.get(i)) {
          this.emit(argHasThrowFunc.get(i));
        } else {
          this.visitExpr(expr, level, env);
        }
        this.emit(')');
      } else {
        if ((expr.expectedType.name === 'number' || expr.expectedType.name === 'integer') && expr.inferred.name === 'int32') {
          this.emit(`dara.ToInt(`);
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, expected);
          }
          this.emit(`)`);
        } else if (expr.type !== 'number' && (expr.inferred.name === 'number' || expr.inferred.name === 'integer') && expr.expectedType.name === 'int32') {
          this.emit(`dara.ToInt32(`);
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, expected);
          }
          this.emit(`)`);
        } else {
          let setFunc;
          if(expr && expr.type === 'call') {
            let dealFunc = this.getVarDealFunc(expr, expected);
            setFunc = dealFunc && dealFunc(_name(expr.inferred));
          }
          if(setFunc) {
            this.emit(`${setFunc}`);
          }
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, expected);
          }
          if(setFunc) {
            this.emit(')');
          }
        }
      }
      if (i !== args.length - 1) {
        this.emit(', ');
      }
    }
    this.emit(')');
  }

  visitMethodCall(ast, level, env, argHasThrowFunc) {
    assert.equal(ast.left.type, 'method_call');
    const name = _format(_name(ast.left.id));
    if (name.startsWith('$') && this.builtin[name]) {
      const method = name.replace('$', '');
      this.builtin[name][method](ast.args, level, env, argHasThrowFunc);
      return;
    } 
    if (!ast.isStatic) {
      this.emit(`client.`);
    }

    this.emit(name);
    this.visitArgs(ast.args, level, env, { pointer: true }, argHasThrowFunc);
  }

  visitInstanceCall(ast, level, env, argHasThrowFunc) {
    assert.equal(ast.left.type, 'instance_call');
    const method = _name(ast.left.propertyPath[0]);
    const builtinInstance = ast.builtinModule && this.builtin[ast.builtinModule];
    if(builtinInstance && this.builtin[ast.builtinModule][method]) {
      this.builtin[ast.builtinModule][method](ast, level, {}, argHasThrowFunc, { pointer: false });
      return;
    }

    if (_name(ast.left.id).indexOf('@') === 0) {
      this.emit(`client.${_format(_name(ast.left.id).substring(1))}.${_format(method)}`);
    } else {
      this.emit(`${_name(ast.left.id)}.${_format(method)}`);
    }
    this.visitArgs(ast.args, level, env, { pointer: builtinInstance ? false : true }, argHasThrowFunc);
  }

  visitCall(ast, level, env, argHasThrowFunc) {
    assert.equal(ast.type, 'call');
    if (ast.left.type === 'method_call') {
      this.visitMethodCall(ast, level, env, argHasThrowFunc);
    } else if (ast.left.type === 'instance_call') {
      this.visitInstanceCall(ast, level, env, argHasThrowFunc);
    } else if (ast.left.type === 'static_call') {
      this.visitStaticCall(ast, level, env, argHasThrowFunc);
    } else {
      throw new Error('un-implemented');
    }
  }

  visitBuiltinStaticCall(ast, level, env, argHasThrowFunc) {
    const moduleName = _name(ast.left.id);

    const builtiner = this.builtin[moduleName];
    if(!builtiner) {
      throw new Error('un-implemented');
    }
    const func = _name(ast.left.propertyPath[0]);
    builtiner[func](ast.args, level, env, argHasThrowFunc);
  }

  visitStaticCall(ast, level, env, argHasThrowFunc) {
    assert.equal(ast.left.type, 'static_call');
    if(ast.left.id.type === 'builtin_module') {
      this.visitBuiltinStaticCall(ast, level, env, argHasThrowFunc);
      return;
    }
    this.visitModuleName(ast.left.id);
    this.emit(`.${_format(_name(ast.left.propertyPath[0]))}(`);
    for (let i = 0; i < ast.args.length; i++) {
      const expr = ast.args[i];
      if(expr.yieldArg) {
        this.visitYieldArgs(expr.ast, level, env);
      } else if (expr.needCast) {
        this.emit('dara.ToMap(');
        if (argHasThrowFunc && argHasThrowFunc.get(i)) {
          this.emit(argHasThrowFunc.get(i));
        } else {
          this.visitExpr(expr, level, env);
        }
        this.emit(')');
      } else {
        if ((expr.expectedType.name === 'number' || expr.expectedType.name === 'integer') && expr.inferred.name === 'int32') {
          this.emit(`dara.ToInt(`);
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, { pointer: true });
          }
          this.emit(`)`);
        } else if (expr.type !== 'number' && (expr.inferred.name === 'number' || expr.inferred.name === 'integer') && expr.expectedType.name === 'int32') {
          this.emit(`dara.ToInt32(`);
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, { pointer: true });
          }
          this.emit(`)`);
        } else {
          let setFunc;
          if(expr && expr.type === 'call') {
            let dealFunc = this.getVarDealFunc(expr, { pointer: true });
            setFunc = dealFunc && dealFunc(_name(expr.inferred));
          }
          if(setFunc) {
            this.emit(`${setFunc}`);
          }
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, { pointer: true });
          }
          if(setFunc) {
            this.emit(')');
          }
        } 
      }
      if (i !== ast.args.length - 1) {
        this.emit(', ');
      }
    }
    this.emit(')');
  }

  visitPropertyAccess(ast, level, env, expected) {
    assert.ok(ast.type === 'property_access' || ast.type === 'property');

    var id = _name(ast.id);

    if(ast.id.tag === Tag.VID){
      id = _vid(ast.id);
    }

    var expr = '';
    if (id === '__response') {
      expr += RESPONSE;
    } else if (id === '__request') {
      expr += REQUEST;
    } else {
      expr += _avoidReserveName(id);
    }
    var fieldType = '';
    var current = ast.id.inferred;
    var isMap = '';
    for (var i = 0; i < ast.propertyPath.length; i++) {
      const name = _name(ast.propertyPath[i]);
      if (current.type === 'model') {
        if (ast.inferred.type === 'array') {
          fieldType = `[]${this.getType(_name(ast.inferred.itemType), false)}`;
        } else {
          fieldType = ast.propertyPathTypes[i].name;
        }
        expr += `.${_format(name)}`;
      } else {
        fieldType = ast.propertyPathTypes[i].name;
        expr += `["${name}"]`;
      }
      if (current.type === 'map') {
        isMap = true;
      }
      current = ast.propertyPathTypes[i];
    }
    if (expected && expected.needCast === 'false') {
      this.emit(expr);

    } else if (expected && expected.pointer) {
      if (ast.id.inferred.type !== 'model' && !isMap && (DSL.util.isBasicType(fieldType) ||
        (current.type === 'array' && current.itemType.type === 'basic')) && !_isFilterType(fieldType)) {
        this.emit(`${_setExtendFunc(fieldType)}${expr})`);
      } else {
        this.emit(`${expr}`);
      }
    } else if ((ast.id.inferred.type === 'model' || isMap) && (DSL.util.isBasicType(fieldType) ||
      (current.type === 'array' && current.itemType.type === 'basic')) && !_isFilterType(fieldType)) {
      this.emit(`${_setValueFunc(fieldType)}${expr})`);
    } else {
      this.emit(expr);
    }
  }

  visitExpr(ast, level, env, expected, argHasThrowFunc) {
    var isPointer = false;
    env = env || {
      pointerParams: [],
      local: new Map()
    };
    if (ast.type === 'boolean') {
      if (expected && expected.pointer) {
        this.emit(`dara.Bool(${ast.value})`);
      } else {
        this.emit(ast.value);
      }
    } else if (ast.type === 'null') {
      this.emit('nil');
    } else if (ast.type === 'property_access') {
      this.visitPropertyAccess(ast, level, env, expected);
    } else if (ast.type === 'string') {
      if (expected && expected.pointer) {
        this.emit(`dara.String("${_string(ast.value).replace(/"/g, '\\"')}")`);
      } else {
        this.emit(`"${_string(ast.value).replace(/"/g, '\\"')}"`);
      }
    } else if (ast.type === 'number') {
      if (expected && expected.pointer) {
        if (_name(expected) && _name(expected) !== 'any') {
          this.emit(`${_setExtendFunc(_name(expected))}${ast.value.value})`);
        } else if (_name(expected) === 'any') {
          this.emit(`${_setExtendFunc(ast.value.type)}${ast.value.value})`);
        } else {
          this.emit(`dara.Int(${ast.value.value})`);
        }
      } else {
        this.emit(ast.value.value);
      }
    } else if (ast.type === 'object') {
      this.visitObject(ast, level, env, expected);
    } else if (ast.type === 'variable') {
      isPointer = false;
      var id = _name(ast.id);
      if ((env.pointerParams && !env.pointerParams.includes(id)) && (expected && expected.pointer) && (ast.inferred.type !== 'model' && 
      ((DSL.util.isBasicType(_name(ast.inferred)) && !_isFilterType(_name(ast.inferred))) &&
      !(ast.inferred.type === 'array' && ast.inferred.itemType.type === 'basic')))) {
        isPointer = true;
        this.emit(_setExtendFunc(_name(ast.inferred)));
      }

      if((env.pointerParams && env.pointerParams.includes(id)) && (!expected || !expected.pointer) && 
      ((ast.inferred.type === 'array' && ast.inferred.itemType.type === 'basic') || 
      (DSL.util.isBasicType(_name(ast.inferred)) && !_isFilterType(_name(ast.inferred))))) {
        isPointer = _setValueFunc(_name(ast.inferred)) ? true : false;
        this.emit(_setValueFunc(_name(ast.inferred)));
      }
      
      if (id === '__response') {
        this.emit(RESPONSE);
      } else if (id === '__request') {
        this.emit(REQUEST);
      } else if (ast.inferred && _name(ast.inferred) === 'class') {
        this.emit('new(' + _avoidReserveName(id) + ')');
      } else {
        this.emit(_avoidReserveName(id));
      }

      if (isPointer) {
        this.emit(')');
      }
    } else if (ast.type === 'virtualVariable') {
      if ((!expected || !expected.pointer) && ((ast.inferred.type === 'array' &&
        ast.inferred.itemType.type === 'basic') || (DSL.util.isBasicType(_name(ast.inferred))
          && !_isFilterType(_name(ast.inferred))))) {
        this.emit(`${_setValueFunc(_name(ast.inferred))}${_vid((_name(ast.vid)))})`);
      } else {
        this.emit(`${_vid(_avoidVariableKeywords(_format(_name(ast.vid))))}`);
      }
    } else if (ast.type === 'decrement') {
      if (ast.position === 'front') {
        this.emit('--');
      }
      this.visitExpr(ast.expr, level, env, { pointer: false }, argHasThrowFunc);
      if (ast.position === 'backend') {
        this.emit('--');
      }
    } else if (ast.type === 'increment') {
      if (ast.position === 'front') {
        this.emit('++');
      }
      this.visitExpr(ast.expr, level, env, { pointer: false }, argHasThrowFunc);
      if (ast.position === 'backend') {
        this.emit('++');
      }
    } else if (ast.type === 'template_string') {
      var j = 0;
      if (expected && expected.pointer) {
        this.emit(`dara.String(`);
      }
      for (let i = 0; i < ast.elements.length; i++) {
        var item = ast.elements[i];
        if (item.type === 'element' && _string(item.value) === '') {
          continue;
        }
        if (j > 0) {
          this.emit(' + ');
        }
        j = j + 1;
        if (item.type === 'element') {
          let val = _string(item.value);
          val = val.replace(/[\n]/g, '" + \n"');
          this.emit(`"${val}"`);
        } else if (item.type === 'expr') {
          const expr = item.expr;
          if (expr.inferred && _name(expr.inferred) !== 'string') {
            this.emit(`dara.ToString(`);
            this.visitExpr(expr, level, env);
            this.emit(`)`);
          } else {
            let pointer = false;
            if(expr.type === 'call' && this.isNotBuiltin(expr)) {
              pointer = _setValueFunc(_name(ast.inferred)) ? true : false;
              this.emit(`${_setValueFunc(_name(ast.inferred))}`);
            }
            this.visitExpr(expr, level, env);
            if(pointer) {
              this.emit(')');
            }
          }
        } else {
          throw new Error('unimpelemented');
        }
      }
      if (expected && expected.pointer) {
        this.emit(`)`);
      }
    } else if (ast.type === 'call') {
      if(_isIterator(ast.inferred)) {
        env.yieldFunc ? this.emit('', level) : this.emit('go ', level);
      }
      this.visitCall(ast, level, env, argHasThrowFunc);
    } else if (ast.type === 'group') {
      this.emit('(');
      this.visitExpr(ast.expr, level, env, expected, argHasThrowFunc);
      this.emit(')');
    } else if (_isBinaryOp(ast.type)) {
      if ((env.pointerParams && env.pointerParams.includes(id)) || (expected && expected.pointer)) {
        this.emit(`${_setExtendFunc(_name(ast.inferred))}`);
      }
      this.visitExpr(ast.left, level, env, { pointer: false }, argHasThrowFunc);
      if (ast.type === 'or') {
        this.emit(' || ');
      } else if (ast.type === 'add') {
        this.emit(' + ');
      } else if (ast.type === 'subtract') {
        this.emit(' - ');
      } else if (ast.type === 'div') {
        this.emit(' / ');
      } else if (ast.type === 'multi') {
        this.emit(' * ');
      } else if (ast.type === 'and') {
        this.emit(' && ');
      } else if (ast.type === 'or') {
        this.emit(' || ');
      } else if (ast.type === 'lte') {
        this.emit(' <= ');
      } else if (ast.type === 'lt') {
        this.emit(' < ');
      } else if (ast.type === 'gte') {
        this.emit(' >= ');
      } else if (ast.type === 'gt') {
        this.emit(' > ');
      } else if (ast.type === 'neq') {
        this.emit(' != ');
      } else if (ast.type === 'eq') {
        this.emit(' == ');
      }
      this.visitExpr(ast.right, level, env, { pointer: false }, argHasThrowFunc);
      if (expected && expected.pointer) {
        this.emit(')');
      }
    } else if (ast.type === 'construct') {
      this.visitConstruct(ast, level, env);
    } else if (ast.type === 'construct_model') {
      this.visitConstructModel(ast, level, env);
    } else if (ast.type === 'not') {
      this.emit(`!`);
      this.visitExpr(ast.expr, level, env, expected, argHasThrowFunc);
    } else if (ast.type === 'array') {
      this.visitArray(ast, level, env, expected);
    } else if (ast.type === 'map_access') {
      this.visitMapAccess(ast, level, env, expected);
    } else if (ast.type === 'array_access') {
      this.visitArrayAccess(ast, level, env, expected);
    } else if (ast.type === 'super') {
      this.visitSuper(ast, level, env);
    } else {
      throw new Error('unimpelemented');
    }
  }

  visitYieldArgs(ast, level, env) {
    let pointer = false;
    let setFunc;
    var hasThrowCall = (ast.expr.type === 'call' &&
      ast.expr.hasThrow) || ast.expr.type === 'construct';
    if (ast.left.type === 'property_assign' || ast.left.type === 'property') {
      pointer = true;
      this.visitPropertyAccess(ast.left, level, env, { needCast: 'false', type: 'pointer', pointer });
    } else if (ast.left.type === 'variable') {
      const name = _name(ast.left.id);
      pointer = env.pointerParams && env.pointerParams.includes(name);
      this.emit(`${name}`);
      if(hasThrowCall){
        let dealFunc = this.getVarDealFunc(ast.expr,  { pointer });
        setFunc = dealFunc && dealFunc(_name(ast.expr.inferred));
        if(setFunc) {
          this.emit('Tmp');
        }
      }
    } else if (ast.left.type === 'virtualVariable') {
      pointer = true;
      this.emit(`${_vid(_avoidVariableKeywords(_format(_name(ast.left.vid))))}`, level);
    } else if (ast.left.type === 'map_access') {
      pointer = true;
      this.visitMapAccess(ast.left, level, env, { pointer });
    } else if (ast.left.type === 'array_access') {
      pointer = true;
      this.visitArrayAccess(ast.left, level, env, { pointer });
    } else {
      throw new Error('unimpelemented');
    }

    if (hasThrowCall) {
      this.emit(', _yieldErr');
    }
  }

  visitSuper(ast, level, env) {
    assert.equal(ast.type, 'super');
    this.emit(`_err = client.${this.structName}.Init(`);
    for (let i = 0; i < ast.args.length; i++) {
      this.visitExpr(ast.args[i], level, env, { pointer: true });
    }
    this.emit(`)\n`, level);
    if(env.runtimeBody){
      this.visitAPIErrCatch(level, env);
    } else {
      this.emit(`if _err != nil {\n`, level);
      this.emit(`return _err\n`, level + 1);
      this.emit(`}`, level);
    }
  }

  visitMapAccess(ast, level, env, expected) {
    assert.equal(ast.type, 'map_access');
    let mapName = _name(ast.id);
    if (ast.id.tag === Tag.VID) {
      mapName = _vid(mapName);
    }
    if (ast.propertyPath && ast.propertyPath.length) {
      var current = ast.id.inferred;
      for (var i = 0; i < ast.propertyPath.length; i++) {
        var name = _name(ast.propertyPath[i]);
        if (current.type === 'model') {
          mapName += `.${_format(name)}`;
        } else {
          mapName += `["${name}"]`;
        }
        current = ast.propertyPathTypes[i];
      }
    }

    let accessKey = ast.accessKey;
    if (!DSL.util.isBasicType(ast.accessKey.type) && !ast.accessKey.inferred) {
      ast.accessKey.inferred = {
        type: 'basic',
        name: 'string'
      };
    }

    if ((!expected || !expected.pointer) && (DSL.util.isBasicType(_name(ast.inferred)) &&
      !_isFilterType(_name(ast.inferred))) && _name(ast.inferred) !== 'model') {
      this.emit(`${_setValueFunc(_name(ast.inferred))}${mapName}[`);
      this.visitExpr(accessKey, level, env);
      this.emit(`])`);
    } else {
      this.emit(`${mapName}[`);
      if (ast.propertyPathTypes && ast.propertyPathTypes.length) {
        if (ast.propertyPathTypes[ast.propertyPathTypes.length - 1].type === 'map') {
          accessKey = {
            inferred: ast.propertyPathTypes[ast.propertyPathTypes.length - 1].keyType,
            ...accessKey
          };
        }
      }
      this.visitExpr(accessKey, level, env);
      this.emit(`]`);
    }
  }

  visitArrayAccess(ast, level, env, expected) {
    assert.equal(ast.type, 'array_access');
    let arrayName = _name(ast.id);
    if (ast.id.tag === DSL.Tag.Tag.VID) {
      arrayName = _vid(arrayName);
    }
    if (ast.propertyPath && ast.propertyPath.length) {
      var current = ast.id.inferred;
      for (var i = 0; i < ast.propertyPath.length; i++) {
        var name = _name(ast.propertyPath[i]);
        if (current.type === 'model') {
          arrayName += `.${_format(name)}`;
        } else {
          arrayName += `["${name}"]`;
        }
        current = ast.propertyPathTypes[i];
      }
    }
    if ((!expected || !expected.pointer) && (DSL.util.isBasicType(_name(ast.inferred)) &&
      !_isFilterType(_name(ast.inferred))) && _name(ast.inferred) !== 'model') {
      this.emit(`${_setValueFunc(_name(ast.inferred))}${arrayName}[`);
      this.visitExpr(ast.accessKey, level, env);
      this.emit(`])`);
    } else {
      this.emit(`${arrayName}[`);
      this.visitExpr(ast.accessKey, level, env);
      this.emit(`]`);
    }
  }

  visitArray(ast, level, env, expected) {
    assert.equal(ast.type, 'array');
    let expectedType;
    if (!expected || !expected.type || (expected && _name(expected) && _name(expected) === 'any')) {
      expectedType = ast.inferred;
    } else {
      expectedType = expected;
    }
    this.visitPointerType(expectedType, level, env);
    let arrayComments = DSL.comment.getBetweenComments(this.comments, ast.tokenRange[0], ast.tokenRange[1]);
    if (ast.items.length === 0) {
      this.emit(`{`);
      if (arrayComments.length > 0) {
        this.emit('\n');
        this.visitComments(arrayComments, level + 1);
        this.emit('', level);
      }
      this.emit('}');
      return;
    }
    let item;
    this.emit(`{`);
    let itemType = expectedType.subType || expectedType.itemType;
    itemType.pointer = true;
    for (let i = 0; i < ast.items.length; i++) {
      item = ast.items[i];
      let comments = DSL.comment.getFrontComments(this.comments, item.tokenRange[0]);
      if (comments.length > 0) {
        this.emit('\n');
        this.visitComments(comments, level + 1);
        this.emit('', level + 1);
      }
      this.visitExpr(item, level + 1, env, itemType);
      if (i < ast.items.length - 1) {
        this.emit(`, `);
      }
    }
    if (item) {
      //find the last item's back comment
      let comments = DSL.comment.getBetweenComments(this.comments, item.tokenRange[0], ast.tokenRange[1]);
      if (comments.length > 0) {
        this.emit('\n');
        this.visitComments(comments, level + 1);
        this.emit(`}`, level);
        return;
      }
    }
    this.emit(`}`);
  }

  visitConstruct(ast, level, env) {
    assert.equal(ast.type, 'construct');
    let moudleName = _name(ast.aliasId);
    let constructFunc = this.constructFunc[moudleName];
    if(moudleName.startsWith('$') && this.builtin[moudleName]) {
      constructFunc = moudleName.replace('$', '');
      moudleName = this.getType(moudleName, false);
      this.emit(moudleName);
    } else {
      this.visitModuleName(ast.aliasId);
    }
    this.emit(`.New${constructFunc}(`);
    for (let i = 0; i < ast.args.length; i++) {
      this.visitExpr(ast.args[i], level, env);
    }
    this.emit(`)`);
  }

  visitConstructModelField(ast, level, env) {
    var str = '';
    var expected = ast.expectedType;
    expected.pointer = true;
    if (ast.expr.type === 'property_access') {
      expected.needCast = 'false';
    }
    let comments = DSL.comment.getFrontComments(this.comments, ast.tokenRange[0]);
    this.visitComments(comments, level);
    this.emit(`${_format(_name(ast.fieldName))}: ${str}`, level);
    if ((_name(ast.expectedType) === 'number' || _name(ast.expectedType) === 'integer') && _name(ast.expr.inferred) === 'int32') {
      this.emit(`dara.ToInt(`);
      this.visitExpr(ast.expr, level, env, expected);
      this.emit(`)`);
    } else if (ast.expr.type !== 'number' && (_name(ast.expr.inferred) === 'number' || _name(ast.expr.inferred) === 'integer') && _name(ast.expectedType) === 'int32') {
      this.emit(`dara.ToInt32(`);
      this.visitExpr(ast.expr, level, env, expected);
      this.emit(`)`);
    } else {
      let setFunc;
      if(ast.expr && ast.expr.type === 'call') {
        let dealFunc = this.getVarDealFunc(ast.expr, expected);
        setFunc = dealFunc && dealFunc(_name(ast.expr.inferred));
      }
      if(setFunc) {
        this.emit(`${setFunc}`);
      }
      this.visitExpr(ast.expr, level, env, expected);
      if(setFunc) {
        this.emit(')');
      }
    }
    if (str !== '') {
      this.emit(`)`);
    }
    this.emit(`,\n`);
  }

  visitConstructModel(ast, level, env) {
    assert.equal(ast.type, 'construct_model');
    let modelName = '';
    this.emit('&');
    if ((ast.inferred && ast.inferred.moduleName) || ast.aliasId.isModule) {
      this.visitModuleName(ast.aliasId);
      for (let i = 0; i < ast.propertyPath.length; i++) {
        const item = ast.propertyPath[i];
        modelName += _format(_name(item));
      }
      const externEx = this.usedExternException.get(_name(ast.aliasId));
      if (externEx && externEx.has(modelName)) {
        modelName += 'Error';
      }
      modelName = `.${modelName}`;
    } else {
      modelName = _modelName(_format(_name(ast.aliasId)));
      if (this.predefined[modelName] && this.predefined[modelName].isException) { 
        modelName += 'Error';
      }
    }
    if (ast.object && ast.object.fields.length > 0) {
      const fields = ast.object.fields;
      this.emit(`${modelName}{\n`);
      for (let i = 0; i < fields.length; i++) {
        const field = fields[i];
        this.visitConstructModelField(field, level + 1, env);  
      }
      this.emit(`}`, level);
    } else {
      this.emit(`${modelName}{`);
      let comments = DSL.comment.getBetweenComments(this.comments, ast.tokenRange[0], ast.tokenRange[1]);
      if (comments.length > 0) {
        this.emit('\n');
        this.visitComments(comments, level + 1);
        this.emit('', level);
      }
      this.emit(`}`);
    }
  }

  visitYield(ast, level, env) {
    assert.equal(ast.type, 'yield');
    // TODO
    if (ast.expr.type === 'null' || ast.expr.type === 'variable' ||
      ast.expr.type === 'property_access' || ast.expr.type === 'virtualVariable' || 
      ast.expr.type === 'decrement' ||  ast.expr.type === 'increment' || ast.expr.type === 'construct_model' ||
      ast.expr.type === 'group' || _isBinaryOp(ast.expr.type)) {
      this.emit(`_yield <- `, level);
      this.visitExpr(ast.expr, level, env, { pointer: true });
      this.emit(`\n`);
      return;
    }

    var returnType = '';
    if (env.returnType.idType === 'module') {
      this.emit(`yield <- ${this.clientName[_name(env.returnType)].replace('*', '&')}{}\n`, level);
    } else if (_name(env.returnType) && !(DSL.util.isBasicType(_name(env.returnType)) && !_isFilterType(_name(env.returnType)))
      && env.returnType.idType !== 'typedef') {
      this.emit(`yield <- ${_initValue(_name(env.returnType))}\n`, level);
    } else if (env.returnType.path && env.returnType.type !== 'moduleTypedef') {
      for (let i = 0; i < env.returnType.path.length; i++) {
        const path = env.returnType.path[i];
        if (i === 0) {
          returnType += _name(path).toLowerCase();
        } else {
          returnType += '.' + _name(path);
        }
      }
      this.emit(`yield <- ${_initValue(returnType)}\n`, level);
    } else if (env.returnType.type === 'map') {
      this.emit(`yield <- make(`, level);
      this.visitPointerType(env.returnType, level, env);
      this.emit(`)\n`);
    } else if (env.returnType.type === 'array') {
      this.emit(`yield <- make(`, level);
      this.visitPointerType(env.returnType, level, env);
      this.emit(`, 0)\n`);
    } 

    if (ast.expr.type === 'call') {
      var argHasThrowFunc = this.visitFunctionNested(ast.expr, level, env);
      var hasThrow = false;
      hasThrow = ast.expr.hasThrow;
      if (hasThrow) {
        this.emit(`_body, _err := `, level);
      } else {
        this.emit(`_body := `, level);
      }
      this.visitExpr(ast.expr, level, env, { pointer: true }, argHasThrowFunc);
      this.emit(`\n`);
      if (hasThrow) {
        this.emit(`if _err != nil {\n`, level);
        this.emit(`_yieldErr <- _err\n`, level + 1);
        this.emit(`return\n`, level + 1);
        this.emit(`}\n`, level);
      }
      if (!ast.needCast) {
        this.emit(`yield <- _body\n`, level);
      } else {
        if (env.hasThrow) {
          this.emit(`_err := dara.ConvertChan(_body, _yield)\n`, level);
          this.emit(`if _err != nil {\n`, level);
          this.emit(`_yieldErr <- _err\n`, level + 1);
          this.emit(`return\n`, level + 1);
          this.emit(`}\n`, level);
        } else {
          this.emit(`dara.ConvertChan(_body, _yield)\n`, level);
        }
      }
    } else if (ast.expr.type === 'template_string') {
      this.emit(`yield <- `, level);
      this.visitExpr(ast.expr, level, env, { pointer: true });
      this.emit(`\n`);
    } else if (ast.expr.fields) {
      // for (let i = 0; i < ast.expr.fields.length; i++) {
      //   const field = ast.expr.fields[i];
      //   if (field.expr.inferred && _name(field.expr.inferred) === 'readable') {
      //     this.emit(
      //       `_result.${_format(_name(field.fieldName))} = `,
      //       level
      //     );
      //     this.visitExpr(field.expr);
      //     this.emit(`\n`);
      //     delete ast.expr.fields[i];
      //   }
      // }
      if (env.hasThrow) {
        this.emit(`_err := dara.ConvertChan(`, level);
      } else {
        this.emit(`dara.ConvertChan(`, level);
      }
      this.visitExpr(ast.expr, level, env, { pointer: true });
      this.emit(`, _yield)\n`);
      if (env.hasThrow) {
        this.emit(`if _err != nil {\n`, level);
        this.emit(`_yieldErr <- _err\n`, level + 1);
        this.emit(`return\n`, level + 1);
        this.emit(`}\n`, level);
      }
    } else if (ast.expr.items) {
      if (env.hasThrow) {
        this.emit(`_err := dara.ConvertChan(`, level);
      } else {
        this.emit(`dara.ConvertChan(`, level);
      }
      this.visitExpr(ast.expr, level, env, { pointer: true });
      this.emit(`, _yield)\n`);
      if (env.hasThrow) {
        this.emit(`if _err != nil {\n`, level);
        this.emit(`_yieldErr <- _err\n`, level + 1);
        this.emit(`return\n`, level + 1);
        this.emit(`}\n`, level);
      }
    } else if (ast.expr.type === 'construct') {
      this.emit(`_result, _err := `, level);
      this.visitConstruct(ast.expr, level, env);
      this.emit(`\n`);
      this.emit(`if _err != nil {\n`, level);
      this.emit(`_yieldErr <- _err\n`, level + 1);
      this.emit(`return\n`, level + 1);
      this.emit(`}\n`, level);
    } else if (ast.expr.type === 'map_access') {
      this.emit(`yield <- `, level);
      this.visitMapAccess(ast.expr, level, env, { pointer: true });
      this.emit(`\n`);
    } else if (DSL.util.isBasicType(ast.expr.type)) {
      this.emit(`yield <- `, level);
      this.visitExpr(ast.expr, level, env, { pointer: true });
      this.emit(`\n`);
    }
  }

  visitReturn(ast, level, env) {
    assert.equal(ast.type, 'return');
    if(env.yieldErrDeal) {
      this.emit(`_err = <- _yieldErr\n`, level);
    }
    if (!ast.expr && !env.finallyBlock) {
      if (env.hasThrow && env.runtimeBody) {
        this.emit(`return _err\n`, level);
      } else {
        this.emit(`return\n`, level);
      }
      return;
    }

    if(_name(env.returnType) === 'void') {
      if (env.hasThrow && env.runtimeBody) {
        this.emit(`return _err\n`, level);
      } else {
        this.emit(`return\n`, level);
      }
      return;
    }

    if (ast.expr.type === 'null' || ast.expr.type === 'variable' ||
      ast.expr.type === 'property_access' || ast.expr.type === 'virtualVariable' || 
      ast.expr.type === 'decrement' ||  ast.expr.type === 'increment' || 
      ast.expr.type === 'group' || _isBinaryOp(ast.expr.type)) {
      this.emit(`_result = `, level);
      this.visitExpr(ast.expr, level, env, { pointer: true });
      this.emit(`\n`);
      if(!env.finallyBlock) {
        if (env.hasThrow) {
          this.emit(`return _result , _err`, level);
        } else {
          this.emit(`return _result`, level);
        }
        this.emit(`\n`);
        env.hasReturn = false;
      }
      return;
    }

    // if (env.returnType.idType === 'module') {
    //   this.emit(`_result = ${this.clientName[_name(env.returnType)].replace('*', '&')}{}\n`, level);
    // } else if (_name(env.returnType) && !(DSL.util.isBasicType(_name(env.returnType)) && !_isFilterType(_name(env.returnType)))
    //   && env.returnType.idType !== 'typedef') {
    //   this.emit(`_result = ${_initValue(_name(env.returnType))}\n`, level);
    // } else if (env.returnType.path && env.returnType.type !== 'moduleTypedef') {
    //   for (let i = 0; i < env.returnType.path.length; i++) {
    //     const path = env.returnType.path[i];
    //     if (i === 0) {
    //       returnType += _name(path).toLowerCase();
    //     } else {
    //       returnType += '.' + _name(path);
    //     }
    //   }
    //   this.emit(`_result = ${_initValue(returnType)}\n`, level);
    // } else if (env.returnType.type === 'map') {
    //   this.emit(`_result = make(`, level);
    //   this.visitPointerType(env.returnType, level, env);
    //   this.emit(`)\n`);
    // } else if (env.returnType.type === 'array') {
    //   this.emit(`_result = make(`, level);
    //   this.visitPointerType(env.returnType, level, env);
    //   this.emit(`, 0)\n`);
    // } 

    if (ast.expr.type === 'call') {
      var argHasThrowFunc = this.visitFunctionNested(ast.expr, level, env);
      var hasThrow = false;
      hasThrow = ast.expr.hasThrow;
      if (hasThrow) {
        this.emit(`_body, _err := `, level);
      } else {
        this.emit(`_body := `, level);
      }
      this.visitExpr(ast.expr, level, env, { pointer: true }, argHasThrowFunc);
      this.emit(`\n`);
      if (hasThrow) {
        this.emit(`if _err != nil {\n`, level);
        this.emit(`return _result, _err\n`, level + 1);
        this.emit(`}\n`, level);
      }
      if (!ast.needCast) {
        this.emit(`_result = _body\n`, level);
      } else {
        if (env.hasThrow) {
          this.emit(`_err = dara.Convert(_body, &_result)\n`, level);
        } else {
          this.emit(`dara.Convert(_body, &_result)\n`, level);
        }
      }
    } else if (ast.expr.type === 'template_string') {
      this.emit(`_result = `, level);
      this.visitExpr(ast.expr, level, env, { pointer: true });
      this.emit(`\n`);
    } else if (ast.expr.fields) {
      for (let i = 0; i < ast.expr.fields.length; i++) {
        const field = ast.expr.fields[i];
        if (field.expr.inferred && _name(field.expr.inferred) === 'readable') {
          this.emit(`_result.${_format(_name(field.fieldName))} = `, level);
          this.visitExpr(field.expr);
          this.emit(`\n`);
          delete ast.expr.fields[i];
        }
      }
      if (env.hasThrow) {
        this.emit(`_err = dara.Convert(`, level);
      } else {
        this.emit(`dara.Convert(`, level);
      }
      
      this.visitExpr(ast.expr, level, env, { pointer: true });
      this.emit(`, &_result)\n`);
      this.emit('\n');
    } else if (ast.expr.items) {
      this.emit('_result = ', level);
      this.visitExpr(ast.expr, level, env, { pointer: true });
      this.emit('\n');
    } else if (ast.expr.type === 'construct') {
      this.emit(`_result, _err = `, level);
      this.visitConstruct(ast.expr, level, env);
      this.emit(`\n`);
    } else if (ast.expr.type === 'map_access') {
      this.emit(`_result = `, level);
      this.visitMapAccess(ast.expr, level, env, { pointer: true });
      this.emit(`\n`);
    } else if (DSL.util.isBasicType(ast.expr.type)) {
      this.emit(`_result = `, level);
      this.visitExpr(ast.expr, level, env, { pointer: true });
      this.emit(`\n`);
    }
    if(!env.finallyBlock) {
      if (env.hasThrow) {
        this.emit(`return _result, _err\n`, level);
      } else {
        this.emit(`return _result\n`, level);
      }
      env.hasReturn = false;
    }
  }

  visitRetry(ast, level) {
    assert.equal(ast.type, 'retry');
    this.emit(`\n`);
  }

  visitIf(ast, level, env) {
    assert.equal(ast.type, 'if');
    let argHasThrowFunc;
    if (ast.condition.type === 'not' && ast.condition.expr && ast.condition.expr.type === 'call') {
      argHasThrowFunc = this.visitFunctionNested(ast.condition.expr, level, env);
    } else if (ast.condition.type === 'call') {
      argHasThrowFunc = this.visitFunctionNested(ast.condition, level, env);
    }

    this.emit('if ', level);
    let setFunc;
    if(ast.condition && ast.condition.type === 'call') {
      let dealFunc = this.getVarDealFunc(ast.condition, true);
      setFunc = dealFunc && dealFunc(_name(ast.condition.inferred));
    }
    if(setFunc) {
      this.emit(`${setFunc}`);
    }
    this.visitExpr(ast.condition, level + 1, env, false, argHasThrowFunc);
    if(setFunc) {
      this.emit(')');
    }
    this.emit(' {\n');
    this.visitStmts(ast.stmts, level + 1, env);
    if (ast.elseIfs) {
      for (let i = 0; i < ast.elseIfs.length; i++) {
        let elseIf = ast.elseIfs[i];
        this.emit(`} else if `, level);
        let setFunc;
        if(elseIf.condition && elseIf.condition.type === 'call') {
          let dealFunc = this.getVarDealFunc(elseIf.condition, true);
          setFunc = dealFunc && dealFunc(_name(elseIf.condition.inferred));
        }
        if(setFunc) {
          this.emit(`${setFunc}`);
        }
        this.visitExpr(elseIf.condition, level + 1, env, false, argHasThrowFunc);
        if(setFunc) {
          this.emit(')');
        }
        this.emit(' {\n');
        this.visitStmts(elseIf.stmts, level + 1, env);
      }
    }
    if (ast.elseStmts) {
      this.emit(`} else {\n`, level);
      this.visitStmts(ast.elseStmts, level + 1, env);
    }
    this.emit('}\n\n', level);
  }

  visitThrow(ast, level, env) {
    if (ast.expr.type === 'construct_model') {
      this.emit(`_err ${env.yieldFunc ? ':' : ''}= `, level);
      this.visitConstructModel(ast.expr, level, env);
      this.emit('\n');
    } else {
      this.emit(`_err ${env.yieldFunc ? ':' : ''}= dara.NewSDKError(`, level);
      this.visitObject(ast.expr, level, env, 'map[string]interface{}');
      this.emit(')\n');
    }
    if(env.runtimeBody && !this.noCompatible) {
      this.emit(`if dara.BoolValue(client.DisableSDKError) != true {\n`, level);
      this.emit(`_err = dara.TeaSDKError(_err)\n`, level + 1);
      this.emit(`}\n`, level);
    }

    if(env.try) {
      const tryStmt = env.try;
      env.try = null;
      this.visitCatch(tryStmt, level, env);
    } else if (!env.returnType) {
      this.emit(`return _err\n`, level);
    } else if (_name(env.returnType) === 'void') {
      if(env.yieldFunc){
        this.emit(`_yieldErr <- _err\n`, level);
        this.emit(`return\n`, level);
      } else if (env.hasThrow) {
        this.emit(`return _err\n`, level);
      } else {
        this.emit(`return\n`, level);
      }
    } else {
      this.emit(`return _result, _err\n`, level);
    }
  }

  visitAssign(ast, level, env) {
    assert.equal(ast.type, 'assign');
    let pointer = false;
    let newVar = false;
    let setFunc;
    var hasThrowCall = (ast.expr.type === 'call' &&
      ast.expr.hasThrow) || ast.expr.type === 'construct';
    const yieldAssign = _isIterator(ast.expr.inferred) && ast.expr.type === 'call';
    env.yieldErrDeal = yieldAssign && !env.yieldFunc;
    if(yieldAssign) {
      ast.expr.args.push({
        yieldArg: true,
        ast,
      });
    } else {
      if (ast.left.type === 'property_assign' || ast.left.type === 'property') {
        this.emit(``, level);
        pointer = true;
        this.visitPropertyAccess(ast.left, level, env, { needCast: 'false', type: 'pointer', pointer });
      } else if (ast.left.type === 'variable') {
        const name = _name(ast.left.id);
        pointer = env.pointerParams && env.pointerParams.includes(name);
        this.emit(`${name}`, level);
        if(hasThrowCall){
          let dealFunc = this.getVarDealFunc(ast.expr,  { pointer });
          setFunc = dealFunc && dealFunc(_name(ast.expr.inferred));
          if(setFunc) {
            this.emit('Tmp');
            if(!env || !env.local || !env.local.has(`${name}Tmp`)) {
              newVar = true;
            }
          }
        }
      } else if (ast.left.type === 'virtualVariable') {
        pointer = true;
        this.emit(`${_vid(_avoidVariableKeywords(_format(_name(ast.left.vid))))}`, level);
      } else if (ast.left.type === 'map_access') {
        this.emit(``, level);
        pointer = true;
        this.visitMapAccess(ast.left, level, env, { pointer });
      } else if (ast.left.type === 'array_access') {
        this.emit(``, level);
        pointer = true;
        this.visitArrayAccess(ast.left, level, env, { pointer });
      } else {
        throw new Error('unimpelemented');
      }
  
      if (hasThrowCall) {
        this.emit(`, _err ${newVar ? ':' : ''}= `);
      } else {
        this.emit(` = `);
      }
    }

    if (ast.expr.needToReadable) {
      this.emit(`dara.ToReader(`);
      this.visitExpr(ast.expr, level, env, { needCast: 'false', pointer: true });
      this.emit(`)`);
    } else if (ast.expr.type === 'object' && ast.left.inferred &&
      ast.left.inferred.type === 'map' &&
      _name(ast.left.inferred.valueType) === 'any') {
      this.visitObject(ast.expr, level, env, 'map[string]interface{}');
    } else {
      if ((ast.left.inferred.name === 'number' || ast.left.inferred.name === 'integer') && ast.expr.inferred.name === 'int32') {
        this.emit(`${pointer ? 'dara.ToInt(' : 'int('}`);
        this.visitExpr(ast.expr, level, env, { pointer });
        this.emit(`)`);
      } else if (ast.expr.type !== 'number' && (ast.expr.inferred.name === 'number' || ast.expr.inferred.name === 'integer') && ast.left.inferred.name === 'int32') {
        this.emit(`${pointer ? 'dara.ToInt32(' : 'int32('}`);
        this.visitExpr(ast.expr, level, env, { pointer });
        this.emit(`)`);
      } else if(ast.expr.type === 'call'){
        if(hasThrowCall && ast.left.type === 'variable') {
          this.visitExpr(ast.expr, level, env, { pointer });
          if(setFunc) {
            this.emit('\n');
            let name = _name(ast.left.id);
            this.emit(`${name} = ${setFunc}${name}Tmp)`, level);
          }
        } else {
          let dealFunc = this.getVarDealFunc(ast.expr,  { pointer });
          setFunc = dealFunc && dealFunc(_name(ast.expr.inferred));
          if(setFunc) {
            this.emit(`${setFunc}`);
          }

          this.visitExpr(ast.expr, level, env, { pointer });

          if(setFunc) { 
            this.emit(')');
          }
        }
      } else {
        this.visitExpr(ast.expr, level, env, { pointer });
      }
    }
    this.emit(`\n`);
    if (hasThrowCall && !yieldAssign) {
      if(env.runtimeBody){
        this.visitAPIErrCatch(level, env);
      } else if(env.yieldFunc) {
        this.emit(`if _err != nil {\n`, level);
        this.emit(`_yieldErr <- _err\n`, level + 1);
        this.emit(`}\n\n`, level);
      } else if (env.returnType && _name(env.returnType) !== 'void') {
        this.emit(`if _err != nil {\n`, level);
        this.emit(`return _result, _err\n`, level + 1);
        this.emit(`}\n\n`, level);
      } else {
        this.emit(`if _err != nil {\n`, level);
        this.emit(`return _err\n`, level + 1);
        this.emit(`}\n\n`, level);
      }
    }
  }

  getVarDealFunc(ast, expected) {
    if((!expected || !expected.pointer) && (DSL.util.isBasicType(_name(ast.inferred))
    && !_isFilterType(_name(ast.inferred))) && this.isNotBuiltin(ast)) {
      return _setValueFunc;
    }

    if(expected && expected.pointer && !this.isNotBuiltin(ast) && !_isFilterType(_name(ast.inferred))) {
      return _setExtendFunc;
    }
  }

  visitDeclare(ast, level, env) {
    var id = _name(ast.id);
    var expected = ast.expectedType || {};
    const expr = ast.expr;
    let varDealFunc, setFunc;
    if (expr.type === 'call') {
      var argHasThrowFunc = this.visitFunctionNested(expr, level, env);
    }
    
    var hasThrowCall = (expr.type === 'call' &&
      expr.hasThrow) || expr.type === 'construct';
    const yieldAssign = _isIterator(ast.expr.inferred) && ast.expr.type === 'call';
    env.yieldErrDeal = yieldAssign && !env.yieldFunc;
    if(yieldAssign) {
      this.emit(`${id} := make(chan `, level);
      this.visitType(ast.expr.inferred.valueType, level);
      this.emit(', 1)\n');
      if((!env.hasThrow && hasThrowCall) || (env.hasThrow && !env.yieldFunc)) {
        this.emit(`_yieldErr := make(chan error, 1)\n`, level);
      }

      ast.expr.args.push({
        yieldArg: true,
        ast: {
          left: {
            type: 'variable',
            id: ast.id
          },
          expr: ast.expr,
        },
      });
    } else {
      if (hasThrowCall) {
        varDealFunc = this.getVarDealFunc(expr, expected);
        let tmpName = id;
        if (varDealFunc) { 
          tmpName = tmpName + 'Tmp';
        }
        if(env && env.local) {
          env.local.set(tmpName, true);
        }
        this.emit(`${tmpName}, _err := `, level);
      } else if (expr.type === 'null') {
        this.emit(`var ${id} `, level);
        this.visitType(ast.expectedType, level);
        this.emit('\n');
        return;
      } else {
        this.emit(`${id} := `, level);
        
        if(expr && expr.type === 'call') {
          let dealFunc = this.getVarDealFunc(expr, expected);
          setFunc = dealFunc && dealFunc(_name(expr.inferred));
        }
        
        if(setFunc) {
          this.emit(`${setFunc}`);
        }
      }
    }
    expected.pointer = false;
    if (ast.expectedType) {
      if ((_name(ast.expectedType) === 'number' || _name(ast.expectedType) === 'integer') && expr.inferred.name === 'int32') {
        this.emit(`int(`);
        this.visitExpr(expr, level, env, expected, argHasThrowFunc);
        this.emit(`)`);
      } else if (
        (expr.type !== 'number' && (expr.inferred.name === 'number' || expr.inferred.name === 'integer') 
        || expr.type === 'number') && _name(ast.expectedType) === 'int32') {
        this.emit(`int32(`);
        this.visitExpr(expr, level, env, expected, argHasThrowFunc);
        this.emit(`)`);
      } else {
        this.visitExpr(expr, level, env, expected, argHasThrowFunc);
      }
    } else {
      this.visitExpr(expr, level, env, expected, argHasThrowFunc);
      if(setFunc) { 
        this.emit(')');
      }
    }
    this.emit('\n');

    if(varDealFunc) {
      this.emit(`${id} := `, level);
      this.emit(`${varDealFunc(_name(expr.inferred))}`);
      this.emit(`${id}Tmp)\n`);
    }

    if (hasThrowCall && !yieldAssign) {
      if(env.runtimeBody){
        this.visitAPIErrCatch(level, env);
      } else if(env.yieldFunc) {
        this.emit(`if _err != nil {\n`, level);
        this.emit(`_yieldErr <- _err\n`, level + 1);
        this.emit(`return\n`, level + 1);
        this.emit(`}\n\n`, level);
      } else if(env.try) {
        const tryStmt = env.try;
        env.try = null;
        this.visitCatch(tryStmt, level, env);
      } else if (env.returnType && _name(env.returnType) !== 'void') {
        this.emit(`if _err != nil {\n`, level);
        this.emit(`return _result, _err\n`, level + 1);
        this.emit(`}\n\n`, level);
      } else {
        this.emit(`if _err != nil {\n`, level);
        this.emit(`return _err\n`, level + 1);
        this.emit(`}\n\n`, level);
      }
    }
  }

  visitStmts(ast, level, env) {
    assert.equal(ast.type, 'stmts');
    let node;
    
    for (var i = 0; i < ast.stmts.length; i++) {
      node = ast.stmts[i];
      this.visitStmt(node, level, env);
    }
    if (node) {
      //find the last node's back comment
      let comments = DSL.comment.getBackComments(this.comments, node.tokenRange[1]);
      this.visitComments(comments, level);
    }

    if (ast.stmts.length === 0) {
      //empty block's comment
      let comments = DSL.comment.getBetweenComments(this.comments, ast.tokenRange[0], ast.tokenRange[1]);
      this.visitComments(comments, level);
    }
  }

  visitYieldReturnBody(ast, funcName, level, env) {
    let args = [];
    args.push({
      name: '_yield chan',
      type: env.returnType.valueType,
    });
    if(env.hasThrow) {
      args.push({
        name: '_yieldErr chan error',
      });
    }
    let vars = this.getStmtsVars(ast.stmts.stmts);
    this.emit(`${funcName}(_yield`, level);
    if(env.hasThrow) {
      this.emit(', _yieldErr, ');
    }

    const argsMap = {};
    let params = new Set();
    for (let i = 0; i < vars.length; i++) {
      const arg = vars[i];
      params.add(arg.name);
      argsMap[arg.name] = arg.type;
    }
    params = [...params];
    this.emit(params.join(', '));

    params.map((name) => {
      const variable = (env.pointerParams && env.pointerParams.includes(name)) ? false : true;
      args.push({
        name,
        type: argsMap[name],
        variable,
      });
    });

    this.emit(')\n');

    if(env.hasThrow) {
      this.emit('_err = <-_yieldErr\n', level);
    }
    this.visitAPIErrCatch(level, env);

    this.emit('return\n', level);
    this.yieldFunc.push({
      args,
      functionBody: ast,
      name: funcName,
      pointerParams: env.pointerParams,
    });
  }

  visitReturnBody(ast, apiName, level, env) {
    assert.equal(ast.type, 'returnBody');
    const funcName = `${_lowerFirst(apiName)}_opResponse`;
    if(_isIterator(env.returnType)) {
      this.visitYieldReturnBody(ast, funcName, level, env);
      return;
    }
    const errCounts = this.stmtsErrCount(ast.stmts.stmts);
    if(errCounts > 1) {
      const args = this.getStmtsVars(ast.stmts.stmts);
      this.tryFunc.push({
        args,
        functionBody: ast.stmts,
        returnType: env.returnType,
        name: funcName,
        pointerParams: env.pointerParams,
      });
      if (_name(env.returnType) && _name(env.returnType) !== 'void') {
        this.emit(`_result, _err = ${funcName}`, level);
      } else {
        this.emit(`_err = ${funcName}`, level);
      }
      this.visitTryArgs(args, level, env);
      this.emit('\n');
      this.visitAPIErrCatch(level, env);
      let err = '_err';
      if(!env.runtimeBody) {
        err = 'nil';
      }
      if(_name(env.returnType) === 'void' ) {
        this.emit(`return ${err}\n`, level);
      } else {
        this.emit(`return _result, ${err}\n`, level);
      }
      return;
    }
    this.visitStmts(ast.stmts, level, env);
  }

  visitYieldFunction(ast, level, env) {
    env.routine = true;
    const args = [];
    const pointerParams = [];
    this.emit('defer close(_yield)\n', level);
    args.push({
      name: '_yield chan',
      type: ast.returnType.valueType,
    });
    if(env.hasThrow) {
      this.emit('defer close(_yieldErr)\n', level);
      args.push({
        name: '_yieldErr chan error',
      });
    }

    const functionName = `${_lowerFirst(env.funcName)}_opYieldFunc`;
    this.emit(`${functionName}(_yield`, level);
    if(env.hasThrow) {
      this.emit(', _yieldErr');
    }
    for (var i = 0; i < ast.params.params.length; i++) {
      this.emit(', ');
      const node = ast.params.params[i];
      assert.equal(node.type, 'param');
      const name = _avoidReserveName(_name(node.paramName));
      pointerParams.push(name);
      this.emit(name);
      if (node.paramType) {
        args.push({
          name,
          type: node.paramType,
        });
      }
    }

    this.emit(')\n');
    this.yieldFunc.push({
      name: functionName,
      args,
      functionBody: ast.functionBody,
      pointerParams,
    });

    this.emit('return\n', level);
  }

  visitFunctionBody(ast, level, env) {
    assert.equal(ast.type, 'functionBody');
    this.visitStmts(ast.stmts, level, env);
    const stmts = ast.stmts.stmts;
    const length = ast.stmts.stmts.length;

    if (_name(env.returnType) === 'void' && env.hasThrow && (length === 0 || (stmts[length - 1].type !== 'return' &&
      stmts[length - 1].type !== 'throw'))) {
      if(env.yieldErrDeal) {
        this.emit(`_err = <- _yieldErr\n`, level);
      }
      this.emit(`return _err\n`, level);
    }

    if (length === 0 || this.functionBodyRetrun(stmts[length - 1], env)) {
      if (_name(env.returnType) !== 'void' && env.hasThrow) {
        if(env.yieldErrDeal) {
          this.emit(`_err = <- _yieldErr\n`, level);
        }
        this.emit(`return _result, _err\n`, level);
      } else if (_name(env.returnType) !== 'void' && !env.hasThrow) {
        this.emit(`return _result\n`, level);
      }
    }
  }

  functionBodyRetrun(stmt, env) {
    if(stmt.type === 'return') {
      return false;
    }

    if(stmt.type === 'throw') {
      return false;
    }

    if(!(stmt.type !== 'if' || !stmt.elseStmts)) {
      return false;
    }

    if(stmt.type === 'try' && (stmt.tryBlock && stmt.tryBlock.stmts.length > 0 
      && !this.functionBodyRetrun(stmt.tryBlock.stmts[stmt.tryBlock.stmts.length - 1])) && !env.hasReturn) {
      return false;
    }

    return true;
  }

  eachFunc(ast, level, predefined, apis) {
    const env = {
      predefined,
      apis,
      local: new Map(),
      returnType: ast.returnType,
      hasThrow: ast.isAsync || ast.hasThrow,
      yieldFunc: _isIterator(ast.returnType),
      nestFuncParamName: new Map(),
      nestFuncParamNameSubscript: { 'count': 0 }
    };
    const functionName = _name(ast.functionName);
    this.visitAnnotation(ast.annotation, level);
    let comments = DSL.comment.getFrontComments(this.comments, ast.tokenRange[0]);
    this.visitComments(comments, level);
    const name = _format(functionName);
    env.funcName = name;
    if (this.exec && name === 'Main') {
      this.emit(`func _main `, level);
    } else if(name.startsWith('$') && this.builtin[name]){
      const method = name.replace('$', '');
      this.builtin[name][method](ast.args, level);
      return;
    } else if (ast.isStatic) {
      this.emit(`func ${name} `, level);
    } else {
      this.emit(`func (client *${this.structName}) ${name} `, level);
    }
    this.visitParams(ast.params, level, env);
    this.visitReturnType(ast, level, env);
    
    if (ast.functionBody) {
      if(_isIterator(ast.returnType)) {
        this.visitYieldFunction(ast, level + 1, env);
      } else {
        this.visitFunctionBody(ast.functionBody, level + 1, env);
      }
    } else {
      this.emit(`panic("No Support!")\n`, level + 1);
    }
    this.emit(`}\n`, level);
    this.emit(`\n`, level);
  }

  visitAPIErrCatch(level, env) {
    this.emit('if _err != nil {\n', level);
    if(env.runtimeBody) {
      this.emit('retriesAttempted++\n', level + 1);
      this.emit('retryPolicyContext = &dara.RetryPolicyContext{\n', level + 1);
      this.emit('RetriesAttempted: retriesAttempted,\n', level + 2);
      this.emit('HttpRequest:      request_,\n', level + 2);
      this.emit('HttpResponse:     response_,\n', level + 2);
      this.emit('Exception:        _err,\n', level + 2);
      this.emit('}\n', level + 1);
      this.emit('_resultErr = _err\n', level + 1);
      this.emit('continue\n', level + 1);
    } else  {
      
      this.emit('return ', level + 1);
      if(_name(env.returnType) !== 'void') {
        this.emit('nil, ');
      }
      this.emit('_err\n');
    }
    this.emit('}\n', level);
    this.emit('\n');
  }

  eachAPI(ast, level, predefined) {
    // if (ast.annotation) {
    //   this.emit(`${_anno(ast.annotation.value)}\n`, level);
    // }
    const env = {
      // params, paramMap, returnType,
      predefined,
      returnType: ast.returnType,
      runtimeBody: ast.runtimeBody,
      local: new Map(),
      hasThrow: true,
      yieldFunc: _isIterator(ast.returnType),
      nestFuncParamName: new Map(),
      nestFuncParamNameSubscript: { 'count': 0 },
    };

    const apiName = _name(ast.apiName);
    this.visitAnnotation(ast.annotation, level);
    let comments = DSL.comment.getFrontComments(this.comments, ast.tokenRange[0]);
    this.visitComments(comments, level);
    // func (b *Buffer) Next(n int) []byte {
    this.emit(`func (client *${this.structName}) ${_format(apiName)}`, level);
    this.visitParams(ast.params, level, env);
    this.visitReturnType(ast, level, env);
    // this.emit(` (*map[string]interface{}, error) {\n`);
    let baseLevel = ast.runtimeBody ? level + 1 : level;
    if(env.yieldFunc) {
      this.emit('defer close(_yield)\n', level + 1);
      if(env.hasThrow) {
        this.emit('defer close(_yieldErr)\n', level + 1);
      }
    }
    // api level
    if (ast.runtimeBody) {
      this.visitRuntimeBefore(ast.runtimeBody, level + 1, env);
    }

    // temp level
    this.visitAPIBody(ast.apiBody, baseLevel + 1, env);

    // if (ast.runtimeBody) {
    //   this.emit(`_lastRequest = ${REQUEST}\n`, baseLevel + 1);
    // }

    this.emit(`${RESPONSE}, _err := dara.DoRequest(${REQUEST}`, baseLevel + 1);

    if (ast.runtimeBody) {
      this.emit(`, _runtime`);
    } else {
      this.emit(`, nil`);
    }
    this.emit(`)\n`);

    this.visitAPIErrCatch(baseLevel + 1, env);

    if (ast.returns) {
      this.visitReturnBody(ast.returns, apiName, baseLevel + 1, env);
    } else {
      this.visitDefaultReturnBody(baseLevel + 1, env);
    }

    if (ast.runtimeBody) {
      this.emit('}\n', level + 1);
      if(env.yieldFunc) {
        this.emit(`_yieldErr <- _resultErr\n`, level + 1);
        this.emit(`return\n`, level + 1);
      } else {
        if(!this.noCompatible) {
          this.emit(`if dara.BoolValue(client.DisableSDKError) != true {\n`, level + 1);
          this.emit(`_resultErr = dara.TeaSDKError(_resultErr)\n`, level + 2);
          this.emit(`}\n`, level + 1);
        }
        this.emit(`return _result, _resultErr\n`, level + 1);
      }
    }
    
    this.emit(`}\n\n`, level);
  }

  visitDefaultReturnBody(level, env) {
    this.emit('\n');
    if (_name(env.returnType) === 'void') {
      this.emit(`return nil\n`, level);
    } else {
      this.emit('return nil, nil\n', level);
    }
    // this.emit(`"statusCode": ${RESPONSE}.statusCode,\n`, level + 1);
    // this.emit(`"statusMessage": ${RESPONSE}.statusMessage,\n`, level + 1);
    // this.emit(`"headers": ${RESPONSE}.headers,\n`, level + 1);
    // this.emit(`"body": client.Auto__(${RESPONSE}),\n`, level + 1);
  }

  moduleBefore(ast, main, level) {
    this.mainModule = main;
    let beginToken = 0;
    if(ast.imports.length > 0) {
      const lastIndex = ast.imports.length - 1;
      beginToken = ast.imports[lastIndex].tokenRange[0];
    }
    const beginNotes = DSL.note.getNotes(this.notes, beginToken, ast.moduleBody.nodes[0].tokenRange[0]);
    const part = beginNotes.find(note => note.note.lexeme === '@go');
    if(part && part.arg.value) {
      this.emit(_string(part.arg.value));
    }
  }

  visitHeader(__module, packageName, level) {
    this.header = '';
    if (!this.editable) {
      this.header += `// This file is auto-generated, don't edit it. Thanks.\n`;
    }
    this.header += `package ${packageName}\n\n`;
    this.header += `import (\n`;
    const imports = [];
    if(this.mainModule) {
      if (this.exec) {
        imports.push(`  "os"`);
      }
      if (this.goPackages.length > 0) {
        this.goPackages.map(pack => {
          imports.push(pack);
        });
      }
    }
    this.imports.map(im => {
      let tmpStr = '  ';
      if(im.aliasId) {
        tmpStr += `${_importFilter(_format(im.aliasId).toLowerCase())} `;
      }
      tmpStr += `"${im.pkgName}"`;
      imports.push(tmpStr);
    });
    imports.push(`  "github.com/alibabacloud-go/tea/dara"`);
    this.builtinModule.forEach(builtinModule => {
      let content = '  ';
      if(builtinModule.name) {
        content += `${builtinModule.name} `;
      }
      content += `"${builtinModule.path}"`;
      imports.push(content);
    });
    this.header += [...new Set(imports)].join('\n') + '\n)\n\n';
  }

  

  modelBefore(filepath) {
    // Nothing
    const targetPath = path.join(this.outputDir, filepath);
    _deleteWithSuffix(path.dirname(targetPath), '_model.go');
    _deleteWithSuffix(path.dirname(targetPath), '_error.go');
  }

  modelAfter() {
    // Nothing
  }

  apiBefore(__module, level) {
    this.emit(`\n`);
  }

  init(level) {
    // Nothing
  }

  apiAfter(__module, level = 0) {
    // Nothing
  }

  wrapBefore(__module, level) {
    this.emit(`\n`);
  }

  wrapAfter(filepath, level) {
    // Nothing
  }

  moduleAfter() {
    this.uselessPack.forEach((pack) => {
      this.header = this.header.replace(`  "${pack}"\n`, '');
    });
  }



  visitTryArgs(args) {
    const params = new Set();
    this.emit('(');
    for (let i = 0; i < args.length; i++) {
      const arg = args[i];
      params.add(arg.name);
    }
    this.emit([...params].join(', '));
    this.emit(')');
  }

  visitTryParams(args, env) {
    const argsMap = {};
    let params = new Set();
    this.emit('(');
    for (let i = 0; i < args.length; i++) {
      const arg = args[i];
      params.add(arg.name);
      argsMap[arg.name] = arg.type;
    }
    params = [...params];
    params.map((param, index) => {
      if(param === '_yield') {
        env.yieldFunc = true;
      }
      this.emit(`${param} `);
      if(env.pointerParams && env.pointerParams.includes(param)){
        this.visitPointerType(argsMap[param]);
      } else {
        this.visitType(argsMap[param]);
      }
      if(index !== params.length - 1) {
        this.emit(', ');
      }
      
    });
    this.emit(')');
  }

  eachYieldFunc(ast, level) {
    const env = {
      local: new Map(),
      returnType: {
        name: 'void'
      },
      hasThrow: false,
      yieldFunc: true,
      pointerParams: ast.pointerParams,
    };
    
    this.emit(`func ${ast.name}(`, level);
    for(let i = 0; i < ast.args.length; i++) {
      const arg = ast.args[i];
      if (arg.name === '_yieldErr chan error') {
        env.hasThrow = true;
      }
      this.emit(arg.name);
      if(arg.type) {
        this.emit(` `);
        if(!arg.variable) {
          this.visitPointerType(arg.type, level, env);
        } else {
          this.visitType(arg.type, level, env);
        }
      }
      if(i !== ast.args.length - 1){
        this.emit(', ');
      }
    }
    this.emit(') {\n');
    
    this.visitStmts(ast.functionBody.stmts, level + 1, env);
    this.emit(`}\n`, level);
    this.emit(`\n`, level);
  }

  eachTryFunc(ast, level) {
    const env = {
      local: new Map(),
      returnType: ast.returnType,
      hasThrow: true,
      pointerParams: ast.pointerParams || [],
    };
    this.emit(`func ${ast.name} `, level);
    this.visitTryParams(ast.args, env);
    this.emit('(');
    if (_name(env.returnType) !== 'void') {
      this.emit(` _result `);
      this.visitPointerType(ast.returnType, level, env);
      this.emit(', ');
    }
    
    this.emit('_err error) {\n');

    this.visitStmts(ast.functionBody, level + 1, env);
    const stmts = ast.functionBody.stmts;
    const length = ast.functionBody.stmts.length;
    if (_name(env.returnType) === 'void') {
      this.emit(`return _err\n`, level + 1);
    } else if ((length === 0 || (stmts[length - 1].type !== 'return' && stmts[length - 1].type !== 'throw' &&
      (stmts[length - 1].type !== 'if' || !stmts[length - 1].elseStmts)))) {
      this.emit(`return _result, _err\n`, level + 1);
    }
    this.emit(`}\n`, level);
    this.emit(`\n`, level);
  }

  isNotBuiltin(ast) {
    if(ast.type === 'call') {
      if (ast.left.type === 'method_call') {
        const name = _format(_name(ast.left.id));
        if (name.startsWith('$') && this.builtin[name]) {
          return false;
        }
      } else if (ast.left.type === 'instance_call') {
        if(ast.builtinModule && this.builtin[ast.builtinModule]) {
          return false;
        }
      } else if (ast.left.type === 'static_call') {
        if(ast.left.id.type === 'builtin_module') {
          return false;
        }
      }
      
    }
    return true;
  }

  typeRelover(type, module) {
    if (module && module.idType === 'module') {
      const aliasId = _name(module);
      if (this.importsTypedef[aliasId] && this.importsTypedef[aliasId][type.lexeme]) {
        let index = this.uselessPack.indexOf(this.importsTypedef[aliasId][type.lexeme].import);
        if (index > -1) {
          this.uselessPack.splice(index, 1);
        }
        this.imports.push({
          pkgName: this.importsTypedef[aliasId][type.lexeme].import,
        });
        return this.importsTypedef[aliasId][type.lexeme].type;
      }
    }
    if (type.idType === 'typedef' && this.typedef[type.lexeme]) {
      if (this.typedef[type.lexeme].import && this.typedef[type.lexeme].type) {
        let index = this.uselessPack.indexOf(this.typedef[type.lexeme].import);
        if (index > -1) {
          this.uselessPack.splice(index, 1);
        }
        this.imports.push({
          pkgName: this.typedef[type.lexeme].import,
        });
        return this.typedef[type.lexeme].type;
      }
    }

    if (this.predefined[_name(type)] && this.predefined[_name(type)].isException) {
      return `${_name(type)}Error`;
    }

    return _name(type);
  }

  getType(name, pointer = true) {
    const type = pointer ? _pointerType(name) : _type(name);
    if(type.includes('io.')) {
      this.imports.push({
        pkgName: 'io',
      });
    }
    return type;
  }
}
module.exports = Visitor;