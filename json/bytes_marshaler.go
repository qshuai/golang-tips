package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type Bytes []byte

func (b Bytes) MarshalJSON() ([]byte, error) {
	if b == nil {
		return nil, nil
	}

	bb := bytes.Buffer{}
	bb.WriteByte('[')
	for idx, item := range b {
		bb.WriteString(strconv.Itoa(int(item)))

		if idx != len(b)-1 {
			bb.WriteByte(',')
		}
	}
	bb.WriteByte(']')
	return bb.Bytes(), nil
}

// 期望: [0,8,101,0]
var data = Bytes{0, 8, 101, 0}

func main() {
	bs, err := json.Marshal(&data)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
}
