package co

import (
    "gin-example/util/logger"
    "runtime"
)

func Go(x func()){
    go func(){
        defer func() {
            if err := recover(); err != nil {
                const size = 64 << 10
                buf := make([]byte, size)
                buf = buf[:runtime.Stack(buf, false)]
                logger.Errorf("co Go panic:服务发生错误\n %v \n %s \n", err, buf)
            }
        }()
        x();
    }()

}
