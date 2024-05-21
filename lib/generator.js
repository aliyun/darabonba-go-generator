'use strict';

const assert = require('assert');

const path = require('path');
const fs = require('fs');

const REQUEST = 'request_';
const RESPONSE = 'response_';
const CORE = 'SDK';
const DSL = require('@darabonba/parser');
const Annotation = require('@darabonba/annotation-parser');
const { Tag } = DSL.Tag;

const {
  _name, _format, _string, _type, _initValue, _avoidReserveName, _importFilter, _avoidVariableKeywords,
  _setExtendFunc, _isFilterType, _getAttr, _setValueFunc, _vid, _pointerType, _lowerFirst, _escape
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
      if (packageName !== 'io') {
        return `  "${packageName}"`;
      }
    }).join('\n');
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

    fs.writeFileSync(targetPath, this.output);
    this.output = '';
  }


  reWriteMod(module, original) {
    const lines = original.split(/\n/g);
    var content = '';
    lines.forEach(function (line, i) {
      if ((/[v][0-9].*/).test(line)) {
        module.importPackages.forEach((importPackage) => {
          const strs = line.split(/ /g);
          if (strs[0].trimLeft('\t') === importPackage.path) {
            strs[1] = importPackage.version;
            line = strs.join(' ');
          }
        });
      }
      if (i < lines.length - 1) {
        content += `${line}\n`;
      } else {
        content += `${line}`;
      }
    });
    return content;
  }

  generatorMod(module) {
    var content = `module ${module.moduleName}\n`;
    content += `\n`;
    content += `require (\n`;
    module.importPackages.forEach((importPackage) => {
      content += '	' + importPackage.path + ' ' + importPackage.version + `\n`;
    });
    content += `)\n`;
    return content;
  }

  emit(str, level) {
    this.output += ' '.repeat(level * 2) + str;
  }

  visit(ast, level = 0) {
    this.visitModule(ast, level);
  }

  visitModule(ast, level) {
    assert.equal(ast.type, 'module');
    this.importBefore(this.__module, level);

    const apis = ast.moduleBody.nodes.filter((item) => {
      return item.type === 'api';
    });

    const models = ast.moduleBody.nodes.filter((item) => {
      return item.type === 'model';
    });

    const nonStaticFuncs = ast.moduleBody.nodes.filter((item) => {
      return item.type === 'function' && !item.isStatic;
    });

    const initParts = ast.moduleBody.nodes.filter((item) => {
      return item.type === 'init';
    });

    this.comments = ast.comments;
    this.visitAnnotation(ast.annotation, level);
    this.importAfter(this.__module, level);
    this.importPackages = '';
    const teafilePath = fs.existsSync(path.join(this.config.pkgDir, 'Teafile')) ?
      path.join(this.config.pkgDir, 'Teafile') : path.join(this.config.pkgDir, 'Darafile');
    const teaFile = JSON.parse(fs.readFileSync(teafilePath, 'utf8'));
    const release = teaFile.releases && teaFile.releases.go || '';
    const strs = release.split(':');
    this.module = {
      importPackages: [],
      moduleName: strs[0].substring(0, strs[0].lastIndexOf('/')) || 'client'
    };
    this.structName = this.config.go ? this.config.go.clientName || 'Client' : 'Client';
    this.clientName = {};
    this.constructFunc = {};
    if (ast.imports.length > 0) {
      const pack = this.goPackages.split('\n');
      const lockPath = path.join(this.config.pkgDir, '.libraries.json');
      const lock = JSON.parse(fs.readFileSync(lockPath, 'utf8'));
      for (let i = 0; i < ast.imports.length; i++) {
        const item = ast.imports[i];
        const aliasId = _name(item);
        const moduleDir = this.config.libraries[aliasId];
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
            if (!pack.includes(`  "${moduleTypedef[types].import}"`)) {
              this.uselessPack.push(moduleTypedef[types].import);
              pack.push(`  "${moduleTypedef[types].import}"`);
            }
          });
        }
        if (importTeaFile.go && importTeaFile.go.clientName) {
          if (importTeaFile.go.clientName.indexOf('*') === 0) {
            this.constructFunc[aliasId] = importTeaFile.go.clientName.substring(1);
            this.clientName[aliasId] = '*' + _importFilter(_format(aliasId).toLowerCase()) + '.' + importTeaFile.go.clientName.substring(1);
          } else {
            this.constructFunc[aliasId] = importTeaFile.go.clientName;
            this.clientName[aliasId] = _importFilter(_format(aliasId).toLowerCase()) + '.' + importTeaFile.go.clientName;
          }
        } else {
          this.constructFunc[aliasId] = 'Client';
          this.clientName[aliasId] = '*' + _importFilter(_format(aliasId).toLowerCase()) + '.Client';
        }
        if (!goPkg) {
          throw new Error(`The '${aliasId}' has no Go supported.`);
        }

        const [pkgName, version] = goPkg.split(':');
        if (version) {
          this.module.importPackages.push({
            path: pkgName.substring(0, pkgName.lastIndexOf('/')),
            version: version,
            name: aliasId,
            goPkg: importTeaFile.go
          });
        }
        this.importPackages += `  ${_importFilter(aliasId.toLowerCase())}  "${pkgName}"\n`;
      }
      this.goPackages = pack.join('\n');
    }

    if (ast.usedTypes.has('readable') || ast.usedTypes.has('writeable')) {
      if (this.goPackages !== '') {
        this.goPackages += `\n`;
      }
      this.goPackages += `  "io"`;
    }
    this.moduleBefore(this.__module, level);

    const __module = this.__module;

    // global definition
    this.modelBefore(__module, level);

    for (let i = 0; i < models.length; i++) {
      this.visitModel(models[i], level);
    }

    this.modelAfter(__module, level);

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
          nestFuncParamNameSubscript: { 'count': 0 }
        };
        this.emit(`${_format(_name(func.functionName))} `, level + 1);
        this.visitParams(func.params, level, env);
        this.visitReturnType(func, level, env, false);
      }
      this.emit(`}\n\n`, level);
    }

    if (!ast.extends && (apis.length > 0 || nonStaticFuncs.length > 0)) {
      this.emit(`type ${this.structName} struct {\n`, level);
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
      this.emit(`${_name(ast.extends).toLowerCase()}.Client\n`, level + 1);
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

    if (this.exec) {
      this.emit(`\nfunc main() {\n`, 0);
      this.emit(`err := _main(tea.StringSlice(os.Args[1:]))\n`, 1);
      this.emit(`if err != nil {\n`, 1);
      this.emit('panic(err)\n', 2);
      this.emit(`}\n`, 1);
      this.emit('}\n', 0);
    }

    this.wrapAfter(__module, level);

    this.moduleAfter(__module, level);

    this.save(this.config.clientPath);
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

    var descriptionText = description ? _escape(description.text.text.trimEnd()) : '';
    var summaryText = summary ? _escape(summary.text.text.trimEnd()) : '';
    var returnText = _return ? _return.text.text.trimEnd() : '';
    let hasNextSection = false;
    if (deprecated) {
      let deprecatedText = deprecated.text.text.trimEnd();
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
      params.forEach((item, index) => {
        this.emit(`// @param ${item.name} - ${item.text}`, level);
        if (index < params.length - 1) {
          this.emit(`// \n`, level);
        }
      });
      hasNextSection = true;
    }
    if (returnText !== '') {
      if (hasNextSection) {
        this.emit(`// \n`, level);
      }
      this.emit(`// @return ${returnText}\n`, level);
      hasNextSection = true;
    }
    if (throws.length > 0) {
      if (hasNextSection) {
        this.emit(`// \n`, level);
      }
      throws.forEach((item, index) => {
        this.emit(`// @throws ${item}`, level);
        if (index < throws.length - 1) {
          this.emit(`// \n`, level);
        }
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
    this.emit('(');
    for (var i = 0; i < ast.params.length; i++) {
      if (i !== 0) {
        this.emit(', ');
      }
      const node = ast.params[i];
      assert.equal(node.type, 'param');
      const name = _name(node.paramName);
      this.emit(`${_avoidReserveName(name)}`);
      if (node.paramType) {
        const paramType = node.paramType;
        this.emit(` `);
        this.visitPointerType(paramType, level, env);
      }
    }

    this.emit(')');
  }

  visitInit(ast, level) {
    const env = {
      local: new Map(),
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

  visitAPIBody(ast, level, env) {
    assert.equal(ast.type, 'apiBody');
    this.emit(`${REQUEST} := tea.NewRequest()\n`, level);

    if (ast.stmts) {
      this.visitStmts(ast.stmts, level, env);
    }
  }

  visitRuntimeBefore(ast, level, env) {
    assert.equal(ast.type, 'object');
    this.emit('_runtime := ', level);
    this.visitObject(ast, level, env, 'map[string]interface{}');
    this.emit('\n\n');
    if (_name(env.returnType) && _name(env.returnType) !== 'void') {
      this.emit(`_resp := ${_initValue(_name(env.returnType))}\n`, level);
    } else if (env.returnType.path) {
      this.emit(`_resp := new(`, level);
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
      this.emit(`_resp := make(`, level);
      this.visitPointerType(env.returnType, level);
      this.emit(`)\n`);
    }
    this.emit(`for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {\n`, level);
    this.emit(`if _retryTimes > 0 {\n`, level + 1);
    this.emit(`_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))\n`, level + 2);
    this.emit(`if tea.IntValue(_backoffTime) > 0 {\n`, level + 2);
    this.emit(`tea.Sleep(_backoffTime)\n`, level + 3);
    this.emit(`}\n`, level + 2);
    this.emit(`}\n`, level + 1);
    this.emit(`\n`);
    if (_name(env.returnType) === 'void') {
      this.emit(`_err = func() error {\n`, level + 1);
    } else if (env.returnType.path) {
      var returnType = '';
      for (let i = 0; i < env.returnType.path.length; i++) {
        const path = env.returnType.path[i];
        if (i === 0) {
          returnType += _name(path).toLowerCase();
        } else {
          returnType += '.' + _name(path);
        }
      }
      this.emit(`_resp, _err = func()(${_pointerType(returnType)}, error){\n`, level + 1);
    } else {
      this.emit(`_resp, _err = func()(`, level + 1);
      this.visitPointerType(env.returnType, level);
      this.emit(`, error){\n`);
    }
  }

  visitStmt(ast, level, env) {
    let comments = DSL.comment.getFrontComments(this.comments, ast.tokenRange[0]);
    this.visitComments(comments, level);
    if (ast.type === 'return') {
      this.visitReturn(ast, level, env);
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
          this.emit(`_, _err = `, level);
        } else {
          this.emit(`_err = `, level);
        }
        this.visitExpr(ast, level, env, { pointer: true });
        this.emit(`\n`);
        this.emit(`if _err != nil {\n`, level);
        if (env.returnType && _name(env.returnType) !== 'void') {
          this.emit(`return _result, _err\n`, level + 1);
        } else {
          this.emit(`return _err\n`, level + 1);
        }
        this.emit(`}\n`, level);
      } else {
        this.emit(``, level);
        this.visitExpr(ast, level, env, { pointer: true });
        this.emit(`\n`);
      }
    }
  }

  visitTry(ast, level, env) {
    assert.equal(ast.type, 'try');
    const tryBlock = ast.tryBlock;
    const catchBlock = ast.catchBlock;
    const finallyBlock = ast.finallyBlock;
    if (finallyBlock && finallyBlock.stmts.length > 0) {
      this.emit(`defer func() {\n`, level);
      this.visitStmts(ast.finallyBlock, level + 1, env);
      this.emit(`}()\n`, level);
    }
    if (tryBlock && tryBlock.stmts.length > 0) {
      if (!env.hasThrow) {
        this.emit(`var _err error\n`, level);
      }
      if (!catchBlock || catchBlock.stmts.length <= 0) {
        if (_name(env.returnType) && _name(env.returnType) !== 'void') {
          this.emit(`func()(_r ${_pointerType(_name(env.returnType))}, _e error) {\n`, level);
        } else if (!_name(env.returnType)) {
          this.emit(`func()(_r `, level);
          this.visitPointerType(env.returnType, level);
          this.emit(`, _e error) {\n`);
        } else {
          this.emit(`func()(_e error) {\n`, level);
        }
      } else {
        if (_name(env.returnType) && _name(env.returnType) !== 'void') {
          this.emit(`_, tryErr := func()(_r ${_pointerType(_name(env.returnType))}, _e error) {\n`, level);
        } else if (!_name(env.returnType)) {
          this.emit(`_, tryErr := func()(_r `, level);
          this.visitPointerType(env.returnType, level);
          this.emit(`, _e error) {\n`);
        } else {
          this.emit(`tryErr := func()(_e error) {\n`, level);
        }
      }

      if (catchBlock && catchBlock.stmts.length > 0) {
        this.emit(`defer func() {\n`, level + 1);
        this.emit(`if r := tea.Recover(recover()); r != nil {\n`, level + 2);
        this.emit(`_e = r\n`, level + 3);
        this.emit(`}\n`, level + 2);
        this.emit(`}()\n`, level + 1);
      }
      const tmp = env.hasThrow;
      env.hasThrow = true;
      this.visitStmts(tryBlock, level + 1, env);
      env.hasThrow = tmp;
      if (tryBlock && tryBlock.stmts[tryBlock.stmts.length - 1].type !== 'return') {
        this.visitDefaultReturnBody(level + 1, env);
      }
      this.emit(`}()\n\n`, level);
      if (catchBlock && catchBlock.stmts.length > 0) {
        this.emit(`if tryErr != nil {\n`, level);
        this.emit(`var ${_name(ast.catchId)} = &tea.SDKError{}\n`, level + 1);
        this.emit(`if _t, ok := tryErr.(*tea.SDKError); ok {\n`, level + 1);
        this.emit(`${_name(ast.catchId)} = _t\n`, level + 2);
        this.emit(`} else {\n`, level + 1);
        this.emit(`${_name(ast.catchId)}.Message = tea.String(tryErr.Error())\n`, level + 2);
        this.emit(`}\n`, level + 1);
        this.visitStmts(catchBlock, level + 1, env);
        this.emit(`}\n`, level);
      }
    }
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
    this.visitExpr(ast.condition, level + 1, env, false, argHasThrowFunc);
    this.emit(' {\n');
    this.visitStmts(ast.stmts, level + 1, env);
    this.emit('}\n', level);
  }

  visitFor(ast, level, env) {
    assert.equal(ast.type, 'for');
    this.emit(`for _, ${_name(ast.id)} := range `, level);
    this.visitExpr(ast.list, level + 1, env, { pointer: true });
    this.emit(' {\n');
    this.visitStmts(ast.stmts, level + 1, env);
    this.emit('}\n', level);
  }

  visitFieldValue(ast, structName, level) {
    if (ast.type === 'fieldType') {
      if (ast.fieldType === 'array') {
        if (ast.fieldItemType.type === 'modelBody') {
          this.emit('{\n', level + 1);
          this.visitModelBody(ast.fieldItemType, structName, level);
        }
        return;
      }
    }

    if (ast.type === 'modelBody') {
      this.emit('{\n');
      this.visitModelBody(ast, structName, level);
      return;
    }

    throw new Error('unimpelemented');
  }

  visitType(ast, level) {
    if (ast.type === 'moduleModel') {
      this.emit(`*${_importFilter(_name(ast.path[0]).toLowerCase())}`);
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
      this.emit(`map[${_type(_name(ast.keyType))}]`);
      this.visitPointerType(ast.valueType, level);
    } else if (ast.fieldType === 'array' || ast.type === 'array') {
      this.emit(`[]`);
      this.visitPointerType(ast.subType || ast.itemType, level);
    } else if (ast.idType === 'module' || this.clientName[_name(ast)]) {
      this.emit(`${this.clientName[_name(ast)]}`);
    } else if (this.typeRelover(ast)) {
      this.emit(_type(this.typeRelover(ast)));
    } else if (ast.fieldType && DSL.util.isBasicType(ast.fieldType)) {
      this.emit(_type(ast.fieldType));
    } else if (ast.fieldType && this.typeRelover(ast.fieldType)) {
      this.emit(_type(this.typeRelover(ast.fieldType)));
    } else {
      this.emit(_type(ast));
    }
  }

  visitPointerType(ast, level) {
    if (ast.type === 'moduleModel') {
      this.emit(`*${_importFilter(_name(ast.path[0]).toLowerCase())}`);
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
      this.emit(`map[${_type(_name(ast.keyType))}]`);
      this.visitPointerType(ast.valueType, level);
    } else if (ast.fieldType === 'array' || ast.type === 'array') {
      this.emit(`[]`);
      this.visitPointerType(ast.subType || ast.itemType, level);
    } else if (ast.idType === 'module' || this.clientName[_name(ast)]) {
      this.emit(`${this.clientName[_name(ast)]}`);
    } else if (ast.type === 'model') {
      this.emit(`*`);
      if (ast.moduleName) {
        this.emit(`${ast.moduleName.replace(/-/g, '_').toLowerCase()}.`);
      }
      let strs = _format(_name(ast)).split('.');
      strs.forEach(str => {
        this.emit(`${_format(str)}`);
      });
    } else if (_name(ast)) {
      this.emit(_pointerType(this.typeRelover(ast)));
    } else if (ast.fieldType && DSL.util.isBasicType(ast.fieldType)) {
      this.emit(_pointerType(ast.fieldType));
    } else if (ast.fieldType && this.typeRelover(ast.fieldTyp)) {
      this.emit(_pointerType(this.typeRelover(ast.fieldType)));
    } else {
      this.emit(_pointerType(ast));
    }
  }

  visitModelField(ast, structName, level) {
    //assert.equal(ast.fieldValue.type, 'fieldType');
    this.emit(`type ${structName} struct `);
    this.visitFieldValue(ast, structName, level);
  }

  visitModelBody(ast, lastName, level) {
    assert.equal(ast.type, 'modelBody');
    var fields = [];
    const structMap = [];
    let node;
    for (let i = 0; i < ast.nodes.length; i++) {
      node = ast.nodes[i];
      let comments = DSL.comment.getFrontComments(this.comments, node.tokenRange[0]);
      this.visitComments(comments, level);
      var fieldName = _name(node.fieldName);
      var type = '';
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
      let omitemptyEnable = true;
      if (node.fieldValue.fieldType === 'array') {
        type = `type:"Repeated"`;
        if (this.config.go && this.config.go.mapAndSliceWithoutOmitempty === true) {
          omitemptyEnable = false;
        }
        if (_name(node.fieldValue.fieldItemType)) {
          this.emit(`[]${_pointerType(_name(node.fieldValue.fieldItemType))} `);
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
    this.addSetFunc(ast, lastName, level - 1);
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

  addSetFunc(ast, structName, level) {
    assert.equal(ast.type, 'modelBody');
    this.emit(`func (s ${structName}) String() string {\n`, level);
    this.emit(`return tea.Prettify(s)\n`, level + 1);
    this.emit(`}\n`, level);
    this.emit(`\n`, level);
    this.emit(`func (s ${structName}) GoString() string {\n`, level);
    this.emit(`return s.String()\n`, level + 1);
    this.emit(`}\n`, level);
    this.emit(`\n`, level);
    for (let i = 0; i < ast.nodes.length; i++) {
      const node = ast.nodes[i];
      const fieldName = _format(_name(node.fieldName));
      const fileldtype = structName + _format(fieldName);
      const itemName = structName + _format(fieldName);
      if (node.fieldValue.fieldType === 'array') {
        if (_name(node.fieldValue.fieldItemType)) {
          this.emit(`func (s *${structName}) Set${fieldName}(v []${_pointerType(_name(node.fieldValue.fieldItemType))}) *${structName} {\n`, level);
          this.emit(`s.${fieldName} = v\n`, level + 1);
          this.emit(`return s\n`, level + 1);
          this.emit(`}\n`, level);
          this.emit(`\n`, level);
        } else if (node.fieldValue.fieldItemType.type === 'map') {
          this.emit(`func (s *${structName}) Set${fieldName}(v []`, level);
          this.visitType(node.fieldValue.fieldItemType);
          this.emit(`) *${structName} {\n`);
          this.emit(`s.${fieldName} = v\n`, level + 1);
          this.emit(`return s\n`, level + 1);
          this.emit(`}\n`, level);
          this.emit(`\n`, level);
        } else if (node.fieldValue.fieldItemType.type === 'modelBody') {
          this.emit(`func (s *${structName}) Set${fieldName}(v []${_pointerType(fileldtype)}) *${structName} {\n`, level);
          this.emit(`s.${fieldName} = v\n`, level + 1);
          this.emit(`return s\n`, level + 1);
          this.emit(`}\n`, level);
          this.emit(`\n`, level);
        } else if (node.fieldValue.fieldItemType.fieldType === 'array') {
          this.emit(`func (s *${structName}) Set${fieldName}(v [][]`, level);
          this.emitFuncArray(node.fieldValue.fieldItemType, itemName);
          this.emit(`) *${structName} {\n`);
          this.emit(`s.${fieldName} = v\n`, level + 1);
          this.emit(`return s\n`, level + 1);
          this.emit(`}\n`, level);
          this.emit(`\n`, level);
        }
      } else if (node.fieldValue.type === 'modelBody') {
        this.emit(`func (s *${structName}) Set${fieldName}(v *${fileldtype}) *${structName} {\n`, level);
        this.emit(`s.${fieldName} = v\n`, level + 1);
        this.emit(`return s\n`, level + 1);
        this.emit(`}\n`, level);
        this.emit(`\n`, level);
      } else if (_name(node.fieldValue.fieldType) && node.fieldValue.fieldType.idType === 'module') {
        const fieldType = node.fieldValue.fieldType;
        this.emit(`func (s *${structName}) Set${fieldName}(v ${this.clientName[_name(fieldType)]}) *${structName} {\n`, level);
        this.emit(`s.${fieldName} = v\n`, level + 1);
        this.emit(`return s\n`, level + 1);
        this.emit(`}\n`, level);
        this.emit(`\n`, level);
      } else if (node.fieldValue.fieldType.type === 'moduleModel' || node.fieldValue.fieldType.type === 'moduleTypedef' || node.fieldValue.fieldType.type === 'subModel') {
        this.emit(`func (s *${structName}) Set${fieldName}(v `, level);
        this.visitType(node.fieldValue.fieldType, level);
        this.emit(`) *${structName} {\n`, level);
        this.emit(`s.${fieldName} = v\n`, level + 1);
        this.emit(`return s\n`, level + 1);
        this.emit(`}\n`, level);
        this.emit(`\n`, level);
      } else {
        var fieldType = '';
        if (!_name(node.fieldValue.fieldType)) {
          fieldType = node.fieldValue.fieldType;
        } else {
          fieldType = _name(node.fieldValue.fieldType);
        }

        this.emit(`func (s *${structName}) Set${fieldName}(v `, level);
        this.visitType(node.fieldValue, level, {});
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
  }

  visitModel(ast, level) {
    assert.equal(ast.type, 'model');
    const modelName = _format(_name(ast.modelName));
    this.visitAnnotation(ast.annotation, level);
    let comments = DSL.comment.getFrontComments(this.comments, ast.tokenRange[0]);
    this.visitComments(comments, level);
    this.emit(`type ${modelName} struct {\n`, level);
    this.visitModelBody(ast.modelBody, modelName, level + 1);
  }

  visitObjectFieldValue(ast, level, env, expected) {
    this.visitExpr(ast, level, env, expected);
  }

  visitObjectField(ast, level, env, expected) {
    assert.equal(ast.type, 'objectField');
    let comments = DSL.comment.getFrontComments(this.comments, ast.tokenRange[0]);
    this.visitComments(comments, level);
    var key = _name(ast.fieldName);
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
          this.emit('tea.Merge(');
        } else {
          this.emit('tea.ToMap(');
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
          this.emit('},', level + 1);
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

  visitMethodCall(ast, level, env, argHasThrowFunc) {
    assert.equal(ast.left.type, 'method_call');
    if (!ast.isStatic) {
      this.emit(`client.`);
    }
    this.emit(`${_format(_name(ast.left.id))}(`);
    for (let i = 0; i < ast.args.length; i++) {
      const expr = ast.args[i];
      if (expr.needCast) {
        this.emit('tea.ToMap(');
        if (argHasThrowFunc && argHasThrowFunc.get(i)) {
          this.emit(argHasThrowFunc.get(i));
        } else {
          this.visitExpr(expr, level, env);
        }
        this.emit(')');
      } else {
        if ((expr.expectedType.name === 'number' || expr.expectedType.name === 'integer') && expr.inferred.name === 'int32') {
          this.emit(`tea.ToInt(`);
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, { pointer: true });
          }
          this.emit(`)`);
        } else if (expr.type !== 'number' && (expr.inferred.name === 'number' || expr.inferred.name === 'integer') && expr.expectedType.name === 'int32') {
          this.emit(`tea.ToInt32(`);
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, { pointer: true });
          }
          this.emit(`)`);
        } else {
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, { pointer: true });
          }
        }
      }
      if (i !== ast.args.length - 1) {
        this.emit(', ');
      }
    }
    this.emit(')');
  }

  visitInstanceCall(ast, level, env, argHasThrowFunc) {
    assert.equal(ast.left.type, 'instance_call');
    const method = ast.left.propertyPath[0];
    if (_name(ast.left.id).indexOf('@') === 0) {
      this.emit(`client.${_format(_name(ast.left.id).substring(1))}.${_format(_name(method))}(`);
    } else {
      this.emit(`${_name(ast.left.id)}.${_format(_name(method))}(`);
    }
    for (let i = 0; i < ast.args.length; i++) {
      const expr = ast.args[i];
      if (expr.needCast) {
        this.emit('tea.ToMap(');
        if (argHasThrowFunc && argHasThrowFunc.get(i)) {
          this.emit(argHasThrowFunc.get(i));
        } else {
          this.visitExpr(expr, level, env);
        }
        this.emit(')');
      } else {
        if ((expr.expectedType.name === 'number' || expr.expectedType.name === 'integer') && expr.inferred.name === 'int32') {
          this.emit(`tea.ToInt(`);
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, { pointer: true });
          }
          this.emit(`)`);
        } else if (expr.type !== 'number' && (expr.inferred.name === 'number' || expr.inferred.name === 'integer') && expr.expectedType.name === 'int32') {
          this.emit(`tea.ToInt32(`);
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, { pointer: true });
          }
          this.emit(`)`);
        } else {
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, { pointer: true });
          }
        }
      }
      if (i !== ast.args.length - 1) {
        this.emit(', ');
      }
    }
    this.emit(')');
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

  visitStaticCall(ast, level, env, argHasThrowFunc) {
    assert.equal(ast.left.type, 'static_call');
    this.emit(`${_importFilter(_format(_name(ast.left.id)).toLowerCase())}.${_format(_name(ast.left.propertyPath[0]))}(`);
    for (let i = 0; i < ast.args.length; i++) {
      const expr = ast.args[i];
      if (expr.needCast) {
        this.emit('tea.ToMap(');
        if (argHasThrowFunc && argHasThrowFunc.get(i)) {
          this.emit(argHasThrowFunc.get(i));
        } else {
          this.visitExpr(expr, level, env);
        }
        this.emit(')');
      } else {
        if ((expr.expectedType.name === 'number' || expr.expectedType.name === 'integer') && expr.inferred.name === 'int32') {
          this.emit(`tea.ToInt(`);
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, { pointer: true });
          }
          this.emit(`)`);
        } else if (expr.type !== 'number' && (expr.inferred.name === 'number' || expr.inferred.name === 'integer') && expr.expectedType.name === 'int32') {
          this.emit(`tea.ToInt32(`);
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, { pointer: true });
          }
          this.emit(`)`);
        } else {
          if (argHasThrowFunc && argHasThrowFunc.get(i)) {
            this.emit(argHasThrowFunc.get(i));
          } else {
            this.visitExpr(expr, level, env, { pointer: true });
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
          fieldType = `[]${_type(_name(ast.inferred.itemType))}`;
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
    if (ast.type === 'boolean') {
      if (expected && expected.pointer) {
        this.emit(`tea.Bool(${ast.value})`);
      } else {
        this.emit(ast.value);
      }
    } else if (ast.type === 'null') {
      if (_name(ast.inferred)) {
        var name = _name(ast.inferred);
        this.emit(`${_initValue(name)}`);
      } else {
        if (ast.inferred.type === 'array') {
          const type = `[ ${_name(ast.inferred.subType)} ]`;
          this.emit(`${_initValue(type)}`);
        }
        if (ast.inferred.type === 'map') {
          const type = `map[${_type(_name(ast.inferred.keyType))}]${_type(_name(ast.inferred.valueType))}`;
          this.emit(`${_initValue(type)}`);
        }
      }
    } else if (ast.type === 'property_access') {
      this.visitPropertyAccess(ast, level, env, expected);
    } else if (ast.type === 'string') {
      if (expected && expected.pointer) {
        this.emit(`tea.String("${_string(ast.value).replace(/"/g, '\\"')}")`);
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
          this.emit(`tea.Int(${ast.value.value})`);
        }
      } else {
        this.emit(ast.value.value);
      }
    } else if (ast.type === 'object') {
      this.visitObject(ast, level, env, expected);
    } else if (ast.type === 'variable') {
      isPointer = true;
      if ((!expected || !expected.pointer) && ((ast.inferred.type === 'array' &&
        ast.inferred.itemType.type === 'basic') || (DSL.util.isBasicType(_name(ast.inferred))
          && !_isFilterType(_name(ast.inferred))))) {
        isPointer = false;
        this.emit(_setValueFunc(_name(ast.inferred)));
      }
      var id = _name(ast.id);
      if (id === '__response') {
        this.emit(RESPONSE);
      } else if (id === '__request') {
        this.emit(REQUEST);
      } else if (ast.inferred && _name(ast.inferred) === 'class') {
        this.emit('new(' + _avoidReserveName(id) + ')');
      } else {
        this.emit(_avoidReserveName(id));
      }

      if (!isPointer) {
        this.emit(`)`);
      }
    } else if (ast.type === 'virtualVariable') {
      if ((!expected || !expected.pointer) && ((ast.inferred.type === 'array' &&
        ast.inferred.itemType.type === 'basic') || (DSL.util.isBasicType(_name(ast.inferred))
          && !_isFilterType(_name(ast.inferred))))) {
        this.emit(`${_setValueFunc(_name(ast.inferred))}${_vid((_name(ast.vid)))})`);
      } else {
        this.emit(`${_vid(_avoidVariableKeywords(_format(_name(ast.vid))))}`);
      }
    } else if (ast.type === 'template_string') {
      var j = 0;
      if (expected && expected.pointer) {
        this.emit(`tea.String(`);
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
            this.emit(`tea.ToString(`);
            this.visitExpr(expr, level, env);
            this.emit(`)`);
          } else {
            this.visitExpr(expr, level, env);
          }
        } else {
          throw new Error('unimpelemented');
        }
      }
      if (expected && expected.pointer) {
        this.emit(`)`);
      }
    } else if (ast.type === 'call') {
      if ((!expected || !expected.pointer) && ((ast.inferred.type === 'array' &&
        ast.inferred.itemType.type === 'basic') || (DSL.util.isBasicType(_name(ast.inferred))
          && !_isFilterType(_name(ast.inferred))))) {
        this.emit(`${_setValueFunc(_name(ast.inferred))}`);
        this.visitCall(ast, level, env, argHasThrowFunc);
        this.emit(`)`);
      } else {
        this.visitCall(ast, level, env, argHasThrowFunc);
      }
    } else if (ast.type === 'and') {
      this.visitExpr(ast.left, level, env);
      this.emit(` && `);
      this.visitExpr(ast.right, level, env);
    } else if (ast.type === 'or') {
      this.visitExpr(ast.left, level, env);
      this.emit(` || `);
      this.visitExpr(ast.right, level, env);
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

  visitSuper(ast, level, env) {
    assert.equal(ast.type, 'super');
    this.emit(`_err = client.${this.structName}.Init(`);
    for (let i = 0; i < ast.args.length; i++) {
      this.visitExpr(ast.args[i], level, env);
    }
    this.emit(`)\n`, level);
    this.emit(`if _err != nil {\n`, level);
    this.emit(`return _err\n`, level + 1);
    this.emit(`}`, level);
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
    this.emit(`${_format(_name(ast.aliasId)).toLowerCase()}.New${this.constructFunc[_name(ast.aliasId)]}(`);
    for (let i = 0; i < ast.args.length; i++) {
      this.visitExpr(ast.args[i], level, env);
    }
    this.emit(`)`);
  }

  visitConstructModel(ast, level, env) {
    assert.equal(ast.type, 'construct_model');
    if (ast.inferred.moduleName) {
      this.emit(`&${_format(_name(ast.aliasId)).toLowerCase()}.`);
    } else {
      this.emit(`&${_format(_name(ast.aliasId))}`);
    }
    for (let i = 0; i < ast.propertyPath.length; i++) {
      const item = ast.propertyPath[i];
      this.emit(`${_format(_name(item))}`);
    }
    if (ast.object && ast.object.fields.length > 0) {
      this.emit(`{\n`);
      const fields = ast.object.fields;
      for (let i = 0; i < fields.length; i++) {
        var field = fields[i];
        var str = '';
        var expected = field.expectedType;
        expected.pointer = true;
        if (field.expr.type === 'property_access') {
          expected.needCast = 'false';
        }
        let comments = DSL.comment.getFrontComments(this.comments, field.tokenRange[0]);
        this.visitComments(comments, level + 1);
        this.emit(`${_format(_name(field.fieldName))}: ${str}`, level + 1);
        if ((_name(field.expectedType) === 'number' || _name(field.expectedType) === 'integer') && _name(field.expr.inferred) === 'int32') {
          this.emit(`tea.ToInt(`);
          this.visitExpr(field.expr, level, env, expected);
          this.emit(`)`);
        } else if (field.expr.type !== 'number' && (_name(field.expr.inferred) === 'number' || _name(field.expr.inferred) === 'integer') && _name(field.expectedType) === 'int32') {
          this.emit(`tea.ToInt32(`);
          this.visitExpr(field.expr, level, env, expected);
          this.emit(`)`);
        } else {
          this.visitExpr(field.expr, level, env, expected);
        }
        if (str !== '') {
          this.emit(`)`);
        }
        this.emit(`,\n`);
      }
      this.emit(`}`, level);
    } else {
      this.emit(`{`);
      let comments = DSL.comment.getBetweenComments(this.comments, ast.tokenRange[0], ast.tokenRange[1]);
      if (comments.length > 0) {
        this.emit('\n');
        this.visitComments(comments, level + 1);
        this.emit('', level);
      }
      this.emit(`}`);
    }
  }

  visitReturn(ast, level, env) {
    assert.equal(ast.type, 'return');
    if (!ast.expr) {
      if (env.hasThrow) {
        this.emit(`return _err\n`, level);
      } else {
        this.emit(`return\n`, level);
      }
      return;
    }

    if (ast.expr.type === 'null' || ast.expr.type === 'variable' ||
      ast.expr.type === 'property_access' || ast.expr.type === 'virtualVariable') {
      this.emit(`_result = `, level);
      this.visitExpr(ast.expr, level, env, { pointer: true });
      this.emit(`\n`);
      if (env.hasThrow) {
        this.emit(`return _result , _err`, level);
      } else {
        this.emit(`return _result`, level);
      }
      this.emit(`\n`);
      return;
    }

    var returnType = '';
    if (env.returnType.idType === 'module') {
      this.emit(`_result = ${this.clientName[_name(env.returnType)].replace('*', '&')}{}\n`, level);
    } else if (_name(env.returnType) && !(DSL.util.isBasicType(_name(env.returnType)) && !_isFilterType(_name(env.returnType)))
      && env.returnType.idType !== 'typedef') {
      this.emit(`_result = ${_initValue(_name(env.returnType))}\n`, level);
    } else if (env.returnType.path && env.returnType.type !== 'moduleTypedef') {
      for (let i = 0; i < env.returnType.path.length; i++) {
        const path = env.returnType.path[i];
        if (i === 0) {
          returnType += _name(path).toLowerCase();
        } else {
          returnType += '.' + _name(path);
        }
      }
      this.emit(`_result = ${_initValue(returnType)}\n`, level);
    } else if (env.returnType.type === 'map') {
      this.emit(`_result = make(`, level);
      this.visitPointerType(env.returnType, level, env);
      this.emit(`)\n`);
    } else if (env.returnType.type === 'array') {
      this.emit(`_result = make(`, level);
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
        this.emit(`return _result, _err\n`, level + 1);
        this.emit(`}\n`, level);
      }
      if (!ast.needCast) {
        this.emit(`_result = _body\n`, level);
      } else {
        if (env.hasThrow) {
          this.emit(`_err = tea.Convert(_body, &_result)\n`, level);
        } else {
          this.emit(`tea.Convert(_body, &_result)\n`, level);
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
          this.emit(
            `_result.${_format(_name(field.fieldName))} = `,
            level
          );
          this.visitExpr(field.expr);
          this.emit(`\n`);
          delete ast.expr.fields[i];
        }
      }
      if (env.hasThrow) {
        this.emit(`_err = tea.Convert(`, level);
      } else {
        this.emit(`tea.Convert(`, level);
      }
      this.visitExpr(ast.expr, level, env, { pointer: true });
      this.emit(`, &_result)\n`);
    } else if (ast.expr.items) {
      if (env.hasThrow) {
        this.emit(`_err = tea.Convert(`, level);
      } else {
        this.emit(`tea.Convert(`, level);
      }
      this.visitExpr(ast.expr, level, env, { pointer: true });
      this.emit(`, &_result)\n`);
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

    if (env.hasThrow) {
      this.emit(`return _result, _err\n`, level);
    } else {
      this.emit(`return _result\n`, level);
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
    this.visitExpr(ast.condition, level + 1, env, false, argHasThrowFunc);
    this.emit(' {\n');
    this.visitStmts(ast.stmts, level + 1, env);
    if (ast.elseIfs) {
      for (let i = 0; i < ast.elseIfs.length; i++) {
        let elseIf = ast.elseIfs[i];
        this.emit(`} else if `, level);
        this.visitExpr(elseIf.condition, level + 1, env);
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
    this.emit(`_err = tea.New${CORE}Error(`, level);
    this.visitObject(ast.expr, level, env, 'map[string]interface{}');
    this.emit(')\n');
    if (!env.returnType) {
      this.emit(`return _err\n`, level);
    } else if (_name(env.returnType) === 'void') {
      if (env.hasThrow) {
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
    if (ast.left.type === 'property_assign' || ast.left.type === 'property') {
      this.emit(``, level);
      this.visitPropertyAccess(ast.left, level, env, { needCast: 'false', type: 'pointer' });
    } else if (ast.left.type === 'variable') {
      this.emit(`${_name(ast.left.id)}`, level);
    } else if (ast.left.type === 'virtualVariable') {
      this.emit(`${_vid(_avoidVariableKeywords(_format(_name(ast.left.vid))))}`, level);
    } else if (ast.left.type === 'map_access') {
      this.emit(``, level);
      this.visitMapAccess(ast.left, level, env, { pointer: true });
    } else if (ast.left.type === 'array_access') {
      this.emit(``, level);
      this.visitArrayAccess(ast.left, level, env, { pointer: true });
    } else {
      throw new Error('unimpelemented');
    }

    var hasThrowCall = (ast.expr.type === 'call' &&
      ast.expr.hasThrow) || ast.expr.type === 'construct';
    if (hasThrowCall) {
      this.emit(`, _err = `);
    } else {
      this.emit(` = `);
    }
    if (ast.expr.needToReadable) {
      this.emit(`tea.ToReader(`);
      this.visitExpr(ast.expr, level, env, { needCast: 'false', pointer: true });
      this.emit(`)`);
    } else if (ast.expr.type === 'object' && ast.left.inferred &&
      ast.left.inferred.type === 'map' &&
      _name(ast.left.inferred.valueType) === 'any') {
      this.visitObject(ast.expr, level, env, 'map[string]interface{}');
    } else {
      if ((ast.left.inferred.name === 'number' || ast.left.inferred.name === 'integer') && ast.expr.inferred.name === 'int32') {
        this.emit(`tea.ToInt(`);
        this.visitExpr(ast.expr, level, env, { pointer: true });
        this.emit(`)`);
      } else if (ast.expr.type !== 'number' && (ast.expr.inferred.name === 'number' || ast.expr.inferred.name === 'integer') && ast.left.inferred.name === 'int32') {
        this.emit(`tea.ToInt32(`);
        this.visitExpr(ast.expr, level, env, { pointer: true });
        this.emit(`)`);
      } else {
        this.visitExpr(ast.expr, level, env, { pointer: true });
      }
    }
    this.emit(`\n`);
    if (hasThrowCall) {
      if (env.returnType && _name(env.returnType) !== 'void') {
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

  visitDeclare(ast, level, env) {
    var id = _name(ast.id);
    var expected = ast.expectedType || {};
    if (ast.expr.type === 'call') {
      var argHasThrowFunc = this.visitFunctionNested(ast.expr, level, env);
    }
    var hasThrowCall = (ast.expr.type === 'call' &&
      ast.expr.hasThrow) || ast.expr.type === 'construct';
    if (hasThrowCall) {
      this.emit(`${id}, _err := `, level);
    } else if (ast.expr.type === 'null') {
      this.emit(`var ${id} `, level);
      this.visitPointerType(ast.expectedType, level);
      this.emit('\n');
      return;
    } else {
      this.emit(`${id} := `, level);
    }
    expected.pointer = true;
    if (ast.expectedType) {
      if ((_name(ast.expectedType) === 'number' || _name(ast.expectedType) === 'integer') && ast.expr.inferred.name === 'int32') {
        this.emit(`tea.ToInt(`);
        this.visitExpr(ast.expr, level, env, expected, argHasThrowFunc);
        this.emit(`)`);
      } else if (ast.expr.type !== 'number' && (ast.expr.inferred.name === 'number' || ast.expr.inferred.name === 'integer') && _name(ast.expectedType) === 'int32') {
        this.emit(`tea.ToInt32(`);
        this.visitExpr(ast.expr, level, env, expected, argHasThrowFunc);
        this.emit(`)`);
      } else {
        this.visitExpr(ast.expr, level, env, expected, argHasThrowFunc);
      }
    } else {
      this.visitExpr(ast.expr, level, env, expected, argHasThrowFunc);
    }
    this.emit('\n');

    if (hasThrowCall) {
      if (_name(env.returnType) !== 'void') {
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

  visitReturnBody(ast, level, env) {
    assert.equal(ast.type, 'returnBody');
    this.visitStmts(ast.stmts, level, env);
  }

  visitFunctionBody(ast, level, env) {
    assert.equal(ast.type, 'functionBody');
    this.visitStmts(ast.stmts, level, env);
    const stmts = ast.stmts.stmts;
    const length = ast.stmts.stmts.length;
    if (_name(env.returnType) === 'void' && env.hasThrow && (length === 0 || (stmts[length - 1].type !== 'return' &&
      stmts[length - 1].type !== 'throw'))) {
      this.emit(`return _err\n`, level);
    }
    if ((length === 0 || (stmts[length - 1].type !== 'return' && stmts[length - 1].type !== 'throw' &&
      (stmts[length - 1].type !== 'if' || !stmts[length - 1].elseStmts)))) {
      if (_name(env.returnType) !== 'void' && env.hasThrow) {
        this.emit(`return _result, _err\n`, level);
      } else if (_name(env.returnType) !== 'void' && !env.hasThrow) {
        this.emit(`return _result\n`, level);
      }
    }
  }

  eachFunc(ast, level, predefined, apis) {
    const env = {
      predefined,
      apis,
      local: new Map(),
      returnType: ast.returnType,
      hasThrow: ast.isAsync || ast.hasThrow,
      nestFuncParamName: new Map(),
      nestFuncParamNameSubscript: { 'count': 0 }
    };
    const functionName = _name(ast.functionName);
    this.visitAnnotation(ast.annotation, level);
    let comments = DSL.comment.getFrontComments(this.comments, ast.tokenRange[0]);
    this.visitComments(comments, level);
    if (this.exec && _format(functionName) === 'Main') {
      this.emit(`func _main `, level);
    } else if (ast.isStatic) {
      this.emit(`func ${_format(functionName)} `, level);
    } else {
      this.emit(`func (client *${this.structName}) ${_format(functionName)} `, level);
    }
    this.visitParams(ast.params, level, env);
    this.visitReturnType(ast, level, env);
    if (ast.functionBody) {
      this.visitFunctionBody(ast.functionBody, level + 1, env);
    } else {
      this.emit(`panic("No Support!")\n`, level + 1);
    }
    this.emit(`}\n`, level);
    this.emit(`\n`, level);
  }

  visitRuntimeAfter(ast, level, env) {
    this.emit(`}()\n`, level + 1);
    this.emit(`if !tea.BoolValue(tea.Retryable(_err)) {\n`, level + 1);
    this.emit(`break\n`, level + 2);
    this.emit(`}\n`, level + 1);
    this.emit(`}\n`, level);
    this.emit(`\n`);
    if (_name(env.returnType) !== 'void') {
      this.emit(`return _resp, _err\n`, level);
    } else {
      this.emit(`return _err\n`, level);
    }
  }

  eachAPI(ast, level, predefined) {
    // if (ast.annotation) {
    //   this.emit(`${_anno(ast.annotation.value)}\n`, level);
    // }
    const env = {
      // params, paramMap, returnType,
      predefined,
      returnType: ast.returnType,
      local: new Map(),
      hasThrow: true,
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
    for (let i = 0; i < ast.params.params.length; i++) {
      const param = ast.params.params[i];
      if (param.needValidate) {
        this.emit(`_err = tea.Validate(${_name(param.paramName)})\n`, level + 1);
        this.emit(`if _err != nil {\n`, level + 1);
        if (_name(ast.returnType) === 'void') {
          this.emit(`return _err\n`, level + 2);
        } else {
          this.emit(`return _result, _err\n`, level + 2);
        }
        this.emit(`}\n`, level + 1);
      }
    }
    // this.emit(` (*map[string]interface{}, error) {\n`);
    let baseLevel = ast.runtimeBody ? level + 2 : level;
    // api level
    if (ast.runtimeBody) {
      this.visitRuntimeBefore(ast.runtimeBody, level + 1, env);
    }

    // temp level
    this.visitAPIBody(ast.apiBody, baseLevel + 1, env);

    // if (ast.runtimeBody) {
    //   this.emit(`_lastRequest = ${REQUEST}\n`, baseLevel + 1);
    // }

    this.emit(`${RESPONSE}, _err := tea.DoRequest(${REQUEST}`, baseLevel + 1);

    if (ast.runtimeBody) {
      this.emit(`, _runtime`);
    } else {
      this.emit(`, nil`);
    }
    this.emit(`)\n`);

    this.emit(`if _err != nil {\n`, baseLevel + 1);
    if (_name(env.returnType) === 'void') {
      this.emit(`return _err\n`, baseLevel + 2);
    } else {
      this.emit(`return _result, _err\n`, baseLevel + 2);
    }
    this.emit(`}\n`, baseLevel + 1);

    if (ast.returns) {
      this.visitReturnBody(ast.returns, baseLevel + 1, env);
    } else {
      this.visitDefaultReturnBody(baseLevel + 1, env);
    }

    if (ast.runtimeBody) {
      this.visitRuntimeAfter(ast.runtimeBody, level + 1, env);
    }

    this.emit(`}\n`, level);
    this.emit(`\n`, level);
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

  importBefore(__module, level) {
    if (!this.editable) {
      this.emit(`// This file is auto-generated, don't edit it. Thanks.\n`, level);
    }
  }

  importAfter(__module, level) {
    // Nothing
  }

  moduleBefore(__module, level) {
    if (this.exec) {
      this.emit(`package main\n\n`);
    } else {
      this.emit(`package client\n\n`);
    }
    this.emit(`import (\n`);
    if (this.exec) {
      this.emit(`  "os"\n`);
    }
    if (this.goPackages !== '') {
      this.emit(`${this.goPackages}\n`);
    }
    if (this.importPackages !== '') {
      this.emit(`${this.importPackages}`);
    }
    this.emit(`  "github.com/alibabacloud-go/tea/tea"\n`);
    this.emit(`)\n\n`);
  }

  modelBefore() {
    // Nothing
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

  wrapAfter(__module, level) {
    // Nothing
  }

  moduleAfter() {
    this.uselessPack.forEach((pack) => {
      this.output = this.output.replace(`  "${pack}"\n`, '');
    });
  }

  typeRelover(type, module) {
    if (module && module.idType === 'module') {
      const aliasId = _name(module);
      if (this.importsTypedef[aliasId] && this.importsTypedef[aliasId][type.lexeme]) {
        let index = this.uselessPack.indexOf(this.importsTypedef[aliasId][type.lexeme].import);
        if (index > -1) {
          this.uselessPack.splice(index, 1);
        }
        return this.importsTypedef[aliasId][type.lexeme].type;
      }
    }
    if (type.idType === 'typedef' && this.typedef[type.lexeme]) {
      if (this.typedef[type.lexeme].import && this.typedef[type.lexeme].type) {
        let index = this.uselessPack.indexOf(this.typedef[type.lexeme].import);
        if (index > -1) {
          this.uselessPack.splice(index, 1);
        }
        return this.typedef[type.lexeme].type;
      }
    }
    return _name(type);
  }
}

module.exports = Visitor;
