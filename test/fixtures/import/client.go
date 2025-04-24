// This file is auto-generated, don't edit it. Thanks.
package client

import (
  string_ "github.com/aliyun/darabonba-go-generator/test"
  map_ "github.com/aliyun/darabonba-go-generator/test"
  localsource "github.com/aliyun/darabonba-go-generator"
  "github.com/alibabacloud-go/tea/dara"
)

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

