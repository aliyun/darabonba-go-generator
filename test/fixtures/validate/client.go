// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type client struct {
  DisableSDKError *bool
}

func NewClient()(*client, error) {
  client := new(client)
  err := client.Init()
  return client, err
}

func (client *client)Init()(_err error) {
  return nil
}


func (client *client) ValidateTest() (_err error) {
  request_ := dara.NewRequest()
  // do nothing
  response_, _err := dara.DoRequest(request_, nil)
  if _err != nil {
    return _err
  }


  return nil
}


