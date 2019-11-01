/**
 * @Author: DollarKiller
 * @Description: 并发锁相关的内容
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 20:06 2019-11-01
 */
package main

import (
	"fmt"
	"go-Interview-book/datasource/pgsql"
	"go-Interview-book/test_model"
	"log"
	"strconv"
	"sync"
	"time"
)

func main() {
	//one()
	two()
}

/*
 * 测试并发内存模型 没有加锁的情况
 */
var ic map[string]string
var oneUn sync.Mutex

func one() {
	ic = make(map[string]string)
	// 并发写数据
	go func() {
		for i := 0; i < 10000; i++ {
			go func(i int) {
				oneUn.Lock()
				defer oneUn.Unlock()
				ic["name"] = strconv.Itoa(i)
			}(i)
		}
	}()
	// 并发读取数据
	go func() {
		for i := 0; i < 10000; i++ {
			go func() {
				oneUn.Lock()
				defer oneUn.Unlock()
				fmt.Println(ic["name"])
			}()
		}
	}()

	time.Sleep(time.Second * 10)
}

func two() {
	//手动初始化 数据库插入初始数据
	//data := test_model.Back{Id:"001",Balance:0}
	//err := mysql.MySQL.Create(&data).Error
	//if err != nil {
	//	panic(err)
	//}

	go func() {
		session := pgsql.PgSQL.Begin()
		data1 := test_model.Back{}
		err := session.Where("id = ?", "001").Find(&data1).Error
		if err != nil {
			log.Println("参数获取失败1")
			panic(err)
		}
		// 模拟获取后已经发生改变值
		go xr()
		time.Sleep(time.Second)

		data1.Balance += 200
		err = session.Model(&data1).Where("id = ?", "001").Update("balance", data1.Balance).Error
		if err != nil {
			log.Println("更新失败")
			panic(err)
		}
		session.Commit()

		//data1 := test_model.Back{}
		//err := pgsql.PgSQL.Where("id = ?", "001").Find(&data1).Error
		//if err != nil {
		//	log.Println("参数获取失败1")
		//	panic(err)
		//}
		//// 模拟获取后已经发生改变值
		//go xr()
		//time.Sleep(time.Second)
		//
		//data1.Balance += 200
		//err = pgsql.PgSQL.Model(&data1).Where("id = ?", "001").Update("balance",data1.Balance).Error
		//if err != nil {
		//	log.Println("更新失败")
		//	panic(err)
		//}

	}()

	//xr()

	time.Sleep(10 * time.Second)
}

func xr() {
	session := pgsql.PgSQL.Begin()
	data1 := test_model.Back{}
	err := session.Where("id = ?", "001").Find(&data1).Error
	if err != nil {
		log.Println("参数获取失败2")
		panic(err)
	}
	data1.Balance += 100
	err = session.Model(&data1).Where("id = ?", "001").Update("balance", data1.Balance).Error
	if err != nil {
		log.Println("更新失败")
		panic(err)
	}
	session.Commit()

	//data1 := test_model.Back{}
	//err := pgsql.PgSQL.Where("id = ?", "001").Find(&data1).Error
	//if err != nil {
	//	log.Println("参数获取失败2")
	//	panic(err)
	//}
	//data1.Balance += 100
	//err = pgsql.PgSQL.Model(&data1).Where("id = ?", "001").Update("balance",data1.Balance).Error
	//if err != nil {
	//	log.Println("更新失败")
	//	panic(err)
	//}
}
