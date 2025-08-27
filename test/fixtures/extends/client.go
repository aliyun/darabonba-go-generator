// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source "github.com/aliyun/darabonba-go-generator/test"
  "github.com/alibabacloud-go/tea/dara"
)

type Client struct {
  source.Client
  DisableSDKError *bool
  EnableValidate *bool
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
  var _resultErr error
  retriesAttempted := int(0)
  retryPolicyContext = &dara.RetryPolicyContext{
    RetriesAttempted: retriesAttempted,
  }

  _result = make(map[string]interface{})
  for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
    _resultErr = nil
    _backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
    dara.Sleep(_backoffDelayTime)

    request_ = dara.NewRequest()
    in := "try"
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


    _result = nil
    return _result , _err
  }
  if dara.BoolValue(client.DisableSDKError) != true {
    _resultErr = dara.TeaSDKError(_resultErr)
  }
  return _result, _resultErr
}


func (client *Client) NewModels () (_err error) {
  s := &Sub{
    Name: dara.String("str"),
    Code: dara.String("str"),
    Age: dara.Int(123),
  }
  sm := &SubModel{
    Name: dara.String("str"),
    MaxAttemp: dara.Int(32),
    MaxRetry: dara.Int(32),
  }
  return _err
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

