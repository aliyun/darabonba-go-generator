// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source "github.com/aliyun/darabonba-go-generator/test"
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
  return dara.Validate(s)
}

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
  return dara.Validate(s)
}

type iTest3 interface {
  dara.Model
  String() string
  GoString() string
}

// Description:
// 
// TestModel3
type Test3 struct {
  // empty comment1
  // empy comment2
}

func (s Test3) String() string {
  return dara.Prettify(s)
}

func (s Test3) GoString() string {
  return s.String()
}

func (s *Test3) Validate() error {
  return dara.Validate(s)
}

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
  return dara.Validate(s)
}

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

type iTestErr3 interface {
  dara.Model
  String() string
  GoString() string
}

// Description:
// 
// TestErr3
type TestErr3 struct {
  // empty comment1
  // empy comment2
}

func (s TestErr3) String() string {
  return dara.Prettify(s)
}

func (s TestErr3) GoString() string {
  return s.String()
}

func (s *TestErr3) Validate() error {
  return dara.Validate(s)
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
  str := "sss"
  // new model instance comment
  modelInstance := &Test1{
    Test: dara.String("test"),
    //test declare back comment
    Test2: dara.String("test2"),
  }
  array := []interface{}{
    // array string comment
    dara.String("string"), 
    // array number comment
    dara.Int(300)
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
  _runtime := dara.NewRuntimeObject(map[string]interface{}{
    // empty runtime comment
    // another runtime comment
  })

  var retryPolicyContext *dara.RetryPolicyContext
  var request_ *dara.Request
  var response_ *dara.Response
  var _resultErr error
  retriesAttempted := int(0)
  retryPolicyContext = &dara.RetryPolicyContext{
    RetriesAttempted: retriesAttempted,
  }

  for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
    _resultErr = nil
    _backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
    dara.Sleep(_backoffDelayTime)

    request_ = dara.NewRequest()
    // new model instance comment
    modelInstance := &Test1{
      // test declare front comment
      Test: dara.String("test"),
      // test2 declare front comment
      Test2: dara.String("test2"),
    }
    // number declare comment
    num := 123
    // static function call comment
    StaticFunc()
    response_, _err := dara.DoRequest(request_, _runtime)
    if _err != nil {
      retriesAttempted++
      retryPolicyContext = &dara.RetryPolicyContext{
        RetriesAttempted: retriesAttempted,
        HttpRequest:      request_,
        HttpResponse:     response_,
        Exception:        _err,
      }
      _resultErr = _err
      continue
    }

    // static async function call
    _err = TestFunc()
    if _err != nil {
      retriesAttempted++
      retryPolicyContext = &dara.RetryPolicyContext{
        RetriesAttempted: retriesAttempted,
        HttpRequest:      request_,
        HttpResponse:     response_,
        Exception:        _err,
      }
      _resultErr = _err
      continue
    }

    // return comment
    return _err
  }
  return _result, _resultErr
}

// testAPI2 comment
func (client *Client) TestAPI2() (_err error) {
  _runtime := dara.NewRuntimeObject(map[string]interface{}{
    // runtime retry comment
    "retry": true,
    // runtime back comment one
    // runtime back comment two
  })

  var retryPolicyContext *dara.RetryPolicyContext
  var request_ *dara.Request
  var response_ *dara.Response
  var _resultErr error
  retriesAttempted := int(0)
  retryPolicyContext = &dara.RetryPolicyContext{
    RetriesAttempted: retriesAttempted,
  }

  for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
    _resultErr = nil
    _backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
    dara.Sleep(_backoffDelayTime)

    request_ = dara.NewRequest()
    // new model instance comment
    modelInstance := &Test3{
      //empty model 
    }
    // boolean declare comment
    bool := true
    if bool {
      //empty if
    } else {
      //empty else
    }

    // api function call comment
    _err = client.TestAPI()
    if _err != nil {
      retriesAttempted++
      retryPolicyContext = &dara.RetryPolicyContext{
        RetriesAttempted: retriesAttempted,
        HttpRequest:      request_,
        HttpResponse:     response_,
        Exception:        _err,
      }
      _resultErr = _err
      continue
    }

    // back comment
    response_, _err := dara.DoRequest(request_, _runtime)
    if _err != nil {
      retriesAttempted++
      retryPolicyContext = &dara.RetryPolicyContext{
        RetriesAttempted: retriesAttempted,
        HttpRequest:      request_,
        HttpResponse:     response_,
        Exception:        _err,
      }
      _resultErr = _err
      continue
    }

    // empty return comment
  }
  return _result, _resultErr
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

