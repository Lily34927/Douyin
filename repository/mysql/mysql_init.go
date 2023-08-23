package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"groupwork/module"
)

var DB *gorm.DB

// Init mysql数据库初始化
func Init() (err error) {
	// 1.连接数据库
	dst := "root:000000@tcp(127.0.0.1:3306)/douyin?parseTime=true&loc=Local"
	DB, err = gorm.Open("mysql", dst)

	if err != nil {
		panic("failed to connect database")
	}

	// 2.自动迁移
	DB.AutoMigrate(&module.User{}, &module.Video{}, &module.FavoriteVideo{}, &module.Comment{})
	return DB.DB().Ping()
}

// Close 关闭mysql数据库
func Close() {
	_ = DB.Close()
}
