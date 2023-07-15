package main

import "fmt"

func main() {
	addr := address{"北京", "北京"}
	printString(addr)
	printString(&addr)
	var si fmt.Stringer = address{"上海", "上海"}
	printString(si)
}

func printString(addr fmt.Stringer) {
	fmt.Println(addr.String())
}

type address struct {
	province string
	city     string
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}
