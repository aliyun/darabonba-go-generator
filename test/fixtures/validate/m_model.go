// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iM interface {
  dara.Model
  String() string
  GoString() string
  SetName(v string) *M
  GetName() *string 
  SetTest(v *MTest) *M
  GetTest() *MTest 
}

type M struct {
  Name *string `json:"name,omitempty" xml:"name,omitempty" maxLength:"500"`
  Test *MTest `json:"test,omitempty" xml:"test,omitempty" require:"true" type:"Struct"`
}

func (s M) String() string {
  return dara.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) GetName() *string  {
  return s.Name
}

func (s *M) GetTest() *MTest  {
  return s.Test
}

func (s *M) SetName(v string) *M {
  s.Name = &v
  return s
}

func (s *M) SetTest(v *MTest) *M {
  s.Test = v
  return s
}

func (s *M) Validate() error {
  if s.Name != nil {
    if err := dara.ValidateMaxLength(s.Name, 500, "Name"); err != nil {
      return err
    }
  }
  if err := dara.ValidateRequired(s.Test, "Test"); err != nil {
    return err
  }
  if s.Test != nil {
    if err := s.Test.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type MTest struct {
  Name *string `json:"name,omitempty" xml:"name,omitempty" maxLength:"500"`
}

func (s MTest) String() string {
  return dara.Prettify(s)
}

func (s MTest) GoString() string {
  return s.String()
}

func (s *MTest) GetName() *string  {
  return s.Name
}

func (s *MTest) SetName(v string) *MTest {
  s.Name = &v
  return s
}

func (s *MTest) Validate() error {
  if s.Name != nil {
    if err := dara.ValidateMaxLength(s.Name, 500, "Name"); err != nil {
      return err
    }
  }
  return nil
}

