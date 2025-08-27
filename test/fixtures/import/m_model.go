// This file is auto-generated, don't edit it. Thanks.
package client

import (
  map_ "github.com/aliyun/darabonba-go-generator/test"
  string_ "github.com/aliyun/darabonba-go-generator/test"
  "github.com/alibabacloud-go/tea/dara"
)

type iM interface {
  dara.Model
  String() string
  GoString() string
  SetA(v *map_.Request) *M
  GetA() *map_.Request 
  SetB(v *string_.Request) *M
  GetB() *string_.Request 
}

type M struct {
  A *map_.Request `json:"a,omitempty" xml:"a,omitempty"`
  B *string_.Request `json:"b,omitempty" xml:"b,omitempty"`
}

func (s M) String() string {
  return dara.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) GetA() *map_.Request  {
  return s.A
}

func (s *M) GetB() *string_.Request  {
  return s.B
}

func (s *M) SetA(v *map_.Request) *M {
  s.A = v
  return s
}

func (s *M) SetB(v *string_.Request) *M {
  s.B = v
  return s
}

func (s *M) Validate() error {
  if s.A != nil {
    if err := s.A.Validate(); err != nil {
      return err
    }
  }
  if s.B != nil {
    if err := s.B.Validate(); err != nil {
      return err
    }
  }
  return nil
}

