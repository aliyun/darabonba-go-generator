// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  source "github.com/aliyun/darabonba-go-generator/test"
  "github.com/alibabacloud-go/tea/dara"
  "fmt"
)

type iM interface {
  dara.Model
  String() string
  GoString() string
  SetSubM(v *MSubM) *M
  GetSubM() *MSubM 
}

type M struct {
  SubM *MSubM `json:"subM,omitempty" xml:"subM,omitempty" require:"true" type:"Struct"`
}

func (s M) String() string {
  return dara.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) GetSubM() *MSubM  {
  return s.SubM
}

func (s *M) SetSubM(v *MSubM) *M {
  s.SubM = v
  return s
}

func (s *M) Validate() error {
  return dara.Validate(s)
}

type MSubM struct {
}

func (s MSubM) String() string {
  return dara.Prettify(s)
}

func (s MSubM) GoString() string {
  return s.String()
}

func (s *MSubM) Validate() error {
  return dara.Validate(s)
}

type iMyErrError interface {
  Error() string
  GetMessage() *string 
  GetCode() *string 
  GetStack() *string 
  GetStringfield() *string 
  GetStringarrayfield() []*string 
  GetMapfield() map[string]*string 
  GetName() *int 
  GetSubmodel() *MyErrSubmodel 
  GetModuleModelMap() map[string]*source.Request 
  GetSubModelMap() map[string]*MSubM 
  GetModelMap() map[string]*M 
  GetModuleMap() map[string]*source.Client 
  GetObject() map[string]interface{} 
  GetNumberfield() *int 
  GetInt64field() *int64 
  GetUint64field() *uint64 
  GetInt32field() *int32 
  GetUint32field() *uint32 
  GetInt16field() *int16 
  GetUint16field() *uint16 
  GetInt8field() *int8 
  GetUint8field() *uint8 
  GetReadable() io.Reader 
  GetRequest() *dara.Request 
  GetLists() [][]*string 
  GetArrays() [][]*MyErrArrays 
  GetComplexList() [][]*string 
}

type MyErrError struct {
  Message *string ``
  Code *string ``
  Stack *string ``
  Stringfield *string ` require:"true"`
  Stringarrayfield []*string ` require:"true" type:"Repeated"`
  Mapfield map[string]*string ` require:"true"`
  Name *int ` require:"true"`
  Submodel *MyErrSubmodel ` require:"true" type:"Struct"`
  ModuleModelMap map[string]*source.Request ` require:"true"`
  SubModelMap map[string]*MSubM ` require:"true"`
  ModelMap map[string]*M ` require:"true"`
  ModuleMap map[string]*source.Client ` require:"true"`
  Object map[string]interface{} ` require:"true"`
  Numberfield *int ` require:"true"`
  Int64field *int64 ` require:"true"`
  Uint64field *uint64 ` require:"true"`
  Int32field *int32 ` require:"true"`
  Uint32field *uint32 ` require:"true"`
  Int16field *int16 ` require:"true"`
  Uint16field *uint16 ` require:"true"`
  Int8field *int8 ` require:"true"`
  Uint8field *uint8 ` require:"true"`
  Readable io.Reader ` require:"true"`
  Request *dara.Request ` require:"true"`
  Lists [][]*string ` require:"true" type:"Repeated"`
  Arrays [][]*MyErrArrays ` require:"true" type:"Repeated"`
  ComplexList [][]*string ` require:"true" type:"Repeated"`
}

func (err MyErrError) Error() string {
  if err.Message == nil {
    str := fmt.Sprintf("MyErrError:\n   Name: %s\n   Code: %s\n",
      dara.StringValue(err.Name), dara.StringValue(err.Code))
    err.Message = dara.String(str)
  }
  return dara.StringValue(err.Message)
}

func (s *MyErrError) GetMessage() *string  {
  return s.Message
}

func (s *MyErrError) GetCode() *string  {
  return s.Code
}

func (s *MyErrError) GetStack() *string  {
  return s.Stack
}

func (s *MyErrError) GetStringfield() *string  {
  return s.Stringfield
}

func (s *MyErrError) GetStringarrayfield() []*string  {
  return s.Stringarrayfield
}

func (s *MyErrError) GetMapfield() map[string]*string  {
  return s.Mapfield
}

func (s *MyErrError) GetName() *int  {
  return s.Name
}

func (s *MyErrError) GetSubmodel() *MyErrSubmodel  {
  return s.Submodel
}

func (s *MyErrError) GetModuleModelMap() map[string]*source.Request  {
  return s.ModuleModelMap
}

func (s *MyErrError) GetSubModelMap() map[string]*MSubM  {
  return s.SubModelMap
}

func (s *MyErrError) GetModelMap() map[string]*M  {
  return s.ModelMap
}

func (s *MyErrError) GetModuleMap() map[string]*source.Client  {
  return s.ModuleMap
}

func (s *MyErrError) GetObject() map[string]interface{}  {
  return s.Object
}

func (s *MyErrError) GetNumberfield() *int  {
  return s.Numberfield
}

func (s *MyErrError) GetInt64field() *int64  {
  return s.Int64field
}

func (s *MyErrError) GetUint64field() *uint64  {
  return s.Uint64field
}

func (s *MyErrError) GetInt32field() *int32  {
  return s.Int32field
}

func (s *MyErrError) GetUint32field() *uint32  {
  return s.Uint32field
}

func (s *MyErrError) GetInt16field() *int16  {
  return s.Int16field
}

func (s *MyErrError) GetUint16field() *uint16  {
  return s.Uint16field
}

func (s *MyErrError) GetInt8field() *int8  {
  return s.Int8field
}

func (s *MyErrError) GetUint8field() *uint8  {
  return s.Uint8field
}

func (s *MyErrError) GetReadable() io.Reader  {
  return s.Readable
}

func (s *MyErrError) GetRequest() *dara.Request  {
  return s.Request
}

func (s *MyErrError) GetLists() [][]*string  {
  return s.Lists
}

func (s *MyErrError) GetArrays() [][]*MyErrArrays  {
  return s.Arrays
}

func (s *MyErrError) GetComplexList() [][]*string  {
  return s.ComplexList
}

type MyErrSubmodel struct {
  Stringfield *string `json:"stringfield,omitempty" xml:"stringfield,omitempty" require:"true"`
}

func (s MyErrSubmodel) String() string {
  return dara.Prettify(s)
}

func (s MyErrSubmodel) GoString() string {
  return s.String()
}

func (s *MyErrSubmodel) GetStringfield() *string  {
  return s.Stringfield
}

func (s *MyErrSubmodel) SetStringfield(v string) *MyErrSubmodel {
  s.Stringfield = &v
  return s
}

func (s *MyErrSubmodel) Validate() error {
  return dara.Validate(s)
}

type MyErrArrays struct {
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s MyErrArrays) String() string {
  return dara.Prettify(s)
}

func (s MyErrArrays) GoString() string {
  return s.String()
}

func (s *MyErrArrays) GetName() *string  {
  return s.Name
}

func (s *MyErrArrays) SetName(v string) *MyErrArrays {
  s.Name = &v
  return s
}

func (s *MyErrArrays) Validate() error {
  return dara.Validate(s)
}

type iSubRespErrError interface {
  Error() string
  GetStatusCode() *int 
  GetRetryAfter() *int64 
  GetName() *string 
  GetMessage() *string 
  GetCode() *string 
  GetStack() *string 
  GetTestField() *string 
  GetRetryAtfter() *string 
}

type SubRespErrError struct {
  StatusCode *int ``
  RetryAfter *int64 ``
  Name *string ``
  Message *string ``
  Code *string ``
  Stack *string ``
  TestField *string ` require:"true"`
  RetryAtfter *string ` require:"true"`
}

func (err SubRespErrError) Error() string {
  if err.Message == nil {
    str := fmt.Sprintf("SubRespErrError:\n   Name: %s\n   Code: %s\n",
      dara.StringValue(err.Name), dara.StringValue(err.Code))
    err.Message = dara.String(str)
  }
  return dara.StringValue(err.Message)
}

func (s *SubRespErrError) GetStatusCode() *int  {
  return s.StatusCode
}

func (s *SubRespErrError) GetRetryAfter() *int64  {
  return s.RetryAfter
}

func (s *SubRespErrError) GetName() *string  {
  return s.Name
}

func (s *SubRespErrError) GetMessage() *string  {
  return s.Message
}

func (s *SubRespErrError) GetCode() *string  {
  return s.Code
}

func (s *SubRespErrError) GetStack() *string  {
  return s.Stack
}

func (s *SubRespErrError) GetTestField() *string  {
  return s.TestField
}

func (s *SubRespErrError) GetRetryAtfter() *string  {
  return s.RetryAtfter
}

type iSubMyErrError interface {
  Error() string
  GetStringfield() *string 
  GetStringarrayfield() []*string 
  GetMapfield() map[string]*string 
  GetName() *int 
  GetSubmodel() *SubMyErrSubmodel 
  GetModuleModelMap() map[string]*source.Request 
  GetSubModelMap() map[string]*MSubM 
  GetModelMap() map[string]*M 
  GetModuleMap() map[string]*source.Client 
  GetObject() map[string]interface{} 
  GetNumberfield() *int 
  GetInt64field() *int64 
  GetUint64field() *uint64 
  GetInt32field() *int32 
  GetUint32field() *uint32 
  GetInt16field() *int16 
  GetUint16field() *uint16 
  GetInt8field() *int8 
  GetUint8field() *uint8 
  GetReadable() io.Reader 
  GetRequest() *dara.Request 
  GetLists() [][]*string 
  GetArrays() [][]*SubMyErrArrays 
  GetComplexList() [][]*string 
  GetTestField() *string 
  GetRetryAtfter() *string 
}

type SubMyErrError struct {
  Stringfield *string ` require:"true"`
  Stringarrayfield []*string ` require:"true" type:"Repeated"`
  Mapfield map[string]*string ` require:"true"`
  Name *int ` require:"true"`
  Submodel *SubMyErrSubmodel ` require:"true" type:"Struct"`
  ModuleModelMap map[string]*source.Request ` require:"true"`
  SubModelMap map[string]*MSubM ` require:"true"`
  ModelMap map[string]*M ` require:"true"`
  ModuleMap map[string]*source.Client ` require:"true"`
  Object map[string]interface{} ` require:"true"`
  Numberfield *int ` require:"true"`
  Int64field *int64 ` require:"true"`
  Uint64field *uint64 ` require:"true"`
  Int32field *int32 ` require:"true"`
  Uint32field *uint32 ` require:"true"`
  Int16field *int16 ` require:"true"`
  Uint16field *uint16 ` require:"true"`
  Int8field *int8 ` require:"true"`
  Uint8field *uint8 ` require:"true"`
  Readable io.Reader ` require:"true"`
  Request *dara.Request ` require:"true"`
  Lists [][]*string ` require:"true" type:"Repeated"`
  Arrays [][]*SubMyErrArrays ` require:"true" type:"Repeated"`
  ComplexList [][]*string ` require:"true" type:"Repeated"`
  TestField *string ` require:"true"`
  RetryAtfter *string ` require:"true"`
}

func (err SubMyErrError) Error() string {
  if err.Message == nil {
    str := fmt.Sprintf("SubMyErrError:\n   Name: %s\n   Code: %s\n",
      dara.StringValue(err.Name), dara.StringValue(err.Code))
    err.Message = dara.String(str)
  }
  return dara.StringValue(err.Message)
}

func (s *SubMyErrError) GetStringfield() *string  {
  return s.Stringfield
}

func (s *SubMyErrError) GetStringarrayfield() []*string  {
  return s.Stringarrayfield
}

func (s *SubMyErrError) GetMapfield() map[string]*string  {
  return s.Mapfield
}

func (s *SubMyErrError) GetName() *int  {
  return s.Name
}

func (s *SubMyErrError) GetSubmodel() *SubMyErrSubmodel  {
  return s.Submodel
}

func (s *SubMyErrError) GetModuleModelMap() map[string]*source.Request  {
  return s.ModuleModelMap
}

func (s *SubMyErrError) GetSubModelMap() map[string]*MSubM  {
  return s.SubModelMap
}

func (s *SubMyErrError) GetModelMap() map[string]*M  {
  return s.ModelMap
}

func (s *SubMyErrError) GetModuleMap() map[string]*source.Client  {
  return s.ModuleMap
}

func (s *SubMyErrError) GetObject() map[string]interface{}  {
  return s.Object
}

func (s *SubMyErrError) GetNumberfield() *int  {
  return s.Numberfield
}

func (s *SubMyErrError) GetInt64field() *int64  {
  return s.Int64field
}

func (s *SubMyErrError) GetUint64field() *uint64  {
  return s.Uint64field
}

func (s *SubMyErrError) GetInt32field() *int32  {
  return s.Int32field
}

func (s *SubMyErrError) GetUint32field() *uint32  {
  return s.Uint32field
}

func (s *SubMyErrError) GetInt16field() *int16  {
  return s.Int16field
}

func (s *SubMyErrError) GetUint16field() *uint16  {
  return s.Uint16field
}

func (s *SubMyErrError) GetInt8field() *int8  {
  return s.Int8field
}

func (s *SubMyErrError) GetUint8field() *uint8  {
  return s.Uint8field
}

func (s *SubMyErrError) GetReadable() io.Reader  {
  return s.Readable
}

func (s *SubMyErrError) GetRequest() *dara.Request  {
  return s.Request
}

func (s *SubMyErrError) GetLists() [][]*string  {
  return s.Lists
}

func (s *SubMyErrError) GetArrays() [][]*SubMyErrArrays  {
  return s.Arrays
}

func (s *SubMyErrError) GetComplexList() [][]*string  {
  return s.ComplexList
}

func (s *SubMyErrError) GetTestField() *string  {
  return s.TestField
}

func (s *SubMyErrError) GetRetryAtfter() *string  {
  return s.RetryAtfter
}

type SubMyErrSubmodel struct {
  Stringfield *string `json:"stringfield,omitempty" xml:"stringfield,omitempty" require:"true"`
}

func (s SubMyErrSubmodel) String() string {
  return dara.Prettify(s)
}

func (s SubMyErrSubmodel) GoString() string {
  return s.String()
}

func (s *SubMyErrSubmodel) GetStringfield() *string  {
  return s.Stringfield
}

func (s *SubMyErrSubmodel) SetStringfield(v string) *SubMyErrSubmodel {
  s.Stringfield = &v
  return s
}

func (s *SubMyErrSubmodel) Validate() error {
  return dara.Validate(s)
}

type SubMyErrArrays struct {
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s SubMyErrArrays) String() string {
  return dara.Prettify(s)
}

func (s SubMyErrArrays) GoString() string {
  return s.String()
}

func (s *SubMyErrArrays) GetName() *string  {
  return s.Name
}

func (s *SubMyErrArrays) SetName(v string) *SubMyErrArrays {
  s.Name = &v
  return s
}

func (s *SubMyErrArrays) Validate() error {
  return dara.Validate(s)
}


