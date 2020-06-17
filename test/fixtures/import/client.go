// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source  "github.com/aliyun/darabonba-go-generator/test"
  "github.com/alibabacloud-go/tea/tea"
)


func Sample (client *source.Client) {
  runtime := &source.RuntimeObject{}
  request := &source.Request{
    Accesskey: tea.String("accesskey"),
    Region: tea.String("region"),
  }
  client.Print(runtime)
}

