// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source "github.com/aliyun/darabonba-go-generator/test"
  "io"
  "github.com/alibabacloud-go/tea/dara"
)

type iMyModel interface {
  dara.Model
  String() string
  GoString() string
  SetStringfield(v string) *MyModel
  GetStringfield() *string 
  SetStringarrayfield(v []*string) *MyModel
  GetStringarrayfield() []*string 
  SetMapfield(v map[string]*string) *MyModel
  GetMapfield() map[string]*string 
  SetName(v string) *MyModel
  GetName() *string 
  SetSubmodel(v *MyModelSubmodel) *MyModel
  GetSubmodel() *MyModelSubmodel 
  SetModuleModelMap(v map[string]*source.Request) *MyModel
  GetModuleModelMap() map[string]*source.Request 
  SetSubModelMap(v map[string]*MSubM) *MyModel
  GetSubModelMap() map[string]*MSubM 
  SetModelMap(v map[string]*M) *MyModel
  GetModelMap() map[string]*M 
  SetModuleMap(v map[string]*source.Client) *MyModel
  GetModuleMap() map[string]*source.Client 
  SetObject(v map[string]interface{}) *MyModel
  GetObject() map[string]interface{} 
  SetNumberfield(v int) *MyModel
  GetNumberfield() *int 
  SetInt64field(v int64) *MyModel
  GetInt64field() *int64 
  SetUint64field(v uint64) *MyModel
  GetUint64field() *uint64 
  SetInt32field(v int32) *MyModel
  GetInt32field() *int32 
  SetUint32field(v uint32) *MyModel
  GetUint32field() *uint32 
  SetInt16field(v int16) *MyModel
  GetInt16field() *int16 
  SetUint16field(v uint16) *MyModel
  GetUint16field() *uint16 
  SetInt8field(v int8) *MyModel
  GetInt8field() *int8 
  SetUint8field(v uint8) *MyModel
  GetUint8field() *uint8 
  SetReadable(v io.Reader) *MyModel
  GetReadable() io.Reader 
  SetRequest(v *dara.Request) *MyModel
  GetRequest() *dara.Request 
  SetLists(v [][]*string) *MyModel
  GetLists() [][]*string 
  SetArrays(v [][]*MyModelArrays) *MyModel
  GetArrays() [][]*MyModelArrays 
  SetComplexList(v [][]*string) *MyModel
  GetComplexList() [][]*string 
}

type MyModel struct {
  Stringfield *string `json:"stringfield,omitempty" xml:"stringfield,omitempty" require:"true"`
  Stringarrayfield []*string `json:"stringarrayfield,omitempty" xml:"stringarrayfield,omitempty" require:"true" type:"Repeated"`
  Mapfield map[string]*string `json:"mapfield,omitempty" xml:"mapfield,omitempty" require:"true"`
  Name *string `json:"realName,omitempty" xml:"realName,omitempty" require:"true"`
  Submodel *MyModelSubmodel `json:"submodel,omitempty" xml:"submodel,omitempty" require:"true" type:"Struct"`
  ModuleModelMap map[string]*source.Request `json:"moduleModelMap,omitempty" xml:"moduleModelMap,omitempty" require:"true"`
  SubModelMap map[string]*MSubM `json:"subModelMap,omitempty" xml:"subModelMap,omitempty" require:"true"`
  ModelMap map[string]*M `json:"modelMap,omitempty" xml:"modelMap,omitempty" require:"true"`
  ModuleMap map[string]*source.Client `json:"moduleMap,omitempty" xml:"moduleMap,omitempty" require:"true"`
  Object map[string]interface{} `json:"object,omitempty" xml:"object,omitempty" require:"true"`
  Numberfield *int `json:"numberfield,omitempty" xml:"numberfield,omitempty" require:"true"`
  Int64field *int64 `json:"int64field,omitempty" xml:"int64field,omitempty" require:"true"`
  Uint64field *uint64 `json:"uint64field,omitempty" xml:"uint64field,omitempty" require:"true"`
  Int32field *int32 `json:"int32field,omitempty" xml:"int32field,omitempty" require:"true"`
  Uint32field *uint32 `json:"uint32field,omitempty" xml:"uint32field,omitempty" require:"true"`
  Int16field *int16 `json:"int16field,omitempty" xml:"int16field,omitempty" require:"true"`
  Uint16field *uint16 `json:"uint16field,omitempty" xml:"uint16field,omitempty" require:"true"`
  Int8field *int8 `json:"int8field,omitempty" xml:"int8field,omitempty" require:"true"`
  Uint8field *uint8 `json:"uint8field,omitempty" xml:"uint8field,omitempty" require:"true"`
  Readable io.Reader `json:"readable,omitempty" xml:"readable,omitempty" require:"true"`
  Request *dara.Request `json:"request,omitempty" xml:"request,omitempty" require:"true"`
  Lists [][]*string `json:"lists,omitempty" xml:"lists,omitempty" require:"true" type:"Repeated"`
  Arrays [][]*MyModelArrays `json:"arrays,omitempty" xml:"arrays,omitempty" require:"true" type:"Repeated"`
  ComplexList [][]*string `json:"complexList,omitempty" xml:"complexList,omitempty" require:"true" type:"Repeated"`
}

func (s MyModel) String() string {
  return dara.Prettify(s)
}

func (s MyModel) GoString() string {
  return s.String()
}

func (s *MyModel) GetStringfield() *string  {
  return s.Stringfield
}

func (s *MyModel) GetStringarrayfield() []*string  {
  return s.Stringarrayfield
}

func (s *MyModel) GetMapfield() map[string]*string  {
  return s.Mapfield
}

func (s *MyModel) GetName() *string  {
  return s.Name
}

func (s *MyModel) GetSubmodel() *MyModelSubmodel  {
  return s.Submodel
}

func (s *MyModel) GetModuleModelMap() map[string]*source.Request  {
  return s.ModuleModelMap
}

func (s *MyModel) GetSubModelMap() map[string]*MSubM  {
  return s.SubModelMap
}

func (s *MyModel) GetModelMap() map[string]*M  {
  return s.ModelMap
}

func (s *MyModel) GetModuleMap() map[string]*source.Client  {
  return s.ModuleMap
}

func (s *MyModel) GetObject() map[string]interface{}  {
  return s.Object
}

func (s *MyModel) GetNumberfield() *int  {
  return s.Numberfield
}

func (s *MyModel) GetInt64field() *int64  {
  return s.Int64field
}

func (s *MyModel) GetUint64field() *uint64  {
  return s.Uint64field
}

func (s *MyModel) GetInt32field() *int32  {
  return s.Int32field
}

func (s *MyModel) GetUint32field() *uint32  {
  return s.Uint32field
}

func (s *MyModel) GetInt16field() *int16  {
  return s.Int16field
}

func (s *MyModel) GetUint16field() *uint16  {
  return s.Uint16field
}

func (s *MyModel) GetInt8field() *int8  {
  return s.Int8field
}

func (s *MyModel) GetUint8field() *uint8  {
  return s.Uint8field
}

func (s *MyModel) GetReadable() io.Reader  {
  return s.Readable
}

func (s *MyModel) GetRequest() *dara.Request  {
  return s.Request
}

func (s *MyModel) GetLists() [][]*string  {
  return s.Lists
}

func (s *MyModel) GetArrays() [][]*MyModelArrays  {
  return s.Arrays
}

func (s *MyModel) GetComplexList() [][]*string  {
  return s.ComplexList
}

func (s *MyModel) SetStringfield(v string) *MyModel {
  s.Stringfield = &v
  return s
}

func (s *MyModel) SetStringarrayfield(v []*string) *MyModel {
  s.Stringarrayfield = v
  return s
}

func (s *MyModel) SetMapfield(v map[string]*string) *MyModel {
  s.Mapfield = v
  return s
}

func (s *MyModel) SetName(v string) *MyModel {
  s.Name = &v
  return s
}

func (s *MyModel) SetSubmodel(v *MyModelSubmodel) *MyModel {
  s.Submodel = v
  return s
}

func (s *MyModel) SetModuleModelMap(v map[string]*source.Request) *MyModel {
  s.ModuleModelMap = v
  return s
}

func (s *MyModel) SetSubModelMap(v map[string]*MSubM) *MyModel {
  s.SubModelMap = v
  return s
}

func (s *MyModel) SetModelMap(v map[string]*M) *MyModel {
  s.ModelMap = v
  return s
}

func (s *MyModel) SetModuleMap(v map[string]*source.Client) *MyModel {
  s.ModuleMap = v
  return s
}

func (s *MyModel) SetObject(v map[string]interface{}) *MyModel {
  s.Object = v
  return s
}

func (s *MyModel) SetNumberfield(v int) *MyModel {
  s.Numberfield = &v
  return s
}

func (s *MyModel) SetInt64field(v int64) *MyModel {
  s.Int64field = &v
  return s
}

func (s *MyModel) SetUint64field(v uint64) *MyModel {
  s.Uint64field = &v
  return s
}

func (s *MyModel) SetInt32field(v int32) *MyModel {
  s.Int32field = &v
  return s
}

func (s *MyModel) SetUint32field(v uint32) *MyModel {
  s.Uint32field = &v
  return s
}

func (s *MyModel) SetInt16field(v int16) *MyModel {
  s.Int16field = &v
  return s
}

func (s *MyModel) SetUint16field(v uint16) *MyModel {
  s.Uint16field = &v
  return s
}

func (s *MyModel) SetInt8field(v int8) *MyModel {
  s.Int8field = &v
  return s
}

func (s *MyModel) SetUint8field(v uint8) *MyModel {
  s.Uint8field = &v
  return s
}

func (s *MyModel) SetReadable(v io.Reader) *MyModel {
  s.Readable = v
  return s
}

func (s *MyModel) SetRequest(v *dara.Request) *MyModel {
  s.Request = v
  return s
}

func (s *MyModel) SetLists(v [][]*string) *MyModel {
  s.Lists = v
  return s
}

func (s *MyModel) SetArrays(v [][]*MyModelArrays) *MyModel {
  s.Arrays = v
  return s
}

func (s *MyModel) SetComplexList(v [][]*string) *MyModel {
  s.ComplexList = v
  return s
}

func (s *MyModel) Validate() error {
  return dara.Validate(s)
}

type MyModelSubmodel struct {
  Stringfield *string `json:"stringfield,omitempty" xml:"stringfield,omitempty" require:"true"`
}

func (s MyModelSubmodel) String() string {
  return dara.Prettify(s)
}

func (s MyModelSubmodel) GoString() string {
  return s.String()
}

func (s *MyModelSubmodel) GetStringfield() *string  {
  return s.Stringfield
}

func (s *MyModelSubmodel) SetStringfield(v string) *MyModelSubmodel {
  s.Stringfield = &v
  return s
}

func (s *MyModelSubmodel) Validate() error {
  return dara.Validate(s)
}

type MyModelArrays struct {
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s MyModelArrays) String() string {
  return dara.Prettify(s)
}

func (s MyModelArrays) GoString() string {
  return s.String()
}

func (s *MyModelArrays) GetName() *string  {
  return s.Name
}

func (s *MyModelArrays) SetName(v string) *MyModelArrays {
  s.Name = &v
  return s
}

func (s *MyModelArrays) Validate() error {
  return dara.Validate(s)
}

