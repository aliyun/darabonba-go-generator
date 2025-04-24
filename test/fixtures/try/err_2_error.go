// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
  "fmt"
)

type iErr2Error interface {
  Error() string
  GetName() *string 
  GetMessage() *string 
  GetCode() *string 
  GetStack() *string 
  GetAccessErrMessage() *string 
}

type Err2Error struct {
  Name *string ``
  Message *string ``
  Code *string ``
  Stack *string ``
  AccessErrMessage *string ` require:"true"`
}

func (err Err2Error) Error() string {
  if err.Message == nil {
    str := fmt.Sprintf("Err2Error:\n   Name: %s\n   Code: %s\n",
      dara.StringValue(err.Name), dara.StringValue(err.Code))
    err.Message = dara.String(str)
  }
  return dara.StringValue(err.Message)
}

func (s *Err2Error) GetName() *string  {
  return s.Name
}

func (s *Err2Error) GetMessage() *string  {
  return s.Message
}

func (s *Err2Error) GetCode() *string  {
  return s.Code
}

func (s *Err2Error) GetStack() *string  {
  return s.Stack
}

func (s *Err2Error) GetAccessErrMessage() *string  {
  return s.AccessErrMessage
}

