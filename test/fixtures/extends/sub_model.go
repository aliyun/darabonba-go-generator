// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iSub interface {
  iBase
  String() string
  GoString() string
  SetName(v string) *Sub
  GetName() *string 
  SetCode(v string) *Sub
  GetCode() *string 
}

type Sub struct {
  Age *int `json:"age,omitempty" xml:"age,omitempty" require:"true"`
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s Sub) String() string {
  return dara.Prettify(s)
}

func (s Sub) GoString() string {
  return s.String()
}

func (s *Sub) GetAge() *int  {
  return s.Age
}

func (s *Sub) GetName() *string  {
  return s.Name
}

func (s *Sub) GetCode() *string  {
  return s.Code
}

func (s *Sub) SetAge(v int) *Sub {
  s.Age = &v
  return s
}

func (s *Sub) SetName(v string) *Sub {
  s.Name = &v
  return s
}

func (s *Sub) SetCode(v string) *Sub {
  s.Code = &v
  return s
}

func (s *Sub) Validate() error {
  if err := dara.ValidateRequired(s.Age, "Age"); err != nil {
    return err
  }
  if err := dara.ValidateRequired(s.Name, "Name"); err != nil {
    return err
  }
  if err := dara.ValidateRequired(s.Code, "Code"); err != nil {
    return err
  }
  return nil
}

