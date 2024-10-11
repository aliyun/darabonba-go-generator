// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea"
  "io"
  source "github.com/aliyun/darabonba-go-generator/test"
  dara "github.com/alibabacloud-go/tea/tea"
  "fmt"
)

type iComplexRequest interface {
  dara.Model
  String() string
  GoString() string
  SetAccessKey(v string) *ComplexRequest
  GetAccessKey() *string 
  SetBody(v io.Reader) *ComplexRequest
  GetBody() io.Reader 
  SetStrs(v []*string) *ComplexRequest
  GetStrs() []*string 
  SetMapList(v []map[string]interface{}) *ComplexRequest
  GetMapList() []map[string]interface{} 
  SetHeader(v *ComplexRequestHeader) *ComplexRequest
  GetHeader() *ComplexRequest 
  SetConfigs(v *ComplexRequestConfigs) *ComplexRequest
  GetConfigs() *ComplexRequest 
  SetNum(v int) *ComplexRequest
  GetNum() *int 
  SetI64(v int64) *ComplexRequest
  GetI64() *int64 
  SetF64(v float64) *ComplexRequest
  GetF64() *float64 
  SetB(v bool) *ComplexRequest
  GetB() *bool 
  SetF32(v float32) *ComplexRequest
  GetF32() *float32 
  SetF64List(v []*float64) *ComplexRequest
  GetF64List() []*float64 
  SetFloatList(v []*float32) *ComplexRequest
  GetFloatList() []*float32 
  SetBooleantList(v []*bool) *ComplexRequest
  GetBooleantList() []*bool 
  SetI32(v int32) *ComplexRequest
  GetI32() *int32 
  SetStringList(v []*string) *ComplexRequest
  GetStringList() []*string 
  SetIntList(v []*int) *ComplexRequest
  GetIntList() []*int 
  SetInt32List(v []*int32) *ComplexRequest
  GetInt32List() []*int32 
  SetInt16List(v []*int16) *ComplexRequest
  GetInt16List() []*int16 
  SetInt64List(v []*int64) *ComplexRequest
  GetInt64List() []*int64 
  SetLongList(v []*int64) *ComplexRequest
  GetLongList() []*int64 
  SetUint64List(v []*uint64) *ComplexRequest
  GetUint64List() []*uint64 
  SetUint32List(v []*uint32) *ComplexRequest
  GetUint32List() []*uint32 
  SetUint16List(v []*uint16) *ComplexRequest
  GetUint16List() []*uint16 
  SetU64(v uint64) *ComplexRequest
  GetU64() *uint64 
  SetU32(v uint32) *ComplexRequest
  GetU32() *uint32 
  SetU16(v uint16) *ComplexRequest
  GetU16() *uint16 
  SetObj(v map[string]interface{}) *ComplexRequest
  GetObj() map[string]interface{} 
  SetAny(v interface{}) *ComplexRequest
  GetAny() interface{} 
  SetByt(v []byte) *ComplexRequest
  GetByt() []byte 
  SetReq(v *dara.Request) *ComplexRequest
  GetReq() *dara.Request 
  SetResp(v *dara.Response) *ComplexRequest
  GetResp() *dara.Response 
  SetMap(v map[string]*string) *ComplexRequest
  GetMap() map[string]*string 
  SetNumMap(v map[string]*int) *ComplexRequest
  GetNumMap() map[string]*int 
  SetModelMap(v map[string]*source.Request) *ComplexRequest
  GetModelMap() map[string]*source.Request 
  SetRequest(v *source.Request) *ComplexRequest
  GetRequest() *source.Request 
  SetClient(v *source.Client) *ComplexRequest
  GetClient() *source.Client 
  SetInstance(v *source.RequestInstance) *ComplexRequest
  GetInstance() *source.RequestInstance 
  SetPart(v []*ComplexRequestPart) *ComplexRequest
  GetPart() []*ComplexRequest 
}

type ComplexRequest struct {
  dara.Model
  AccessKey *string `json:"accessKey,omitempty" xml:"accessKey,omitempty" require:"true"`
  // Body
  // 
  // example:
  // 
  // Body
  Body io.Reader `json:"Body,omitempty" xml:"Body,omitempty" require:"true"`
  // Strs
  // 
  // example:
  // 
  // Strs
  Strs []*string `json:"Strs,omitempty" xml:"Strs,omitempty" require:"true" type:"Repeated"`
  // mapList
  // 
  // example:
  // 
  // mapList
  MapList []map[string]interface{} `json:"mapList,omitempty" xml:"mapList,omitempty" require:"true" type:"Repeated"`
  // header
  Header *ComplexRequestHeader `json:"header,omitempty" xml:"header,omitempty" require:"true" type:"Struct"`
  Configs *ComplexRequestConfigs `json:"configs,omitempty" xml:"configs,omitempty" require:"true" type:"Struct"`
  // check if is blank:
  // false
  // 
  // if can be null:
  // false
  // 
  // if sensitive:
  // true
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
  LongList []*int64 `json:"longList,omitempty" xml:"longList,omitempty" require:"true" type:"Repeated"`
  Uint64List []*uint64 `json:"uint64List,omitempty" xml:"uint64List,omitempty" require:"true" type:"Repeated"`
  Uint32List []*uint32 `json:"uint32List,omitempty" xml:"uint32List,omitempty" require:"true" type:"Repeated"`
  Uint16List []*uint16 `json:"uint16List,omitempty" xml:"uint16List,omitempty" require:"true" type:"Repeated"`
  U64 *uint64 `json:"u64,omitempty" xml:"u64,omitempty" require:"true"`
  U32 *uint32 `json:"u32,omitempty" xml:"u32,omitempty" require:"true"`
  U16 *uint16 `json:"u16,omitempty" xml:"u16,omitempty" require:"true"`
  Obj map[string]interface{} `json:"obj,omitempty" xml:"obj,omitempty" require:"true"`
  Any interface{} `json:"any,omitempty" xml:"any,omitempty" require:"true"`
  Byt []byte `json:"byt,omitempty" xml:"byt,omitempty" require:"true"`
  Req *dara.Request `json:"req,omitempty" xml:"req,omitempty" require:"true"`
  Resp *dara.Response `json:"resp,omitempty" xml:"resp,omitempty" require:"true"`
  Map map[string]*string `json:"map,omitempty" xml:"map,omitempty" require:"true"`
  NumMap map[string]*int `json:"numMap,omitempty" xml:"numMap,omitempty" require:"true"`
  ModelMap map[string]*source.Request `json:"modelMap,omitempty" xml:"modelMap,omitempty" require:"true"`
  Request *source.Request `json:"request,omitempty" xml:"request,omitempty" require:"true"`
  Client *source.Client `json:"client,omitempty" xml:"client,omitempty" require:"true"`
  Instance *source.RequestInstance `json:"instance,omitempty" xml:"instance,omitempty" require:"true"`
  // Deprecated
  // 
  // Part
  Part []*ComplexRequestPart `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s ComplexRequest) String() string {
  return dara.Prettify(s)
}

func (s ComplexRequest) GoString() string {
  return s.String()
}

func (s *ComplexRequest) GetAccessKey() *string  {
  return s.AccessKey
}

func (s *ComplexRequest) GetBody() io.Reader  {
  return s.Body
}

func (s *ComplexRequest) GetStrs() []*string  {
  return s.Strs
}

func (s *ComplexRequest) GetMapList() []map[string]interface{}  {
  return s.MapList
}

func (s *ComplexRequest) GetHeader() *ComplexRequest  {
  return s.Header
}

func (s *ComplexRequest) GetConfigs() *ComplexRequest  {
  return s.Configs
}

func (s *ComplexRequest) GetNum() *int  {
  return s.Num
}

func (s *ComplexRequest) GetI64() *int64  {
  return s.I64
}

func (s *ComplexRequest) GetF64() *float64  {
  return s.F64
}

func (s *ComplexRequest) GetB() *bool  {
  return s.B
}

func (s *ComplexRequest) GetF32() *float32  {
  return s.F32
}

func (s *ComplexRequest) GetF64List() []*float64  {
  return s.F64List
}

func (s *ComplexRequest) GetFloatList() []*float32  {
  return s.FloatList
}

func (s *ComplexRequest) GetBooleantList() []*bool  {
  return s.BooleantList
}

func (s *ComplexRequest) GetI32() *int32  {
  return s.I32
}

func (s *ComplexRequest) GetStringList() []*string  {
  return s.StringList
}

func (s *ComplexRequest) GetIntList() []*int  {
  return s.IntList
}

func (s *ComplexRequest) GetInt32List() []*int32  {
  return s.Int32List
}

func (s *ComplexRequest) GetInt16List() []*int16  {
  return s.Int16List
}

func (s *ComplexRequest) GetInt64List() []*int64  {
  return s.Int64List
}

func (s *ComplexRequest) GetLongList() []*int64  {
  return s.LongList
}

func (s *ComplexRequest) GetUint64List() []*uint64  {
  return s.Uint64List
}

func (s *ComplexRequest) GetUint32List() []*uint32  {
  return s.Uint32List
}

func (s *ComplexRequest) GetUint16List() []*uint16  {
  return s.Uint16List
}

func (s *ComplexRequest) GetU64() *uint64  {
  return s.U64
}

func (s *ComplexRequest) GetU32() *uint32  {
  return s.U32
}

func (s *ComplexRequest) GetU16() *uint16  {
  return s.U16
}

func (s *ComplexRequest) GetObj() map[string]interface{}  {
  return s.Obj
}

func (s *ComplexRequest) GetAny() interface{}  {
  return s.Any
}

func (s *ComplexRequest) GetByt() []byte  {
  return s.Byt
}

func (s *ComplexRequest) GetReq() *dara.Request  {
  return s.Req
}

func (s *ComplexRequest) GetResp() *dara.Response  {
  return s.Resp
}

func (s *ComplexRequest) GetMap() map[string]*string  {
  return s.Map
}

func (s *ComplexRequest) GetNumMap() map[string]*int  {
  return s.NumMap
}

func (s *ComplexRequest) GetModelMap() map[string]*source.Request  {
  return s.ModelMap
}

func (s *ComplexRequest) GetRequest() *source.Request  {
  return s.Request
}

func (s *ComplexRequest) GetClient() *source.Client  {
  return s.Client
}

func (s *ComplexRequest) GetInstance() *source.RequestInstance  {
  return s.Instance
}

func (s *ComplexRequest) GetPart() []*ComplexRequest  {
  return s.Part
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

func (s *ComplexRequest) SetLongList(v []*int64) *ComplexRequest {
  s.LongList = v
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

func (s *ComplexRequest) SetReq(v *dara.Request) *ComplexRequest {
  s.Req = v
  return s
}

func (s *ComplexRequest) SetResp(v *dara.Response) *ComplexRequest {
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
  // The ID of the security group to which you want to assign the instance. Instances in the same security group can communicate with each other. The maximum number of instances that a security group can contain depends on the type of the security group. For more information, see the "Security group limits" section in [Limits](https://help.aliyun.com/document_detail/25412.html#SecurityGroupQuota).
  // 
  // 	Notice:  The network type of the new instance must be the same as that of the security group specified by the `SecurityGroupId` parameter. For example, if the specified security group is of the VPC type, the new instance is also of the VPC type and you must specify `VSwitchId`.
  // 
  // If you do not use `LaunchTemplateId` or `LaunchTemplateName` to specify a launch template, you must specify SecurityGroupId. Take note of the following items:
  // 
  // 	- You can set `SecurityGroupId` to specify a single security group or set `SecurityGroupIds.N` to specify one or more security groups. However, you cannot specify both `SecurityGroupId` and `SecurityGroupIds.N`.
  // 
  // 	- If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot specify `SecurityGroupId` or `SecurityGroupIds.N` but can specify `NetworkInterface.N.SecurityGroupId` or `NetworkInterface.N.SecurityGroupIds.N`.
  // 
  // check if is blank:
  // true
  // 
  // if can be null:
  // true
  // 
  // if sensitive:
  // false
  // 
  // example:
  // 
  // Content
  // 
  // test example
  // 
  // test example11
  Content *string `json:"Content,omitempty" xml:"Content,omitempty" require:"true" signed:"true"`
}

func (s ComplexRequestHeader) String() string {
  return dara.Prettify(s)
}

func (s ComplexRequestHeader) GoString() string {
  return s.String()
}

func (s *ComplexRequestHeader) GetContent() *string  {
  return s.Content
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
  return dara.Prettify(s)
}

func (s ComplexRequestConfigs) GoString() string {
  return s.String()
}

func (s *ComplexRequestConfigs) GetKey() *string  {
  return s.Key
}

func (s *ComplexRequestConfigs) GetValue() []*string  {
  return s.Value
}

func (s *ComplexRequestConfigs) GetExtra() map[string]*string  {
  return s.Extra
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

type ComplexRequestPart struct {
  // PartNumber
  PartNumber *string `json:"PartNumber,omitempty" xml:"PartNumber,omitempty"`
}

func (s ComplexRequestPart) String() string {
  return dara.Prettify(s)
}

func (s ComplexRequestPart) GoString() string {
  return s.String()
}

func (s *ComplexRequestPart) GetPartNumber() *string  {
  return s.PartNumber
}

func (s *ComplexRequestPart) SetPartNumber(v string) *ComplexRequestPart {
  s.PartNumber = &v
  return s
}

type iResponse interface {
  dara.Model
  String() string
  GoString() string
  SetInstance(v *ComplexRequestPart) *Response
  GetInstance() *ComplexRequestPart 
}

type Response struct {
  dara.Model
  Instance *ComplexRequestPart `json:"instance,omitempty" xml:"instance,omitempty" require:"true"`
}

func (s Response) String() string {
  return dara.Prettify(s)
}

func (s Response) GoString() string {
  return s.String()
}

func (s *Response) GetInstance() *ComplexRequestPart  {
  return s.Instance
}

func (s *Response) SetInstance(v *ComplexRequestPart) *Response {
  s.Instance = v
  return s
}

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
    if _t, ok := _err.(dara.BaseError); ok {
    }
  }
  return _err
}

func TestTryWithComplexReturnType () (_result *source.Request, _err error) {
  strTmp, _err := client.TemplateString()
  str := dara.StringValue(strTmp)
  final := "ok"
  if _err != nil {
    if _t, ok := _err.(dara.BaseError); ok {
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
    if _t, ok := _err.(dara.BaseError); ok {
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

func (client *Client) ReturnFunc () (_result *string) {
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
  if true {
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
    if _t, ok := _err.(*Err1); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
    }
    if _t, ok := _err.(*Err2); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
    }
    if _t, ok := _err.(*source.Err3); ok {
      err := _t;
      fmt.Printf("[LOG] %s\n", dara.StringValue(err.Name))
    }
    if _t, ok := _err.(dara.BaseError); ok {
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
  } else if true || false {
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
    _err = &Err1{
      Name: dara.String("str"),
      Code: dara.String("str"),
      Data: map[string]*string{
        "key1": dara.String("str"),
      },
    }
    return _err
  } else if dara.IntValue(a) == 0 {
    _err = &Err2{
      Name: dara.String("str"),
      Code: dara.String("str"),
      AccessErrMessage: dara.String("str2"),
    }
    return _err
  } else if dara.IntValue(a) == -10 {
    _err = &source.Err3{
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

