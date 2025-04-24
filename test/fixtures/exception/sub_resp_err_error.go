// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
  "fmt"
)

type iSubRespErrError interface {
  Error() string
  GetStatusCode() *int 
  GetRetryAfter() *int64 
  GetDescription() *string 
  GetData() map[string]interface{} 
  GetAccessDeniedDetail() map[string]interface{} 
  GetName() *string 
  GetMessage() *string 
  GetCode() *string 
  GetStack() *string 
  GetTestField() *string 
  GetRetryAtfter() *string 
}

type SubRespErrError struct {
  StatusCode *int ``
  RetryAfter *int64 ``
  Description *string ``
  Data map[string]interface{} ``
  AccessDeniedDetail map[string]interface{} ``
  Name *string ``
  Message *string ``
  Code *string ``
  Stack *string ``
  TestField *string ` require:"true"`
  RetryAtfter *string ` require:"true"`
}

func (err SubRespErrError) Error() string {
  if err.Message == nil {
    str := fmt.Sprintf("SubRespErrError:\n   Name: %s\n   Code: %s\n",
      dara.StringValue(err.Name), dara.StringValue(err.Code))
    err.Message = dara.String(str)
  }
  return dara.StringValue(err.Message)
}

func (s *SubRespErrError) GetStatusCode() *int  {
  return s.StatusCode
}

func (s *SubRespErrError) GetRetryAfter() *int64  {
  return s.RetryAfter
}

func (s *SubRespErrError) GetDescription() *string  {
  return s.Description
}

func (s *SubRespErrError) GetData() map[string]interface{}  {
  return s.Data
}

func (s *SubRespErrError) GetAccessDeniedDetail() map[string]interface{}  {
  return s.AccessDeniedDetail
}

func (s *SubRespErrError) GetName() *string  {
  return s.Name
}

func (s *SubRespErrError) GetMessage() *string  {
  return s.Message
}

func (s *SubRespErrError) GetCode() *string  {
  return s.Code
}

func (s *SubRespErrError) GetStack() *string  {
  return s.Stack
}

func (s *SubRespErrError) GetTestField() *string  {
  return s.TestField
}

func (s *SubRespErrError) GetRetryAtfter() *string  {
  return s.RetryAtfter
}

