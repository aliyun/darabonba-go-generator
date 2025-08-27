// This file is auto-generated, don't edit it. Thanks.
package client

import (
  source "github.com/aliyun/darabonba-go-generator/test"
  "github.com/alibabacloud-go/tea/dara"
)

type iValidateModel interface {
  dara.Model
  String() string
  GoString() string
  SetName(v string) *ValidateModel
  GetName() *string 
  SetAge(v int) *ValidateModel
  GetAge() *int 
  SetEmail(v string) *ValidateModel
  GetEmail() *string 
  SetTags(v []*string) *ValidateModel
  GetTags() []*string 
  SetMetadata(v map[string]*string) *ValidateModel
  GetMetadata() map[string]*string 
  GetRequests() 
  SetModelMap(v map[string]*M) *ValidateModel
  GetModelMap() map[string]*M 
  SetRequest(v *source.Request) *ValidateModel
  GetRequest() *source.Request 
  SetM(v *M) *ValidateModel
  GetM() *M 
  SetMm(v *MTest) *ValidateModel
  GetMm() *MTest 
  SetProfile(v *ValidateModelProfile) *ValidateModel
  GetProfile() *ValidateModelProfile 
}

type ValidateModel struct {
  // User name
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true" pattern:"[a-zA-Z][a-zA-Z0-9]*" maxLength:"10" minLength:"2"`
  // User age
  Age *int `json:"age,omitempty" xml:"age,omitempty" minimum:"0" maximum:"150"`
  // User email
  Email *string `json:"email,omitempty" xml:"email,omitempty" pattern:"\w+@\w+\.\w+"`
  // 基本类型数组，无验证逻辑
  Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
  // 基本类型map，无验证逻辑  
  Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
  // 需要验证的数组
  Requests `json:"requests,omitempty" xml:"requests,omitempty" type:"Repeated"`
  // 需要验证的map
  ModelMap map[string]*M `json:"modelMap,omitempty" xml:"modelMap,omitempty"`
  Request *source.Request `json:"request,omitempty" xml:"request,omitempty" require:"true"`
  M *M `json:"m,omitempty" xml:"m,omitempty" require:"true"`
  Mm *MTest `json:"mm,omitempty" xml:"mm,omitempty" require:"true"`
  Profile *ValidateModelProfile `json:"profile,omitempty" xml:"profile,omitempty" type:"Struct"`
}

func (s ValidateModel) String() string {
  return dara.Prettify(s)
}

func (s ValidateModel) GoString() string {
  return s.String()
}

func (s *ValidateModel) GetName() *string  {
  return s.Name
}

func (s *ValidateModel) GetAge() *int  {
  return s.Age
}

func (s *ValidateModel) GetEmail() *string  {
  return s.Email
}

func (s *ValidateModel) GetTags() []*string  {
  return s.Tags
}

func (s *ValidateModel) GetMetadata() map[string]*string  {
  return s.Metadata
}

func (s *ValidateModel) GetRequests()  {
  return s.Requests
}

func (s *ValidateModel) GetModelMap() map[string]*M  {
  return s.ModelMap
}

func (s *ValidateModel) GetRequest() *source.Request  {
  return s.Request
}

func (s *ValidateModel) GetM() *M  {
  return s.M
}

func (s *ValidateModel) GetMm() *MTest  {
  return s.Mm
}

func (s *ValidateModel) GetProfile() *ValidateModelProfile  {
  return s.Profile
}

func (s *ValidateModel) SetName(v string) *ValidateModel {
  s.Name = &v
  return s
}

func (s *ValidateModel) SetAge(v int) *ValidateModel {
  s.Age = &v
  return s
}

func (s *ValidateModel) SetEmail(v string) *ValidateModel {
  s.Email = &v
  return s
}

func (s *ValidateModel) SetTags(v []*string) *ValidateModel {
  s.Tags = v
  return s
}

func (s *ValidateModel) SetMetadata(v map[string]*string) *ValidateModel {
  s.Metadata = v
  return s
}

func (s *ValidateModel) SetModelMap(v map[string]*M) *ValidateModel {
  s.ModelMap = v
  return s
}

func (s *ValidateModel) SetRequest(v *source.Request) *ValidateModel {
  s.Request = v
  return s
}

func (s *ValidateModel) SetM(v *M) *ValidateModel {
  s.M = v
  return s
}

func (s *ValidateModel) SetMm(v *MTest) *ValidateModel {
  s.Mm = v
  return s
}

func (s *ValidateModel) SetProfile(v *ValidateModelProfile) *ValidateModel {
  s.Profile = v
  return s
}

func (s *ValidateModel) Validate() error {
  if err := dara.ValidateRequired(s.Name, "Name"); err != nil {
    return err
  }
  if err := dara.ValidatePattern(s.Name, `[a-zA-Z][a-zA-Z0-9]*`, "Name"); err != nil {
    return err
  }
  if err := dara.ValidateMaxLength(s.Name, 10, "Name"); err != nil {
    return err
  }
  if err := dara.ValidateMinLength(s.Name, 2, "Name"); err != nil {
    return err
  }
  if s.Age != nil {
    if err := dara.ValidateMaximum(s.Age, 150.0, "Age"); err != nil {
      return err
    }
    if err := dara.ValidateMinimum(s.Age, 0.0, "Age"); err != nil {
      return err
    }
  }
  if s.Email != nil {
    if err := dara.ValidatePattern(s.Email, `\w+@\w+\.\w+`, "Email"); err != nil {
      return err
    }
  }
  if s.Requests != nil {
    for _, item := range s.Requests {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if err := dara.ValidateRequired(s.Request, "Request"); err != nil {
    return err
  }
  if s.Request != nil {
    if err := s.Request.Validate(); err != nil {
      return err
    }
  }
  if err := dara.ValidateRequired(s.M, "M"); err != nil {
    return err
  }
  if s.M != nil {
    if err := s.M.Validate(); err != nil {
      return err
    }
  }
  if err := dara.ValidateRequired(s.Mm, "Mm"); err != nil {
    return err
  }
  if s.Mm != nil {
    if err := s.Mm.Validate(); err != nil {
      return err
    }
  }
  if s.Profile != nil {
    if err := s.Profile.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type ValidateModelProfile struct {
  Bio *string `json:"bio,omitempty" xml:"bio,omitempty" maxLength:"500"`
  Score *int `json:"score,omitempty" xml:"score,omitempty" minimum:"0" maximum:"100"`
}

func (s ValidateModelProfile) String() string {
  return dara.Prettify(s)
}

func (s ValidateModelProfile) GoString() string {
  return s.String()
}

func (s *ValidateModelProfile) GetBio() *string  {
  return s.Bio
}

func (s *ValidateModelProfile) GetScore() *int  {
  return s.Score
}

func (s *ValidateModelProfile) SetBio(v string) *ValidateModelProfile {
  s.Bio = &v
  return s
}

func (s *ValidateModelProfile) SetScore(v int) *ValidateModelProfile {
  s.Score = &v
  return s
}

func (s *ValidateModelProfile) Validate() error {
  if s.Bio != nil {
    if err := dara.ValidateMaxLength(s.Bio, 500, "Bio"); err != nil {
      return err
    }
  }
  if s.Score != nil {
    if err := dara.ValidateMaximum(s.Score, 100.0, "Score"); err != nil {
      return err
    }
    if err := dara.ValidateMinimum(s.Score, 0.0, "Score"); err != nil {
      return err
    }
  }
  return nil
}

