// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
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

