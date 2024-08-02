package dao

import (
	"log"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/MyDataBase?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    // 获取通用数据库对象 sql.DB，以便配置连接池
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("failed to get sql.DB: %v", err)
    }

    // 设置连接池参数
    sqlDB.SetMaxIdleConns(10)           // 设置空闲连接池中的最大连接数
    sqlDB.SetMaxOpenConns(100)          // 设置数据库的最大连接数
    sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接的最大可复用时间

    DB = db;
    log.Println("connect to database success.")
}
