package recovery

import (
    "fmt"
    "gin-example/util/logger"
    "github.com/gin-gonic/gin"
    "runtime"
)

// 捕获系统致命错误 放置进程推出
func Recovery() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                const size = 64 << 10
                buf := make([]byte, size)
                buf = buf[:runtime.Stack(buf, false)]
                pl := fmt.Sprint("http call panic:服务发生错误\n%v\n%s\n", err, buf)
                logger.Panic(pl)
                c.AbortWithStatus(500)
                return
            }
        }()
        c.Next()
    }
}
