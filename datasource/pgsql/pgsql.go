/**
 * @Author: DollarKiller
 * @Description: 做并发数据库测试提供的链接
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 22:46 2019-11-01
 */
package pgsql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-Interview-book/test_model"
	"time"
)

var PgSQL *gorm.DB

func init() {
	var err error
	PgSQL, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=test password=123456 sslmode=disable")
	if err != nil {
		panic(err)
	}
	PgSQL.LogMode(true)

	PgSQL.DB().SetConnMaxLifetime(500 * time.Millisecond)
	PgSQL.DB().SetMaxIdleConns(20)
	PgSQL.DB().SetMaxOpenConns(100)

	padding()
}

func padding() {
	PgSQL.AutoMigrate(&test_model.Back{})
}
