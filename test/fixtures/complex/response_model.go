// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iResponse interface {
  dara.Model
  String() string
  GoString() string
  SetInstance(v *ComplexRequestPart) *Response
  GetInstance() *ComplexRequestPart 
}

type Response struct {
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

func (s *Response) Validate() error {
  return dara.Validate(s)
}

