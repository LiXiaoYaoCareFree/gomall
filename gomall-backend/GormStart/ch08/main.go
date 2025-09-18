package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type NewUser struct {
	ID           uint
	MyName       string `gorm:"column:name"`
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActiveAt     sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Deleted      gorm.DeletedAt
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&NewUser{})
	if err != nil {
		return
	}
	//var users = []NewUser{{MyName: "jinzhu1"}, {MyName: "jinzhu2"}, {MyName: "jinzhu3"}}
	//db.Create(&users)

	// 硬删除
	db.Unscoped().Delete(&NewUser{ID: 2})

	// 软删除
	//db.Delete(&NewUser{}, 1)
	//var users []NewUser
	//db.Find(&users)
	//for _, user := range users{
	//	fmt.Println(user.ID)
	//}

}
