// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)


// Quote and backslash escape cases for Go string literals.
func Main (args []*string) (_err error) {
  s := "a\"b"
  t := "{\"enableHistoryServer\":false}"
  u := "a\\b"
  v := "a\\\"b"
  return _err
}

func Build () (_result *M) {
  m := &M{
    Config: dara.String("{\"enableHistoryServer\":false,\"disableHistoryServerAuth\":false,\"useVPCEndpoint\":false}"),
  }
  _result = m
  return _result
}

