// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source "github.com/aliyun/darabonba-go-generator/test"
  "github.com/alibabacloud-go/tea/dara"
  "fmt"
)

type Client struct {
  DisableSDKError *bool
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



func Test (str *string) {
  panic("No Support!")
}

func (client *Client) TryMultiCatch (a *int, client *source.Client, b *string, c *int, m *source.Config) (_result *int, _err error) {
  inc := 0
  _result, _err  = tryMultiCatch_opTryFunc(client, b, m, c, a, inc)
  final := "ok"
  if _err != nil {
    if _t, ok := _err.(*Err1Error); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
      _result = nil
      return _result , _err
    }
    if _t, ok := _err.(*Err2Error); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
      _result = nil
      return _result , _err
    }
    if _t, ok := _err.(*source.Err3Error); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
      _result = nil
      return _result , _err
    }
    if _t, ok := _err.(*dara.SDKError); ok {
      err := _t;
      Test(err.Message)
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
    _err = &Err1Error{
      Name: dara.String(name),
      Code: dara.String(code),
      Data: data,
    }
    _result = data
    if _err != nil {
      if _t, ok := _err.(*Err1Error); ok {
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
    _err = &Err2Error{
      Name: dara.String(name),
      Code: dara.String(code),
      AccessErrMessage: dara.String(accessErrMessage),
    }
    _result = data
    if _err != nil {
      if _t, ok := _err.(*Err2Error); ok {
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
    if _t, ok := _err.(*source.Err3Error); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
    }
    if _t, ok := _err.(*dara.SDKError); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Message))
    }
  }
  return _result, _err
}

func tryMultiCatch_opTryFunc (client *source.Client, b *string, m *Config, c *int, a *int, inc int)( _result *int, _err error) {
  obj := map[string]interface{}{}
  client.Print(obj, dara.String("test"))
  b.Split(",")
  _err = m.Validate()
  if _err != nil {
    return _result, _err
  }
  int(c)
  req := &source.Request{
    Accesskey: dara.String(dara.Stringify([]*string{b})),
  }
  if dara.IntValue(a) > 0 {
    a = dara.Int(20)
    _err = &Err1Error{
      Name: dara.String("str"),
      Code: dara.String("str"),
      Data: map[string]*string{
        "key1": dara.String("str"),
      },
    }
    return _result, _err
  } else if dara.IntValue(a) == 0 {
    _err = &Err2Error{
      Name: dara.String("str"),
      Code: dara.String("str"),
      AccessErrMessage: dara.String("str2"),
    }
    return _result, _err
  } else if dara.IntValue(a) == -10 {
    _err = &source.Err3Error{
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

  inc++
  _result = dara.Int(dara.IntValue(a) + 100)
  return _result , _err
}

func multiTryCatch_opTryFunc (a *int, name string, code string, data map[string]*string)(_err error) {
  if dara.IntValue(a) == -10 {
    _err = &source.Err3Error{
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

