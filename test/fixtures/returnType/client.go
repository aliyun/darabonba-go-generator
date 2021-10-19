// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source  "github.com/aliyun/darabonba-go-generator/test"
  "github.com/alibabacloud-go/tea/tea"
)

type ComplexRequest struct {
  AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty" require:"true"`
}

func (s ComplexRequest) String() string {
  return tea.Prettify(s)
}

func (s ComplexRequest) GoString() string {
  return s.String()
}

func (s *ComplexRequest) SetAccessKey(v string) *ComplexRequest {
  s.AccessKey = &v
  return s
}

type Client struct {
  // 返回一个模版字符串
  Protocol  *string
}

func NewClient()(*Client, error) {
  client := new(Client)
  err := client.Init()
  return client, err
}

func (client *Client)Init()(_err error) {
  return nil
}


// api返回一个引用的模型构造
func (client *Client) ApiReturnConstruct() (_result *source.RuntimeObject, _err error) {
  _runtime := map[string]interface{}{
    "timeouted": "retry",
  }

  _resp := new(source.RuntimeObject)
  for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
    if _retryTimes > 0 {
      _backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
      if tea.IntValue(_backoffTime) > 0 {
        tea.Sleep(_backoffTime)
      }
    }

    _resp, _err = func()(*source.RuntimeObject, error){
      request_ := tea.NewRequest()
      response_, _err := tea.DoRequest(request_, _runtime)
      if _err != nil {
        return _result, _err
      }
      _result = &source.RuntimeObject{}
      return _result, _err
    }()
    if !tea.BoolValue(tea.Retryable(_err)) {
      break
    }
  }

  return _resp, _err
}


// 返回一个其他方法的调用结果
func (client *Client) ReturnInString (a *string) (_result *string) {
  _result = a
  return _result
}

func (client *Client) ReturnFuncResult (a *string, b *string) (_result *bool) {
  _body := client.ReturnBoolean(client.ReturnInString(a), client.ReturnInString(b))
  _result = _body
  return _result
}

func (client *Client) ReturnTemplateString () (_result *string) {
  _result = tea.String("/" + tea.StringValue(client.Protocol))
  return _result
}

// 返回一个匿名对象
func (client *Client) ReturnMap () (_result map[string]*string) {
  m := make(map[string]*string)
  tea.Convert(tea.Merge(map[string]*string{
    "key": tea.String("value"),
    "key-1": tea.String("value-1"),
    },m), &_result)
  return _result
}

// 返回一个集合类型
func (client *Client) ReturnItems () (_result []*string) {
  tea.Convert([]*string{tea.String("1")}, &_result)
  return _result
}

// 返回一个构造器
// fixme 这里如果不加throws，生成的代码会有编译错误
// 确定一下 如果没有throws的话是painc中断，还是要忽略err
func (client *Client) ReturnConstruct (cfg *source.Config) (_result *source.Client) {
  config := &source.Config{}
  _result, _err := source.NewClient(config)
  if _err != nil {
    panic(_err)
  }
  return _result
}

// 返回一个map的值
// fixme 如果这里定义map时 只有一种类型的值。会导致生成的golang map类型不是interface，产生编译错误。
func (client *Client) ReturnMapValue () (_result map[string]interface{}) {
  tmp := map[string]interface{}{
    "str": "ok",
  }
  mapVal := map[string]map[string]interface{}{
    "test": tmp,
  }
  _result = mapVal["test"]
  return _result
}

// 返回一个基本类型
func (client *Client) ReturnBoolean (a *string, b *string) (_result *bool) {
  _result = tea.Bool(true)
  return _result
}

// 返回一个模型的构造
func (client *Client) ReturnModel () (_result *ComplexRequest) {
  _result = &ComplexRequest{}
  return _result
}

