// https://www.programiz.com/golang/pointers-struct
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{
		name: "Jack",
		age:  32,
	}
	fmt.Println(p1)
	fmt.Println(p1.name)
	p1.age = 46
	fmt.Println(p1.age)
	p2 := &Person{
		name: "Rose",
		age:  31,
	}
	fmt.Println(p2)
	fmt.Println((*p2).name)
	fmt.Println(p2.name)

	var p3 *Person
	p3 = &Person{
		name: "Tim",
		age:  36,
	}
	fmt.Println(p3)
	fmt.Println(reflect.TypeOf(p3))
}
