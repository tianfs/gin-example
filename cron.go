package main

import (
    "github.com/robfig/cron"
    "log"
    "time"
)

func main() {
    log.Println("Starting...")

    c := cron.New()
    c.AddFunc("*/5 * * * * *", func() {
        log.Println("Run models.CleanAllTag...")
    })


    c.Start()

    t1 := time.NewTimer(time.Second * 10)
    for {
        select {
        case <-t1.C:
            t1.Reset(time.Second * 10)
        }
    }
}


