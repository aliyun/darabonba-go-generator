package client

import (
  dara "github.com/alibabacloud-go/tea/tea"
  
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
  dara.Model
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

type Client struct {
  A  *string
}

// Description:
// 
// Init Func
func NewClient()(*Client, error) {
  client := new(Client)
  err := client.Init()
  return client, err
}

func (client *Client)Init()(_err error) {
  return nil
}


// Description:
// 
// testAPI
func (client *Client) TestAPI() (_err error) {
  _runtime := dara.NewRuntimeObject(map[string]interface{}{})

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

    return _err
  }
  return _result, _resultErr
}


// Description:
// 
// testFunc
func TestFunc () (_err error) {
  return _err
}

// Deprecated: annotation test deprecated
// 
// Summary:
// 
// annotation test summary
// 
// Description:
// 
// annotation test description
// 
// 	- description1 test for typescript
// 
// 	- description2 test for typescript
// 
// 	- test link: [Limits](https://help.aliyun.com/document_detail/25412.html#SecurityGroupQuota).
// 
// @param test - string param1
// 
// @param _test - string param2
// 
// @return void
// 
// @throws InternalError Server error. 500 服务器端出现未知异常。
// 
// @throws StackNotFound The Stack (%(stack_name)s) could not be found.  404 资源栈不存在。
func TestFuncWithAnnotation1 (test *string, _test *string) (_err error) {
  // empty comment1
  // empty comment2
  return _err
}

// Deprecated: test is deprecated, use xxx instead.
//
// deprecated description1
//
// deprecated description2
// 
// Summary:
// 
// annotation test summary
// 
// summary description1
// 
// summary description2
// 
// @param test - string param1
// 
// @param _test - string param2
// 
// @return void
// 
// @throws InternalError Server error. 500 服务器端出现未知异常。
// 
func TestFuncWithAnnotation2 (test *string, _test *string) (_err error) {
  // empty comment1
  // empty comment2
  return _err
}

// Deprecated: deprecated test for line break.
// 
// @param test - string param1
// 
// param test for line break.
// 
// @param _test - string param2
// 
// @return void
// 
// return test for line break.
// 
// @throws InternalError Server error. 500 服务器端出现未知异常。
// 
// throws test for line break.
// 
func LineBreakAnnotation (test *string, _test *string) (_err error) {
  return _err
}

