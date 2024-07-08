// This file is auto-generated, don't edit it. Thanks.
package user

import (
  darautil "github.com/alibabacloud-go/tea-utils/v2/service"
  util "darabonba.com/multi/lib/util"
  dara "github.com/alibabacloud-go/tea/tea"
  
)

type Info struct {
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  Age *int `json:"age,omitempty" xml:"age,omitempty" require:"true"`
}

func (s Info) String() string {
  return dara.Prettify(s)
}

func (s Info) GoString() string {
  return s.String()
}

func (s *Info) SetName(v string) *Info {
  s.Name = &v
  return s
}

func (s *Info) SetAge(v int) *Info {
  s.Age = &v
  return s
}


func Test () (_result <-chan *string, _err error) {
  _yield := make(chan *string)
  _yieldErr := make(chan error, 1)
  go test_opYieldFunc(_yield, _yieldErr)
  _result = _yield
  _err = <-_yieldErr
  return _result, _err
}

func test_opYieldFunc(_yield chan<- *string, _yieldErr chan<- error) {
  defer close(_yield)
  defer close(_yieldErr)
  it := util.Test1()
  for test := range it {
    _yield <- dara.String(test)
  }
}

