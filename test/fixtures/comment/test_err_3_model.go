// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTestErr3 interface {
  dara.Model
  String() string
  GoString() string
}

// Description:
// 
// TestErr3
type TestErr3 struct {
  // empty comment1
  // empy comment2
}

func (s TestErr3) String() string {
  return dara.Prettify(s)
}

func (s TestErr3) GoString() string {
  return s.String()
}

func (s *TestErr3) Validate() error {
  return dara.Validate(s)
}

