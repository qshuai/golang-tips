package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func main() {
	checkUnmarshalJson([]byte(`{"register_at": "2020-01-01T15:08:23"}`))
}

type DefaultDatetime time.Time

func (d *DefaultDatetime) UnmarshalJSON(bs []byte) error {
	if len(bs) == 0 {
		return nil
	}

	t, err := time.ParseInLocation("2006-01-02 15:04:05", strings.Trim(string(bs), "\""), time.Local)
	if err != nil {
		return err
	}

	*d = DefaultDatetime(t)
	return nil
}

type SearchUserV2Request struct {
	RegisterAt *DefaultDatetime `json:"register_at"`
}

func checkUnmarshalJson(body []byte) {
	var req SearchUserV2Request
	err := json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success")
}
