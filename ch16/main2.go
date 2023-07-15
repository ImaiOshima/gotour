package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	p := person{Name: "飞雪无情", Age: 20}
	jsonB, err := json.Marshal(p)
	if err == nil {
		fmt.Println(string(jsonB))
	}
	responseJson := "{\"Name\":\"李四\",\"Age\":40}"
	json.Unmarshal([]byte(responseJson), &p)
	fmt.Println(p)
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
