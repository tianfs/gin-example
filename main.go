package main

import (
    "fmt"
    "gin-example/config"
    "gin-example/cron"
    "gin-example/database/mysql"
    "gin-example/database/redis"
    "gin-example/router"
    "gin-example/util/logger"
    "github.com/gin-gonic/gin"

    "net/http"
)

func init() {

    config.Setup()
    logger.Setup()
    mysql.Setup()
    redis.Setup()
    cron.Setup()
}
func main() {
    // 进程诊断工具 gops
    //if err := agent.Listen(agent.Options{}); err != nil {
    //    log.Fatalf("agent.Listen err: %v", err)
    //}

    if config.App.Env == "prod" {
        gin.SetMode(gin.ReleaseMode)
    }
    fmt.Println(fmt.Sprintf("%s:%d", config.Http.HttpHost, config.Http.HttpPort))
    server := &http.Server{
        Addr:           fmt.Sprintf("%s:%d", config.Http.HttpHost, config.Http.HttpPort),
        Handler:        router.SetRouter(),
        ReadTimeout:    config.Http.ReadTimeout,
        WriteTimeout:   config.Http.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }
    server.ListenAndServe()



}
