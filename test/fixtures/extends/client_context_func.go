// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source "github.com/aliyun/darabonba-go-generator/test"
  "context"
  "github.com/alibabacloud-go/tea/dara"
)


func (client *Client) NewModelsCtx (ctx context.Context) (_err error) {
  s := &Sub{
    Name: dara.String("str"),
    Code: dara.String("str"),
    Age: dara.Int(123),
  }
  sm := &SubModel{
    Name: dara.String("str"),
    MaxAttemp: dara.Int(32),
    MaxRetry: dara.Int(32),
  }
  sc := &source.Config{
    MaxAttemp: dara.Int(32),
    MaxRetry: dara.Int(32),
  }
  return _err
}

