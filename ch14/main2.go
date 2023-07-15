package main

import "fmt"

func main() {
	p := person{name: "zhangsan", age: 18}
	fmt.Println(p.name, p.age)
	modify(&p)
	fmt.Println("person name:", p.name, "age:", p.age)
	name := "wyh"
	modifyString(&name)
	fmt.Println("name:", name)
}

func modify(p *person) {
	p.name = "lisi"
	p.age = 20
}

func modifyString(name *string) {
	*name = "wyh2"
}

type person struct {
	name string
	age  int
}
