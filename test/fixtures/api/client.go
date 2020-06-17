// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/tea"
)

type Config struct {
}

func (s Config) String() string {
  return tea.Prettify(s)
}

func (s Config) GoString() string {
  return s.String()
}

type M struct {
}

func (s M) String() string {
  return tea.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

type Client struct {
}

func NewClient(config *Config)(*Client, error) {
  client := new(Client)
  err := client.Init(config)
  return client, err
}

func (client *Client)Init(config *Config)(_err error) {
  return nil
}


func (client *Client) Hello() (_err error) {
  request_ := tea.NewRequest()
  request_.Method = tea.String("GET")
  request_.Pathname = tea.String("/")
  request_.Headers = map[string]*string{
    "host": tea.String("www.test.com"),
  }
  response_, _err := tea.DoRequest(request_, nil)
  if _err != nil {
    return _err
  }
  return _err
}

func (client *Client) HelloRuntime() (_err error) {
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
      request_.Method = tea.String("GET")
      request_.Pathname = tea.String("/")
      request_.Headers = map[string]*string{
        "host": tea.String("www.test.com"),
      }
      var test map[string]interface{}
      test = map[string]interface{}{
        "key": "value",
      }
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

func (client *Client) HelloVirtualCall(m *M) (_err error) {
  _err = tea.Validate(m)
  if _err != nil {
    return _err
  }
  request_ := tea.NewRequest()
  request_.Method = tea.String("GET")
  request_.Pathname = tea.String("/")
  request_.Headers = map[string]*string{
    "key": tea.String(""),
  }
  response_, _err := tea.DoRequest(request_, nil)
  if _err != nil {
    return _err
  }
  return _err
}


