// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iM interface {
  dara.Model
  String() string
  GoString() string
}

type M struct {
}

func (s M) String() string {
  return dara.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) Validate() error {
  return dara.Validate(s)
}

