package cron

import (
    "fmt"
    "github.com/robfig/cron"


)

func Setup(){
    c := cron.New()
    c.AddFunc("*/5 * * * * *", testCron1)
    c.AddFunc("*/6 * * * * *", testCron2)
    c.Start()

}

func testCron1(){
    fmt.Println("testCron1 Run models.CleanAllTag...")
}
func testCron2(){
    fmt.Println("testCron2 Run models.CleanAllTag...")
}