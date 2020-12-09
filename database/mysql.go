package database

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "log"
    "gin-example/config"
)

var Mysql *gorm.DB

func init() {
    fmt.Println("数据库链接初始化")

    fmt.Println(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%s&parseTime=true",
        config.Mysql.Username,
        config.Mysql.Password,
        config.Mysql.Host,
        config.Mysql.Port,
        config.Mysql.Database,
        config.Mysql.Charset,
        config.Mysql.Timeout))
    var db, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%s&parseTime=true",
        config.Mysql.Username,
        config.Mysql.Password,
        config.Mysql.Host,
        config.Mysql.Port,
        config.Mysql.Database,
        config.Mysql.Charset,
        config.Mysql.Timeout))
    if err != nil {
        fmt.Println("数据库链接失败")
        log.Println(err)
    }
    db.DB().SetMaxIdleConns(config.Mysql.MaxIdle)
    db.DB().SetMaxOpenConns(config.Mysql.MaxOpen)
    db.DB().Ping()
    Mysql = db
    // 全局禁用表名复数
    db.SingularTable(true)
}


