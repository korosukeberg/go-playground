package main

import "fmt"

var packageLevelVal bool

func main() {
	fmt.Printf("hello, kotaro\n")

	a, b := addAndSubtract(1, 1)

	fmt.Println(a, b, packageLevelVal)

	//書式指定し %T:type of the variable, %v:value of the varioable
	c := 1
	fmt.Printf("%T %v\n", c, c)

	//Type Conversion
	fmt.Printf(string(c))

	//Constant
	const Pi = 3.14

}

//複数の戻り値を返せる
func addAndSubtract(x, y int) (int, int) {
	return x + y, x - y
}

//戻り値宣言の時点で変数に名前をつけておける。そのままreturnできる。naked return
func sum(x, y int) (answer int) {
	answer = x + y
	return
}
