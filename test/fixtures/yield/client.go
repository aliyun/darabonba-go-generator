// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
  "fmt"
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


func (client *Client) Test3(name *string, _yield chan interface{}, _yieldErr chan error) {
  defer close(_yield)
  defer close(_yieldErr)
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
    request_.Query = map[string]*string{
      "nextToken": dara.String("100"),
      "maxResults": dara.String("200"),
    }
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

    test3_opResponse(_yield, _yieldErr, response_, name)
    _err = <-_yieldErr
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

    return
  }
  _yieldErr <- _resultErr
  return
}


func (client *Client) Test1 () (_result []*string) {
  _result = []*string{dara.String("string"), dara.String("string1"), dara.String("string2")}
  return _result
}

func (client *Client) Test2 (name *string, _yield chan *string) {
  defer close(_yield)
  test2_opYieldFunc(_yield, name)
  return
}

func (client *Client) Test4 (name *string, _yield chan interface{}, _yieldErr chan error) {
  defer close(_yield)
  defer close(_yieldErr)
  test4_opYieldFunc(_yield, _yieldErr, name)
  return
}

func (client *Client) Test5 (name *string, _yield chan *string) {
  defer close(_yield)
  test5_opYieldFunc(_yield, name)
  return
}

func (client *Client) Test6 (name *string) (_err error) {
  arr := make(chan interface{}, 1)
  _yieldErr := make(chan error, 1)
  go client.Test3(name, arr, _yieldErr)
  for data := range arr {
    fmt.Printf("[INFO] %s\n", dara.Stringify(data))
  }
  _err = <- _yieldErr
  return _err
}

func test3_opResponse(_yield chan interface{}, _yieldErr chan error, response_ *dara.Response, name *string) {
  resp := map[string]interface{}{
    "nextToken": "100",
    "truncated": false,
    "replicaPairs": "sdfs",
  }
  if dara.IntValue(response_.StatusCode) > 400 {
    _err := dara.NewSDKError(map[string]interface{}{
      "code": "sdfsd",
      "message": "sdfs",
    })
    _yieldErr <- _err
    return
  }

  name = dara.String("test")
  it := make(chan *dara.SSEEvent, 1)
  dara.ReadAsSSE(response_.Body, it, _yieldErr)
  for i := range it {
    _body := dara.ParseJSON(dara.StringValue(i.Data))
    yield <- _body
  }
}

func test2_opYieldFunc(_yield chan *string, name *string) {
  arr := client.Test1()
  name = dara.String("test")
  for _, str := range arr {
    _yield <- dara.String(str)
  }
}

func test4_opYieldFunc(_yield chan interface{}, _yieldErr chan error, name *string) {
  arr := make(chan interface{}, 1)
  client.Test3(name, arr, _yieldErr)
  for data := range arr {
    _yield <- data
  }
}

func test5_opYieldFunc(_yield chan *string, name *string) {
  arr := make(chan string, 1)
  client.Test2(name, arr)
  for data := range arr {
    _yield <- dara.String(data)
  }
}

