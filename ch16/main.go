package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 3
	iv := reflect.ValueOf(i)
	it := reflect.TypeOf(i)
	fmt.Println(iv, it)
}
