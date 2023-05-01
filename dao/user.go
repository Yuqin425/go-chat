package dao

import (
	"IM_QQ/models"
	"IM_QQ/settings"
	"IM_QQ/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitMysql(cfg *settings.MySQLConfig) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	var err error

	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	sqlDB, _ := db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func FindUser(username string) bool {
	// fmt.Println(222)
	user := models.User{}
	var userCount int64
	// fmt.Println(111)
	db.Model(user).Where("username", username).Count(&userCount)
	// fmt.Println(333)
	if userCount > 0 {
		// fmt.Println(userCount)
		return true
	}
	return false
}

func CheckPwd(username, password string) bool {
	user := models.User{}
	db.Model(user).Where("username = ?", username).First(&user)
	if !utils.ValidPassword(password, user.Password) {
		return false
	}
	return true
}

func CreateUser(user *models.User) *gorm.DB {
	return db.Create(user)
}
