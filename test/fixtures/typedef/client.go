// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "net/http"
  "net/url"
  oss  "github.com/aliyun/darabonba-go-generator"
  "github.com/alibabacloud-go/tea/tea"
)

type M struct {
  A *http.Request `json:"a,omitempty" xml:"a,omitempty"`
  B *url.URL `json:"b,omitempty" xml:"b,omitempty"`
}

func (s M) String() string {
  return tea.Prettify(s)
}

func (s M) GoString() string {
  return s.String()
}

func (s *M) SetA(v *http.Request) *M {
  s.A = v
  return s
}

func (s *M) SetB(v *url.URL) *M {
  s.B = v
  return s
}

type Client struct {
  Vid  *http.Request
  Url  *url.URL
}

func NewClient(request *http.Request, url *url.URL)(*Client, error) {
  client := new(Client)
  err := client.Init(request, url)
  return client, err
}

func (client *Client)Init(request *http.Request, url *url.URL)(_err error) {
  client.Vid = request
  client.Url = url
  return nil
}



func (client *Client) Main (test1 *http.Request, test2 *url.URL) (_err error) {
  oss, _err := oss.NewClient(test1)
  if _err != nil {
    return _err
  }

  m := &M{
    A: test1,
    B: test2,
  }
  client.Vid = test1
  client.Url = test2
  return _err
}

func (client *Client) TestHttpRequest (req *http.Request) (_result *http.Response, _err error) {
  _body := TestHttpRequestWith(tea.String("test"), req)
  _result = _body
  return _result, _err
}

func TestHttpRequestWith (method *string, req *http.Request) (_result *http.Response) {
  panic("No Support!")
}

func TestHttpHeader (method *string, headers *http.Header) (_result *http.Request) {
  panic("No Support!")
}

func (client *Client) TestHttpHeaderWith (headers *http.Header) (_result *http.Request, _err error) {
  _body := TestHttpHeader(tea.String("test"), headers)
  _result = _body
  return _result, _err
}

