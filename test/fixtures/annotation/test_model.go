package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

// Description:
// 
// top annotation
type iTest interface {
  dara.Model
  String() string
  GoString() string
  SetTest(v string) *Test
  GetTest() *string 
}

// Description:
// 
// TestModel
type Test struct {
  // Alichange app id 
  Test *string `json:"test,omitempty" xml:"test,omitempty" require:"true"`
}

func (s Test) String() string {
  return dara.Prettify(s)
}

func (s Test) GoString() string {
  return s.String()
}

func (s *Test) GetTest() *string  {
  return s.Test
}

func (s *Test) SetTest(v string) *Test {
  s.Test = &v
  return s
}

func (s *Test) Validate() error {
  return dara.Validate(s)
}

