// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "net/http"
  "net/url"
  "github.com/alibabacloud-go/tea/dara"
  
)

type iM interface {
  dara.Model
  String() string
  GoString() string
  SetA(v *http.Request) *M
  GetA() *http.Request 
  SetB(v *url.URL) *M
  GetB() *url.URL 
}

type M struct {
  A *http.Request `json:"a,omitempty" xml:"a,omitempty"`
  B *url.URL `json:"b,omitempty" xml:"b,omitempty"`
}

func (s M) String() string {
  return dara.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) GetA() *http.Request  {
  return s.A
}

func (s *M) GetB() *url.URL  {
  return s.B
}

func (s *M) SetA(v *http.Request) *M {
  s.A = v
  return s
}

func (s *M) SetB(v *url.URL) *M {
  s.B = v
  return s
}

func (s *M) Validate() error {
  return dara.Validate(s)
}

