// This file is auto-generated, don't edit it. Thanks.
package client

import (
  user "darabonba.com/multi/model/user"
  api "darabonba.com/multi/api"
  util "darabonba.com/multi/lib/util"
  "github.com/alibabacloud-go/tea/dara"
)

type Client struct {
  DisableSDKError *bool
  EnableValidate *bool
  User  *user.Info
}

func NewClient()(*Client, error) {
  client := new(Client)
  err := client.Init()
  return client, err
}

func (client *Client)Init()(_err error) {
  client.User = &user.Info{
    Name: dara.String("test"),
    Age: dara.Int(124),
    MaxAttemp: dara.Int(3),
    Autoretry: dara.Bool(true),
  }
  return nil
}



func (client *Client) Test3 (_yield chan *string, _yieldErr chan error) {
  defer close(_yield)
  client.test3_opYieldFunc(_yield, _yieldErr)
  return
}

func (client *Client) Test4 () (_result *int, _err error) {
  api, _err := api.NewClient()
  if _err != nil {
    return _result, _err
  }

  statusTmp, _err := api.Test3()
  status := dara.IntValue(statusTmp)
  if _err != nil {
    return _result, _err
  }

  _result = dara.Int(status)
  return _result , _err
}

func (client *Client) test3_opYieldFunc(_yield chan *string, _yieldErr chan error) {
  it := make(chan string, 1)
  util.Test1(it)
  for test := range it {
    _yield <- dara.String(test)
  }
}

