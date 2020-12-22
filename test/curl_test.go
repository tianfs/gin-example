package test

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "testing"
)

type base struct {
    Code    int     `json:"code"`
    Message string  `json:"message"`
    Data    []testD `json:"data"`
}
type testD struct {
    Id      int    `json:"Id"`
    InvCode string `json:"InvCode"`
    InvId   int    `json:"InvId"`
}

func TestCurl(t *testing.T) {
    // get_struct();
    get_map()

}

func get_struct() {
    res, error := http.Get("http://127.0.0.1:8083/saleOrder/list")
    defer res.Body.Close()
    if error != nil {
        fmt.Println("失败", error)
        return
    }
    body, _ := ioutil.ReadAll(res.Body)
    var aaa base
    json.Unmarshal(body, &aaa)
    fmt.Println("get_struct", aaa.Data)
    fmt.Println(res.StatusCode)
    if res.StatusCode == 200 {
        fmt.Println("ok")
    }
}
func get_map() {
    res, error := http.Get("http://127.0.0.1:8083/saleOrder/list")
    defer res.Body.Close()
    if error != nil {
        fmt.Println("失败", error)
        return
    }
    body, _ := ioutil.ReadAll(res.Body)

    jsonMap := map[string]interface{}{}
    json.Unmarshal(body, &jsonMap)
    aa := jsonMap["data"].([]interface{})
    for k, v := range aa {
        fmt.Println(k, v.(map[string]interface{}))
    }

    fmt.Println(res.StatusCode)
    if res.StatusCode == 200 {
        fmt.Println("ok")
    }
}
