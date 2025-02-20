// This file is auto-generated, don't edit it. Thanks.
package client

import (
  string_ "github.com/aliyun/darabonba-go-generator/test"
  map_ "github.com/aliyun/darabonba-go-generator/test"
  localsource "github.com/aliyun/darabonba-go-generator"
  "github.com/alibabacloud-go/tea/dara"
  
)

type iM interface {
  dara.Model
  String() string
  GoString() string
  SetA(v *map_.Request) *M
  GetA() *map_.Request 
  SetB(v *string_.Request) *M
  GetB() *string_.Request 
}

type M struct {
  A *map_.Request `json:"a,omitempty" xml:"a,omitempty"`
  B *string_.Request `json:"b,omitempty" xml:"b,omitempty"`
}

func (s M) String() string {
  return dara.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) GetA() *map_.Request  {
  return s.A
}

func (s *M) GetB() *string_.Request  {
  return s.B
}

func (s *M) SetA(v *map_.Request) *M {
  s.A = v
  return s
}

func (s *M) SetB(v *string_.Request) *M {
  s.B = v
  return s
}

func (s *M) Validate() error {
  return dara.Validate(s)
}

type Client struct {
  DisableSDKError *bool
  SourceClient  *string_.Client
  SourceMap  *map_.Client
}

func NewClient()(*Client, error) {
  client := new(Client)
  err := client.Init()
  return client, err
}

func (client *Client)Init()(_err error) {
  return nil
}



func (client *Client) Sample (str *string_.Client, m *map_.Client) {
  runtime := &string_.RuntimeObject{}
  request := &localsource.Request{
    Accesskey: dara.String("accesskey"),
    Region: dara.String("region"),
  }
  string_.StaticCall()
  map_.StaticCall()
  str.Print(runtime)
  client.SourceClient.Print(runtime)
}

