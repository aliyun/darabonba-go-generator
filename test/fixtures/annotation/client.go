// Description:
// 
// top annotation
package client

import (
  "github.com/alibabacloud-go/tea/tea"
)

// Description:
// 
// TestModel
type Test struct {
  // Alichange app id 
  Test *string `json:"test,omitempty" xml:"test,omitempty" require:"true"`
}

func (s Test) String() string {
  return tea.Prettify(s)
}

func (s Test) GoString() string {
  return s.String()
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
  _runtime := map[string]interface{}{}

  for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
    if _retryTimes > 0 {
      _backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
      if tea.IntValue(_backoffTime) > 0 {
        tea.Sleep(_backoffTime)
      }
    }

    _err = func() error {
      request_ := tea.NewRequest()
      response_, _err := tea.DoRequest(request_, _runtime)
      if _err != nil {
        return _err
      }
      return _err
    }()
    if !tea.BoolValue(tea.Retryable(_err)) {
      break
    }
  }

  return _err
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
func TestFuncWithAnnotation2 (test *string, _test *string) (_err error) {
  // empty comment1
  // empty comment2
  return _err
}

