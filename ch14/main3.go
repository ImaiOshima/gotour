package main

import "fmt"

func main() {
	p := person{name: "zhangsan", age: 18}
	p.name = "lisi"
	p.age = 20
	fmt.Println(p.name, p.age)
}
