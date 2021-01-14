// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "io"
  source  "github.com/aliyun/darabonba-go-generator/test"
  "github.com/alibabacloud-go/tea/tea"
)

type M struct {
  SubM *MSubM `json:"subM,omitempty" xml:"subM,omitempty" require:"true" type:"Struct"`
}

func (s M) String() string {
  return tea.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) SetSubM(v *MSubM) *M {
  s.SubM = v
  return s
}

type MSubM struct {
}

func (s MSubM) String() string {
  return tea.Prettify(s)
}

func (s MSubM) GoString() string {
  return s.String()
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
  Readable io.Reader `json:"readable,omitempty" xml:"readable,omitempty" require:"true"`
  Request *tea.Request `json:"request,omitempty" xml:"request,omitempty" require:"true"`
  Lists [][]*string `json:"lists,omitempty" xml:"lists,omitempty" require:"true" type:"Repeated"`
  Arrays [][]*MyModelArrays `json:"arrays,omitempty" xml:"arrays,omitempty" require:"true" type:"Repeated"`
  ComplexList [][]*string `json:"complexList,omitempty" xml:"complexList,omitempty" require:"true" type:"Repeated"`
}

func (s MyModel) String() string {
  return tea.Prettify(s)
}

func (s MyModel) GoString() string {
  return s.String()
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

func (s *MyModel) SetReadable(v io.Reader) *MyModel {
  s.Readable = v
  return s
}

func (s *MyModel) SetRequest(v *tea.Request) *MyModel {
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

type MyModelSubmodel struct {
  Stringfield *string `json:"stringfield,omitempty" xml:"stringfield,omitempty" require:"true"`
}

func (s MyModelSubmodel) String() string {
  return tea.Prettify(s)
}

func (s MyModelSubmodel) GoString() string {
  return s.String()
}

func (s *MyModelSubmodel) SetStringfield(v string) *MyModelSubmodel {
  s.Stringfield = &v
  return s
}

type MyModelArrays struct     {
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s MyModelArrays) String() string {
  return tea.Prettify(s)
}

func (s MyModelArrays) GoString() string {
  return s.String()
}

func (s *MyModelArrays) SetName(v string) *MyModelArrays {
  s.Name = &v
  return s
}


