// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTest2 interface {
  dara.Model
  String() string
  GoString() string
  SetTest(v string) *Test2
  GetTest() *string 
  SetTest2(v string) *Test2
  GetTest2() *string 
}

// Description:
// 
// TestModel2
type Test2 struct {
  // model的test front comment
  // test desc
  Test *string `json:"test,omitempty" xml:"test,omitempty" require:"true"`
  // model的test front comment
  // test2 desc
  Test2 *string `json:"test2,omitempty" xml:"test2,omitempty" require:"true"`
}

func (s Test2) String() string {
  return dara.Prettify(s)
}

func (s Test2) GoString() string {
  return s.String()
}

func (s *Test2) GetTest() *string  {
  return s.Test
}

func (s *Test2) GetTest2() *string  {
  return s.Test2
}

func (s *Test2) SetTest(v string) *Test2 {
  s.Test = &v
  return s
}

func (s *Test2) SetTest2(v string) *Test2 {
  s.Test2 = &v
  return s
}

func (s *Test2) Validate() error {
  if err := dara.ValidateRequired(s.Test, "Test"); err != nil {
    return err
  }
  if err := dara.ValidateRequired(s.Test2, "Test2"); err != nil {
    return err
  }
  return nil
}

