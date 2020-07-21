// This file is auto-generated, don't edit it. Thanks.
/**
 top annotation
*/
package client

import (
  "github.com/alibabacloud-go/tea/tea"
)

/**
  TestModel
*/
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

/**
  Init Func
*/
func NewClient()(*Client, error) {
  client := new(Client)
  err := client.Init()
  return client, err
}

func (client *Client)Init()(_err error) {
  return nil
}


/**
  testAPI
*/
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


/**
  testFunc
*/
func TestFunc () (_err error) {
  return _err
}

