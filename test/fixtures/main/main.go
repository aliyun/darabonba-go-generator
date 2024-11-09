// This file is auto-generated, don't edit it. Thanks.
package main

import (
  "os"
  "github.com/alibabacloud-go/tea/dara"
  
)


func _main (args []*string) (_err error) {
  _err = ThrowError()
  if _err != nil {
    return _err
  }
  _err = dara.NewSDKError(map[string]interface{}{
    "code": "error",
  })
  return _err
  tmpTmp, _err := ThrowError0()
  tmp := dara.StringValue(tmpTmp)
  if _err != nil {
    return _err
  }

  tmpTmp, _err = ThrowError1()
  tmp = dara.StringValue(tmpTmp)
  if _err != nil {
    return _err
  }

  return _err
}

func ThrowError () (_err error) {
  panic("No Support!")
}

func ThrowError0 () (_result *string, _err error) {
  panic("No Support!")
}

func ThrowError1 () (_result *string, _err error) {
  panic("No Support!")
}


func main() {
  err := _main(dara.StringSlice(os.Args[1:]))
  if err != nil {
    panic(err)
  }
}
