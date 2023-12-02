package main

import (
	"fmt"
	"time"
)

type SearchUserV1Request struct {
	RegisterAt string `json:"register_at"`
}

func main() {
	checkTimeParse(&SearchUserV1Request{
		RegisterAt: "2020-01-01T15:08:23",
	})
}

func checkTimeParse(req *SearchUserV1Request) {
	_, err := time.Parse("2006-01-02 15:04:05", req.RegisterAt)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success")
}
