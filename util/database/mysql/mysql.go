package mysql

import (
    "fmt"
    "gin-example/config"
    "gin-example/util/logger"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
)

var Mysql *gorm.DB

func Setup() {
    fmt.Println("数据库链接初始化")
    openStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&timeout=%s&parseTime=true",
        config.Mysql.Username,
        config.Mysql.Password,
        config.Mysql.Host,
        config.Mysql.Port,
        config.Mysql.Database,
        config.Mysql.Charset,
        config.Mysql.Timeout)

    var db, err = gorm.Open("mysql", openStr)
    if err != nil {
        fmt.Println("数据库链接失败:", openStr)
        logger.Error("数据库链接失败:", openStr)
    }
    db.DB().SetMaxIdleConns(config.Mysql.MaxIdle)
    db.DB().SetMaxOpenConns(config.Mysql.MaxOpen)
    // 全局禁用表名复数
    db.SingularTable(true)
    // 启用Logger，显示详细日志
    if config.Mysql.RunMode == "debug" {
        db.LogMode(true)
    }
    db.DB().Ping()
    Mysql = db

}
