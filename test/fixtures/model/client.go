// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/tea"
)

type MyModel struct {
  Stringfield *string `json:"stringfield,omitempty" xml:"stringfield,omitempty" require:"true"`
  Stringarrayfield []*string `json:"stringarrayfield,omitempty" xml:"stringarrayfield,omitempty" require:"true" type:"Repeated"`
  Mapfield map[string]*string `json:"mapfield,omitempty" xml:"mapfield,omitempty" require:"true"`
  Name *string `json:"realName,omitempty" xml:"realName,omitempty" require:"true"`
  Submodel *MyModelSubmodel `json:"submodel,omitempty" xml:"submodel,omitempty" require:"true" type:"Struct"`
  Object map[string]interface{} `json:"object,omitempty" xml:"object,omitempty" require:"true"`
  Numberfield *int `json:"numberfield,omitempty" xml:"numberfield,omitempty" require:"true"`
  Readable io.Reader `json:"readable,omitempty" xml:"readable,omitempty" require:"true"`
  Request *tea.Request `json:"request,omitempty" xml:"request,omitempty" require:"true"`
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


