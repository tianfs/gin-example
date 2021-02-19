package test

import (
    "gin-example/util/request/gorequest"
    "testing"
)

type httpResponse struct {
    Code    int
    Message string
}

func TestGoRequestPost(t *testing.T) {
    url := "http://local.user.api.meidaifu.com/instSpaceSchedule/getScheduleStatus?test1=11"
    params := map[string]interface{}{
        "Safari": "12321323",
    }
    params["type"]="Get";
    gorequest.Get(url, &params)
    params["type"]="PostForm";
    gorequest.PostForm(url, &params)
    params["type"]="PostJson";
    gorequest.PostJson(url, &params)

}
