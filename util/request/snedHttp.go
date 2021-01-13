package request

import (
    "encoding/json"
    "errors"
    "gin-example/util/e"
    "io/ioutil"
    "net/http"
)

type httpResponse struct {
    Code    int                    `json:"errorCode"`
    Message string                 `json:"errorMsg"`
    Data    map[string]interface{} `json:"data"`
}

func Get(url string, params *map[string]interface{}) (*map[string]interface{}, error) {

    res, error := http.Get(url)
    defer res.Body.Close()
    if error != nil {
        return nil, error
    }
    body, error := ioutil.ReadAll(res.Body)
    if error != nil {
        return nil, error
    }
    var result httpResponse
    json.Unmarshal(body, &result)
    if res.StatusCode != e.HTTP_SUCCESS {
        return nil, errors.New("网络请求错误")
    }

    if result.Code != 0 {
        return nil, errors.New("w:" + result.Message)
    }
    return &result.Data, nil
}

type httpRequest struct {
    Body *[]byte
}

func NewGet(url string, params *map[string]interface{}) (*httpRequest,error) {
    request := new(httpRequest)
    res, error := http.Get(url)

    if error != nil {
        return nil,error
    }
    defer res.Body.Close()
    body, error := ioutil.ReadAll(res.Body)
    if error != nil {
        return nil,error
    }

    if res.StatusCode != e.HTTP_SUCCESS {
        return nil,errors.New("网络请求错误")
    }
    request.Body = &body
    return request,nil
}
func (this *httpRequest) ToJson(dataStruct interface{}) error {
    json.Unmarshal(*this.Body, dataStruct)
    return nil
}


