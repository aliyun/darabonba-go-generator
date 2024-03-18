'use strict';

const filterTypes = ['readable', 'writeable', 'map', 'object', 'any', 'bytes', 'class', '$Error', '$Model'];

const keyWords = ['String', 'string', 'Prettify', 'prettify', 'map'];

function _name(str) {
  return str.lexeme || str.name;
}

function _importFilter(name) {
  if (keyWords.indexOf(name) !== -1) {
    name = name + '_';
  }
  return name;
}

function _isKeyWord(name) {
  return keyWords.indexOf(name) !== -1;
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
    expr = `tea.Int(`;
  } else if (name === 'long' || name === 'int64') {
    expr = `tea.Int64(`;
  } else if (name === 'double') {
    expr = `tea.Float64(`;
  } else if (name === 'float') {
    expr = `tea.Float32(`;
  } else if (name === '[]float64') {
    expr = `tea.Float64Slice(`;
  } else if (name === '[]float32') {
    expr = `tea.Float32Slice(`;
  } else if (name === 'boolean') {
    expr = `tea.Bool(`;
  } else if (name === '[]bool') {
    expr = `tea.BoolSlice(`;
  } else if (name === 'string') {
    expr = `tea.String(`;
  } else if (name === 'int16') {
    expr = `tea.Int16(`;
  } else if (name === 'int32') {
    expr = `tea.Int32(`;
  } else if (name === '[]string') {
    expr = `tea.StringSlice(`;
  } else if (name === '[]int') {
    expr = `tea.IntSlice(`;
  } else if (name === '[]int32') {
    expr = `tea.Int32Slice(`;
  } else if (name === '[]int64') {
    expr = `tea.Int64Slice(`;
  } else if (name === '[]uint') {
    expr = `tea.UintSlice(`;
  } else if (name === '[]uint32') {
    expr = `tea.Uint32Slice(`;
  } else if (name === '[]uint64') {
    expr = `tea.Uint64Slice(`;
  } else if (name === 'uint') {
    expr = `tea.Uint(`;
  } else if (name === 'uint32') {
    expr = `tea.Uint32(`;
  } else if (name === 'uint64') {
    expr = `tea.Uint64(`;
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
    return '&tea.SDKError{}';
  }

  if (type === 'float') {
    return `${_setExtendFunc(type)}0.00)`;
  }

  if (type === 'null') {
    return 'nil';
  }

  if (type === '$Response') {
    return '&tea.Response{}';
  }

  if (type === '$Request') {
    return '&tea.Request{}';
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
  return str.string;
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
    return '*tea.SDKError';
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
    return '*tea.Response';
  }

  if (name === '$Request') {
    return '*tea.Request';
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
    return '*tea.Response';
  }

  if (name === '$Request') {
    return '*tea.Request';
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

  if (name === 'any' || name === 'class' || name === '$Model') {
    return 'interface{}';
  }

  if (name === '$Error') {
    return '*tea.SDKError';
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
    expr = `tea.IntValue(`;
  } else if (name === 'long' || name === 'int64') {
    expr = `tea.Int64Value(`;
  } else if (name === 'double') {
    expr = `tea.Float64Value(`;
  } else if (name === 'float') {
    expr = `tea.Float32Value(`;
  } else if (name === '[]float64') {
    expr = `tea.Float64SliceValue(`;
  } else if (name === '[]float32') {
    expr = `tea.Float32SliceValue(`;
  } else if (name === 'boolean') {
    expr = `tea.BoolValue(`;
  } else if (name === '[]bool') {
    expr = `tea.BoolSliceValue(`;
  } else if (name === 'string') {
    expr = `tea.StringValue(`;
  } else if (name === 'int32') {
    expr = `tea.Int32Value(`;
  } else if (name === '[]string') {
    expr = `tea.StringSliceValue(`;
  } else if (name === '[]int') {
    expr = `tea.IntSliceValue(`;
  } else if (name === '[]int32') {
    expr = `tea.Int32SliceValue(`;
  } else if (name === '[]int16') {
    expr = `tea.Int16SliceValue(`;
  } else if (name === '[]int64') {
    expr = `tea.Int64SliceValue(`;
  } else if (name === 'uint64') {
    expr = `tea.Uint64Value(`;
  } else if (name === 'uint32') {
    expr = `tea.Uint32Value(`;
  } else if (name === 'uint16') {
    expr = `tea.Uint16Value(`;
  } else if (name === '[]uint64') {
    expr = `tea.Uint64SliceValue(`;
  } else if (name === '[]uint32') {
    expr = `tea.Uint32SliceValue(`;
  } else if (name === '[]uint16') {
    expr = `tea.Uint16SliceValue(`;
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

module.exports = {
  _name, _string, _type, _format, _initValue, _avoidReserveName, _importFilter, _upperFirst,
  _setExtendFunc, _isFilterType, _getAttr, _setValueFunc, _vid, _pointerType, _lowerFirst, _escape
};