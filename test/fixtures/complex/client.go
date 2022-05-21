// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea"
  "io"
  source  "github.com/aliyun/darabonba-go-generator/test"
  "github.com/alibabacloud-go/tea/tea"
)

type ComplexRequest struct {
  AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty" require:"true"`
  // Body
  Body io.Reader `json:"Body,omitempty" xml:"Body,omitempty" require:"true"`
  // Strs
  Strs []*string `json:"Strs,omitempty" xml:"Strs,omitempty" require:"true" type:"Repeated"`
  // mapList
  MapList []map[string]interface{} `json:"mapList,omitempty" xml:"mapList,omitempty" require:"true" type:"Repeated"`
  // header
  Header *ComplexRequestHeader `json:"header,omitempty" xml:"header,omitempty" require:"true" type:"Struct"`
  Configs *ComplexRequestConfigs `json:"configs,omitempty" xml:"configs,omitempty" require:"true" type:"Struct"`
  Num *int `json:"num,omitempty" xml:"num,omitempty" require:"true"`
  I64 *int64 `json:"i64,omitempty" xml:"i64,omitempty" require:"true"`
  F64 *float64 `json:"f64,omitempty" xml:"f64,omitempty" require:"true"`
  B *bool `json:"b,omitempty" xml:"b,omitempty" require:"true"`
  F32 *float32 `json:"f32,omitempty" xml:"f32,omitempty" require:"true"`
  F64List []*float64 `json:"f64List,omitempty" xml:"f64List,omitempty" require:"true" type:"Repeated"`
  FloatList []*float32 `json:"floatList,omitempty" xml:"floatList,omitempty" require:"true" type:"Repeated"`
  BooleantList []*bool `json:"booleantList,omitempty" xml:"booleantList,omitempty" require:"true" type:"Repeated"`
  I32 *int32 `json:"i32,omitempty" xml:"i32,omitempty" require:"true"`
  StringList []*string `json:"stringList,omitempty" xml:"stringList,omitempty" require:"true" type:"Repeated"`
  IntList []*int `json:"intList,omitempty" xml:"intList,omitempty" require:"true" type:"Repeated"`
  Int32List []*int32 `json:"int32List,omitempty" xml:"int32List,omitempty" require:"true" type:"Repeated"`
  Int16List []*int16 `json:"int16List,omitempty" xml:"int16List,omitempty" require:"true" type:"Repeated"`
  Int64List []*int64 `json:"int64List,omitempty" xml:"int64List,omitempty" require:"true" type:"Repeated"`
  Uint64List []*uint64 `json:"uint64List,omitempty" xml:"uint64List,omitempty" require:"true" type:"Repeated"`
  Uint32List []*uint32 `json:"uint32List,omitempty" xml:"uint32List,omitempty" require:"true" type:"Repeated"`
  Uint16List []*uint16 `json:"uint16List,omitempty" xml:"uint16List,omitempty" require:"true" type:"Repeated"`
  U64 *uint64 `json:"u64,omitempty" xml:"u64,omitempty" require:"true"`
  U32 *uint32 `json:"u32,omitempty" xml:"u32,omitempty" require:"true"`
  U16 *uint16 `json:"u16,omitempty" xml:"u16,omitempty" require:"true"`
  Obj map[string]interface{} `json:"obj,omitempty" xml:"obj,omitempty" require:"true"`
  Any interface{} `json:"any,omitempty" xml:"any,omitempty" require:"true"`
  Byt []byte `json:"byt,omitempty" xml:"byt,omitempty" require:"true"`
  Req *tea.Request `json:"req,omitempty" xml:"req,omitempty" require:"true"`
  Resp *tea.Response `json:"resp,omitempty" xml:"resp,omitempty" require:"true"`
  Map map[string]*string `json:"map,omitempty" xml:"map,omitempty" require:"true"`
  NumMap map[string]*int `json:"numMap,omitempty" xml:"numMap,omitempty" require:"true"`
  ModelMap map[string]*source.Request `json:"modelMap,omitempty" xml:"modelMap,omitempty" require:"true"`
  Request *source.Request `json:"request,omitempty" xml:"request,omitempty" require:"true"`
  Client *source.Client `json:"client,omitempty" xml:"client,omitempty" require:"true"`
  Instance *source.RequestInstance `json:"instance,omitempty" xml:"instance,omitempty" require:"true"`
  // Deprecated
  // Part
  Part []*ComplexRequestPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
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

func (s *ComplexRequest) SetBody(v io.Reader) *ComplexRequest {
  s.Body = v
  return s
}

func (s *ComplexRequest) SetStrs(v []*string) *ComplexRequest {
  s.Strs = v
  return s
}

func (s *ComplexRequest) SetMapList(v []map[string]interface{}) *ComplexRequest {
  s.MapList = v
  return s
}

func (s *ComplexRequest) SetHeader(v *ComplexRequestHeader) *ComplexRequest {
  s.Header = v
  return s
}

func (s *ComplexRequest) SetConfigs(v *ComplexRequestConfigs) *ComplexRequest {
  s.Configs = v
  return s
}

func (s *ComplexRequest) SetNum(v int) *ComplexRequest {
  s.Num = &v
  return s
}

func (s *ComplexRequest) SetI64(v int64) *ComplexRequest {
  s.I64 = &v
  return s
}

func (s *ComplexRequest) SetF64(v float64) *ComplexRequest {
  s.F64 = &v
  return s
}

func (s *ComplexRequest) SetB(v bool) *ComplexRequest {
  s.B = &v
  return s
}

func (s *ComplexRequest) SetF32(v float32) *ComplexRequest {
  s.F32 = &v
  return s
}

func (s *ComplexRequest) SetF64List(v []*float64) *ComplexRequest {
  s.F64List = v
  return s
}

func (s *ComplexRequest) SetFloatList(v []*float32) *ComplexRequest {
  s.FloatList = v
  return s
}

func (s *ComplexRequest) SetBooleantList(v []*bool) *ComplexRequest {
  s.BooleantList = v
  return s
}

func (s *ComplexRequest) SetI32(v int32) *ComplexRequest {
  s.I32 = &v
  return s
}

func (s *ComplexRequest) SetStringList(v []*string) *ComplexRequest {
  s.StringList = v
  return s
}

func (s *ComplexRequest) SetIntList(v []*int) *ComplexRequest {
  s.IntList = v
  return s
}

func (s *ComplexRequest) SetInt32List(v []*int32) *ComplexRequest {
  s.Int32List = v
  return s
}

func (s *ComplexRequest) SetInt16List(v []*int16) *ComplexRequest {
  s.Int16List = v
  return s
}

func (s *ComplexRequest) SetInt64List(v []*int64) *ComplexRequest {
  s.Int64List = v
  return s
}

func (s *ComplexRequest) SetUint64List(v []*uint64) *ComplexRequest {
  s.Uint64List = v
  return s
}

func (s *ComplexRequest) SetUint32List(v []*uint32) *ComplexRequest {
  s.Uint32List = v
  return s
}

func (s *ComplexRequest) SetUint16List(v []*uint16) *ComplexRequest {
  s.Uint16List = v
  return s
}

func (s *ComplexRequest) SetU64(v uint64) *ComplexRequest {
  s.U64 = &v
  return s
}

func (s *ComplexRequest) SetU32(v uint32) *ComplexRequest {
  s.U32 = &v
  return s
}

func (s *ComplexRequest) SetU16(v uint16) *ComplexRequest {
  s.U16 = &v
  return s
}

func (s *ComplexRequest) SetObj(v map[string]interface{}) *ComplexRequest {
  s.Obj = v
  return s
}

func (s *ComplexRequest) SetAny(v interface{}) *ComplexRequest {
  s.Any = v
  return s
}

func (s *ComplexRequest) SetByt(v []byte) *ComplexRequest {
  s.Byt = v
  return s
}

func (s *ComplexRequest) SetReq(v *tea.Request) *ComplexRequest {
  s.Req = v
  return s
}

func (s *ComplexRequest) SetResp(v *tea.Response) *ComplexRequest {
  s.Resp = v
  return s
}

func (s *ComplexRequest) SetMap(v map[string]*string) *ComplexRequest {
  s.Map = v
  return s
}

func (s *ComplexRequest) SetNumMap(v map[string]*int) *ComplexRequest {
  s.NumMap = v
  return s
}

func (s *ComplexRequest) SetModelMap(v map[string]*source.Request) *ComplexRequest {
  s.ModelMap = v
  return s
}

func (s *ComplexRequest) SetRequest(v *source.Request) *ComplexRequest {
  s.Request = v
  return s
}

func (s *ComplexRequest) SetClient(v *source.Client) *ComplexRequest {
  s.Client = v
  return s
}

func (s *ComplexRequest) SetInstance(v *source.RequestInstance) *ComplexRequest {
  s.Instance = v
  return s
}

func (s *ComplexRequest) SetPart(v []*ComplexRequestPart) *ComplexRequest {
  s.Part = v
  return s
}

type ComplexRequestHeader struct {
  // Body
  Content *string `json:"Content,omitempty" xml:"Content,omitempty" require:"true" signed:"true"`
}

func (s ComplexRequestHeader) String() string {
  return tea.Prettify(s)
}

func (s ComplexRequestHeader) GoString() string {
  return s.String()
}

func (s *ComplexRequestHeader) SetContent(v string) *ComplexRequestHeader {
  s.Content = &v
  return s
}

type ComplexRequestConfigs struct {
  Key *string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
  Value []*string `json:"value,omitempty" xml:"value,omitempty" require:"true" type:"Repeated"`
  Extra map[string]*string `json:"extra,omitempty" xml:"extra,omitempty" require:"true"`
}

func (s ComplexRequestConfigs) String() string {
  return tea.Prettify(s)
}

func (s ComplexRequestConfigs) GoString() string {
  return s.String()
}

func (s *ComplexRequestConfigs) SetKey(v string) *ComplexRequestConfigs {
  s.Key = &v
  return s
}

func (s *ComplexRequestConfigs) SetValue(v []*string) *ComplexRequestConfigs {
  s.Value = v
  return s
}

func (s *ComplexRequestConfigs) SetExtra(v map[string]*string) *ComplexRequestConfigs {
  s.Extra = v
  return s
}

type ComplexRequestPart struct     {
  // PartNumber
  PartNumber *string `json:"PartNumber,omitempty" xml:"PartNumber,omitempty"`
}

func (s ComplexRequestPart) String() string {
  return tea.Prettify(s)
}

func (s ComplexRequestPart) GoString() string {
  return s.String()
}

func (s *ComplexRequestPart) SetPartNumber(v string) *ComplexRequestPart {
  s.PartNumber = &v
  return s
}

type Response struct {
  Instance *ComplexRequestPart `json:"instance,omitempty" xml:"instance,omitempty" require:"true"`
}

func (s Response) String() string {
  return tea.Prettify(s)
}

func (s Response) GoString() string {
  return s.String()
}

func (s *Response) SetInstance(v *ComplexRequestPart) *Response {
  s.Instance = v
  return s
}

type Client struct {
  Protocol  *string
  Pathname  *string
  Strs  []*string
  CompleList  [][]*string
  EndpointMap  map[string]*string
  Configs  []*source.Config
}

func NewClient(config *source.Config)(*Client, error) {
  client := new(Client)
  err := client.Init(config)
  return client, err
}

func (client *Client)Init(config *source.Config)(_err error) {
  client.Configs[0] = config
  return nil
}


func (client *Client) Complex1(request *ComplexRequest, client *source.Client) (_result *source.RuntimeObject, _err error) {
  _err = tea.Validate(request)
  if _err != nil {
    return _result, _err
  }
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
      name := tea.String("complex")
      var read io.Reader
      var byt []byte
      moduleModelMapVal := make(map[string]*source.RuntimeObject)
      moduleMapVal := make(map[string]*source.Client)
      modelMapVal := make(map[string]*ComplexRequest)
      subModelMapVal := make(map[string]*ComplexRequestHeader)
      var reqMap map[string]*ComplexRequest
      mapString := map[string]*string{
        "str": request.AccessKey,
      }
      inte := tea.Int(1)
      a := tea.Int(1)
      var b *int32
      b = tea.ToInt32(a)
      c := a
      IntToInt32(tea.ToInt32(a))
      mapVal := map[string]interface{}{
        "read": read,
        "test": "{\"test\":\"ok\"}",
        "b": tea.BoolValue(request.B),
        "num": tea.IntValue(request.Num),
        "u16": tea.Uint16Value(request.U16),
        "u32": tea.Uint32Value(request.U32),
        "u64": tea.Uint64Value(request.U64),
        "u16List": tea.Uint16SliceValue(request.Uint16List),
        "u32List": tea.Uint32SliceValue(request.Uint32List),
        "u64List": tea.Uint64SliceValue(request.Uint64List),
        "i64List": tea.Int64SliceValue(request.Int64List),
        "i16List": tea.Int16SliceValue(request.Int16List),
        "i32List": tea.Int32SliceValue(request.Int32List),
        "intList": tea.IntSliceValue(request.IntList),
        "stringList": tea.StringSliceValue(request.StringList),
        "i32": tea.Int32Value(request.I32),
        "booleantList": tea.BoolSliceValue(request.BooleantList),
        "floatList": tea.Float32SliceValue(request.FloatList),
        "float64List": tea.Float64SliceValue(request.F64List),
        "f32": tea.Float32Value(request.F32),
        "f64": tea.Float64Value(request.F64),
        "i64": tea.Int64Value(request.I64),
      }
      req := &ComplexRequest{
        B: tea.Bool(false),
        Num: tea.Int(10),
        I32: tea.ToInt32(a),
        IntList: []*int{tea.Int(10), tea.Int(11)},
        StringList: []*string{tea.String("10"), tea.String("11")},
        BooleantList: []*bool{tea.Bool(true), tea.Bool(false)},
      }
      client.Strs = request.Strs
      client.EndpointMap[tea.StringValue(client.Protocol)]
      client.EndpointMap["test"] = tea.String("ok")
      request.Strs = client.Strs
      request_.Protocol = client.Protocol
      request_.Port = request.Num
      request_.Method = tea.String("GET")
      request_.Pathname = tea.String("/" + tea.StringValue(client.Pathname))
      request_.Query = map[string]*string{
        "date": tea.String("2019"),
        "name": request_.Method,
      }
      tmp := tea.ToMap(request_.Query,
        request_.Headers,
        request_)
      response_, _err := tea.DoRequest(request_, _runtime)
      if _err != nil {
        return _result, _err
      }
      if true && true {
        _result = nil
        return _result , _err
      } else if true || false {
        _result = &source.RuntimeObject{}
        return _result, _err
      }

      client.Print(tea.ToMap(request), tea.String("1"))
      _, _err = client.Hello(tea.ToMap(request), []*string{tea.String("1"), tea.String("2")}, nil)
      if _err != nil {
        return _result, _err
      }
      _, _err = client.Hello(nil, nil, nil)
      if _err != nil {
        return _result, _err
      }
      _result = &source.RuntimeObject{}
      _err = tea.Convert(map[string]interface{}{}, &_result)
      return _result, _err
      _, _err = client.Complex3(nil)
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

func (client *Client) Complex2(request *ComplexRequest, str []*string, val map[string]*string) (_result map[string]interface{}, _err error) {
  _err = tea.Validate(request)
  if _err != nil {
    return _result, _err
  }
  request_ := tea.NewRequest()
  name := tea.String("complex")
  config := &source.Config{}
  client, _err := source.NewClient(config)
  if _err != nil {
    return _result, _err
  }

  configArray := []*source.Config{config}
  request_.Protocol = tea.String("HTTP")
  request_.Port = tea.Int(80)
  request_.Method = tea.String("GET")
  request_.Pathname = tea.String("/")
  request_.Query = map[string]*string{
    "date": tea.String("2019"),
    "protocol": request_.Protocol,
  }
  response_, _err := tea.DoRequest(request_, nil)
  if _err != nil {
    return _result, _err
  }

  return nil, nil
}

func (client *Client) ComplexMap() (_result map[string]interface{}, _err error) {
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
      response_, _err := tea.DoRequest(request_, _runtime)
      if _err != nil {
        return _result, _err
      }

      return nil, nil
    }()
    if !tea.BoolValue(tea.Retryable(_err)) {
      break
    }
  }

  return _resp, _err
}

func (client *Client) Complex3(request *ComplexRequest) (_result *ComplexRequest, _err error) {
  _err = tea.Validate(request)
  if _err != nil {
    return _result, _err
  }
  request_ := tea.NewRequest()
  name := tea.String("complex")
  request_.Protocol, _err = client.TemplateString()
  if _err != nil {
    return _result, _err
  }

  request_.Port = tea.Int(80)
  request_.Method = tea.String("GET")
  request_.Pathname = tea.String("/")
  request_.Body = tea.ToReader(tea.String("body"))
  request_.Query = map[string]*string{
    "date": tea.String("2019"),
  }
  var tmp *ComplexRequest
  tmp = client.ReturnModel()
  response_, _err := tea.DoRequest(request_, nil)
  if _err != nil {
    return _result, _err
  }
  resp := response_
  req := &source.Request{
    Accesskey: request.AccessKey,
    Region: resp.StatusMessage,
  }
  Array0(tea.ToMap(request))
  req.Accesskey = tea.String("accesskey")
  req.Accesskey = request.AccessKey
  _err = PrintNull()
  if _err != nil {
    return _result, _err
  }
  _, _err = client.ThrowsFunc()
  if _err != nil {
    return _result, _err
  }
  response_.StatusCode
  source.Array(tea.ToMap(request), tea.String("1"))
  _result = &ComplexRequest{}
  _err = tea.Convert(tea.Merge(request_.Query), &_result)
  return _result, _err
}

func (client *Client) NoReturn() (_err error) {
  request_ := tea.NewRequest()
  response_, _err := tea.DoRequest(request_, nil)
  if _err != nil {
    return _err
  }

  return nil
}


func (client *Client) Hello (request map[string]interface{}, strs []*string, complexList [][]*string) (_result []*string, _err error) {
  var a [][]*string
  _result = make([]*string, 0)
  _body := Array1()
  _result = _body
  return _result, _err
}

func Print (reqeust *tea.Request, reqs []*ComplexRequest, response *tea.Response, val map[string]*string) (_result *source.Request, _err error) {
  panic("No Support!")
}

func IntToInt32 (a *int32) {
  panic("No Support!")
}

func AssignWithArray () (_err error) {
  var list []*string
  list = []*string{tea.String("test")}
  var str *string
  str, _err = client.ThrowsFunc()
  if _err != nil {
    return _err
  }

  return _err
}

func (client *Client) MapAcess () {
  tmp := map[string]interface{}{
    "protocol": tea.StringValue(client.EndpointMap[tea.StringValue(client.Protocol)]),
  }
}

func (client *Client) ExprFunc () (_result []*string, _err error) {
  if !true {
  }

  num := tea.Int(10)
  req := &ComplexRequest{}
  mapVal := map[string]interface{}{
    "num": 10,
    "client": new(Source),
    "strs": Array1()),
    "str": "string" + tea.ToString(tea.IntValue(num)),
    "str1": "string" + tea.StringValue(req.AccessKey),
  }
  _result = nil
  return _result , _err
}

func PrintNull () (_err error) {
  defer func() {
    final := tea.String("ok")
  }()
  func()(_e error) {
    str, _err := client.TemplateString()
    if _err != nil {
      return _err
    }


    return nil
  }()

  return _err
}

func TestTryWithComplexReturnType () (_result *source.Request, _err error) {
  defer func() {
    final := tea.String("ok")
  }()
  func()(_r *source.Request, _e error) {
    str, _err := client.TemplateString()
    if _err != nil {
      return _result, _err
    }


    return nil, nil
  }()

  _result = nil
  return _result , _err
}

func TestTryWithComplexReturnTypeWithOutCat () (_result *source.Request, _err error) {
  defer func() {
    final := tea.String("ok")
  }()
  _, tryErr := func()(_r *source.Request, _e error) {
    defer func() {
      if r := tea.Recover(recover()); r != nil {
        _e = r
      }
    }()
    str, _err := client.TemplateString()
    if _err != nil {
      return _result, _err
    }


    return nil, nil
  }()

  if tryErr != nil {
    var e = &tea.SDKError{}
    if _t, ok := tryErr.(*tea.SDKError); ok {
      e = _t
    } else {
      e.Message = tea.String(tryErr.Error())
    }
    sim := tea.String("a")
  }
  _result = nil
  return _result , _err
}

func Array0 (req map[string]interface{}) (_result []interface{}) {
  _result = make([]interface{}, 0)
  tea.Convert([]interface{}{}, &_result)
  return _result
}

func Array1 () (_result []*string) {
  _result = make([]*string, 0)
  tea.Convert([]*string{tea.String("1")}, &_result)
  return _result
}

func (client *Client) TemplateString () (_result *string, _err error) {
  _result = tea.String("/" + tea.StringValue(client.Protocol))
  return _result, _err
}

func (client *Client) ThrowsFunc () (_result *string, _err error) {
  _result = tea.String("/" + tea.StringValue(client.Protocol))
  return _result, _err
}

func (client *Client) ThrowsFunc1 () (_result *string, _err error) {
  _result = tea.String("")
  return _result, _err
}

func (client *Client) ThrowsFunc2 () (_err error) {
  _err = tea.NewSDKError(map[string]interface{}{
    "code": "",
  })
  return _err
}

func (client *Client) ThrowsFunc3 () (_result *string, _err error) {
  _err = tea.NewSDKError(map[string]interface{}{
    "code": "",
  })
  return _result, _err
}

func (client *Client) ReturnFunc () (_result *string) {
  _result = nil
  return _result
}

func (client *Client) ReturnFunc1 (cfg *source.Config) (_result *source.Client) {
  config := &source.Config{}
  _result = &source.Client{}
  _result, _err = source.NewClient(config)
  return _result
}

func (client *Client) ReturnFunc2 () (_result map[string]interface{}) {
  tmp := map[string]*string{
    "subMap": tea.String("ok"),
  }
  mapVal := map[string]map[string]*string{
    "test": tmp,
  }
  if true {
    _result = make(map[string]interface{})
    _result = mapVal["test"]
    return _result
  } else {
    var body io.Reader
    _result = make(map[string]interface{})
    _result.Body = body
    tea.Convert(tea.ToMap(tmp), &_result)
    return _result
  }

}

func (client *Client) ReturnModel () (_result *ComplexRequest) {
  _result = &ComplexRequest{}
  return _result
}

func (client *Client) EmptyFunc () {
  panic("No Support!")
}

func (client *Client) Error (e *tea.SDKError) (_result *tea.SDKError) {
  var tmp *tea.SDKError
  var c interface{}
  _result = e
  return _result
}

func ArrayAccess () (_result *string) {
  configs := []*string{tea.String("a"), tea.String("b"), tea.String("c")}
  config := configs[0]
  _result = config
  return _result
}

func ArrayAccess2 () (_result *string) {
  data := map[string][]*string{
    "configs": []*string{tea.String("a"), tea.String("b"), tea.String("c")},
  }
  config := data["configs"][0]
  _result = config
  return _result
}

func ArrayAccess3 (request *ComplexRequest) (_result *string) {
  req := &source.Request{}
  ArrayAccess4([]*source.Request{req})
  configVal := request.Configs.Value[0]
  _result = configVal
  return _result
}

func ArrayAccess4 (requests []*source.Request) (_result *string) {
  _result = tea.String("")
  return _result
}

func ArrayAssign (config *string) (_result []*string) {
  configs := []*string{tea.String("a"), tea.String("b"), tea.String("c")}
  configs[3] = config
  _result = configs
  return _result
}

func ArrayAssign2 (config *string) (_result []*string) {
  data := map[string][]*string{
    "configs": []*string{tea.String("a"), tea.String("b"), tea.String("c")},
  }
  data["configs"][3] = config
  _result = data["configs"]
  return _result
}

func ArrayAssign3 (request *ComplexRequest, config *string) {
  request.Configs.Value[0] = config
}

func MapAccess (request *ComplexRequest) (_result *string) {
  configInfo := request.Configs.Extra["name"]
  _result = configInfo
  return _result
}

func MapAccess2 (request *source.Request) (_result *string) {
  configInfo := request.Configs.Extra["name"]
  _result = configInfo
  return _result
}

func MapAccess3 () (_result *string) {
  data := map[string]map[string]*string{
    "configs": map[string]*string{
      "value": tea.String("string"),
    },
  }
  _result = data["configs"]["value"]
  return _result
}

func MapAccess4 (request *ComplexRequest) (_result *string) {
  key := tea.String("name")
  model := request.ModelMap[tea.StringValue(key)]
  configInfo := request.Configs.Extra[tea.StringValue(key)]
  _result = configInfo
  return _result
}

func MapAssign (request *ComplexRequest, name *string) {
  request.Configs.Extra["name"] = name
  key := tea.String("name")
  request.Configs.Extra[tea.StringValue(key)] = name
  request.Map[tea.StringValue(key)] = name
  request.NumMap[tea.StringValue(key)] = tea.Int(1)
}

func Arrayimport2 (request []*source.Request) (_result *string) {
  s := tea.String("{" + 
"    "a": "test"," + 
"    "b": "ok"" + 
"  }")
  _result = tea.String("")
  return _result
}

func DefaultReturn () (_err error) {
  if true {
  } else {
  }

  return _err
}

