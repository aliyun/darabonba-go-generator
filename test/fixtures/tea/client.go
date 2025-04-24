// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source "github.com/aliyun/darabonba-go-generator"
  localsource "github.com/aliyun/darabonba-go-generator"
  "github.com/alibabacloud-go/tea/dara"
)


func Sample (client *source.Client) {
  runtime := &source.RuntimeObject{}
  request := &localsource.Request{
    Accesskey: dara.String("accesskey"),
    Region: dara.String("region"),
  }
  client.Print(runtime)
}

