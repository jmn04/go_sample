package main

import (
	m "mypkg" // import mypkg as m
	t "time"
	f "fmt"
)

// 初期化関数
func init() {
	m.Hello()
}

// func funcA() {
// 	for i := 0; i < 10; i++ {
// 		f.Print("A")
// 		t.Sleep(10 * t.Millisecond) // t.Millisecond : 1ms
// 	}
// }

// チャネルの送信
func funcB(chB chan <- string) {
	t.Sleep(3 * t.Second)
	chB <- "funcB Finished"
}

func funcC(chC chan <- string) {
	t.Sleep(1 * t.Second)
	chC <- "funcC Finished"
}

func main() {
	// go funcA()
	// for i:= 0; i < 10; i++ {
	// 	f.Print("M")
	// 	t.Sleep(20 * t.Millisecond)
	// }

	chB := make(chan string)
	chC := make(chan string)
	defer close(chB)
	defer close(chC)
	finFlagB := false
	finFlagC := false
	go funcB(chB)
	go funcC(chC)
	for {
		select {
		case msg := <- chB:
			finFlagB = true
			f.Println(msg)
		case msg := <- chC:
			finFlagC = true
			f.Println(msg)
		}
		if finFlagB && finFlagC {
			f.Println("All Finished")
			break
		}
	}
}