// This file is auto-generated, don't edit it. Thanks.
package client

import (
  user "darabonba.com/multi/model/user"
  "github.com/alibabacloud-go/tea/dara"
)

type iTestModelDIR interface {
  dara.Model
  String() string
  GoString() string
  SetTest(v string) *TestModelDIR
  GetTest() *string 
  SetA(v interface{}) *TestModelDIR
  GetA() interface{} 
}

type TestModelDIR struct {
  Test *string `json:"test,omitempty" xml:"test,omitempty" require:"true"`
  A interface{} `json:"a,omitempty" xml:"a,omitempty" require:"true"`
}

func (s TestModelDIR) String() string {
  return dara.Prettify(s)
}

func (s TestModelDIR) GoString() string {
  return s.String()
}

func (s *TestModelDIR) GetTest() *string  {
  return s.Test
}

func (s *TestModelDIR) GetA() interface{}  {
  return s.A
}

func (s *TestModelDIR) SetTest(v string) *TestModelDIR {
  s.Test = &v
  return s
}

func (s *TestModelDIR) SetA(v interface{}) *TestModelDIR {
  s.A = v
  return s
}

func (s *TestModelDIR) Validate() error {
  if err := dara.ValidateRequired(s.Test, "Test"); err != nil {
    return err
  }
  if err := dara.ValidateRequired(s.A, "A"); err != nil {
    return err
  }
  return nil
}

type iTestModelDir interface {
  dara.Model
  String() string
  GoString() string
  SetTest(v int) *TestModelDir
  GetTest() *int 
  SetM(v *user.Info) *TestModelDir
  GetM() *user.Info 
}

type TestModelDir struct {
  Test *int `json:"test,omitempty" xml:"test,omitempty" require:"true"`
  M *user.Info `json:"m,omitempty" xml:"m,omitempty" require:"true"`
}

func (s TestModelDir) String() string {
  return dara.Prettify(s)
}

func (s TestModelDir) GoString() string {
  return s.String()
}

func (s *TestModelDir) GetTest() *int  {
  return s.Test
}

func (s *TestModelDir) GetM() *user.Info  {
  return s.M
}

func (s *TestModelDir) SetTest(v int) *TestModelDir {
  s.Test = &v
  return s
}

func (s *TestModelDir) SetM(v *user.Info) *TestModelDir {
  s.M = v
  return s
}

func (s *TestModelDir) Validate() error {
  if err := dara.ValidateRequired(s.Test, "Test"); err != nil {
    return err
  }
  if err := dara.ValidateRequired(s.M, "M"); err != nil {
    return err
  }
  if s.M != nil {
    if err := s.M.Validate(); err != nil {
      return err
    }
  }
  return nil
}

