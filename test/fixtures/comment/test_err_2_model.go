// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iTestErr2 interface {
  dara.Model
  String() string
  GoString() string
  SetTest(v string) *TestErr2
  GetTest() *string 
  SetTest2(v string) *TestErr2
  GetTest2() *string 
}

// Description:
// 
// TestErr2
type TestErr2 struct {
  // model的test front comment
  // test desc
  Test *string `json:"test,omitempty" xml:"test,omitempty" require:"true"`
  // model的test front comment
  // test2 desc
  Test2 *string `json:"test2,omitempty" xml:"test2,omitempty" require:"true"`
}

func (s TestErr2) String() string {
  return dara.Prettify(s)
}

func (s TestErr2) GoString() string {
  return s.String()
}

func (s *TestErr2) GetTest() *string  {
  return s.Test
}

func (s *TestErr2) GetTest2() *string  {
  return s.Test2
}

func (s *TestErr2) SetTest(v string) *TestErr2 {
  s.Test = &v
  return s
}

func (s *TestErr2) SetTest2(v string) *TestErr2 {
  s.Test2 = &v
  return s
}

func (s *TestErr2) Validate() error {
  return dara.Validate(s)
}

