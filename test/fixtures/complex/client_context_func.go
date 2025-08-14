// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source "github.com/aliyun/darabonba-go-generator/test"
  "io"
  "context"
  "github.com/alibabacloud-go/tea/dara"
)

func (client *Client) Complex1WithCtx(ctx context.Context, request *ComplexRequest, client *source.Client) (_result *source.RuntimeObject, _err error) {
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
    client.PrintWithCtx(ctx, dara.ToMap(request), dara.String("1"))
    go client.PrintSSE(dara.ToMap(request), dara.String("1"))
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
    response_, _err := dara.DoRequestWithCtx(ctx, request_, _runtime)
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

    _err = complex1WithCtx_opResponse(ctx, request, client)
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


func complex1WithCtx_opResponse (ctx *context.Context, request *ComplexRequest, client *Client)( _result *source.RuntimeObject, _err error) {
  if true && true {
    _result = nil
    return _result , _err
  } else if source.JudgeStr(dara.String("test")) || false {
    _result = &source.RuntimeObject{}
    return _result, _err
  }

  client.PrintWithCtx(ctx, dara.ToMap(request), dara.String("1"))
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

