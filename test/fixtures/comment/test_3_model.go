// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTest3 interface {
  dara.Model
  String() string
  GoString() string
}

// Description:
// 
// TestModel3
type Test3 struct {
  // empty comment1
  // empy comment2
}

func (s Test3) String() string {
  return dara.Prettify(s)
}

func (s Test3) GoString() string {
  return s.String()
}

func (s *Test3) Validate() error {
  return dara.Validate(s)
}

