// https://replit.com/@MonguDykrai/golangfor-loop#main.go
// https://freshman.tech/snippets/go/check-type-of-value/
// https://www.socketloop.com/tutorials/golang-how-to-convert-character-to-ascii-and-back
// https://www.codecademy.com/courses/learn-go-loops-arrays-maps-and-structs/lessons/go-loops/exercises/the-classic-for-loop
package main

import (
	"fmt"
	"reflect"
)

func main() {
	greeting := "Hello, welcome!"
	// string
	fmt.Println("reflect.TypeOf(greeting): ", reflect.TypeOf(greeting))
	for i := 0; i < len(greeting); i++ {
		// 打印 ASCII 码，72 101 108 ...
		fmt.Println(greeting[i], "")
		// 打印 ASCII 码，72 101 108 ...
		fmt.Println(rune(greeting[i]), "")
		// 打印字符，H e l l ...
		fmt.Println(string(greeting[i]), "")
	}
}
