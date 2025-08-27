'use strict';
const DSL = require('@darabonba/parser');
const  { _vid, _name, _upperFirst, _setValueFunc, _isFilterType } = require('./helper');

const types = [
  'integer', 'int8', 'int16', 'int32', 
  'int64', 'long', 'ulong', 'string',
  'uint8', 'uint16', 'uint32', 'uint64',
  'number', 'float', 'double', 'boolean',
  'bytes', 'readable', 'writable', 'object', 'any'
];

const integers = [
  'int8', 'int16', 'int32', 
  'int64', 'uint32', 'uint64',
  'uint8', 'uint16', 
];

class Builtin {
  constructor(generator, module = '', methods = []){
    this.generator = generator;
    this.moduleName = module;
    
    methods.forEach(method => {
      this[method] = function(args, level, env, argHasThrowFunc) {
        this.generator.emit(`dara.${_upperFirst(method)}`);
        this.generator.visitArgs(args, level, env, argHasThrowFunc);
      };
    });
  }

  getInstanceName(ast) {
    if (ast.left.id.tag === DSL.Tag.Tag.VID) {
      this.generator.emit(`${_vid(ast.left.id)}`);
    } else {
      this.generator.emit(`${_name(ast.left.id)}`);
    }
  }

  visitExpr(args, level, env, argHasThrowFunc) {
    for(let i = 0; i < args.length; i++) {
      const expr = args[i];
      if (argHasThrowFunc && argHasThrowFunc.get(i)) {
        this.generator.emit(argHasThrowFunc.get(i));
      } else {
        this.generator.visitExpr(expr, level, env);
      }
      if(i !== args.length - 1) {
        this.generator.emit(', ');
      }
    }
  }
}

class Env extends Builtin {
  constructor(generator){
    super(generator);
  }

  get(args, level, env, argHasThrowFunc){
    this.generator.emit('os.Getenv(');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  set(args, level, env, argHasThrowFunc){
    this.generator.emit(`os.Setenv(`);
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }
}

class Logger extends Builtin {
  constructor(generator){
    const methods = ['log', 'info', 'debug', 'error', 'warning'];
    super(generator, 'fmt', []);
    methods.forEach(method => {
      this[method] = function(args, level, env, argHasThrowFunc) {
        this.generator.builtinModule.push({
          path: 'fmt'
        });
        this.generator.emit(`fmt.Printf("[${method.toUpperCase()}] %s\\n", `);
        this.visitExpr(args, level, env, argHasThrowFunc);
        this.generator.emit(')');
      };
    });
  }

  error(args, level, env, argHasThrowFunc){
    this.generator.emit(`dara.Fprintf(os.Stderr, "[ERROR] %s\\n", `);
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }
}

class XML extends Builtin {
  constructor(generator){
    const methods = ['parseXml', 'toXML'];
    super(generator, 'Dara', methods);
  }
}

class URL extends Builtin {
  constructor(generator){
    const methods = ['parse', 'urlEncode', 'percentEncode', 'pathEncode'];
    super(generator, 'Dara', methods);
  }

  parse(args, level, env, argHasThrowFunc) {
    this.generator.emit(`dara.ParseURL`);
    this.generator.visitArgs(args, level, env, argHasThrowFunc);
  }
}

class Stream extends Builtin {
  constructor(generator){
    const methods = ['readAsBytes', 'readAsJSON', 'readAsString', 'readAsSSE'];
    super(generator, 'DaraStream', methods);
  }
}

class Number extends Builtin {
  constructor(generator){
    const methods = ['random', 'floor', 'round', 'min', 'max'];
    super(generator, 'DaraMath', methods);
  }

  parseInt(ast) {
    this.generator.emit('int(');
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  parseLong(ast) {
    this.generator.emit('int64(');
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  parseFloat(ast) {
    this.generator.emit('float32(');
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  parseDouble(ast) {
    this.generator.emit('float64(');
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  itol(ast) {
    this.generator.emit('int64(');
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  ltoi(ast) {
    this.generator.emit('int(');
    this.getInstanceName(ast);
    this.generator.emit(')');
  }
}

class JSON extends Builtin {
  constructor(generator){
    const methods = ['stringify', 'parseJSON'];
    super(generator, 'DaraJSON', methods);
  }

}

class Form extends Builtin {
  constructor(generator){
    const methods = ['toFormString', 'getBoundary', 'toFileForm'];
    super(generator, `DaraForm`, methods);
  }
}

class File extends Builtin {
  constructor(generator){
    const methods = ['createReadStream', 'createWriteStream', 'exists'];
    super(generator, 'DaraFile', methods);
  }
}

class Bytes extends Builtin {
  constructor(generator){
    const methods = [];
    super(generator, 'DaraBytes', methods);
  }

  from(args, level, env, argHasThrowFunc) {
    this.generator.emit(`dara.BytesFromString`);
    this.generator.visitArgs(args, level, env, argHasThrowFunc);
  }

  toString(ast, level){ 
    this.generator.emit('dara.ToString(');
    this.getInstanceName(ast, level);
    this.generator.emit(')');
  }

  toHex(ast, level){ 
    this.generator.builtinModule.push({
      path: 'encoding/hex'
    });
    this.generator.emit('hex.EncodeToString(');
    this.getInstanceName(ast, level);
    this.generator.emit(')');
  }

  toBase64(ast, level){ 
    this.generator.builtinModule.push({
      path: 'encoding/base64'
    });
    this.generator.emit('base64.StdEncoding.EncodeToString(');
    this.getInstanceName(ast, level);
    this.generator.emit(')');
  }

  toJSON(ast, level) {
    this.toString(ast, level);
  }

  length(ast){ 
    this.generator.emit('len(');
    this.getInstanceName(ast);
    this.generator.emit(')');
  }
}


class Converter extends Builtin {
  constructor(generator){
    super(generator);
    this.generator = generator;
    integers.forEach(type => {
      this[type] = function(args, level) {
        const expr = args[0];
        generator.emit(`dara.Force${_upperFirst(type)}(`);
        generator.visitExpr(expr, level, {
          pointer: false,
          pointerParams: []
        });
        generator.emit(')');
      };
    });
  }

  integer(args, level, env, argHasThrowFunc) {
    this.generator.emit('dara.ForceInt(');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  long(args, level, env, argHasThrowFunc) {
    this.generator.emit('dara.ForceInt64(');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  ulong(args, level, env, argHasThrowFunc) {
    this.generator.emit('dara.ForceUint64(');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  float(args, level, env, argHasThrowFunc) {
    this.generator.emit('dara.ForceFloat32(');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  double(args, level, env, argHasThrowFunc) {
    this.generator.emit('dara.ForceFloat64(');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  string(args, level, env, argHasThrowFunc) {
    this.generator.emit('dara.ToString(');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  number(args, level, env, argHasThrowFunc) {
    this.generator.emit('dara.ForceInt(');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  boolean(args, level, env, argHasThrowFunc) {
    this.generator.emit('dara.ForceBoolean(');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  bytes(args, level, env, argHasThrowFunc) {
    this.generator.emit('[]byte(dara.ToString(');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit('))');
  }

  any(args, level, env, argHasThrowFunc){
    this.visitExpr(args, level, env, argHasThrowFunc);
  }

  object(args, level, env, argHasThrowFunc){
    this.generator.emit(`dara.ToMap(`);
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');

  }

  readable(args, level, env, argHasThrowFunc){
    this.generator.emit(`dara.ToReader(`);
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  writable(args, level, env, argHasThrowFunc){
    this.generator.emit(`dara.ToWriter(`);
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }
}

class Func {
  constructor(generator){
    this.generator = generator;
  }
  
  default(args, level, env, argHasThrowFunc) {
    this.generator.emit('dara.Default(');
    this.generator.visitExpr(args[0], level, env, argHasThrowFunc);
    this.generator.emit(', ');
    this.generator.visitExpr(args[1], level, env, argHasThrowFunc);
    this.generator.emit(')');
  }
  
  isNull(args, level, env, argHasThrowFunc) {
    this.generator.emit('dara.IsNil(');
    this.generator.visitExpr(args[0], level, env, {pointer: true, pointerParams: []});
    this.generator.emit(')');
  }

  sleep(args, level, env, argHasThrowFunc) {
    this.generator.builtinModule.push({
      path: 'time'
    });
    this.generator.emit('time.Sleep(time.Duration(');
    this.generator.visitExpr(args[0], level, env, argHasThrowFunc);
    this.generator.emit(') * time.Millisecond)');
  }

  equal(args, level, env, argHasThrowFunc) {
    this.generator.visitExpr(args[0], level, env, argHasThrowFunc);
    this.generator.emit(' == ');
    this.generator.visitExpr(args[1], level, env, argHasThrowFunc);
  }
}

function convertJsRegexStrToRegExp(jsRegexStr) {
  const matches = jsRegexStr.match(/^\/(.+)\/([gimuy]*)$/);
  if (!matches) {
    throw new Error('正则表达式格式错误应为 /pattern/flags');
  }
  const [, pattern, flags] = matches;
  
  let goRegexStr = pattern;
  if (flags.includes('i')) {
    goRegexStr = '(?i)' + goRegexStr;
  }

  return goRegexStr;
}

class String extends Builtin {

  constructor(generator){
    super(generator, 'DaraString');
  }

  replace(ast, level, env, argHasThrowFunc) {
    this.generator.builtinModule.push({
      path: 'regexp'
    });
    const args = ast.args;
    const regex = convertJsRegexStrToRegExp(args[0].value.string);
    this.generator.emit(`regexp.MustCompile(\`${regex}\`).ReplaceAllString(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[1], level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  contains(ast, level, env, argHasThrowFunc) {
    this.generator.emit(`dara.Contains(`);
    this.getInstanceName(ast);
    const args = ast.args;
    this.generator.emit(', ');
    this.generator.visitExpr(args[0], level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  length(ast) {
    this.generator.emit(`dara.Length(`);
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  hasPrefix(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.generator.emit(`dara.HasPrefix(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[0], level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  hasSuffix(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.generator.emit(`dara.HasSuffix(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[0], level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  index(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.generator.emit(`dara.Index(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[0], level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  subString(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.getInstanceName(ast);
    this.generator.emit('[');
    this.generator.visitExpr(args[0], level, env, argHasThrowFunc);
    this.generator.emit(': ');
    this.generator.visitExpr(args[1], level, env, argHasThrowFunc);
    this.generator.emit(']');
  }

  trim(ast) {
    this.generator.emit(`dara.TrimSpace(`);
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  toLower(ast) {
    this.generator.emit(`dara.ToLower(`);
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  toUpper(ast) {
    this.generator.emit(`dara.ToUpper(`);
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  equals(ast, level, env, argHasThrowFunc) {
    this.getInstanceName(ast);
    const args = ast.args;
    const expr = args[0];
    this.generator.emit(' == ');
    this.generator.visitExpr(expr, level, env, argHasThrowFunc);

  }

  empty(ast) {
    this.generator.emit('len(');
    this.getInstanceName(ast);
    this.generator.emit(')');
    this.generator.emit(' == 0');
  }

  toBytes(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.generator.emit('dara.ToBytes(');
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  parseInt(ast) {
    this.generator.builtinModule.push({
      path: 'strconv'
    });
    this.generator.emit('strconv.Atoi(');
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  parseLong(ast) {
    this.generator.builtinModule.push({
      path: 'strconv'
    });
    this.generator.emit('strconv.ParseInt(');
    this.getInstanceName(ast);
    this.generator.emit(', 10, 64)');
  }

  parseFloat(ast) {
    this.generator.builtinModule.push({
      path: 'strconv'
    });
    this.generator.emit('strconv.ParseFloat(');
    this.getInstanceName(ast);
    this.generator.emit(', 32)');
  }

  parseDouble(ast) {
    this.generator.builtinModule.push({
      path: 'strconv'
    });
    this.generator.emit('strconv.ParseFloat(');
    this.getInstanceName(ast);
    this.generator.emit(', 64)');
  }
}

class Array extends Builtin {

  constructor(generator){
    super(generator, 'DaraArray');
  }

  join(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.generator.emit(`dara.ArrJoin(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  full(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.generator.emit(`dara.ArrFull(&`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  shift(ast) {
    this.generator.emit(`dara.ArrShift(&`);
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  pop(ast) {
    this.generator.emit(`dara.ArrPop(&`);
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  push(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.generator.emit(`dara.ArrPush(&`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  unshift(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.generator.emit(`dara.ArrUnshift(&`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  contains(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.generator.emit(`dara.ArrContains(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  length(ast) {
    this.generator.emit('len(');
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  index(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.generator.emit(`dara.ArrIndex(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  get(ast, level, env, argHasThrowFunc) {
    this.getInstanceName(ast);
    const args = ast.args;
    this.generator.emit(`[`);
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(`]`);
  }

  sort(ast, level, env, argHasThrowFunc) {
    this.generator.emit('dara.SortArr(');
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.visitExpr(ast.args, level, env, argHasThrowFunc);
    this.generator.emit(').(');
    this.generator.emit('[]');
    this.generator.visitPointerType(ast.inferred.itemType);
    this.generator.emit(')');
  }

  concat(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.generator.emit(`dara.ConcatArr(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(').(');
    this.generator.emit('[]');
    this.generator.visitPointerType(ast.inferred.itemType);
    this.generator.emit(')');
  }

  append(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.generator.emit(`dara.ArrAppend(&`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }

  remove(ast, level, env, argHasThrowFunc) {
    const args = ast.args;
    this.generator.emit(`dara.ArrRemove(&`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.visitExpr(args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }
}

class Map extends Builtin {

  constructor(generator){
    super(generator, 'DaraMap');
  }

  length(ast) {
    this.generator.emit(`len(`);
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  keySet(ast) {
    this.generator.emit(`dara.KeySet(`);
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  entries(ast) {
    this.generator.emit(`dara.Entries(`);
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  toJSON(ast) {
    this.generator.emit(`dara.Stringify(`);
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  merge(ast, level, env, argHasThrowFunc) {
    this.generator.emit('dara.ToMap(');
    this.getInstanceName(ast);
    this.generator.emit(' , ');
    this.visitExpr(ast.args, level, env, argHasThrowFunc);
    this.generator.emit(')');
  }
}

class Entry extends Builtin {

  key(ast) {
    this.getInstanceName(ast);
    this.generator.emit('.Key');
  }

  value(ast) {
    let setValue = false;
    if(((ast.inferred.type === 'array' && ast.inferred.itemType.type === 'basic') || 
      (DSL.util.isBasicType(_name(ast.inferred)) && !_isFilterType(_name(ast.inferred))))) {
      setValue = true;
      this.generator.emit(_setValueFunc(_name(ast.inferred)));
    }
    this.getInstanceName(ast);
    this.generator.emit('.Value');
    if(ast.inferred.type !== 'basic' || ast.inferred.name !== 'any') {
      this.generator.emit('.(');
      this.generator.visitPointerType(ast.inferred);
      this.generator.emit(')');
    }
    if(setValue) {
      this.generator.emit(')');
    }
  }
}

class ModelInstance extends Builtin {

  validate(ast) {
    this.getInstanceName(ast);
    this.generator.emit('.Validate(');
    this.generator.emit(')');
  }

  toMap(ast) {
    this.generator.emit('dara.ToMap(');
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

}


module.exports = (generator) => {
  const builtin = {};
  builtin['$Env'] = new Env(generator);
  builtin['$Logger'] = new Logger(generator);
  builtin['$XML'] = new XML(generator);
  builtin['$URL'] = new URL(generator);
  builtin['$Stream'] = new Stream(generator);
  builtin['$Number'] = new Number(generator);
  builtin['$JSON'] = new JSON(generator);
  builtin['$Form'] = new Form(generator);
  builtin['$File'] = new File(generator);
  builtin['$Bytes'] = new Bytes(generator);
  const converter = new Converter(generator);
  types.map(type => {
    builtin[`$${type}`] = converter;
  });

  const func = new Func(generator);
  builtin['$isNull'] = func;
  builtin['$sleep'] = func;
  builtin['$default'] = func;
  builtin['$equal'] = func;

  builtin['$String'] = new String(generator);
  builtin['$Array'] = new Array(generator);
  builtin['$Date'] = new Date(generator);
  builtin['$Map'] = new Map(generator);
  builtin['$Entry'] = new Entry(generator);
  builtin['$ModelInstance'] = new ModelInstance(generator);

  return builtin;
};