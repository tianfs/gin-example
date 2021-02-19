package test

import (
    "gin-example/util/request/resty"
    "testing"
)

type httpResponse struct {
    Code    int
    Message string
}

func TestGoRequestPost(t *testing.T) {
    url := "http://local.user.api.meidaifu.com/instSpaceSchedule/getScheduleStatus?test1=11"
    params := map[string]string{
        "Safari": "12321323",
    }
    params["type"] = "Get"
    resty.Get(url, params)
    params["type"] = "PostForm"
    resty.PostForm(url, params)
    params["type"] = "PostJson"
    resty.PostJson(url, params)

}
