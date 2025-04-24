// This file is auto-generated, don't edit it. Thanks.
package api

import (
  user "darabonba.com/multi/model/user"
  util "darabonba.com/multi/lib/util"
  "github.com/alibabacloud-go/tea/dara"
)

type Client struct {
  DisableSDKError *bool
}

func NewClient()(*Client, error) {
  client := new(Client)
  err := client.Init()
  return client, err
}

func (client *Client)Init()(_err error) {
  return nil
}


func (client *Client) Test3() (_result *int, _err error) {
  _runtime := dara.NewRuntimeObject(map[string]interface{}{
    "timeouted": "retry",
  })

  var retryPolicyContext *dara.RetryPolicyContext
  var request_ *dara.Request
  var response_ *dara.Response
  var _resultErr error
  retriesAttempted := int(0)
  retryPolicyContext = &dara.RetryPolicyContext{
    RetriesAttempted: retriesAttempted,
  }

  _result = dara.Int(0)
  for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
    _resultErr = nil
    _backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
    dara.Sleep(_backoffDelayTime)

    request_ = dara.NewRequest()
    request_.Protocol = dara.String("https")
    request_.Method = dara.String("DELETE")
    request_.Pathname = dara.String("/")
    request_.Headers = map[string]*string{
      "host": dara.String("test.aliyun.com"),
      "accept": dara.String("application/json"),
    }
    request_.Query = util.GetQuery()
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

    _result = response_.StatusCode
    return _result , _err
  }
  if dara.BoolValue(client.DisableSDKError) != true {
    _resultErr = dara.TeaSDKError(_resultErr)
  }
  return _result, _resultErr
}


