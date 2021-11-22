// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/tea"
)

type SPIInterface interface {
  STest (a *string) (_result *string) {
  Test (a *string) (_result *string) {
}

type Client struct {
}

func NewClient()(*Client, error) {
  client := new(Client)
  err := client.Init()
  return client, err
}

func (client *Client)Init()(_err error) {
  return nil
}



func SATest (a *string) (_result *string, _err error) {
  _result = a
  return _result , _err
}

func (client *Client) STest (a *string) (_result *string, _err error) {
  _result = a
  return _result , _err
}

func (client *Client) Test (a *string) (_result *string) {
  _result = a
  return _result
}

