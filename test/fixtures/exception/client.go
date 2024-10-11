// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  source "github.com/aliyun/darabonba-go-generator/test"
  dara "github.com/alibabacloud-go/tea/tea"
  "fmt"
)

type iM interface {
  dara.Model
  String() string
  GoString() string
  SetSubM(v *MSubM) *M
  GetSubM() *M 
}

type M struct {
  dara.Model
  SubM *MSubM `json:"subM,omitempty" xml:"subM,omitempty" require:"true" type:"Struct"`
}

func (s M) String() string {
  return dara.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) GetSubM() *M  {
  return s.SubM
}

func (s *M) SetSubM(v *MSubM) *M {
  s.SubM = v
  return s
}

type MSubM struct {
}

func (s MSubM) String() string {
  return dara.Prettify(s)
}

func (s MSubM) GoString() string {
  return s.String()
}

type iMyErr interface {
  dara.BaseError
  GetStringfield() *string 
  GetStringarrayfield() []*string 
  GetMapfield() map[string]*string 
  GetName() *int 
  GetSubmodel() *MyErr 
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
  GetArrays() [][]*MyErr 
  GetComplexList() [][]*string 
}

type MyErr struct {
  dara.BaseError
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

func (err MyErr) Error() string {
  if err.Message == nil {
    str := fmt.Sprintf("MyErr:\n   Name: %s\n   Code: %s\n",
      dara.StringValue(err.Name), dara.StringValue(err.Code))
    err.Message = dara.String(str)
  }
  return dara.StringValue(err.Message)
}

func (s *MyErr) GetMessage() *string  {
  return s.Message
}

func (s *MyErr) GetCode() *string  {
  return s.Code
}

func (s *MyErr) GetStack() *string  {
  return s.Stack
}

func (s *MyErr) GetStringfield() *string  {
  return s.Stringfield
}

func (s *MyErr) GetStringarrayfield() []*string  {
  return s.Stringarrayfield
}

func (s *MyErr) GetMapfield() map[string]*string  {
  return s.Mapfield
}

func (s *MyErr) GetName() *int  {
  return s.Name
}

func (s *MyErr) GetSubmodel() *MyErr  {
  return s.Submodel
}

func (s *MyErr) GetModuleModelMap() map[string]*source.Request  {
  return s.ModuleModelMap
}

func (s *MyErr) GetSubModelMap() map[string]*MSubM  {
  return s.SubModelMap
}

func (s *MyErr) GetModelMap() map[string]*M  {
  return s.ModelMap
}

func (s *MyErr) GetModuleMap() map[string]*source.Client  {
  return s.ModuleMap
}

func (s *MyErr) GetObject() map[string]interface{}  {
  return s.Object
}

func (s *MyErr) GetNumberfield() *int  {
  return s.Numberfield
}

func (s *MyErr) GetInt64field() *int64  {
  return s.Int64field
}

func (s *MyErr) GetUint64field() *uint64  {
  return s.Uint64field
}

func (s *MyErr) GetInt32field() *int32  {
  return s.Int32field
}

func (s *MyErr) GetUint32field() *uint32  {
  return s.Uint32field
}

func (s *MyErr) GetInt16field() *int16  {
  return s.Int16field
}

func (s *MyErr) GetUint16field() *uint16  {
  return s.Uint16field
}

func (s *MyErr) GetInt8field() *int8  {
  return s.Int8field
}

func (s *MyErr) GetUint8field() *uint8  {
  return s.Uint8field
}

func (s *MyErr) GetReadable() io.Reader  {
  return s.Readable
}

func (s *MyErr) GetRequest() *dara.Request  {
  return s.Request
}

func (s *MyErr) GetLists() [][]*string  {
  return s.Lists
}

func (s *MyErr) GetArrays() [][]*MyErr  {
  return s.Arrays
}

func (s *MyErr) GetComplexList() [][]*string  {
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

type iSubRespErr interface {
  dara.IResponseError
  GetTestField() *string 
  GetRetryAtfter() *string 
}

type SubRespErr struct {
  dara.IResponseError
  StatusCode *int ``
  RetryAfter *int ``
  Name *string ``
  Message *string ``
  Code *string ``
  Stack *string ``
  TestField *string ` require:"true"`
  RetryAtfter *string ` require:"true"`
}

func (err SubRespErr) Error() string {
  if err.Message == nil {
    str := fmt.Sprintf("SubRespErr:\n   Name: %s\n   Code: %s\n",
      dara.StringValue(err.Name), dara.StringValue(err.Code))
    err.Message = dara.String(str)
  }
  return dara.StringValue(err.Message)
}

func (s *SubRespErr) GetStatusCode() *int  {
  return s.StatusCode
}

func (s *SubRespErr) GetRetryAfter() *int  {
  return s.RetryAfter
}

func (s *SubRespErr) GetName() *string  {
  return s.Name
}

func (s *SubRespErr) GetMessage() *string  {
  return s.Message
}

func (s *SubRespErr) GetCode() *string  {
  return s.Code
}

func (s *SubRespErr) GetStack() *string  {
  return s.Stack
}

func (s *SubRespErr) GetTestField() *string  {
  return s.TestField
}

func (s *SubRespErr) GetRetryAtfter() *string  {
  return s.RetryAtfter
}

type iSubMyErr interface {
  iMyErr
  GetTestField() *string 
  GetRetryAtfter() *string 
}

type SubMyErr struct {
  iMyErr
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

func (err SubMyErr) Error() string {
  if err.Message == nil {
    str := fmt.Sprintf("SubMyErr:\n   Name: %s\n   Code: %s\n",
      dara.StringValue(err.Name), dara.StringValue(err.Code))
    err.Message = dara.String(str)
  }
  return dara.StringValue(err.Message)
}

func (s *SubMyErr) GetStringfield() *string  {
  return s.Stringfield
}

func (s *SubMyErr) GetStringarrayfield() []*string  {
  return s.Stringarrayfield
}

func (s *SubMyErr) GetMapfield() map[string]*string  {
  return s.Mapfield
}

func (s *SubMyErr) GetName() *int  {
  return s.Name
}

func (s *SubMyErr) GetSubmodel() *SubMyErr  {
  return s.Submodel
}

func (s *SubMyErr) GetModuleModelMap() map[string]*source.Request  {
  return s.ModuleModelMap
}

func (s *SubMyErr) GetSubModelMap() map[string]*MSubM  {
  return s.SubModelMap
}

func (s *SubMyErr) GetModelMap() map[string]*M  {
  return s.ModelMap
}

func (s *SubMyErr) GetModuleMap() map[string]*source.Client  {
  return s.ModuleMap
}

func (s *SubMyErr) GetObject() map[string]interface{}  {
  return s.Object
}

func (s *SubMyErr) GetNumberfield() *int  {
  return s.Numberfield
}

func (s *SubMyErr) GetInt64field() *int64  {
  return s.Int64field
}

func (s *SubMyErr) GetUint64field() *uint64  {
  return s.Uint64field
}

func (s *SubMyErr) GetInt32field() *int32  {
  return s.Int32field
}

func (s *SubMyErr) GetUint32field() *uint32  {
  return s.Uint32field
}

func (s *SubMyErr) GetInt16field() *int16  {
  return s.Int16field
}

func (s *SubMyErr) GetUint16field() *uint16  {
  return s.Uint16field
}

func (s *SubMyErr) GetInt8field() *int8  {
  return s.Int8field
}

func (s *SubMyErr) GetUint8field() *uint8  {
  return s.Uint8field
}

func (s *SubMyErr) GetReadable() io.Reader  {
  return s.Readable
}

func (s *SubMyErr) GetRequest() *dara.Request  {
  return s.Request
}

func (s *SubMyErr) GetLists() [][]*string  {
  return s.Lists
}

func (s *SubMyErr) GetArrays() [][]*SubMyErr  {
  return s.Arrays
}

func (s *SubMyErr) GetComplexList() [][]*string  {
  return s.ComplexList
}

func (s *SubMyErr) GetTestField() *string  {
  return s.TestField
}

func (s *SubMyErr) GetRetryAtfter() *string  {
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


