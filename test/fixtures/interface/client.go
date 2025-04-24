// This file is auto-generated, don't edit it. Thanks.
package client

import (
  openapi "github.com/alibabacloud-go/darabonba-openapi/client"
  spi "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
  gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-sls/client"
  "github.com/alibabacloud-go/tea/dara"
)

type Client struct {
  openapi.Client
  DisableSDKError *bool
  Client_  spi.Client
  A  *string
  B  *int
}

func NewClient(config *openapi.Config)(*Client, error) {
  client := new(Client)
  err := client.Init(config)
  return client, err
}

func (client *Client)Init(config *openapi.Config)(_err error) {
  _err = client.Client.Init(config  )
  if _err != nil {
    return _err
  }
  client.Client_, _err = gatewayclient.NewClient()
  if _err != nil {
    return _err
  }

  client.Spi = client.Client_
  client.A = dara.String("test")
  client.B = dara.Int(1)
  return nil
}



func (client *Client) Test (project *string, logstore *string) (_result *string, _err error) {
  _result = dara.String("")
  return _result, _err
}

