// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  source "github.com/aliyun/darabonba-go-generator/test"
  "github.com/alibabacloud-go/tea/dara"
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
  GetHeader() *ComplexRequestHeader 
  SetConfigs(v *ComplexRequestConfigs) *ComplexRequest
  GetConfigs() *ComplexRequestConfigs 
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
  GetPart() []*ComplexRequestPart 
}

type ComplexRequest struct {
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

func (s *ComplexRequest) GetHeader() *ComplexRequestHeader  {
  return s.Header
}

func (s *ComplexRequest) GetConfigs() *ComplexRequestConfigs  {
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

func (s *ComplexRequest) GetPart() []*ComplexRequestPart  {
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

func (s *ComplexRequest) Validate() error {
  return dara.Validate(s)
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

func (s *ComplexRequestHeader) Validate() error {
  return dara.Validate(s)
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

func (s *ComplexRequestConfigs) Validate() error {
  return dara.Validate(s)
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

func (s *ComplexRequestPart) Validate() error {
  return dara.Validate(s)
}

