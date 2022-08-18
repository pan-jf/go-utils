package main

import (
	"encoding/json"
	"fmt"
	"github.com/pan-jf/go-utils/t2t"
)

// KindGetter ...
type KindGetter struct {
	Kind string `json:"kind"`
}

// Car ...
type Car struct {
	Kind   string  `json:"kind"`
	Doors  float64 `json:"doors"`
	Wheels float64 `json:"wheels"`
}

func main() {
	jsonData := `
{
	"kind": "Car",
	"doors": 4,
	"wheels": 4
}`

	var data interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Printf("json unmarshal fail: %v", err)
		return
	}

	tf := t2t.NewTransformer()
	tf.TagName = t2t.JsonFormat

	var kindGetter KindGetter
	err = tf.Transfer(data, &kindGetter)
	if err != nil {
		fmt.Printf("transfer to KindGetter fail: %v", err)
		return
	}
	fmt.Println(kindGetter)

	var car Car
	err = tf.Transfer(data, &car)
	if err != nil {
		fmt.Printf("transfer to Car fail: %v", err)
		return
	}
	fmt.Println(car)
}
