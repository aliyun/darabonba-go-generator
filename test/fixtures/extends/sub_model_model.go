// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iSubModel interface {
  source.iConfig
  String() string
  GoString() string
  SetName(v string) *SubModel
  GetName() *string 
}

type SubModel struct {
  MaxAttemp *int `json:"maxAttemp,omitempty" xml:"maxAttemp,omitempty" require:"true"`
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
}

func (s SubModel) String() string {
  return dara.Prettify(s)
}

func (s SubModel) GoString() string {
  return s.String()
}

func (s *SubModel) GetMaxAttemp() *int  {
  return s.MaxAttemp
}

func (s *SubModel) GetName() *string  {
  return s.Name
}

func (s *SubModel) SetMaxAttemp(v int) *SubModel {
  s.MaxAttemp = &v
  return s
}

func (s *SubModel) SetName(v string) *SubModel {
  s.Name = &v
  return s
}

func (s *SubModel) Validate() error {
  if err := dara.ValidateRequired(s.MaxAttemp, "MaxAttemp"); err != nil {
    return err
  }
  if err := dara.ValidateRequired(s.Name, "Name"); err != nil {
    return err
  }
  return nil
}

