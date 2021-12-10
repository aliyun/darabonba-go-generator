// This file is auto-generated, don't edit it. Thanks.
package client

import (
  string_  "github.com/aliyun/darabonba-go-generator/test"
  map_  "github.com/aliyun/darabonba-go-generator/test"
  localsource  "github.com/aliyun/darabonba-go-generator"
  "github.com/alibabacloud-go/tea/tea"
)

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



func (client *Client) Sample (client *string_.Client, test *map_.Client) {
  runtime := &string_.RuntimeObject{}
  request := &localsource.Request{
    Accesskey: tea.String("accesskey"),
    Region: tea.String("region"),
  }
  string_.StaticCall()
  client.Print(runtime)
  client.SourceClient.Print(runtime)
}

