package main

import (
	"encoding/json"
	"fmt"
)

type Object struct {
	Name string `json:"name"`
	Vip  bool   `json:"vip,string"`
}

func main() {
	printJson(Object{"golang", true})
	printJson([]int{1, 2, 3})
	printJson("hello world")
	printJson(1024)
	printJson(true)
	printJson(false)
	printJson(nil)
}

func printJson(v interface{}) {
	bs, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}
