package main

import "fmt"

// 循环时，每次获取的 p 的地址都一样
func forAddr() {
	nums := []int{1, 3, 5}
	for _, n := range nums {
		// 每次获取的 p 的地址都一样
		fmt.Printf("地址：%p\n", &n)
	}
}
