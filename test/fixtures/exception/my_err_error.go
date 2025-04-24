// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source "github.com/aliyun/darabonba-go-generator/test"
  "io"
  "github.com/alibabacloud-go/tea/dara"
  "fmt"
)

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

