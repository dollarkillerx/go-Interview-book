/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 23:44 2019-11-01
 */
package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-Interview-book/test_model"
	"time"
)

var MySQL *gorm.DB

func init() {
	var err error
	MySQL, err = gorm.Open("mysql", "root:123456@/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	MySQL.LogMode(true)

	MySQL.DB().SetConnMaxLifetime(500 * time.Millisecond)
	MySQL.DB().SetMaxIdleConns(20)
	MySQL.DB().SetMaxOpenConns(100)

	padding()
}

func padding() {
	MySQL.AutoMigrate(&test_model.Back{})
}
