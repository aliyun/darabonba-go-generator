// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type Client struct {
  DisableSDKError *bool
  EnableValidate *bool
  // type's comment
  A  []*string
}

// Description:
// 
// Init Func
// comment between init and annotation
func NewClient()(*Client, error) {
  client := new(Client)
  err := client.Init()
  return client, err
}

func (client *Client)Init()(_err error) {
  // string declate comment
  str := "sss"
  // new model instance comment
  modelInstance := &Test1{
    Test: dara.String("test"),
    //test declare back comment
    Test2: dara.String("test2"),
  }
  array := []interface{}{
    // array string comment
    dara.String("string"), 
    // array number comment
    dara.Int(300)
    // array back comment
  }
  return nil
}


// Description:
// 
// testAPI
//testAPI comment one
//testAPI comment two
func (client *Client) TestAPI() (_err error) {
  _runtime := dara.NewRuntimeObject(map[string]interface{}{
    // empty runtime comment
    // another runtime comment
  })

  var retryPolicyContext *dara.RetryPolicyContext
  var request_ *dara.Request
  var response_ *dara.Response
  var _resultErr error
  retriesAttempted := int(0)
  retryPolicyContext = &dara.RetryPolicyContext{
    RetriesAttempted: retriesAttempted,
  }

  for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
    _resultErr = nil
    _backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
    dara.Sleep(_backoffDelayTime)

    request_ = dara.NewRequest()
    // new model instance comment
    modelInstance := &Test1{
      // test declare front comment
      Test: dara.String("test"),
      // test2 declare front comment
      Test2: dara.String("test2"),
    }
    // number declare comment
    num := 123
    // static function call comment
    StaticFunc()
    response_, _err = dara.DoRequest(request_, _runtime)
    if _err != nil {
      retriesAttempted++
      retryPolicyContext = &dara.RetryPolicyContext{
        RetriesAttempted: retriesAttempted,
        HttpRequest:      request_,
        HttpResponse:     response_,
        Exception:        _err,
      }
      _resultErr = _err
      continue
    }

    // static async function call
    _err = TestFunc()
    if _err != nil {
      retriesAttempted++
      retryPolicyContext = &dara.RetryPolicyContext{
        RetriesAttempted: retriesAttempted,
        HttpRequest:      request_,
        HttpResponse:     response_,
        Exception:        _err,
      }
      _resultErr = _err
      continue
    }

    // return comment
    return _err
  }
  if dara.BoolValue(client.DisableSDKError) != true {
    _resultErr = dara.TeaSDKError(_resultErr)
  }
  return _result, _resultErr
}

// testAPI2 comment
func (client *Client) TestAPI2() (_err error) {
  _runtime := dara.NewRuntimeObject(map[string]interface{}{
    // runtime retry comment
    "retry": true,
    // runtime back comment one
    // runtime back comment two
  })

  var retryPolicyContext *dara.RetryPolicyContext
  var request_ *dara.Request
  var response_ *dara.Response
  var _resultErr error
  retriesAttempted := int(0)
  retryPolicyContext = &dara.RetryPolicyContext{
    RetriesAttempted: retriesAttempted,
  }

  for dara.ShouldRetry(_runtime.RetryOptions, retryPolicyContext) {
    _resultErr = nil
    _backoffDelayTime := dara.GetBackoffDelay(_runtime.RetryOptions, retryPolicyContext)
    dara.Sleep(_backoffDelayTime)

    request_ = dara.NewRequest()
    // new model instance comment
    modelInstance := &Test3{
      //empty model 
    }
    // boolean declare comment
    bool := true
    if bool {
      //empty if
    } else {
      //empty else
    }

    // api function call comment
    _err = client.TestAPI()
    if _err != nil {
      retriesAttempted++
      retryPolicyContext = &dara.RetryPolicyContext{
        RetriesAttempted: retriesAttempted,
        HttpRequest:      request_,
        HttpResponse:     response_,
        Exception:        _err,
      }
      _resultErr = _err
      continue
    }

    // back comment
    response_, _err = dara.DoRequest(request_, _runtime)
    if _err != nil {
      retriesAttempted++
      retryPolicyContext = &dara.RetryPolicyContext{
        RetriesAttempted: retriesAttempted,
        HttpRequest:      request_,
        HttpResponse:     response_,
        Exception:        _err,
      }
      _resultErr = _err
      continue
    }

    // empty return comment
  }
  if dara.BoolValue(client.DisableSDKError) != true {
    _resultErr = dara.TeaSDKError(_resultErr)
  }
  return _result, _resultErr
}


func StaticFunc () {
  a := []interface{}{
    // empty annotation comment
  }
}

// Description:
// 
// testFunc
func TestFunc () (_err error) {
  // empty comment1
  // empty comment2
  return _err
}

