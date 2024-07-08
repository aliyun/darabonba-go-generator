// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source "github.com/aliyun/darabonba-go-generator/test"
  dara "github.com/alibabacloud-go/tea/tea"
  "fmt"
)

type iErr1 interface {
  dara.BaseError
  GetData() map[string]*string 
}

type Err1 struct {
  dara.BaseError
  Name *string ``
  Message *string ``
  Code *string ``
  Stack *string ``
  Data map[string]*string ` require:"true"`
}

func (err Err1) Error() string {
  if err.Message == nil {
    str := fmt.Sprintf("Err1:\n   Name: %s\n   Code: %s\n",
      dara.StringValue(err.Name), dara.StringValue(err.Code))
    err.Message = dara.String(str)
  }
  return dara.StringValue(err.Message)
}

func (s *Err1) GetName() *string  {
  return s.Name
}

func (s *Err1) GetMessage() *string  {
  return s.Message
}

func (s *Err1) GetCode() *string  {
  return s.Code
}

func (s *Err1) GetStack() *string  {
  return s.Stack
}

func (s *Err1) GetData() map[string]*string  {
  return s.Data
}

type iErr2 interface {
  dara.BaseError
  GetAccessErrMessage() *string 
}

type Err2 struct {
  dara.BaseError
  Name *string ``
  Message *string ``
  Code *string ``
  Stack *string ``
  AccessErrMessage *string ` require:"true"`
}

func (err Err2) Error() string {
  if err.Message == nil {
    str := fmt.Sprintf("Err2:\n   Name: %s\n   Code: %s\n",
      dara.StringValue(err.Name), dara.StringValue(err.Code))
    err.Message = dara.String(str)
  }
  return dara.StringValue(err.Message)
}

func (s *Err2) GetName() *string  {
  return s.Name
}

func (s *Err2) GetMessage() *string  {
  return s.Message
}

func (s *Err2) GetCode() *string  {
  return s.Code
}

func (s *Err2) GetStack() *string  {
  return s.Stack
}

func (s *Err2) GetAccessErrMessage() *string  {
  return s.AccessErrMessage
}

type Client struct {
  Configs  []*source.Config
}

func NewClient(config *source.Config)(*Client, error) {
  client := new(Client)
  err := client.Init(config)
  return client, err
}

func (client *Client)Init(config *source.Config)(_err error) {
  client.Configs[0] = config
  test, _err := source.NewClient(config)
  if _err != nil {
    return _err
  }

  return nil
}



func (client *Client) TryMultiCatch (a *int) (_result *int, _err error) {
  _result, _err  = tryMultiCatch_opTryFunc(a)
  final := "ok"
  if _err != nil {
    if _t, ok := _err.(*Err1); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
      _result = nil
      return _result , _err
    }
    if _t, ok := _err.(*Err2); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
      _result = nil
      return _result , _err
    }
    if _t, ok := _err.(*source.Err3); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
      _result = nil
      return _result , _err
    }
    if _t, ok := _err.(dara.BaseError); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
      _result = nil
      return _result , _err
    }
  }
  return _result, _err
}

func (client *Client) MultiTryCatch (a *int) (_result map[string]*string, _err error) {
  name := "str"
  code := "str"
  data := map[string]*string{
    "key1": dara.String("str"),
  }
  if dara.IntValue(a) > 0 {
    _err = &Err1{
      Name: dara.String(name),
      Code: dara.String(code),
      Data: data,
    }
    _result = data
    if _err != nil {
      if _t, ok := _err.(*Err1); ok {
        err := _t;
        name = "str1"
        code = "str1"
        data = map[string]*string{
          "key1": dara.String("str1"),
        }
        fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
      }
    }
  }

  accessErrMessage := "str2"
  if dara.IntValue(a) == 0 {
    _err = &Err2{
      Name: dara.String(name),
      Code: dara.String(code),
      AccessErrMessage: dara.String(accessErrMessage),
    }
    _result = data
    if _err != nil {
      if _t, ok := _err.(*Err2); ok {
        err := _t;
        name = "str2"
        code = "str2"
        data = map[string]*string{
          "key1": dara.String("str2"),
        }
        fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
      }
    }
  }

  _result, _err  = multiTryCatch_opTryFunc(a, name, code, data)
  final := "ok"
  _result = data
  if _err != nil {
    if _t, ok := _err.(*source.Err3); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
    }
    if _t, ok := _err.(dara.BaseError); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
    }
  }
  return _result, _err
}

func tryMultiCatch_opTryFunc (a *int)( _result *int, _err error) {
  if dara.IntValue(a) > 0 {
    a = dara.Int(20)
    _err = &Err1{
      Name: dara.String("str"),
      Code: dara.String("str"),
      Data: map[string]*string{
        "key1": dara.String("str"),
      },
    }
    return _result, _err
  } else if dara.IntValue(a) == 0 {
    _err = &Err2{
      Name: dara.String("str"),
      Code: dara.String("str"),
      AccessErrMessage: dara.String("str2"),
    }
    return _result, _err
  } else if dara.IntValue(a) == -10 {
    _err = &source.Err3{
      Name: dara.String("str"),
      Code: dara.String("str"),
    }
    return _result, _err
  } else {
    _err = &dara.SDKError{
      Name: dara.String("str"),
      Code: dara.String("str"),
    }
    return _result, _err
  }

  _result = dara.Int(dara.IntValue(a) + 100)
  return _result , _err
}

func multiTryCatch_opTryFunc (a *int, name string, code string, data map[string]*string)(_err error) {
  if dara.IntValue(a) == -10 {
    _err = &source.Err3{
      Name: dara.String(name),
      Code: dara.String(code),
      AccessErrMessage: data["key1"],
    }
    return _err
  } else if dara.IntValue(a) == -100 {
    _err = &dara.SDKError{
      Name: dara.String(name),
      Code: dara.String(code),
    }
    return _err
  }

  return _err
}

