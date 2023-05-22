package test

import "fmt"
import "github.com/OneHundred86/golang/test/aa"

func Print(str string) {
	fmt.Printf("[test]: %s \n", str)
}

// 调用aa模块的函数或者资源，需要引入完整的aa模块目录：github.com/OneHundred86/golang/test/aa
func callFuncInaa() {
	aa.Show("abc")
}
