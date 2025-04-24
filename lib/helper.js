'use strict';

const DSL = require('@darabonba/parser');
const fs = require('fs');
const path = require('path');
const { Tag } = DSL.Tag;
const filterTypes = ['readable', 'writeable', 'map', 'object', 'any', 'bytes', 'class', '$Error', '$Model'];

const keyWords = ['String', 'string', 'Prettify', 'prettify', 'map'];
const variableKeyWords = ['Client', 'client','@Client', '@client'];

function _name(str) {
  return str.lexeme || str.name;
}

function _snakeCase(str) {
  if (!str) {
    return '';
  }

  let res = '';
  let tmp = '';
  
  for (const [, c] of [...str].entries()) {
    // 检查字符是否为大写字母或数字
    if (/[A-Z0-9]/.test(c)) {
      tmp += c;
    } else {
      if (tmp.length > 0) {
        res += (res === '' ? tmp.toLowerCase() : '_' + tmp.toLowerCase());
        tmp = '';
      }
      res += c;
    }
  }
  
  if (tmp.length > 0) {
    res += (res === '' ? tmp.toLowerCase() : '_' + tmp.toLowerCase());
  }

  return res;
}

function _importFilter(name) {
  if (keyWords.indexOf(name) !== -1) {
    name = name + '_';
  }
  return name;
}



function _deleteWithSuffix(directory, suffix) {
  if(!fs.existsSync(directory)) {
    return;
  }
  const files = fs.readdirSync(directory, { withFileTypes: true });
  
  for (const file of files) {
    const filePath = path.join(directory, file.name);

    if (file.isDirectory()) {
      _deleteWithSuffix(filePath, suffix);
    } else if (file.name.endsWith(suffix)) {
      fs.unlinkSync(filePath);
    }
  }
}

function _isIterator(type) {
  if (type.type === 'iterator' || type.type === 'asyncIterator') {
    return true;
  }
  return false;
}

function _isKeyWord(name) {
  return keyWords.indexOf(name) !== -1;
}

function _avoidVariableKeywords(name) {
  if (variableKeyWords.indexOf(name) !== -1) {
    return name + '_';
  }
  return name;
}

function _upperFirst(str) {
  str = str.replace(/-/g, '_');
  return str[0].toUpperCase() + str.substring(1);
}

function _lowerFirst(str) {
  str = str.replace(/-/g, '_');
  return str[0].toLowerCase() + str.substring(1);
}

function _vid(id) {
  return `client.` + _upperFirst(id.substr(1));
}

function _setExtendFunc(name) {
  var expr = '';
  if (name === 'number' || name === 'integer' || name === 'int') {
    expr = `dara.Int(`;
  } else if (name === 'long' || name === 'int64') {
    expr = `dara.Int64(`;
  } else if (name === 'double') {
    expr = `dara.Float64(`;
  } else if (name === 'float') {
    expr = `dara.Float32(`;
  } else if (name === '[]float64') {
    expr = `dara.Float64Slice(`;
  } else if (name === '[]float32') {
    expr = `dara.Float32Slice(`;
  } else if (name === 'boolean') {
    expr = `dara.Bool(`;
  } else if (name === '[]bool') {
    expr = `dara.BoolSlice(`;
  } else if (name === 'string') {
    expr = `dara.String(`;
  } else if (name === 'int16') {
    expr = `dara.Int16(`;
  } else if (name === 'int32') {
    expr = `dara.Int32(`;
  } else if (name === '[]string') {
    expr = `dara.StringSlice(`;
  } else if (name === '[]int') {
    expr = `dara.IntSlice(`;
  } else if (name === '[]int32') {
    expr = `dara.Int32Slice(`;
  } else if (name === '[]int64') {
    expr = `dara.Int64Slice(`;
  } else if (name === '[]uint') {
    expr = `dara.UintSlice(`;
  } else if (name === '[]uint32') {
    expr = `dara.Uint32Slice(`;
  } else if (name === '[]uint64') {
    expr = `dara.Uint64Slice(`;
  } else if (name === 'uint') {
    expr = `dara.Uint(`;
  } else if (name === 'uint32') {
    expr = `dara.Uint32(`;
  } else if (name === 'uint64') {
    expr = `dara.Uint64(`;
  }

  return expr;
}

function _setArrayFunc(name) {
  var expr = '';
  if (name === '[]float64') {
    expr = `dara.ToFloat64Slice(`;
  } else if (name === '[]float32') {
    expr = `dara.Float32Slice(`;
  } else if (name === '[]bool') {
    expr = `dara.ToBoolSlice(`;
  } else if (name === '[]string') {
    expr = `dara.ToStringSlice(`;
  } else if (name === '[]int') {
    expr = `dara.ToIntSlice(`;
  } else if (name === '[]int32') {
    expr = `dara.ToInt32Slice(`;
  } else if (name === '[]int64') {
    expr = `dara.ToInt64Slice(`;
  } else if (name === '[]uint') {
    expr = `dara.ToUintSlice(`;
  } else if (name === '[]uint32') {
    expr = `dara.ToUint32Slice(`;
  } else if (name === '[]uint64') {
    expr = `dara.ToUint64Slice(`;
  }

  return expr;
}

function _initValue(type) {
  if (type === 'number' || type === 'integer') {
    return `${_setExtendFunc(type)}0)`;
  }
  if (type === 'string') {
    return `${_setExtendFunc(type)}"")`;
  }
  if (type === 'boolean') {
    return `${_setExtendFunc(type)}false)`;
  }
  if (type === 'bytes') {
    return 'make([]byte, 0)';
  }
  if (type === 'any' || type === 'class' || type === '$Model') {
    return 'interface{}(nil)';
  }

  if (type === '$Error') {
    return '&dara.SDKError{}';
  }

  if (type === 'float') {
    return `${_setExtendFunc(type)}0.00)`;
  }

  if (type === 'null') {
    return 'nil';
  }

  if (type === '$Response') {
    return '&dara.Response{}';
  }

  if (type === '$Request') {
    return '&dara.Request{}';
  }

  if (type === '$SSEEvent') {
    return '&dara.SSEEvent{}';
  }

  if (type === '$Date') {
    return '&dara.Date{}';
  }

  if (type === '$File') {
    return '&dara.File{}';
  }

  if (type === '$RetryOptions') {
    return '&dara.RetryOptions{}';
  }

  if (type === '$RuntimeOptions') {
    return '&dara.RuntimeOptions{}';
  }

  if (type === '$ExtendsParameters') {
    return '&dara.ExtendsParameters{}';
  }

  if (type === '$URL') {
    return '&dara.URL{}';
  }

  if (type === '$Stream') {
    return '&dara.Stream{}';
  }

  if (type === 'object') {
    return 'make(map[string]interface{})';
  }

  if (type.startsWith('map')) {
    return `make(${type})`;
  }

  if (type.startsWith('[]')) {
    return `make(${type}, 1)`;
  }

  return `&${type}{}`;
}

function _string(str) {
  if (str.string === '""') {
    return '\\"\\"';
  }
  return str.string.replace(/([^\\])"+|^"/g, function(str){
    return str.replace(/"/g, '\\"');
  });
}

function _format(name) {
  var strs = name.split('-');
  name = _upperFirst(strs[0]);
  for (let i = 1; i < strs.length; i++) {
    name = name + _upperFirst(strs[i]);
  }
  if (_isKeyWord(name)) {
    name = name + '_';
  }
  return name;
}

function _extendFieldName(str){
  return str.split('.').map(name => {
    return _format(name);
  }).join('');
}

function _modelName(name) {
  if (name === '$Error') {
    return 'dara.SDKError';
  }

  if (name === '$ResponseError') {
    return 'dara.ResponseError';
  }

  if (name === '$SSEEvent') {
    return 'dara.SSEEvent';
  }

  if (name === '$RetryOptions') {
    return 'dara.RetryOptions';
  }

  if (name === '$RuntimeOptions') {
    return 'dara.RuntimeOptions';
  }

  if (name === '$ExtendsParameters') {
    return 'dara.ExtendsParameters';
  }

  if (name === '$Date') {
    return 'dara.Date';
  }

  if (name === '$URL') {
    return 'dara.URL';
  }

  if (name === '$File') {
    return 'dara.File';
  }

  if (name === '$Response') {
    return 'dara.Response';
  }

  if (name === '$Request') {
    return 'dara.Request';
  }

  if (name === 'writeable') {
    return 'io.Writer';
  }

  if (name === '$Model') {
    return 'dara.Model';
  }

  return name;
}

function _field(name, type, required = false) {
  return {
    'attrs': [],
    'fieldName': {
      'lexeme': name,
      'tag': Tag.ID,
    },
    'fieldValue': {
      'fieldType': type,
      'type': 'fieldType'
    },
    'tokenRange': [0, 0],
    'required': required,
    'type': 'modelField'
  };
}

function _type(name) {
  if (name === 'object') {
    return 'map[string]interface{}';
  }

  if (name === 'integer' || name === 'number') {
    return 'int';
  }

  if (name === 'readable') {
    return 'io.Reader';
  }

  if (name === '$Error') {
    return 'dara.BaseError';
  }

  if (name === '$ResponseError') {
    return 'dara.ResponseError';
  }

  if (name === '$SSEEvent') {
    return '*dara.SSEEvent';
  }

  if (name === '$RetryOptions') {
    return '*dara.RetryOptions';
  }

  if (name === '$RuntimeOptions') {
    return '*dara.RuntimeOptions';
  }

  if (name === '$ExtendsParameters') {
    return '*dara.ExtendsParameters';
  }

  if (name === '$Date') {
    return '*dara.Date';
  }

  if (name === '$URL') {
    return '*dara.URL';
  }

  if (name === '$File') {
    return '*dara.File';
  }

  if (name === 'bytes') {
    return '[]byte';
  }

  if (name === 'int64') {
    return 'int64';
  }

  if (name === 'uint64') {
    return 'uint64';
  }

  if (name === 'int32') {
    return 'int32';
  }

  if (name === 'uint32') {
    return 'uint32';
  }

  if (name === 'int16') {
    return 'int16';
  }

  if (name === 'uint16') {
    return 'uint16';
  }

  if (name === 'int8') {
    return 'int8';
  }

  if (name === 'uint8') {
    return 'uint8';
  }

  if (name === '$Response') {
    return '*dara.Response';
  }

  if (name === '$Request') {
    return '*dara.Request';
  }

  if (name === 'writeable') {
    return 'io.Writer';
  }

  if (name === 'double') {
    return 'float64';
  }

  if (name === 'long' || name === 'int64') {
    return 'int64';
  }

  if (name === 'float') {
    return 'float32';
  }

  if (name === 'boolean') {
    return 'bool';
  }

  if (name === 'any' || name === 'class' || name === '$Model') {
    return 'interface{}';
  }

  if (name === 'string') {
    return 'string';
  }
  name = '*' + name;
  return name;
}

function _pointerType(name) {
  if (name === 'object') {
    return 'map[string]interface{}';
  }

  if (name === 'integer' || name === 'number') {
    return '*int';
  }

  if (name === 'readable') {
    return 'io.Reader';
  }

  if (name === 'bytes') {
    return '[]byte';
  }

  if (name === 'uint64') {
    return '*uint64';
  }

  if (name === 'int32') {
    return '*int32';
  }

  if (name === 'uint32') {
    return '*uint32';
  }

  if (name === '$Response') {
    return '*dara.Response';
  }

  if (name === '$Request') {
    return '*dara.Request';
  }

  if (name === '$RetryOptions') {
    return '*dara.RetryOptions';
  }

  if (name === '$RuntimeOptions') {
    return '*dara.RuntimeOptions';
  }

  if (name === '$ExtendsParameters') {
    return '*dara.ExtendsParameters';
  }

  if (name === 'writeable') {
    return 'io.Writer';
  }

  if (name === 'double') {
    return '*float64';
  }

  if (name === 'long' || name === 'int64') {
    return '*int64';
  }

  if (name === 'float') {
    return '*float32';
  }

  if (name === 'boolean') {
    return '*bool';
  }

  if (name === 'any' || name === 'class') {
    return 'interface{}';
  }

  if (name === '$Model') {
    return 'dara.Model';
  }

  if (name === '$Error') {
    return '*dara.SDKError';
  }

  if (name === 'string') {
    return '*string';
  }
  name = '*' + name;
  return name;
}

function _avoidReserveName(name) {
  const reserves = [
    'function'
  ];
  if (reserves.indexOf(name) !== -1) {
    return `_${name}`;
  }

  return name;
}

function _setValueFunc(name) {
  var expr = '';
  if (name === 'number' || name === 'integer' || name === 'int') {
    expr = `dara.IntValue(`;
  } else if (name === 'long' || name === 'int64') {
    expr = `dara.Int64Value(`;
  } else if (name === 'double') {
    expr = `dara.Float64Value(`;
  } else if (name === 'float') {
    expr = `dara.Float32Value(`;
  } else if (name === '[]float64') {
    expr = `dara.Float64SliceValue(`;
  } else if (name === '[]float32') {
    expr = `dara.Float32SliceValue(`;
  } else if (name === 'boolean') {
    expr = `dara.BoolValue(`;
  } else if (name === '[]bool') {
    expr = `dara.BoolSliceValue(`;
  } else if (name === 'string') {
    expr = `dara.StringValue(`;
  } else if (name === 'int32') {
    expr = `dara.Int32Value(`;
  } else if (name === '[]string') {
    expr = `dara.StringSliceValue(`;
  } else if (name === '[]int') {
    expr = `dara.IntSliceValue(`;
  } else if (name === '[]int32') {
    expr = `dara.Int32SliceValue(`;
  } else if (name === '[]int16') {
    expr = `dara.Int16SliceValue(`;
  } else if (name === '[]int64') {
    expr = `dara.Int64SliceValue(`;
  } else if (name === 'uint64') {
    expr = `dara.Uint64Value(`;
  } else if (name === 'uint32') {
    expr = `dara.Uint32Value(`;
  } else if (name === 'uint16') {
    expr = `dara.Uint16Value(`;
  } else if (name === '[]uint64') {
    expr = `dara.Uint64SliceValue(`;
  } else if (name === '[]uint32') {
    expr = `dara.Uint32SliceValue(`;
  } else if (name === '[]uint16') {
    expr = `dara.Uint16SliceValue(`;
  }

  return expr;
}

function _isFilterType(fieldType) {
  return filterTypes.indexOf(fieldType) !== -1;
}

function _getAttr(node, attrName) {
  for (let i = 0; i < node.attrs.length; i++) {
    if (_name(node.attrs[i].attrName) === attrName) {
      return node.attrs[i].attrValue.string;
    }
  }
}

function _escape(str) {
  return str.replace(/>(Note:|Notice:|Warning:|Danger:)/g, '\t$1')
    .replace(/(\*) +/g, '\t- ');
}

function _isBinaryOp(type){
  const op = [
    'or', 'eq', 'neq',
    'gt', 'gte', 'lt',
    'lte', 'add', 'subtract',
    'div', 'multi', 'and'
  ];
  return op.includes(type);
}

module.exports = {
  _name, _string, _type, _format, _initValue, _avoidReserveName, _importFilter, _upperFirst, _snakeCase,
  _setExtendFunc, _isFilterType, _getAttr, _setValueFunc, _vid, _pointerType, _lowerFirst, _escape, _deleteWithSuffix,
  _avoidVariableKeywords, _isBinaryOp, _setArrayFunc, _isIterator, _modelName, _extendFieldName, _field
};