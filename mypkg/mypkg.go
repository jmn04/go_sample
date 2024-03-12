package mypkg

import (
	"fmt"
)

// 大文字で始まる関数は外部からアクセス可能
func MyFunc() {
	fmt.Println("Hello, mypkg")
}

// // 小文字で始まる関数は外部からアクセス不可
// func my_func() {
// 	fmt.Println("Hello, mypkg")
// }