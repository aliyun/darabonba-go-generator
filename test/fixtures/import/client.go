// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source  "github.com/aliyun/darabonba-go-generator/test"
  localsource  "github.com/aliyun/darabonba-go-generator"
  "github.com/alibabacloud-go/tea/tea"
)

type Client struct {
  SourceClient  *source.Client
}

func NewClient()(*Client, error) {
  client := new(Client)
  err := client.Init()
  return client, err
}

func (client *Client)Init()(_err error) {
  return nil
}



func (client *Client) Sample (client *source.Client) {
  runtime := &source.RuntimeObject{}
  request := &localsource.Request{
    Accesskey: tea.String("accesskey"),
    Region: tea.String("region"),
  }
  client.Print(runtime)
  client.SourceClient.Print(runtime)
}

