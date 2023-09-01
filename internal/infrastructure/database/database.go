package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"demo-gorm-gen/internal/pkg/gu"
)

func Init() *gorm.DB {
	return gu.Must(gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local",
	}), &gorm.Config{}))
}
