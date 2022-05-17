// This file is auto-generated, don't edit it. Thanks.
package client

import (
  string_  "github.com/aliyun/darabonba-go-generator/test"
  map_  "github.com/aliyun/darabonba-go-generator/test"
  localsource  "github.com/aliyun/darabonba-go-generator"
  "github.com/alibabacloud-go/tea/tea"
)

type M struct {
  A *map_.Request `json:"a,omitempty" xml:"a,omitempty"`
  B *string_.Request `json:"b,omitempty" xml:"b,omitempty"`
}

func (s M) String() string {
  return tea.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) SetA(v *map_.Request) *M {
  s.A = v
  return s
}

func (s *M) SetB(v *string_.Request) *M {
  s.B = v
  return s
}

type Client struct {
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
    Accesskey: tea.String("accesskey"),
    Region: tea.String("region"),
  }
  string_.StaticCall()
  map_.StaticCall()
  str.Print(runtime)
  client.SourceClient.Print(runtime)
}

