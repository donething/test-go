package main

import "fmt"

// defer 普通函数和匿名函数时，传递的参数值
func deferAnonymous() {
	name := "Li"

	// name 的参数值为 "Li"
	defer rename(name)

	defer func() {
		// name 的参数值为 "Zhang"
		fmt.Printf("main: %s\n", name)
	}()

	name = "Zhang"
}

func rename(name string) {
	// 打印 "Li"
	fmt.Printf("rename: %s\n", name)
}
