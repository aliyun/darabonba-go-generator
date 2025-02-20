// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
  
)


func Hello () {
  return
}

func HelloMap () (_result map[string]*string) {
  m := make(map[string]*string)
  dara.Convert(dara.Merge(map[string]*string{
    "key": dara.String("value"),
    "key-1": dara.String("value-1"),
  }, m), &_result)

  return _result
}

func HelloArrayMap () (_result []map[string]*string) {
  _result = []map[string]*string{map[string]*string{
      "key": dara.String("value"),
    }}
  return _result
}

func HelloParams (a *string, b *string) (_err error) {
  return
}

func HelloTest (a *string) (_result *string, _err error) {
  _result = a
  return _result , _err
}

func EqualString (a *string, b *string) (_result *bool, _err error) {
  _result = dara.Bool(true)
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
  tmpTmp, _err := EqualString(helloTestTmp, helloTestTmp1)
  tmp := dara.BoolValue(tmpTmp)
  if _err != nil {
    return _result, _err
  }

  _result = dara.Bool(tmp)
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
  if dara.BoolValue(EqualString(helloTestTmp, helloTestTmp1)) {
    _result = dara.Bool(true)
    return _result, _err
  }

  _result = dara.Bool(false)
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
  for dara.BoolValue(EqualString(helloTestTmp, helloTestTmp1)) {
    _result = dara.Bool(true)
    return _result, _err
  }
  _result = dara.Bool(false)
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
  for dara.BoolValue(EqualString(helloTestTmp, helloTestTmp1)) {
    _result = dara.Bool(true)
    return _result, _err
  }
  _result = dara.Bool(false)
  return _result, _err
}

