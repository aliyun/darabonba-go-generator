// This file is auto-generated, don't edit it. Thanks.
package user

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iInfo interface {
  iBaseInfo
  String() string
  GoString() string
  SetName(v string) *Info
  GetName() *string 
  SetAge(v int) *Info
  GetAge() *int 
}

type Info struct {
  MaxAttemp *int `json:"maxAttemp,omitempty" xml:"maxAttemp,omitempty" require:"true"`
  // whether to try again
  Autoretry *bool `json:"autoretry,omitempty" xml:"autoretry,omitempty"`
  // ignore SSL validation
  IgnoreSSL *bool `json:"ignoreSSL,omitempty" xml:"ignoreSSL,omitempty"`
  // privite key for client certificate
  Key *string `json:"key,omitempty" xml:"key,omitempty"`
  // client certificate
  Cert *string `json:"cert,omitempty" xml:"cert,omitempty"`
  // server certificate
  Ca *string `json:"ca,omitempty" xml:"ca,omitempty"`
  // maximum number of retries
  MaxAttempts *int `json:"max_attempts,omitempty" xml:"max_attempts,omitempty"`
  // backoff policy
  BackoffPolicy *string `json:"backoff_policy,omitempty" xml:"backoff_policy,omitempty"`
  // backoff period
  BackoffPeriod *int `json:"backoff_period,omitempty" xml:"backoff_period,omitempty"`
  // read timeout
  ReadTimeout *int `json:"readTimeout,omitempty" xml:"readTimeout,omitempty"`
  // connect timeout
  ConnectTimeout *int `json:"connectTimeout,omitempty" xml:"connectTimeout,omitempty"`
  // http proxy url
  HttpProxy *string `json:"httpProxy,omitempty" xml:"httpProxy,omitempty"`
  // https Proxy url
  HttpsProxy *string `json:"httpsProxy,omitempty" xml:"httpsProxy,omitempty"`
  // agent blacklist
  NoProxy *string `json:"noProxy,omitempty" xml:"noProxy,omitempty"`
  // maximum number of connections
  MaxIdleConns *int `json:"maxIdleConns,omitempty" xml:"maxIdleConns,omitempty"`
  // local addr
  LocalAddr *string `json:"localAddr,omitempty" xml:"localAddr,omitempty"`
  // SOCKS5 proxy
  Socks5Proxy *string `json:"socks5Proxy,omitempty" xml:"socks5Proxy,omitempty"`
  // SOCKS5 netWork
  Socks5NetWork *string `json:"socks5NetWork,omitempty" xml:"socks5NetWork,omitempty"`
  // whether to enable keep-alive
  KeepAlive *bool `json:"keepAlive,omitempty" xml:"keepAlive,omitempty"`
  Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
  Age *int `json:"age,omitempty" xml:"age,omitempty" require:"true"`
}

func (s Info) String() string {
  return dara.Prettify(s)
}

func (s Info) GoString() string {
  return s.String()
}

func (s *Info) GetMaxAttemp() *int  {
  return s.MaxAttemp
}

func (s *Info) GetAutoretry() *bool  {
  return s.Autoretry
}

func (s *Info) GetIgnoreSSL() *bool  {
  return s.IgnoreSSL
}

func (s *Info) GetKey() *string  {
  return s.Key
}

func (s *Info) GetCert() *string  {
  return s.Cert
}

func (s *Info) GetCa() *string  {
  return s.Ca
}

func (s *Info) GetMaxAttempts() *int  {
  return s.MaxAttempts
}

func (s *Info) GetBackoffPolicy() *string  {
  return s.BackoffPolicy
}

func (s *Info) GetBackoffPeriod() *int  {
  return s.BackoffPeriod
}

func (s *Info) GetReadTimeout() *int  {
  return s.ReadTimeout
}

func (s *Info) GetConnectTimeout() *int  {
  return s.ConnectTimeout
}

func (s *Info) GetHttpProxy() *string  {
  return s.HttpProxy
}

func (s *Info) GetHttpsProxy() *string  {
  return s.HttpsProxy
}

func (s *Info) GetNoProxy() *string  {
  return s.NoProxy
}

func (s *Info) GetMaxIdleConns() *int  {
  return s.MaxIdleConns
}

func (s *Info) GetLocalAddr() *string  {
  return s.LocalAddr
}

func (s *Info) GetSocks5Proxy() *string  {
  return s.Socks5Proxy
}

func (s *Info) GetSocks5NetWork() *string  {
  return s.Socks5NetWork
}

func (s *Info) GetKeepAlive() *bool  {
  return s.KeepAlive
}

func (s *Info) GetName() *string  {
  return s.Name
}

func (s *Info) GetAge() *int  {
  return s.Age
}

func (s *Info) SetMaxAttemp(v int) *Info {
  s.MaxAttemp = &v
  return s
}

func (s *Info) SetAutoretry(v bool) *Info {
  s.Autoretry = &v
  return s
}

func (s *Info) SetIgnoreSSL(v bool) *Info {
  s.IgnoreSSL = &v
  return s
}

func (s *Info) SetKey(v string) *Info {
  s.Key = &v
  return s
}

func (s *Info) SetCert(v string) *Info {
  s.Cert = &v
  return s
}

func (s *Info) SetCa(v string) *Info {
  s.Ca = &v
  return s
}

func (s *Info) SetMaxAttempts(v int) *Info {
  s.MaxAttempts = &v
  return s
}

func (s *Info) SetBackoffPolicy(v string) *Info {
  s.BackoffPolicy = &v
  return s
}

func (s *Info) SetBackoffPeriod(v int) *Info {
  s.BackoffPeriod = &v
  return s
}

func (s *Info) SetReadTimeout(v int) *Info {
  s.ReadTimeout = &v
  return s
}

func (s *Info) SetConnectTimeout(v int) *Info {
  s.ConnectTimeout = &v
  return s
}

func (s *Info) SetHttpProxy(v string) *Info {
  s.HttpProxy = &v
  return s
}

func (s *Info) SetHttpsProxy(v string) *Info {
  s.HttpsProxy = &v
  return s
}

func (s *Info) SetNoProxy(v string) *Info {
  s.NoProxy = &v
  return s
}

func (s *Info) SetMaxIdleConns(v int) *Info {
  s.MaxIdleConns = &v
  return s
}

func (s *Info) SetLocalAddr(v string) *Info {
  s.LocalAddr = &v
  return s
}

func (s *Info) SetSocks5Proxy(v string) *Info {
  s.Socks5Proxy = &v
  return s
}

func (s *Info) SetSocks5NetWork(v string) *Info {
  s.Socks5NetWork = &v
  return s
}

func (s *Info) SetKeepAlive(v bool) *Info {
  s.KeepAlive = &v
  return s
}

func (s *Info) SetName(v string) *Info {
  s.Name = &v
  return s
}

func (s *Info) SetAge(v int) *Info {
  s.Age = &v
  return s
}

func (s *Info) Validate() error {
  return dara.Validate(s)
}

