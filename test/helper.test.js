'use strict';

const expect = require('expect.js');

const {
  _name, _string, _type, _format, _initValue, _avoidReserveName, 
  _setExtendFunc, _pointerType
} = require('../lib//helper');

describe('helper', function () {
  it('_name should ok', function () {
    var name =  _name({lexeme: 'helper'});
    expect(name).to.equal('helper');

    name =  _name({name: 'name'});
    expect(name).to.equal('name');
  });

  it('_string should ok', function () {
    var name =  _string({string: 'helper'});
    expect(name).to.equal('helper');
  });

  it('_type should ok', function () {
    var type =  _type('object');
    expect(type).to.equal('map[string]interface{}');

    type =  _type('integer');
    expect(type).to.equal('int');

    type =  _type('number');
    expect(type).to.equal('int');

    type =  _type('readable');
    expect(type).to.equal('io.Reader');

    type =  _type('bytes');
    expect(type).to.equal('[]byte');

    type =  _type('uint64');
    expect(type).to.equal('uint64');

    type =  _type('int32');
    expect(type).to.equal('int32');

    type =  _type('int16');
    expect(type).to.equal('int16');

    type =  _type('uint16');
    expect(type).to.equal('uint16');

    type =  _type('uint32');
    expect(type).to.equal('uint32');

    type =  _type('$Response');
    expect(type).to.equal('*tea.Response');

    type =  _type('$Request');
    expect(type).to.equal('*tea.Request');

    type =  _type('writeable');
    expect(type).to.equal('io.Writer');

    type =  _type('double');
    expect(type).to.equal('float64');

    type =  _type('long');
    expect(type).to.equal('int64');

    type =  _type('int64');
    expect(type).to.equal('int64');

    type =  _type('float');
    expect(type).to.equal('float32');

    type =  _type('boolean');
    expect(type).to.equal('bool');

    type =  _type('any');
    expect(type).to.equal('interface{}');

    type =  _type('string');
    expect(type).to.equal('string');

    type =  _type('struct');
    expect(type).to.equal('*struct');
  });

  it('_format should ok', function () {
    var name =  _format('a-b');
    expect(name).to.equal('AB');
  });

  it('_initValue should ok', function () {
    var val =  _initValue('number');
    expect(val).to.equal('tea.Int(0)');

    val =  _initValue('integer');
    expect(val).to.equal('tea.Int(0)');

    val =  _initValue('string');
    expect(val).to.equal('tea.String("")');

    val =  _initValue('boolean');
    expect(val).to.equal('tea.Bool(false)');

    val =  _initValue('bytes');
    expect(val).to.equal('make([]byte, 0)');

    val =  _initValue('any');
    expect(val).to.equal('interface{}(nil)');

    val =  _initValue('float');
    expect(val).to.equal('tea.Float32(0.00)');

    val =  _initValue('null');
    expect(val).to.equal('nil');

    val =  _initValue('$Response');
    expect(val).to.equal('&tea.Response{}');

    val =  _initValue('$Request');
    expect(val).to.equal('&tea.Request{}');

    val =  _initValue('object');
    expect(val).to.equal('make(map[string]interface{})');

    val =  _initValue('map[string]string');
    expect(val).to.equal('make(map[string]string)');

    val =  _initValue('[]byte');
    expect(val).to.equal('make([]byte, 1)');

    val =  _initValue('struct');
    expect(val).to.equal('&struct{}');
  });

  it('_setExtendFunc should ok', function () {
    var val =  _setExtendFunc('number');
    expect(val).to.equal('tea.Int(');

    val =  _setExtendFunc('integer');
    expect(val).to.equal('tea.Int(');

    val =  _setExtendFunc('int');
    expect(val).to.equal('tea.Int(');

    val =  _setExtendFunc('long');
    expect(val).to.equal('tea.Int64(');

    val =  _setExtendFunc('int64');
    expect(val).to.equal('tea.Int64(');

    val =  _setExtendFunc('double');
    expect(val).to.equal('tea.Float64(');

    val =  _setExtendFunc('float');
    expect(val).to.equal('tea.Float32(');

    val =  _setExtendFunc('[]float64');
    expect(val).to.equal('tea.Float64Slice(');

    val =  _setExtendFunc('[]float32');
    expect(val).to.equal('tea.Float32Slice(');

    val =  _setExtendFunc('boolean');
    expect(val).to.equal('tea.Bool(');

    val =  _setExtendFunc('[]bool');
    expect(val).to.equal('tea.BoolSlice(');

    val =  _setExtendFunc('string');
    expect(val).to.equal('tea.String(');

    val =  _setExtendFunc('int32');
    expect(val).to.equal('tea.Int32(');
    
    val =  _setExtendFunc('[]string');
    expect(val).to.equal('tea.StringSlice(');

    val =  _setExtendFunc('[]int');
    expect(val).to.equal('tea.IntSlice(');

    val =  _setExtendFunc('[]int32');
    expect(val).to.equal('tea.Int32Slice(');

    val =  _setExtendFunc('[]int64');
    expect(val).to.equal('tea.Int64Slice(');

    val =  _setExtendFunc('[]uint');
    expect(val).to.equal('tea.UintSlice(');

    val =  _setExtendFunc('[]uint32');
    expect(val).to.equal('tea.Uint32Slice(');

    val =  _setExtendFunc('[]uint64');
    expect(val).to.equal('tea.Uint64Slice(');
    
    val =  _setExtendFunc('uint');
    expect(val).to.equal('tea.Uint(');

    val =  _setExtendFunc('uint32');
    expect(val).to.equal('tea.Uint32(');

    val =  _setExtendFunc('uint64');
    expect(val).to.equal('tea.Uint64(');

    val =  _setExtendFunc('struct');
    expect(val).to.equal('');
  });

  it('_pointerType should ok', function () {
    var val =  _pointerType('object');
    expect(val).to.equal('map[string]interface{}');

    val =  _pointerType('integer');
    expect(val).to.equal('*int');

    val =  _pointerType('number');
    expect(val).to.equal('*int');

    val =  _pointerType('readable');
    expect(val).to.equal('io.Reader');

    val =  _pointerType('bytes');
    expect(val).to.equal('[]byte');

    val =  _pointerType('uint64');
    expect(val).to.equal('*uint64');

    val =  _pointerType('int32');
    expect(val).to.equal('*int32');

    val =  _pointerType('uint32');
    expect(val).to.equal('*uint32');

    val =  _pointerType('$Response');
    expect(val).to.equal('*tea.Response');

    val =  _pointerType('$Request');
    expect(val).to.equal('*tea.Request');

    val =  _pointerType('writeable');
    expect(val).to.equal('io.Writer');

    val =  _pointerType('double');
    expect(val).to.equal('*float64');

    val =  _pointerType('long');
    expect(val).to.equal('*int64');
    
    val =  _pointerType('int64');
    expect(val).to.equal('*int64');

    val =  _pointerType('float');
    expect(val).to.equal('*float32');

    val =  _pointerType('boolean');
    expect(val).to.equal('*bool');

    val =  _pointerType('any');
    expect(val).to.equal('interface{}');

    val =  _pointerType('string');
    expect(val).to.equal('*string');

    val =  _pointerType('struct');
    expect(val).to.equal('*struct');
  });

  it('_avoidReserveName should ok', function () {
    var val =  _avoidReserveName('function');
    expect(val).to.equal('_function');
  });
});