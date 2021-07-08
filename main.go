package main

import (
    "context"
    "fmt"
    "gin-example/config"
    "gin-example/cron"
    "gin-example/router"
    "gin-example/util/cache/redis"
    "gin-example/util/database/mysql"
    "gin-example/util/logger"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"

    "os"
    "os/signal"
    "syscall"
    "time"
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
    /*if err := agent.Listen(agent.Options{
        ShutdownCleanup: true,
    }); err != nil {
        log.Fatalf("agent.Listen err: %v", err)
    }*/

    gin.SetMode(config.Http.RunMode)
    server := &http.Server{
        Addr:           fmt.Sprintf("%s:%d", config.Http.HttpHost, config.Http.HttpPort),
        Handler:        router.SetRouter(),
        ReadTimeout:    config.Http.ReadTimeout,
        WriteTimeout:   config.Http.WriteTimeout,
        MaxHeaderBytes: 1 << 20,
    }
    go func() {
        // 服务连接
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Errorf("listen: %s\n", err)
        }
    }()
    // 平滑重启
    listenSignal(server)

}

// 平滑重启
func listenSignal(httpSrv *http.Server) {
    //接收信号
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
    sig := <-sigs
    log.Println("接收到退出信号", sig)

    // 设置超时
    ctxWt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // 运行退出
    log.Println("run Server Shutdown ....")
    err := httpSrv.Shutdown(ctxWt)
    if err != nil {
        log.Println("Server Shutdown error", err)
    }

    // 监听超时
    log.Println("context.WithTimeout超时监控中,time=", time.Now().Unix())
    select {
    case <-ctxWt.Done():
        log.Println("timeout of context.WithTimeout.")

    }

    log.Println("Server exiting")

}
