package main

import "fmt"

func main() {
	nameAgemap := make(map[string]int)
	nameAgemap["wyh"] = 20

	age, ok := nameAgemap["wyh"]
	if ok {
		fmt.Println(age)
	}
}
