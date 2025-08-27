// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

// top comment
// Description:
// 
// top annotation
type iTest1 interface {
  dara.Model
  String() string
  GoString() string
  SetTest(v string) *Test1
  GetTest() *string 
  SetTest2(v string) *Test1
  GetTest2() *string 
}

// Description:
// 
// TestModel
type Test1 struct {
  // test desc
  Test *string `json:"test,omitempty" xml:"test,omitempty" require:"true"`
  //model的test back comment
  // test2 desc
  Test2 *string `json:"test2,omitempty" xml:"test2,omitempty" require:"true"`
  //model的test2 back comment
}

func (s Test1) String() string {
  return dara.Prettify(s)
}

func (s Test1) GoString() string {
  return s.String()
}

func (s *Test1) GetTest() *string  {
  return s.Test
}

func (s *Test1) GetTest2() *string  {
  return s.Test2
}

func (s *Test1) SetTest(v string) *Test1 {
  s.Test = &v
  return s
}

func (s *Test1) SetTest2(v string) *Test1 {
  s.Test2 = &v
  return s
}

func (s *Test1) Validate() error {
  if err := dara.ValidateRequired(s.Test, "Test"); err != nil {
    return err
  }
  if err := dara.ValidateRequired(s.Test2, "Test2"); err != nil {
    return err
  }
  return nil
}

