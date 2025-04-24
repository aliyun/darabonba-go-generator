// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea"
  source "github.com/aliyun/darabonba-go-generator/test"
  "io"
  "github.com/alibabacloud-go/tea/dara"
  "fmt"
)

type Client struct {
  DisableSDKError *bool
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
  test, _err := source.NewClient(config)
  if _err != nil {
    return _err
  }

  return nil
}


func (client *Client) Complex1(request *ComplexRequest, client *source.Client) (_result *source.RuntimeObject, _err error) {
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

  _result = new(source.RuntimeObject)
  for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
    _resultErr = nil
    _backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
    dara.Sleep(_backoffDelayTime)

    request_ = dara.NewRequest()
    name := "complex"
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
    inte := 1
    a := 1
    var b int32
    b = int32(a)
    c := a
    IntToInt32(dara.ToInt32(dara.Int(a)))
    mapVal := map[string]interface{}{
      "read": read,
      "test": "{\\"test\\":\\"ok\\"}",
      "b": dara.BoolValue(request.B),
      "num": dara.IntValue(request.Num),
      "u16": dara.Uint16Value(request.U16),
      "u32": dara.Uint32Value(request.U32),
      "u64": dara.Uint64Value(request.U64),
      "u16List": dara.Uint16SliceValue(request.Uint16List),
      "u32List": dara.Uint32SliceValue(request.Uint32List),
      "u64List": dara.Uint64SliceValue(request.Uint64List),
      "i64List": dara.Int64SliceValue(request.Int64List),
      "i16List": dara.Int16SliceValue(request.Int16List),
      "i32List": dara.Int32SliceValue(request.Int32List),
      "intList": dara.IntSliceValue(request.IntList),
      "stringList": dara.StringSliceValue(request.StringList),
      "i32": dara.Int32Value(request.I32),
      "booleantList": dara.BoolSliceValue(request.BooleantList),
      "floatList": dara.Float32SliceValue(request.FloatList),
      "float64List": dara.Float64SliceValue(request.F64List),
      "f32": dara.Float32Value(request.F32),
      "f64": dara.Float64Value(request.F64),
      "i64": dara.Int64Value(request.I64),
    }
    req := &ComplexRequest{
      B: dara.Bool(false),
      Num: dara.Int(10),
      I32: dara.ToInt32(dara.Int(a)),
      IntList: []*int{dara.Int(10), dara.Int(11)},
      Int16List: []*int16{dara.Int16(10), dara.Int16(11)},
      Int32List: []*int32{dara.Int32(10), dara.Int32(11)},
      Int64List: []*int64{dara.Int64(10), dara.Int64(11)},
      LongList: []*int64{dara.Int64(10), dara.Int64(11)},
      FloatList: []*float32{dara.Float32(0.1), dara.Float32(0.2)},
      StringList: []*string{dara.String("10"), dara.String("11")},
      BooleantList: []*bool{dara.Bool(true), dara.Bool(false)},
    }
    longList := []*int64{dara.Int64(432435)}
    anyList := []interface{}{dara.Int64(432435), dara.String("str"), dara.Bool(true), dara.Int(10), dara.Float32(0.1)}
    floatMap := map[string]*float32{
      "key1": dara.Float32(0.1),
      "key2": dara.Float32(0.2),
    }
    doubleMap := map[string]*float64{
      "key1": dara.Float64(0.1),
      "key2": dara.Float64(0.2),
    }
    intMap := map[string]*int{
      "key1": dara.Int(1),
      "key2": dara.Int(2),
    }
    longMap := map[string]*int64{
      "key1": dara.Int64(1),
      "key2": dara.Int64(2),
    }
    int16Map := map[string]*int16{
      "key1": dara.Int16(1),
      "key2": dara.Int16(2),
    }
    int32Map := map[string]*int32{
      "key1": dara.Int32(1),
      "key2": dara.Int32(2),
    }
    int64Map := map[string]*int64{
      "key1": dara.Int64(1),
      "key2": dara.Int64(2),
    }
    anyMap := map[string]interface{}{
      "key1": 0.1,
      "key2": 1,
      "key3": "test",
      "key4": true,
      "key5": []interface{}{dara.String("test"), dara.Int(1), dara.Bool(true), []*string{dara.String("test")}},
      "key6": []map[string]interface{}{map[string]interface{}{
          "a": "test",
          "b": 1,
          "c": true,
          "d": []*string{dara.String("test")},
        }},
    }
    for _, item := range []*string{dara.String("1"), dara.String("2")} {
      anyMap := make(map[string]*string)
      anyMap[item] = dara.String("test")
      break
    }
    client.Strs = request.Strs
    client.Protocol = dara.String("test")
    dara.StringValue(client.EndpointMap[dara.StringValue(client.Protocol)])
    client.EndpointMap["test"] = dara.String("ok")
    request.Strs = client.Strs
    request_.Protocol = client.Protocol
    request_.Port = request.Num
    request_.Method = dara.String("GET")
    request_.Pathname = dara.String("/" + dara.StringValue(client.Pathname))
    request_.Query = map[string]*string{
      "date": dara.String("2019"),
      "name": request_.Method,
    }
    tmp := dara.ToMap(request_.Query,
      request_.Headers,
      request_)
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

    _err = complex1_opResponse(request, client)
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

func (client *Client) Complex2(request *ComplexRequest, str []*string, val map[string]*string) (_result map[string]interface{}, _err error) {
  request_ := dara.NewRequest()
  name := "complex"
  config := &source.Config{}
  client, _err := source.NewClient(config)
  if _err != nil {
    return _result, _err
  }

  configArray := []*source.Config{config}
  request_.Protocol = dara.String("HTTP")
  request_.Port = dara.Int(80)
  request_.Method = dara.String("GET")
  request_.Pathname = dara.String("/")
  request_.Query = map[string]*string{
    "date": dara.String("2019"),
    "protocol": request_.Protocol,
  }
  response_, _err := dara.DoRequest(request_, nil)
  if _err != nil {
    return nil, _err
  }


  return nil, nil
}

func (client *Client) ComplexMap() (_result map[string]interface{}, _err error) {
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


    return nil, nil
  }
  if dara.BoolValue(client.DisableSDKError) != true {
    _resultErr = dara.TeaSDKError(_resultErr)
  }
  return _result, _resultErr
}

func (client *Client) Complex3(request *ComplexRequest, name *string) (_result *ComplexRequest, _err error) {
  request_ := dara.NewRequest()
  name = dara.String("complex")
  request_.Protocol, _err = client.TemplateString()
  if _err != nil {
    return _result, _err
  }

  request_.Port = dara.Int(80)
  request_.Method = dara.String("GET")
  request_.Pathname = dara.String("/")
  request_.Body = dara.ToReader(dara.String("body"))
  request_.Query = map[string]*string{
    "date": dara.String("2019"),
  }
  var tmp *ComplexRequest
  tmp = client.ReturnModel()
  name = client.Protocol
  response_, _err := dara.DoRequest(request_, nil)
  if _err != nil {
    return nil, _err
  }

  _result, _err = complex3_opResponse(response_, request, client)
  if _err != nil {
    return nil, _err
  }

  return _result, nil
}

func (client *Client) NoReturn() (_err error) {
  request_ := dara.NewRequest()
  response_, _err := dara.DoRequest(request_, nil)
  if _err != nil {
    return _err
  }


  return nil
}


func (client *Client) Hello (request map[string]interface{}, strs []*string, complexList [][]*string) (_result []*string, _err error) {
  var a [][]*string
  _body := Array1()
  _result = _body
  return _result, _err
}

func Print (reqeust *dara.Request, reqs []*ComplexRequest, response *dara.Response, val map[string]*string) (_result *source.Request, _err error) {
  panic("No Support!")
}

func IntToInt32 (a *int32) {
  panic("No Support!")
}

func AssignWithArray () (_err error) {
  var list []*string
  list = []*string{dara.String("test")}
  var str string
  strTmp, _err := client.ThrowsFunc()
  str = dara.StringValue(strTmp)
  if _err != nil {
    return _err
  }

  return _err
}

func (client *Client) MapAcess () {
  tmp := map[string]interface{}{
    "protocol": dara.StringValue(client.EndpointMap[dara.StringValue(client.Protocol)]),
  }
}

func (client *Client) ExprFunc () (_result []*string, _err error) {
  if !true {
  }

  num := 10
  req := &ComplexRequest{}
  mapVal := map[string]interface{}{
    "num": 10,
    "client": new(Source),
    "strs": Array1(),
    "str": "string" + dara.ToString(num),
    "str1": "string" + dara.StringValue(req.AccessKey),
  }
  _result = nil
  return _result , _err
}

func PrintNull () (_err error) {
  strTmp, _err := client.TemplateString()
  str := dara.StringValue(strTmp)
  final := "ok"
  if _err != nil {
    if _t, ok := _err.(*dara.SDKError); ok {
    }
  }
  return _err
}

func TestTryWithComplexReturnType () (_result *source.Request, _err error) {
  strTmp, _err := client.TemplateString()
  str := dara.StringValue(strTmp)
  final := "ok"
  if _err != nil {
    if _t, ok := _err.(*dara.SDKError); ok {
    }
  }
  _result = nil
  return _result , _err
}

func TestTryWithComplexReturnTypeWithOutCat () (_result *source.Request, _err error) {
  strTmp, _err := client.TemplateString()
  str := dara.StringValue(strTmp)
  final := "ok"
  if _err != nil {
    if _t, ok := _err.(*dara.SDKError); ok {
      e := _t;
      sim := "a"
    }
  }
  _result = nil
  return _result , _err
}

func Array0 (req map[string]interface{}) (_result []interface{}) {
  _result = []interface{}{}
  return _result
}

func Array1 () (_result []*string) {
  _result = []*string{dara.String("1")}
  return _result
}

func (client *Client) TemplateString () (_result *string, _err error) {
  _result = dara.String("/" + dara.StringValue(client.Protocol))
  return _result, _err
}

func (client *Client) IntOp (a *int) {
  b := dara.IntValue(a)
  b++
  ++b
  b--
  --b
}

func (client *Client) ThrowsFunc () (_result *string, _err error) {
  _result = dara.String("/" + dara.StringValue(client.Protocol))
  return _result, _err
}

func (client *Client) ThrowsFunc1 () (_result *string, _err error) {
  _result = dara.String("")
  return _result, _err
}

func (client *Client) ThrowsFunc2 () (_err error) {
  _err = dara.NewSDKError(map[string]interface{}{
    "code": "",
  })
  return _err
}

func (client *Client) ThrowsFunc3 () (_result *string, _err error) {
  _err = dara.NewSDKError(map[string]interface{}{
    "code": "",
  })
  return _result, _err
}

func (client *Client) GetInt (num *int32) (_result *int32) {
  _result = num
  return _result
}

func (client *Client) ReturnFunc () (_result *string) {
  index := int32(0)
  i := dara.Int32Value(client.GetInt(dara.Int32(index)))
  _result = nil
  return _result
}

func (client *Client) ReturnFunc1 (cfg *source.Config) (_result *source.Client) {
  config := &source.Config{}
  _result, _err = source.NewClient(config)
  return _result
}

func (client *Client) ReturnFunc2 () (_result map[string]interface{}) {
  tmp := map[string]*string{
    "subMap": dara.String("ok"),
  }
  mapVal := map[string]map[string]*string{
    "test": tmp,
  }
  if dara.BoolValue(source.JudgeStr(dara.String("test"))) {
    _result = mapVal["test"]
    return _result
  } else {
    var body io.Reader
    _result.Body = body
    dara.Convert(dara.ToMap(tmp), &_result)

    return _result
  }

}

func (client *Client) ReturnModel () (_result *ComplexRequest) {
  return _result
}

func (client *Client) EmptyFunc () {
  panic("No Support!")
}

func (client *Client) Error (e dara.BaseError) (_result dara.BaseError) {
  var tmp dara.BaseError
  var c interface{}
  _result = e
  return _result
}

func ArrayAccess () (_result *string) {
  configs := []*string{dara.String("a"), dara.String("b"), dara.String("c")}
  config := dara.StringValue(configs[0])
  _result = dara.String(config)
  return _result
}

func ArrayAccess2 () (_result *string) {
  data := map[string][]*string{
    "configs": []*string{dara.String("a"), dara.String("b"), dara.String("c")},
  }
  config := dara.StringValue(data["configs"][0])
  _result = dara.String(config)
  return _result
}

func ArrayAccess3 (request *ComplexRequest) (_result *string) {
  req := &source.Request{}
  ArrayAccess4([]*source.Request{req})
  configVal := dara.StringValue(request.Configs.Value[0])
  _result = dara.String(configVal)
  return _result
}

func ArrayAccess4 (requests []*source.Request) (_result *string) {
  _result = dara.String("")
  return _result
}

func ArrayAssign (config *string) (_result []*string) {
  configs := []*string{dara.String("a"), dara.String("b"), dara.String("c")}
  configs[3] = config
  _result = configs
  return _result
}

func ArrayAssign2 (config *string) (_result []*string) {
  data := map[string][]*string{
    "configs": []*string{dara.String("a"), dara.String("b"), dara.String("c")},
  }
  data["configs"][3] = config
  _result = data["configs"]
  return _result
}

func ArrayAssign3 (request *ComplexRequest, config *string) {
  request.Configs.Value[0] = config
}

func MapAccess (request *ComplexRequest) (_result *string) {
  configInfo := dara.StringValue(request.Configs.Extra["name"])
  _result = dara.String(configInfo)
  return _result
}

func MapAccess2 (request *source.Request) (_result *string) {
  configInfo := dara.StringValue(request.Configs.Extra["name"])
  _result = dara.String(configInfo)
  return _result
}

func MapAccess3 () (_result *string) {
  data := map[string]map[string]*string{
    "configs": map[string]*string{
      "value": dara.String("string"),
    },
  }
  _result = data["configs"]["value"]
  return _result
}

func MapAccess4 (request *ComplexRequest) (_result *string) {
  key := "name"
  model := request.ModelMap[key]
  configInfo := dara.StringValue(request.Configs.Extra[key])
  _result = dara.String(configInfo)
  return _result
}

func MapAssign (request *ComplexRequest, name *string) {
  request.Configs.Extra["name"] = name
  key := "name"
  name = dara.String(key)
  request.Configs.Extra[key] = name
  name = request.Configs.Extra["name"]
  request.Map[key] = name
  request.NumMap[key] = dara.Int(1)
}

func Arrayimport2 (request []*source.Request) (_result *string) {
  s := "{" + 
"    \"a\": \"test\"," + 
"    \"b\": \"ok\"" + 
"  }"
  _result = dara.String("")
  return _result
}

func DefaultReturn () (_err error) {
  if true {
  } else {
  }

  return _err
}

func (client *Client) MultiTryCatch (a *int) (_err error) {
  _result, _err  = multiTryCatch_opTryFunc(a)
  final := "ok"
  if _err != nil {
    if _t, ok := _err.(*Err1Error); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
    }
    if _t, ok := _err.(*Err2Error); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
    }
    if _t, ok := _err.(*source.Err3Error); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
    }
    if _t, ok := _err.(*dara.SDKError); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
    }
  }
  return _err
}

func complex1_opResponse (request *ComplexRequest, client *Client)( _result *source.RuntimeObject, _err error) {
  if true && true {
    _result = nil
    return _result , _err
  } else if source.JudgeStr(dara.String("test")) || false {
    return _result, _err
  }

  client.Print(dara.ToMap(request), dara.String("1"))
  _, _err = client.Hello(dara.ToMap(request), []*string{dara.String("1"), dara.String("2")}, nil)
  if _err != nil {
    return _result, _err
  }
  _, _err = client.Hello(nil, nil, nil)
  if _err != nil {
    return _result, _err
  }
  _err = dara.Convert(map[string]interface{}{}, &_result)

  return _result, _err
  _, _err = client.Complex3(nil, dara.String("test"))
  if _err != nil {
    return _result, _err
  }
  _result = nil
  return _result , _err
}

func complex3_opResponse (response_ *dara.Response, request *ComplexRequest, client *Client)( _result *ComplexRequest, _err error) {
  resp := response_
  req := &source.Request{
    Accesskey: request.AccessKey,
    Region: resp.StatusMessage,
  }
  Array0(dara.ToMap(request))
  req.Accesskey = dara.String("accesskey")
  req.Accesskey = request.AccessKey
  _err = PrintNull()
  if _err != nil {
    return _result, _err
  }
  _, _err = client.ThrowsFunc()
  if _err != nil {
    return _result, _err
  }
  dara.IntValue(response_.StatusCode)
  source.Array(dara.ToMap(request), dara.String("1"))
  _err = dara.Convert(dara.Merge(request_.Query), &_result)

  return _result, _err
}

func multiTryCatch_opTryFunc (a *int)(_err error) {
  if dara.IntValue(a) > 0 {
    _err = &Err1Error{
      Name: dara.String("str"),
      Code: dara.String("str"),
      Data: map[string]*string{
        "key1": dara.String("str"),
      },
    }
    return _err
  } else if dara.IntValue(a) == 0 {
    _err = &Err2Error{
      Name: dara.String("str"),
      Code: dara.String("str"),
      AccessErrMessage: dara.String("str2"),
    }
    return _err
  } else if dara.IntValue(a) == -10 {
    _err = &source.Err3Error{
      Name: dara.String("str"),
      Code: dara.String("str"),
    }
    return _err
  } else {
    _err = &dara.SDKError{
      Name: dara.String("str"),
      Code: dara.String("str"),
    }
    return _err
  }

  return _err
}

