// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iBase interface {
  dara.Model
  String() string
  GoString() string
  SetName(v string) *Base
  GetName() *string 
  SetAge(v int) *Base
  GetAge() *int 
}

type Base struct {
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  Age *int `json:"age,omitempty" xml:"age,omitempty" require:"true"`
}

func (s Base) String() string {
  return dara.Prettify(s)
}

func (s Base) GoString() string {
  return s.String()
}

func (s *Base) GetName() *string  {
  return s.Name
}

func (s *Base) GetAge() *int  {
  return s.Age
}

func (s *Base) SetName(v string) *Base {
  s.Name = &v
  return s
}

func (s *Base) SetAge(v int) *Base {
  s.Age = &v
  return s
}

func (s *Base) Validate() error {
  return dara.Validate(s)
}

