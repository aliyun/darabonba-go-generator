// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iM interface {
  dara.Model
  String() string
  GoString() string
  SetConfig(v string) *M
  GetConfig() *string 
}

type M struct {
  Config *string `json:"config,omitempty" xml:"config,omitempty"`
}

func (s M) String() string {
  return dara.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) GetConfig() *string  {
  return s.Config
}

func (s *M) SetConfig(v string) *M {
  s.Config = &v
  return s
}

func (s *M) Validate() error {
  return dara.Validate(s)
}

