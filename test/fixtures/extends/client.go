// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source "github.com/aliyun/darabonba-go-generator/test"
  dara "github.com/alibabacloud-go/tea/tea"
  
)

type Base struct {
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s Base) String() string {
  return dara.Prettify(s)
}

func (s Base) GoString() string {
  return s.String()
}

func (s *Base) SetName(v string) *Base {
  s.Name = &v
  return s
}

type Sub struct {
  Base
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s Sub) String() string {
  return dara.Prettify(s)
}

func (s Sub) GoString() string {
  return s.String()
}

func (s *Sub) SetName(v string) *Sub {
  s.Name = &v
  return s
}

type SubModel struct {
  source.Config
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s SubModel) String() string {
  return dara.Prettify(s)
}

func (s SubModel) GoString() string {
  return s.String()
}

func (s *SubModel) SetName(v string) *SubModel {
  s.Name = &v
  return s
}

type Client struct {
  source.Client
}

func NewClient(config *source.Config)(*Client, error) {
  client := new(Client)
  err := client.Init(config)
  return client, err
}

func (client *Client)Init(config *source.Config)(_err error) {
  _err = client.Client.Init(config  )
  if _err != nil {
    return _err
  }
  return nil
}


func (client *Client) _request() (_result map[string]interface{}, _err error) {
  _runtime := dara.NewRuntimeObject(map[string]interface{}{})

  var retryPolicyContext *dara.RetryPolicyContext
  var request_ *dara.Request
  var response_ *dara.Response
  retriesAttempted := int(0)
  retryPolicyContext = &dara.RetryPolicyContext{
    RetriesAttempted: retriesAttempted,
  }

  _resp := make(map[string]interface{})
  for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
    _backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
    dara.Sleep(_backoffDelayTime)

    request_ = dara.NewRequest()
    in := "try"
    response_, _err := dara.DoRequest(request_, _runtime)
    if _err != nil {
      retriesAttempted++
      retryPolicyContext = &dara.RetryPolicyContext{
        RetriesAttempted: retriesAttempted,
        Request:          request_,
        Response:         response_,
        Error:            _err,
      }
      continue
    }


    _result = nil
    return _result , _err
  }
}


func (client *Client) TryCatch () {
  in := "try"
}

func (client *Client) TryCatchWithReturn () (_result *string) {
  in := "try"
  _result = dara.String(in)
  return _result
  _result = dara.String("")
  return _result
}

