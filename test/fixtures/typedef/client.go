// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "net/http"
  "net/url"
  oss "github.com/aliyun/darabonba-go-generator"
  "github.com/alibabacloud-go/tea/dara"
)

type Client struct {
  DisableSDKError *bool
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
  _body := TestHttpRequestWith(dara.String("test"), req)
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
  _body := TestHttpHeader(dara.String("test"), headers)
  _result = _body
  return _result, _err
}

