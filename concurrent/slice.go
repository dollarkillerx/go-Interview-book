/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 09:33 2019-11-02
 */
package main

import "log"

func main() {

	//ts1()
	//ts2()
	//ts3()

	ts4()

	//dsa1 = append(dsa1, 6, 7, 8, 9) 添
	//dsa1[0] = 188
}

func ts1() {
	data := [5]int{1, 2, 3, 4, 5}

	dsa1 := data[1:3] // 切片是[]的视图  当修改视图中的内容  原数组也会被修改
	dsa1[0] = 188
	log.Println(data)
}

func ts2() {
	data := [5]int{1, 2, 3, 4, 5}

	dsa1 := data[1:3]

	log.Println(len(dsa1), cap(dsa1))

	dsa1 = append(dsa1, 112) // 当添加的数组没有操作cap的时候  原数组会被修改

	log.Println(data)
}

func ts3() {
	data := [5]int{1, 2, 3, 4, 5}

	dsa1 := data[1:3]

	log.Println(len(dsa1), cap(dsa1))

	dsa1 = append(dsa1, 123, 3434, 554, 565, 76)
	// 当填充 超过cap的时候 生成新的数组
	log.Println(data)
}

func ts4() {
	data := [5]int{1,2,3,4,5}
	ts5(data[:])
	log.Println(data)
}

func ts5(ac []int) {
	ac[1] = 122
}
