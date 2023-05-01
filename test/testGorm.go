package main

import (
	"IM_QQ/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:zyq425ZYQ@tcp(127.0.0.1:3306)/QQ?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.UserFriend{})
	db.AutoMigrate(&models.GroupMember{})
	db.AutoMigrate(&models.Group{})

	//user := &models.UserInfo{}
	//user.Name = "zyq"
	//
	//// Create
	//db.Create(user)
	//
	//// Read
	//db.First(user, 1) // 根据整型主键查找

	// Update - 将 product 的 price 更新为 200
	//db.Model(user).Update("Password", "1234")

	// Delete - 删除 product
	//db.Delete(user, 1)
}
