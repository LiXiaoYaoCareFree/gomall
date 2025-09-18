package main

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActiveAt     sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func main() {
	//a := []int{1, 2, 3}

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

	_ = db.AutoMigrate(&User{}) //此处应该有sql语句
	user := User{
		Name: "lly",
	}

	fmt.Println(user.ID)
	result := db.Create(&user)
	fmt.Println(user.ID)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected)
	db.Model(&User{ID: 1}).Update("Name", "")
	// updates语句不会更新零值，但是update语句会更新
	empty := ""
	db.Model(&User{ID: 1}).Updates(User{Name: empty, Age: 18})
	//解决仅更新非零值字段的方法有两种
	/*
		1. 将string 设置为 *string
		2. 使用sql的NULLxxx来解决
	*/

}
