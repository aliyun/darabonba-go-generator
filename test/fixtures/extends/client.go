// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source  "github.com/aliyun/darabonba-go-generator/test"
  "github.com/alibabacloud-go/tea/tea"
)

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
  _runtime := map[string]interface{}{}

  _resp := make(map[string]interface{})
  for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
    if _retryTimes > 0 {
      _backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
      if tea.IntValue(_backoffTime) > 0 {
        tea.Sleep(_backoffTime)
      }
    }

    _resp, _err = func()(map[string]interface{}, error){
      request_ := tea.NewRequest()
      _, tryErr := func()(_r map[string]interface{}, _e error) {
        defer func() {
          if r := tea.Recover(recover()); r != nil {
            _e = r
          }
        }()
        in := tea.String("try")

        return nil, nil
      }()

      if tryErr != nil {
        var e = &tea.SDKError{}
        if _t, ok := tryErr.(*tea.SDKError); ok {
          e = _t
        } else {
          e.Message = tea.String(tryErr.Error())
        }
        tmp := e.Message
      }
      response_, _err := tea.DoRequest(request_, _runtime)
      if _err != nil {
        return _result, _err
      }

      _result = nil
      return _result , _err
    }()
    if !tea.BoolValue(tea.Retryable(_err)) {
      break
    }
  }

  return _resp, _err
}


func (client *Client) TryCatch () {
  var _err error
  tryErr := func()(_e error) {
    defer func() {
      if r := tea.Recover(recover()); r != nil {
        _e = r
      }
    }()
    in := tea.String("try")

    return nil
  }()

  if tryErr != nil {
    var e = &tea.SDKError{}
    if _t, ok := tryErr.(*tea.SDKError); ok {
      e = _t
    } else {
      e.Message = tea.String(tryErr.Error())
    }
    tmp := e.Message
  }
}

func (client *Client) TryCatchWithReturn () (_result *string) {
  var _err error
  _, tryErr := func()(_r *string, _e error) {
    defer func() {
      if r := tea.Recover(recover()); r != nil {
        _e = r
      }
    }()
    in := tea.String("try")
    _result = in
    return _result , _err
  }()

  if tryErr != nil {
    var e = &tea.SDKError{}
    if _t, ok := tryErr.(*tea.SDKError); ok {
      e = _t
    } else {
      e.Message = tea.String(tryErr.Error())
    }
    tmp := e.Message
  }
  _result = tea.String("")
  return _result
}

