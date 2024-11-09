// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
  
)

type iConfig interface {
  dara.Model
  String() string
  GoString() string
}

type Config struct {
  dara.Model
}

func (s Config) String() string {
  return dara.Prettify(s)
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
  request_ := dara.NewRequest()
  request_.Method = dara.String("GET")
  request_.Pathname = dara.String("/")
  request_.Headers = map[string]*string{
    "host": dara.String("www.test.com"),
  }
  if true {
    request_.Headers["host"] = dara.String("www.test2.com")
  }

  response_, _err := dara.DoRequest(request_, nil)
  if _err != nil {
    return _err
  }

  HelloIf()
  return
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
  _err = dara.NewSDKError(map[string]interface{}{})
  return
}

func HelloForBreak () {
  for _, item := range []*string{dara.String("1"), dara.String("2")} {
    break
  }
}

func HelloWhile () {
  for true {
    break
  }
}

func HelloDeclare () {
  hello := "world"
  var helloNull string
  hello = "hehe"
}

