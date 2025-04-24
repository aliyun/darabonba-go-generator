// This file is auto-generated, don't edit it. Thanks.
package user

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iBaseInfo interface {
  darautil.iRuntimeOptions
  String() string
  GoString() string
  SetMaxAttemp(v int) *BaseInfo
  GetMaxAttemp() *int 
}

type BaseInfo struct {
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
  MaxAttemp *int `json:"maxAttemp,omitempty" xml:"maxAttemp,omitempty" require:"true"`
}

func (s BaseInfo) String() string {
  return dara.Prettify(s)
}

func (s BaseInfo) GoString() string {
  return s.String()
}

func (s *BaseInfo) GetAutoretry() *bool  {
  return s.Autoretry
}

func (s *BaseInfo) GetIgnoreSSL() *bool  {
  return s.IgnoreSSL
}

func (s *BaseInfo) GetKey() *string  {
  return s.Key
}

func (s *BaseInfo) GetCert() *string  {
  return s.Cert
}

func (s *BaseInfo) GetCa() *string  {
  return s.Ca
}

func (s *BaseInfo) GetMaxAttempts() *int  {
  return s.MaxAttempts
}

func (s *BaseInfo) GetBackoffPolicy() *string  {
  return s.BackoffPolicy
}

func (s *BaseInfo) GetBackoffPeriod() *int  {
  return s.BackoffPeriod
}

func (s *BaseInfo) GetReadTimeout() *int  {
  return s.ReadTimeout
}

func (s *BaseInfo) GetConnectTimeout() *int  {
  return s.ConnectTimeout
}

func (s *BaseInfo) GetHttpProxy() *string  {
  return s.HttpProxy
}

func (s *BaseInfo) GetHttpsProxy() *string  {
  return s.HttpsProxy
}

func (s *BaseInfo) GetNoProxy() *string  {
  return s.NoProxy
}

func (s *BaseInfo) GetMaxIdleConns() *int  {
  return s.MaxIdleConns
}

func (s *BaseInfo) GetLocalAddr() *string  {
  return s.LocalAddr
}

func (s *BaseInfo) GetSocks5Proxy() *string  {
  return s.Socks5Proxy
}

func (s *BaseInfo) GetSocks5NetWork() *string  {
  return s.Socks5NetWork
}

func (s *BaseInfo) GetKeepAlive() *bool  {
  return s.KeepAlive
}

func (s *BaseInfo) GetMaxAttemp() *int  {
  return s.MaxAttemp
}

func (s *BaseInfo) SetAutoretry(v bool) *BaseInfo {
  s.Autoretry = &v
  return s
}

func (s *BaseInfo) SetIgnoreSSL(v bool) *BaseInfo {
  s.IgnoreSSL = &v
  return s
}

func (s *BaseInfo) SetKey(v string) *BaseInfo {
  s.Key = &v
  return s
}

func (s *BaseInfo) SetCert(v string) *BaseInfo {
  s.Cert = &v
  return s
}

func (s *BaseInfo) SetCa(v string) *BaseInfo {
  s.Ca = &v
  return s
}

func (s *BaseInfo) SetMaxAttempts(v int) *BaseInfo {
  s.MaxAttempts = &v
  return s
}

func (s *BaseInfo) SetBackoffPolicy(v string) *BaseInfo {
  s.BackoffPolicy = &v
  return s
}

func (s *BaseInfo) SetBackoffPeriod(v int) *BaseInfo {
  s.BackoffPeriod = &v
  return s
}

func (s *BaseInfo) SetReadTimeout(v int) *BaseInfo {
  s.ReadTimeout = &v
  return s
}

func (s *BaseInfo) SetConnectTimeout(v int) *BaseInfo {
  s.ConnectTimeout = &v
  return s
}

func (s *BaseInfo) SetHttpProxy(v string) *BaseInfo {
  s.HttpProxy = &v
  return s
}

func (s *BaseInfo) SetHttpsProxy(v string) *BaseInfo {
  s.HttpsProxy = &v
  return s
}

func (s *BaseInfo) SetNoProxy(v string) *BaseInfo {
  s.NoProxy = &v
  return s
}

func (s *BaseInfo) SetMaxIdleConns(v int) *BaseInfo {
  s.MaxIdleConns = &v
  return s
}

func (s *BaseInfo) SetLocalAddr(v string) *BaseInfo {
  s.LocalAddr = &v
  return s
}

func (s *BaseInfo) SetSocks5Proxy(v string) *BaseInfo {
  s.Socks5Proxy = &v
  return s
}

func (s *BaseInfo) SetSocks5NetWork(v string) *BaseInfo {
  s.Socks5NetWork = &v
  return s
}

func (s *BaseInfo) SetKeepAlive(v bool) *BaseInfo {
  s.KeepAlive = &v
  return s
}

func (s *BaseInfo) SetMaxAttemp(v int) *BaseInfo {
  s.MaxAttemp = &v
  return s
}

func (s *BaseInfo) Validate() error {
  return dara.Validate(s)
}

