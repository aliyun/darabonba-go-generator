// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/tea"
)


func Hello () {
  return
}

func HelloMap () (_result map[string]*string) {
  m := make(map[string]*string)
  tea.Convert(tea.Merge(map[string]*string{
    "key": tea.String("value"),
    "key-1": tea.String("value-1"),
    },m), &_result)
  return _result
}

func HelloArrayMap () (_result []map[string]*string) {
  tea.Convert([]map[string]*string{map[string]*string{
      "key": tea.String("value"),
    }}, &_result)
  return _result
}

func HelloParams (a *string, b *string) (_err error) {
  return _err
}

func HelloTest (a *string) (_result *string, _err error) {
  _result = a
  return _result , _err
}

func EqualString (a *string, b *string) (_result *bool, _err error) {
  _result = tea.Bool(true)
  return _result, _err
}

func HelloTestNestReturn (a *string, b *string) (_result *bool, _err error) {
  helloTestTmp, err := HelloTest(a)
  if err != nil {
    _err = err
    return _result, _err
  }
  helloTestTmp1, err := HelloTest(b)
  if err != nil {
    _err = err
    return _result, _err
  }
  _body, _err := EqualString(helloTestTmp, helloTestTmp1)
  if _err != nil {
    return _result, _err
  }
  _result = _body
  return _result, _err
}

func HelloTestNestDeclar (a *string, b *string) (_result *bool, _err error) {
  helloTestTmp, err := HelloTest(a)
  if err != nil {
    _err = err
    return _result, _err
  }
  helloTestTmp1, err := HelloTest(b)
  if err != nil {
    _err = err
    return _result, _err
  }
  tmp, _err := EqualString(helloTestTmp, helloTestTmp1)
  if _err != nil {
    return _result, _err
  }

  _result = tmp
  return _result , _err
}

func HelloTestNestIf (a *string, b *string) (_result *bool, _err error) {
  helloTestTmp, err := HelloTest(a)
  if err != nil {
    _err = err
    return _result, _err
  }
  helloTestTmp1, err := HelloTest(b)
  if err != nil {
    _err = err
    return _result, _err
  }
  if tea.BoolValue(EqualString(helloTestTmp, helloTestTmp1)) {
    _result = tea.Bool(true)
    return _result, _err
  }

  _result = tea.Bool(false)
  return _result, _err
}

func HelloTestNestFor (a *string, b *string) (_result *bool, _err error) {
  helloTestTmp, err := HelloTest(a)
  if err != nil {
    _err = err
    return _result, _err
  }
  helloTestTmp1, err := HelloTest(b)
  if err != nil {
    _err = err
    return _result, _err
  }
  for tea.BoolValue(EqualString(helloTestTmp, helloTestTmp1)) {
    _result = tea.Bool(true)
    return _result, _err
  }
  _result = tea.Bool(false)
  return _result, _err
}

func HelloTestNestFor1 (a *string, b *string) (_result *bool, _err error) {
  helloTestTmp, err := HelloTest(a)
  if err != nil {
    _err = err
    return _result, _err
  }
  helloTestTmp1, err := HelloTest(b)
  if err != nil {
    _err = err
    return _result, _err
  }
  for tea.BoolValue(EqualString(helloTestTmp, helloTestTmp1)) {
    _result = tea.Bool(true)
    return _result, _err
  }
  _result = tea.Bool(false)
  return _result, _err
}

