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
}

func (s Config) String() string {
  return dara.Prettify(s)
}

func (s Config) GoString() string {
  return s.String()
}

func (s *Config) Validate() error {
  return dara.Validate(s)
}

type iM interface {
  dara.Model
  String() string
  GoString() string
}

type M struct {
}

func (s M) String() string {
  return dara.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) Validate() error {
  return dara.Validate(s)
}

type Client struct {
  DisableSDKError *bool
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
  response_, _err := dara.DoRequest(request_, nil)
  if _err != nil {
    return _err
  }

  return
}

func (client *Client) HelloRuntime() (_err error) {
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
    request_.Method = dara.String("GET")
    request_.Pathname = dara.String("/")
    request_.Headers = map[string]*string{
      "host": dara.String("www.test.com"),
    }
    var test map[string]interface{}
    test = map[string]interface{}{
      "key": "value",
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

    testStrTmp, _err := client.GetHost()
    testStr := dara.StringValue(testStrTmp)
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
  if dara.BoolValue(client.DisableSDKError) != true {
    _resultErr = dara.TeaSDKError(_resultErr)
  }
  return _result, _resultErr
}

func (client *Client) HelloVirtualCall(m *M) (_err error) {
  request_ := dara.NewRequest()
  request_.Method = dara.String("GET")
  request_.Pathname = dara.String("/")
  request_.Headers = map[string]*string{
    "key": dara.String(""),
  }
  response_, _err := dara.DoRequest(request_, nil)
  if _err != nil {
    return _err
  }

  return
}

func (client *Client) HelloComplex() (_result interface{}, _err error) {
  _runtime := dara.NewRuntimeObject(map[string]interface{}{})

  var retryPolicyContext *dara.RetryPolicyContext
  var request_ *dara.Request
  var response_ *dara.Response
  var _resultErr error
  retriesAttempted := int(0)
  retryPolicyContext = &dara.RetryPolicyContext{
    RetriesAttempted: retriesAttempted,
  }

  _result = interface{}(nil)
  for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
    _resultErr = nil
    _backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
    dara.Sleep(_backoffDelayTime)

    request_ = dara.NewRequest()
    hostTmp, _err := client.GetHost()
    host := dara.StringValue(hostTmp)
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

    request_.Method = dara.String("GET")
    request_.Pathname = dara.String("/")
    request_.Headers = map[string]*string{
      "host": dara.String(host),
    }
    var test map[string]interface{}
    test = map[string]interface{}{
      "key": "value",
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

    _result, _err = helloComplex_opResponse(response_)
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

    return _result, _err
  }
  if dara.BoolValue(client.DisableSDKError) != true {
    _resultErr = dara.TeaSDKError(_resultErr)
  }
  return _result, _resultErr
}


func (client *Client) GetHost () (_result *string, _err error) {
  _result = dara.String("www.test.com")
  return _result, _err
}

func helloComplex_opResponse (response_ *dara.Response)( _result interface{}, _err error) {
  if dara.IntValue(response_.StatusCode) != 200 {
    _err = dara.NewSDKError(map[string]interface{}{
      "code": dara.ToString(dara.IntValue(response_.StatusCode)),
      "message": "httpCode: " + dara.ToString(dara.IntValue(response_.StatusCode)) + " ",
    })
    return _result, _err
  }

  obj, _err := dara.ReadAsJSON(response_.Body)
  if _err != nil {
    return _result, _err
  }

  _result = obj
  return _result , _err
}

