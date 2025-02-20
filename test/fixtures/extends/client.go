// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source "github.com/aliyun/darabonba-go-generator/test"
  "github.com/alibabacloud-go/tea/dara"
  
)

type iBase interface {
  dara.Model
  String() string
  GoString() string
  SetName(v string) *Base
  GetName() *string 
  SetAge(v int) *Base
  GetAge() *int 
}

type Base struct {
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  Age *int `json:"age,omitempty" xml:"age,omitempty" require:"true"`
}

func (s Base) String() string {
  return dara.Prettify(s)
}

func (s Base) GoString() string {
  return s.String()
}

func (s *Base) GetName() *string  {
  return s.Name
}

func (s *Base) GetAge() *int  {
  return s.Age
}

func (s *Base) SetName(v string) *Base {
  s.Name = &v
  return s
}

func (s *Base) SetAge(v int) *Base {
  s.Age = &v
  return s
}

func (s *Base) Validate() error {
  return dara.Validate(s)
}

type iSub interface {
  iBase
  String() string
  GoString() string
  SetName(v string) *Sub
  GetName() *string 
  SetCode(v string) *Sub
  GetCode() *string 
}

type Sub struct {
  Age *int `json:"age,omitempty" xml:"age,omitempty" require:"true"`
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s Sub) String() string {
  return dara.Prettify(s)
}

func (s Sub) GoString() string {
  return s.String()
}

func (s *Sub) GetAge() *int  {
  return s.Age
}

func (s *Sub) GetName() *string  {
  return s.Name
}

func (s *Sub) GetCode() *string  {
  return s.Code
}

func (s *Sub) SetAge(v int) *Sub {
  s.Age = &v
  return s
}

func (s *Sub) SetName(v string) *Sub {
  s.Name = &v
  return s
}

func (s *Sub) SetCode(v string) *Sub {
  s.Code = &v
  return s
}

func (s *Sub) Validate() error {
  return dara.Validate(s)
}

type iSubModel interface {
  source.iConfig
  String() string
  GoString() string
  SetName(v string) *SubModel
  GetName() *string 
}

type SubModel struct {
  MaxAttemp *int `json:"maxAttemp,omitempty" xml:"maxAttemp,omitempty" require:"true"`
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s SubModel) String() string {
  return dara.Prettify(s)
}

func (s SubModel) GoString() string {
  return s.String()
}

func (s *SubModel) GetMaxAttemp() *int  {
  return s.MaxAttemp
}

func (s *SubModel) GetName() *string  {
  return s.Name
}

func (s *SubModel) SetMaxAttemp(v int) *SubModel {
  s.MaxAttemp = &v
  return s
}

func (s *SubModel) SetName(v string) *SubModel {
  s.Name = &v
  return s
}

func (s *SubModel) Validate() error {
  return dara.Validate(s)
}

type Client struct {
  source.Client
  DisableSDKError *bool
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

