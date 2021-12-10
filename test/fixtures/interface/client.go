// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/tea"
)

type SPIInterface interface {
  STest (a *string) (_result *string) 
  Test (a *string) (_result *string) 
}

type SPI struct {
}

func NewClient()(*SPI, error) {
  client := new(SPI)
  err := client.Init()
  return client, err
}

func (client *SPI)Init()(_err error) {
  return nil
}



func SATest (a *string) (_result *string, _err error) {
  subResourcesMap := make(map[string]*string)
  key := tea.String("key")
  subResourcesMap[tea.StringValue(key)] = tea.String("value")
  _result = a
  return _result , _err
}

func (client *SPI) STest (a *string) (_result *string, _err error) {
  _result = a
  return _result , _err
}

func (client *SPI) Test (a *string) (_result *string) {
  _result = a
  return _result
}

