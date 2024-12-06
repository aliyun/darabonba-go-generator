// This file is auto-generated, don't edit it. Thanks.
package client

import (
  user "darabonba.com/multi/model/user"
  util "darabonba.com/multi/lib/util"
  api "darabonba.com/multi/api"
  "github.com/alibabacloud-go/tea/dara"
  
)

type Client struct {
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



func (client *Client) Test3 () (_result <-chan *string, _err error) {
  _yield := make(chan *string)
  _yieldErr := make(chan error, 1)
  go test3_opYieldFunc(_yield, _yieldErr)
  _result = _yield
  _err = <-_yieldErr
  return _result, _err
}

func (client *Client) Test4 () (_result *int, _err error) {
  api, _err := api.NewClient()
  if _err != nil {
    return nil, _err
  }

  statusTmp, _err := api.Test3()
  status := dara.IntValue(statusTmp)
  if _err != nil {
    return nil, _err
  }

  _result = dara.Int(status)
  return _result , nil
}

func test3_opYieldFunc(_yield chan<- *string, _yieldErr chan<- error) {
  defer close(_yield)
  defer close(_yieldErr)
  it := util.Test1()
  for test := range it {
    _yield <- dara.String(test)
  }
}

