// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iConfig interface {
  dara.Model
  String() string
  GoString() string
}

type Config struct {
}

func (s Config) String() string {
  return dara.Prettify(s)
}

func (s Config) GoString() string {
  return s.String()
}

func (s *Config) Validate() error {
  return dara.Validate(s)
}

