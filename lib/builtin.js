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
      this[method] = function(args, level, env) {
        this.generator.emit(`dara.${_upperFirst(method)}`);
        this.generator.visitArgs(args, level, env);
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
}

class Env extends Builtin {
  constructor(generator){
    super(generator);
  }

  get(args, level, env){
    const key = args[0];
    this.generator.emit('os.Getenv(');
    this.generator.visitExpr(key, level, env);
    this.generator.emit(')');
  }

  set(args, level, env){
    const key = args[0];
    this.generator.emit(`os.Setenv(`);
    this.generator.visitExpr(key, level, env);
    this.generator.emit(', ');
    const value = args[1];
    this.generator.visitExpr(value, level, env);
    this.generator.emit(')');
  }
}

class Logger extends Builtin {
  constructor(generator){
    const methods = ['log', 'info', 'debug', 'error', 'warning'];
    super(generator, 'fmt', []);
    methods.forEach(method => {
      this[method] = function(args, level, env) {
        this.generator.builtinModule.push({
          path: 'fmt'
        });
        this.generator.emit(`fmt.Printf("[${method.toUpperCase()}] %s\\n", `);
        const key = args[0];
        this.generator.visitExpr(key, level, env);
        this.generator.emit(')');
      };
    });
  }

  error(args, level, env){
    this.generator.emit(`dara.Fprintf(os.Stderr, "[ERROR] %s\\n", `);
    const key = args[0];
    this.generator.visitExpr(key, level, env);
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

  parse(args, level, env) {
    this.generator.emit(`dara.ParseURL`);
    this.generator.visitArgs(args, level, env);
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

  from(args, level, env) {
    this.generator.emit(`dara.BytesFromString`);
    this.generator.visitArgs(args, level, env);
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


class Converter {
  constructor(generator){
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

  integer(args, level, env) {
    const expr = args[0];
    this.generator.emit('dara.ForceInt(');
    this.generator.visitExpr(expr, level, env);
    this.generator.emit(')');
  }

  long(args, level, env) {
    const expr = args[0];
    this.generator.emit('dara.ForceInt64(');
    this.generator.visitExpr(expr, level, env);
    this.generator.emit(')');
  }

  ulong(args, level, env) {
    const expr = args[0];
    this.generator.emit('dara.ForceUint64(');
    this.generator.visitExpr(expr, level, env);
    this.generator.emit(')');
  }

  float(args, level, env) {
    const expr = args[0];
    this.generator.emit('dara.ForceFloat32(');
    this.generator.visitExpr(expr, level, env);
    this.generator.emit(')');
  }

  double(args, level, env) {
    const expr = args[0];
    this.generator.emit('dara.ForceFloat64(');
    this.generator.visitExpr(expr, level, env);
    this.generator.emit(')');
  }

  string(args, level, env) {
    const expr = args[0];
    this.generator.emit('dara.ToString(');
    this.generator.visitExpr(expr, level, env);
    this.generator.emit(')');
  }

  number(args, level, env) {
    const expr = args[0];
    this.generator.emit('dara.ForceInt(');
    this.generator.visitExpr(expr, level, env);
    this.generator.emit(')');
  }

  boolean(args, level, env) {
    const expr = args[0];
    this.generator.emit('dara.ForceBoolean(');
    this.generator.visitExpr(expr, level, env);
    this.generator.emit(')');
  }

  bytes(args, level, env) {
    const expr = args[0];
    this.generator.emit('[]byte(dara.ToString(');
    this.generator.visitExpr(expr, level, env);
    this.generator.emit('))');
  }

  any(args, level, env){
    const expr = args[0];
    this.generator.visitExpr(expr, level, env);
  }

  object(args, level, env){
    const expr = args[0];
    this.generator.emit(`dara.ToMap(`);
    this.generator.visitExpr(expr, level, env);
    this.generator.emit(')');

  }

  readable(args, level, env){
    const expr = args[0];
    this.generator.emit(`dara.ToReadable(`);
    this.generator.visitExpr(expr, level, env);
    this.generator.emit(')');
  }

  writable(args, level, env){
    const expr = args[0];
    this.generator.emit(`dara.ToWritable(`);
    this.generator.visitExpr(expr, level, env);
    this.generator.emit(')');
  }
}

class Func {
  constructor(generator){
    this.generator = generator;
  }
  
  default(args, level, env) {
    this.generator.emit('dara.Default(');
    this.generator.visitExpr(args[0], level, env);
    this.generator.emit(', ');
    this.generator.visitExpr(args[1], level, env);
    this.generator.emit(')');
  }
  
  isNull(args, level, env) {
    this.generator.emit('dara.IsNil(');
    this.generator.visitExpr(args[0], level, env, {pointer: true, pointerParams: []});
    this.generator.emit(')');
  }

  sleep(args, level, env) {
    this.generator.builtinModule.push({
      path: 'time'
    });
    this.generator.emit('time.Sleep(time.Duration(');
    this.generator.visitExpr(args[0], level, env);
    this.generator.emit(') * time.Millisecond)');
  }

  equal(args, level, env) {
    this.generator.visitExpr(args[0], level, env);
    this.generator.emit(' == ');
    this.generator.visitExpr(args[1], level, env);
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

  replace(ast, level, env) {
    this.generator.builtinModule.push({
      path: 'regexp'
    });
    const args = ast.args;
    const regex = convertJsRegexStrToRegExp(args[0].value.string);
    this.generator.emit(`regexp.MustCompile(\`${regex}\`).ReplaceAllString(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[1], level, env);
    this.generator.emit(')');
  }

  contains(ast, level, env) {
    this.generator.emit(`dara.Contains(`);
    this.getInstanceName(ast);
    const args = ast.args;
    this.generator.emit(', ');
    this.generator.visitExpr(args[0], level, env);
    this.generator.emit(')');
  }

  length(ast) {
    this.generator.emit(`dara.Length(`);
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  hasPrefix(ast, level, env) {
    const args = ast.args;
    this.generator.emit(`dara.HasPrefix(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[0], level, env);
    this.generator.emit(')');
  }

  hasSuffix(ast, level, env) {
    const args = ast.args;
    this.generator.emit(`dara.HasSuffix(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[0], level, env);
    this.generator.emit(')');
  }

  index(ast, level, env) {
    const args = ast.args;
    this.generator.emit(`dara.Index(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[0], level, env);
    this.generator.emit(')');
  }

  subString(ast, level, env) {
    const args = ast.args;
    this.getInstanceName(ast);
    this.generator.emit('[');
    this.generator.visitExpr(args[0], level, env);
    this.generator.emit(': ');
    this.generator.visitExpr(args[1], level, env);
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

  equals(ast, level, env) {
    this.getInstanceName(ast);
    const args = ast.args;
    const expr = args[0];
    this.generator.emit(' == ');
    this.generator.visitExpr(expr, level, env);

  }

  empty(ast) {
    this.generator.emit('len(');
    this.getInstanceName(ast);
    this.generator.emit(')');
    this.generator.emit(' == 0');
  }

  toBytes(ast, level, env) {
    const args = ast.args;
    const expr = args[0];
    this.generator.emit('dara.ToBytes(');
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(expr, level, env);
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

  join(ast) {
    const args = ast.args;
    this.generator.emit(`dara.ArrJoin(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    const expr = args[0];
    this.generator.visitExpr(expr);
    this.generator.emit(')');
  }

  full(ast) {
    const args = ast.args;
    this.generator.emit(`dara.ArrFull(&`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    const expr = args[0];
    this.generator.visitExpr(expr);
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

  push(ast) {
    const args = ast.args;
    this.generator.emit(`dara.ArrPush(&`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[0]);
    this.generator.emit(')');
  }

  unshift(ast) {
    const args = ast.args;
    this.generator.emit(`dara.ArrUnshift(&`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[0]);
    this.generator.emit(')');
  }

  contains(ast) {
    const args = ast.args;
    this.generator.emit(`dara.ArrContains(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[0]);
    this.generator.emit(')');
  }

  length(ast) {
    this.generator.emit('len(');
    this.getInstanceName(ast);
    this.generator.emit(')');
  }

  index(ast) {
    const args = ast.args;
    this.generator.emit(`dara.ArrIndex(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[0]);
    this.generator.emit(')');
  }

  get(ast) {
    this.getInstanceName(ast);
    const args = ast.args;
    this.generator.emit(`[`);
    const expr = args[0];
    this.generator.visitExpr(expr);
    this.generator.emit(`]`);
  }

  sort(ast) {
    this.generator.emit('dara.SortArr(');
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(ast.args[0]);
    this.generator.emit(').(');
    this.generator.emit('[]');
    this.generator.visitPointerType(ast.inferred.itemType);
    this.generator.emit(')');
  }

  concat(ast) {
    const args = ast.args;
    this.generator.emit(`dara.ConcatArr(`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[0]);
    this.generator.emit(').(');
    this.generator.emit('[]');
    this.generator.visitPointerType(ast.inferred.itemType);
    this.generator.emit(')');
  }

  append(ast) {
    const args = ast.args;
    this.generator.emit(`dara.ArrAppend(&`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[0]);
    this.generator.emit(')');
  }

  remove(ast) {
    const args = ast.args;
    this.generator.emit(`dara.ArrRemove(&`);
    this.getInstanceName(ast);
    this.generator.emit(', ');
    this.generator.visitExpr(args[0]);
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

  merge(ast, level, env) {
    this.generator.emit('dara.ToMap(');
    this.getInstanceName(ast);
    this.generator.emit(' , ');
    this.generator.visitExpr(ast.args[0], level, env);
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

  return builtin;
};