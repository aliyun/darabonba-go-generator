// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iAttributeMap interface {
  dara.Model
  String() string
  GoString() string
  SetAttributes(v map[string]interface{}) *AttributeMap
  GetAttributes() map[string]interface{} 
  SetKey(v map[string]*string) *AttributeMap
  GetKey() map[string]*string 
}

type AttributeMap struct {
  Attributes map[string]interface{} `json:"attributes,omitempty" xml:"attributes,omitempty" require:"true"`
  Key map[string]*string `json:"key,omitempty" xml:"key,omitempty" require:"true"`
}

func (s AttributeMap) String() string {
  return dara.Prettify(s)
}

func (s AttributeMap) GoString() string {
  return s.String()
}

func (s *AttributeMap) GetAttributes() map[string]interface{}  {
  return s.Attributes
}

func (s *AttributeMap) GetKey() map[string]*string  {
  return s.Key
}

func (s *AttributeMap) SetAttributes(v map[string]interface{}) *AttributeMap {
  s.Attributes = v
  return s
}

func (s *AttributeMap) SetKey(v map[string]*string) *AttributeMap {
  s.Key = v
  return s
}

func (s *AttributeMap) Validate() error {
  if err := dara.ValidateRequired(s.Attributes, "Attributes"); err != nil {
    return err
  }
  if err := dara.ValidateRequired(s.Key, "Key"); err != nil {
    return err
  }
  return nil
}

