// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
  "fmt"
)

type iErr1Error interface {
  Error() string
  GetName() *string 
  GetMessage() *string 
  GetCode() *string 
  GetStack() *string 
  GetData() map[string]*string 
}

type Err1Error struct {
  Name *string ``
  Message *string ``
  Code *string ``
  Stack *string ``
  Data map[string]*string ` require:"true"`
}

func (err Err1Error) Error() string {
  if err.Message == nil {
    str := fmt.Sprintf("Err1Error:\n   Name: %s\n   Code: %s\n",
      dara.StringValue(err.Name), dara.StringValue(err.Code))
    err.Message = dara.String(str)
  }
  return dara.StringValue(err.Message)
}

func (s *Err1Error) GetName() *string  {
  return s.Name
}

func (s *Err1Error) GetMessage() *string  {
  return s.Message
}

func (s *Err1Error) GetCode() *string  {
  return s.Code
}

func (s *Err1Error) GetStack() *string  {
  return s.Stack
}

func (s *Err1Error) GetData() map[string]*string  {
  return s.Data
}

