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


