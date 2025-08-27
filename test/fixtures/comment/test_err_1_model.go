// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTestErr1 interface {
  dara.Model
  String() string
  GoString() string
  SetTest(v string) *TestErr1
  GetTest() *string 
  SetTest2(v string) *TestErr1
  GetTest2() *string 
}

// Description:
// 
// TestErr
type TestErr1 struct {
  // test desc
  Test *string `json:"test,omitempty" xml:"test,omitempty" require:"true"`
  //error的test back comment
  // test2 desc
  Test2 *string `json:"test2,omitempty" xml:"test2,omitempty" require:"true"`
  //error的test2 back comment
}

func (s TestErr1) String() string {
  return dara.Prettify(s)
}

func (s TestErr1) GoString() string {
  return s.String()
}

func (s *TestErr1) GetTest() *string  {
  return s.Test
}

func (s *TestErr1) GetTest2() *string  {
  return s.Test2
}

func (s *TestErr1) SetTest(v string) *TestErr1 {
  s.Test = &v
  return s
}

func (s *TestErr1) SetTest2(v string) *TestErr1 {
  s.Test2 = &v
  return s
}

func (s *TestErr1) Validate() error {
  if err := dara.ValidateRequired(s.Test, "Test"); err != nil {
    return err
  }
  if err := dara.ValidateRequired(s.Test2, "Test2"); err != nil {
    return err
  }
  return nil
}

