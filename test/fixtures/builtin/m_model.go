// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iM interface {
  dara.Model
  String() string
  GoString() string
  SetA(v string) *M
  GetA() *string 
  SetB(v int) *M
  GetB() *int 
}

type M struct {
  A *string `json:"a,omitempty" xml:"a,omitempty" require:"true"`
  B *int `json:"b,omitempty" xml:"b,omitempty" require:"true"`
}

func (s M) String() string {
  return dara.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) GetA() *string  {
  return s.A
}

func (s *M) GetB() *int  {
  return s.B
}

func (s *M) SetA(v string) *M {
  s.A = &v
  return s
}

func (s *M) SetB(v int) *M {
  s.B = &v
  return s
}

func (s *M) Validate() error {
  return dara.Validate(s)
}

