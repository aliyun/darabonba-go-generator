// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
  spi  "github.com/alibabacloud-go/alibabacloud-gateway-spi/client"
  gatewayclient  "github.com/alibabacloud-go/alibabacloud-gateway-sls/client"
  openapi  "github.com/alibabacloud-go/darabonba-openapi/client"
  "github.com/alibabacloud-go/tea/tea"
)

type Client struct {
  openapi.Client
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
  interfaceSPI, _err := gatewayclient.NewClient()
  if _err != nil {
    return _err
  }

  client.Spi =   interfaceSPI.Client
  client.EndpointRule = tea.String("central")
  client.EndpointMap = map[string]*string{
    "cn-hangzhou": tea.String("sls.cn-hangzhou.aliyuncs.com"),
  }
  return nil
}



func (client *Client) Test (project *string, logstore *string) (_result *string, _err error) {
  _result = tea.String("")
  return _result, _err
}

