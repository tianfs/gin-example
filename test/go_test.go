package test

import (
    "fmt"
    "gin-example/config"
    "gin-example/util/co"
    "gin-example/util/logger"
    "sync"
    "testing"
    "time"
)

func TestTo(t *testing.T){
    config.Setup()
    logger.Setup()
   var wg  sync.WaitGroup;
   wg.Add(2);
    co.Go(func(){
        fmt.Println("1")
        panic("报错了")
        wg.Done();
    })
    co.Go(func(){
        fmt.Println("2")
        wg.Done();
    })
    wg.Wait()
    time.Sleep(5*time.Second)

}