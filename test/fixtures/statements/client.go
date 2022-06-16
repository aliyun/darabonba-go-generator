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
  if true {
    request_.Headers["host"] = tea.String("www.test2.com")
  }

  response_, _err := tea.DoRequest(request_, nil)
  if _err != nil {
    return _err
  }
  HelloIf()
  return _err
}


func HelloIf () {
  if true {
  }

  if true {
  } else if true {
  } else {
  }

}

func HelloThrow () {
  _err = tea.NewSDKError(map[string]interface{}{})
  return
}

func HelloForBreak () {
  for _, item := range []*string{tea.String("1"), tea.String("2")} {
    break
  }
}

func HelloWhile () {
  for true {
    break
  }
}

func HelloDeclare () {
  hello := tea.String("world")
  var helloNull *string
  hello = tea.String("hehe")
  num := []*int{tea.Int(1234567890)}
  num = []*int64{tea.Int64(1234567890)}
  longNum := []*int64{tea.Int64(1234567890)}
  intNum := []*int32{tea.Int(1234567890)}
  floatNum := []*float32{tea.Float32(0.123456789)}
  doubleNum := []*float64{tea.Float64(0.123456789)}
}

