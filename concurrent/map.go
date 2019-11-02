/**
 * @Author: DollarKiller
 * @Description:
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 10:10 2019-11-02
 */
package main

import "log"

func main() {
	m1()
}

func m1() {
	map1 := make(map[int]int, 3)

	map1[1] = 1
	map1[2] = 2
	map1[3] = 3
	log.Println(map1)
	mt1(map1)

	log.Println(map1)
	// 禁止对map元素取址的原因是map可能随着元素数量的增长而重新分配更大的内存空间，从而可能导致之前的地址无效。
}

func mt1(map1 map[int]int) {
	map1[5] = 6
}
