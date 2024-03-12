package main

import (
	"fmt"
	"time"
	"go_sample/mypkg" // 自作パッケージのインポート
)

// 構造体
type Person struct {
	// public
	Age int
	Height int
	Weight int
	// private
	name string
}

func (p *Person) setPerson(Age int, Height int, Weight int, name string) {
	p.Age = Age
	p.Height = Height
	p.Weight = Weight
	p.name = name
}

func (p *Person) getPerson() (int, int, int, string) {
	return p.Age, p.Height, p.Weight, p.name
}

func (p Person) ToString() (int, int, int, string) {
	return p.Age, p.Height, p.Weight, p.name
}

func (p Person) PrintOut() {
	fmt.Println(p.ToString())
}

type Book struct {
	Title string
}

func (b Book) ToString() string {
	return b.Title
}

func (b Book) PrintOut() {
	fmt.Println(b.ToString())
}

// インターフェース
// ToStringは中身が異なるので仕方ないが、PrintOutは同じなのでインターフェースを使う
/* Printable というインタフェースは、ToString() というメソッドを
サポートしている構造体であれば、自動的に適用することが可能 */
type Printable interface {
	ToString() string
}
func PrintOut (p Printable) {
	fmt.Println(p.ToString())
}

func main() {
	num := 1
	str := "Hello, World!"

	// 変数群
	var (
		al int = 0
		// bd string = "Apple"
		// cm int = 8
	)

	// 定数群
	const (
		// pi = 3.14
		// e = 2.71
	)

	// 配列
	// ar0 := [3]int{}
	// ar1 := [3]int{1, 3, 3}
	// ar2 := [3]string{"Apple", "Banana", "Cherry"}
	// ar4 := [...]string{"Ken", "John", "Mike", "Tom", "Bob"}

	// スライス
	// sl := []int{}
	// sl = append(sl, 0)
	// sl = append(sl, 2)
	// sl = append(sl, 3)
	// sl = append(sl, 4, 5, 6)

	// マップ
	mp := map[string]string{
		"Apple": "Red",
		"Banana": "Yellow",
		"Cherry": "Red",
		"Durian": "Green",
	}
	fmt.Println(mp["Apple"])

	// loop
	sl0 := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(sl0); i++ {
		if i % 2 == 0 {
			fmt.Println(sl0[i])
		} else if i % 3 == 0 {
			continue
		} else if i % 8 == 0 {
			break
		} else {
			fmt.Println("odd")
		}
	}

	// マップに要素があるか確認
	_, ok := mp["Orange"]
	if ok {
		fmt.Println("Apple is in the map")
	} else {
		fmt.Println("Apple is not in the map")
	}

	// マップのループ処理
	for key, value := range mp {
		fmt.Println(key, value)
	}

	// switch文
	num0 := 8
	// case文で値を使う場合
	switch num0 {
	case 1:
		fmt.Println("num0 is 1")
	case 2:
		fmt.Println("num0 is 2")
	default:
		fmt.Println("num0 is default")
	}
	// case文で条件式を使う場合
	switch {
	case num0 == 1:
		fmt.Println("num0 is 1")
	case num0 < 5:
		fmt.Println("num0 is smaller than 5")
	default:
		fmt.Println("num0 is default")
	}
	// case文で次の条件の処理も行う場合 この場合、cloudyの場合もrainyの処理が行われる
	weather := "rainy"
	switch weather {
	case "cloudy":
		fallthrough
	case "rainy":
		fmt.Println("It's not sunny")
	default:
		fmt.Println("It's sunny")
	}

	// comment
	/*
	This is comment block
	This is comment block
	This is comment block
	 */

	al = 5
	fmt.Println("num + al:", num + al)
	fmt.Println(str)
	fmt.Println("The time is", time.Now())

	// 関数呼び出し
	add, min := addMinus(5, 3)
	fmt.Println(add, min)

	_, min2 := addMinus(1, 3) // 1つ目の戻り値を無視 _はブランク変数という
	fmt.Println(min2)

	// Pythonの*argsのような可変長引数
	// def funcA()
	funcA(1, 2, 3, 4, 5)

	// 構造体
	p1 := Person{}
	p1.setPerson(30, 180, 70, "Ken")
	_, height, weight, name := p1.getPerson()
	fmt.Printf("%sは身長が%d, 体重が%dです\n", name, height, weight)

	// ポインタ
	// a1 := 456
	// var ptr *int
	// ptr = &a1
	// *ptr = 123
	// fmt.Println(a1)

	// 自作パッケージ
	mypkg.MyFunc()
}

// 関数定義
func addMinus(a int, b int) (int, int) {
	return a + b, a - b
}

func funcA(a int, b ... int) {
	fmt.Printf("a=%d\n", a)
	for i, num := range b {
		fmt.Printf("b[%d]=%d\n", i, num)
	}
}