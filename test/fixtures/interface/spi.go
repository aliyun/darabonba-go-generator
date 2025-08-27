// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type ClientInterface interface {
  ModifyConfiguration (context *InterceptorContext, attributeMap *AttributeMap) (_err error) 
  ModifyRequest (context *InterceptorContext, attributeMap *AttributeMap) (_err error) 
  ModifyResponse (context *InterceptorContext, attributeMap *AttributeMap) (_err error) 
  Test1 () (_result *string, _err error) 
  Test2 () (_result *int, _err error) 
  Test3 () (_result interface{}, _err error) 
}

type Client struct {
  DisableSDKError *bool
  EnableValidate *bool
}

func NewClient()(*Client, error) {
  client := new(Client)
  err := client.Init()
  return client, err
}

func (client *Client)Init()(_err error) {
  return nil
}



func (client *Client) ModifyConfiguration (context *InterceptorContext, attributeMap *AttributeMap) (_err error) {
  panic("No Support!")
}

func (client *Client) ModifyRequest (context *InterceptorContext, attributeMap *AttributeMap) (_err error) {
  panic("No Support!")
}

func (client *Client) ModifyResponse (context *InterceptorContext, attributeMap *AttributeMap) (_err error) {
  panic("No Support!")
}

func (client *Client) Test1 () (_result *string, _err error) {
  panic("No Support!")
}

func (client *Client) Test2 () (_result *int, _err error) {
  panic("No Support!")
}

func (client *Client) Test3 () (_result interface{}, _err error) {
  panic("No Support!")
}

