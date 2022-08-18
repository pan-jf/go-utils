package main

import (
	"fmt"
	"github.com/pan-jf/go-utils/t2t"
)

// SA ...
type SA struct {
	Data int64  `t2t:"data"`
	Str  string `t2t:"str2"`
}

func main() {
	data := map[string]interface{}{
		"data": 999,
		"str2": "hello",
	}

	var sa SA
	err := t2t.Transfer(data, &sa)
	if err != nil {
		fmt.Printf("transfer error: %v\n", err)
		return
	}

	fmt.Printf("transfer data: %+v", sa)
}
