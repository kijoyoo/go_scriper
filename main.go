/*go에서 compile 하기 위해서는 반드시 main.go package main  fun main가 있어야한다.*/
package main

import (
	"fmt"
	"strings"
)

func multyply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(name ...string) {
	fmt.Println(name)
}

func main() {
	const name1 string = "kijo"
	name2 := "sol"
	answer := false
	fmt.Println(name2, answer)

	multyAnswer := multyply(2, 2)
	fmt.Println(multyAnswer)

	totalLength, upperName := lenAndUpper("kijo")
	totalLength2, _ := lenAndUpper("kijo")
	fmt.Println("hi this is multy function", totalLength, upperName)
	fmt.Println("_ is ignore var ", totalLength2)

	repeatMe("kijo", "sol", "dame", "jean")
}
