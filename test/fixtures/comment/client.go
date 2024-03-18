// This file is auto-generated, don't edit it. Thanks.
// top comment
// Description:
// 
// top annotation
package client

import (
  source  "github.com/aliyun/darabonba-go-generator/test"
  "github.com/alibabacloud-go/tea/tea"
)

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
  return tea.Prettify(s)
}

func (s Test1) GoString() string {
  return s.String()
}

func (s *Test1) SetTest(v string) *Test1 {
  s.Test = &v
  return s
}

func (s *Test1) SetTest2(v string) *Test1 {
  s.Test2 = &v
  return s
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
  return tea.Prettify(s)
}

func (s Test2) GoString() string {
  return s.String()
}

func (s *Test2) SetTest(v string) *Test2 {
  s.Test = &v
  return s
}

func (s *Test2) SetTest2(v string) *Test2 {
  s.Test2 = &v
  return s
}

// Description:
// 
// TestModel3
type Test3 struct {
  // empty comment1
  // empy comment2
}

func (s Test3) String() string {
  return tea.Prettify(s)
}

func (s Test3) GoString() string {
  return s.String()
}

type Client struct {
  // type's comment
  A  []*string
}

// Description:
// 
// Init Func
// comment between init and annotation
func NewClient()(*Client, error) {
  client := new(Client)
  err := client.Init()
  return client, err
}

func (client *Client)Init()(_err error) {
  // string declate comment
  str := tea.String("sss")
  // new model instance comment
  modelInstance := &Test1{
    Test: tea.String("test"),
    //test declare back comment
    Test2: tea.String("test2"),
  }
  array := []interface{}{
    // array string comment
    tea.String("string"), 
    // array number comment
    tea.Int(300)
    // array back comment
  }
  return nil
}


// Description:
// 
// testAPI
//testAPI comment one
//testAPI comment two
func (client *Client) TestAPI() (_err error) {
  _runtime := map[string]interface{}{
    // empty runtime comment
    // another runtime comment
  }

  for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
    if _retryTimes > 0 {
      _backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
      if tea.IntValue(_backoffTime) > 0 {
        tea.Sleep(_backoffTime)
      }
    }

    _err = func() error {
      request_ := tea.NewRequest()
      // new model instance comment
      modelInstance := &Test1{
        // test declare front comment
        Test: tea.String("test"),
        // test2 declare front comment
        Test2: tea.String("test2"),
      }
      // number declare comment
      num := tea.Int(123)
      // static function call comment
      StaticFunc()
      response_, _err := tea.DoRequest(request_, _runtime)
      if _err != nil {
        return _err
      }
      // static async function call
      _err = TestFunc()
      if _err != nil {
        return _err
      }
      // return comment
      return _err
    }()
    if !tea.BoolValue(tea.Retryable(_err)) {
      break
    }
  }

  return _err
}

// testAPI2 comment
func (client *Client) TestAPI2() (_err error) {
  _runtime := map[string]interface{}{
    // runtime retry comment
    "retry": true,
    // runtime back comment one
    // runtime back comment two
  }

  for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
    if _retryTimes > 0 {
      _backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
      if tea.IntValue(_backoffTime) > 0 {
        tea.Sleep(_backoffTime)
      }
    }

    _err = func() error {
      request_ := tea.NewRequest()
      // new model instance comment
      modelInstance := &Test3{
        //empty model 
      }
      // boolean declare comment
      bool := tea.Bool(true)
      if tea.BoolValue(bool) {
        //empty if
      } else {
        //empty else
      }

      // api function call comment
      _err = client.TestAPI()
      if _err != nil {
        return _err
      }
      // back comment
      response_, _err := tea.DoRequest(request_, _runtime)
      if _err != nil {
        return _err
      }
      // empty return comment
    }()
    if !tea.BoolValue(tea.Retryable(_err)) {
      break
    }
  }

  return _err
}


func StaticFunc () {
  a := []interface{}{
    // empty annotation comment
  }
}

// Description:
// 
// testFunc
func TestFunc () (_err error) {
  // empty comment1
  // empty comment2
  return _err
}

