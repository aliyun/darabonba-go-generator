// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
  "fmt"
)

type iERRError interface {
  Error() string
  GetName() *string 
  GetMessage() *string 
  GetCode() *string 
  GetStack() *string 
  GetTest() *int 
}

type ERRError struct {
  Name *string ``
  Message *string ``
  Code *string ``
  Stack *string ``
  Test *int ` require:"true"`
}

func (err ERRError) Error() string {
  if err.Message == nil {
    str := fmt.Sprintf("ERRError:\n   Name: %s\n   Code: %s\n",
      dara.StringValue(err.Name), dara.StringValue(err.Code))
    err.Message = dara.String(str)
  }
  return dara.StringValue(err.Message)
}

func (s *ERRError) GetName() *string  {
  return s.Name
}

func (s *ERRError) GetMessage() *string  {
  return s.Message
}

func (s *ERRError) GetCode() *string  {
  return s.Code
}

func (s *ERRError) GetStack() *string  {
  return s.Stack
}

func (s *ERRError) GetTest() *int  {
  return s.Test
}

type iErrError interface {
  Error() string
  GetName() *string 
  GetMessage() *string 
  GetCode() *string 
  GetStack() *string 
  GetTest() *string 
}

type ErrError struct {
  Name *string ``
  Message *string ``
  Code *string ``
  Stack *string ``
  Test *string ` require:"true"`
}

func (err ErrError) Error() string {
  if err.Message == nil {
    str := fmt.Sprintf("ErrError:\n   Name: %s\n   Code: %s\n",
      dara.StringValue(err.Name), dara.StringValue(err.Code))
    err.Message = dara.String(str)
  }
  return dara.StringValue(err.Message)
}

func (s *ErrError) GetName() *string  {
  return s.Name
}

func (s *ErrError) GetMessage() *string  {
  return s.Message
}

func (s *ErrError) GetCode() *string  {
  return s.Code
}

func (s *ErrError) GetStack() *string  {
  return s.Stack
}

func (s *ErrError) GetTest() *string  {
  return s.Test
}

