// This file is auto-generated, don't edit it. Thanks.
package main

import (
  "os"
  "github.com/alibabacloud-go/tea/dara"
)


func _main (args []*string) (_err error) {
  s := "a\"b"
  config := "{\"enableHistoryServer\":false,\"disableHistoryServerAuth\":false,\"useVPCEndpoint\":false}"
  return _err
}


func main() {
  err := _main(dara.StringSlice(os.Args[1:]))
  if err != nil {
    panic(err)
  }
}
