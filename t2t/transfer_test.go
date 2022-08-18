package t2t_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/pan-jf/go-utils/t2t"
	"gopkg.in/mgo.v2/bson"
)

// MapKey ...
type MapKey struct {
	Key string
}

// MarshalText ...
func (m *MapKey) MarshalText() (text []byte, err error) {
	return []byte(m.Key), nil
}

// TestBoolTransfer ...
func TestBoolTransfer(t *testing.T) {
	var output interface{}

	err := t2t.Transfer(true, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, true) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(),
			output,
			reflect.TypeOf(true).String(),
			true)
		return
	}
}

// TestIntTransfer ...
func TestIntTransfer(t *testing.T) {
	input := 999
	var output interface{}
	expect := input

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestStringTransfer ...
func TestStringTransfer(t *testing.T) {
	input := "hello"
	var output interface{}
	expect := input

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestFloatTransfer ...
func TestFloatTransfer(t *testing.T) {
	input := float32(0.5)
	var output interface{}
	expect := input

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestInvalidTypeCase ...
func TestInvalidTypeCase(t *testing.T) {
	input := 999
	var output bool

	err := t2t.Transfer(input, &output)
	if err == nil {
		t.Error("invalid type case check fail")
		return
	}
	t.Log("Success: error message ->", err)
}

// TestInvalidTypeCase2 ...
func TestInvalidTypeCase2(t *testing.T) {
	input := float32(0.5)
	var output int

	err := t2t.Transfer(input, &output)
	if err == nil {
		t.Error("invalid type case check fail")
		return
	}
	t.Log("Success: error message ->", err)
}

// TestArrayToInterfaceTransfer ...
func TestArrayToInterfaceTransfer(t *testing.T) {
	input := [5]int{1, 2, 3, 4, 5}
	var output interface{}
	expect := []interface{}{1, 2, 3, 4, 5}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestSliceToInterfaceTransfer2 ...
func TestSliceToInterfaceTransfer2(t *testing.T) {
	input := []struct {
		A int
	}{
		{1},
		{2},
	}

	var output interface{}

	expect := []interface{}{
		map[string]interface{}{
			"A": 1,
		},
		map[string]interface{}{
			"A": 2,
		},
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestArrayLessLengthTransfer ...
func TestArrayLessLengthTransfer(t *testing.T) {
	input := [5]float32{1.1, 2.2, 3.3, 4.4, 5.5}
	var output [4]float32
	expect := [4]float32{1.1, 2.2, 3.3, 4.4}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestArrayGreaterLengthTransfer ...
func TestArrayGreaterLengthTransfer(t *testing.T) {
	input := [5]string{"a", "b", "c", "d", "e"}

	output := [6]string{}
	output[5] = "hello world!"

	expect := [6]string{"a", "b", "c", "d", "e", ""}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestArrayTransferInvalidCase ...
func TestArrayTransferInvalidCase(t *testing.T) {
	input := [5]string{"a", "b", "c", "d", "e"}
	var output [5]int

	err := t2t.Transfer(input, &output)
	if err == nil {
		t.Error("invalid type case check fail")
		return
	}
	t.Log("Success: error message ->", err)
}

// TestArrayToStructTransfer ...
func TestArrayToStructTransfer(t *testing.T) {
	input := []interface{}{
		1,
		"hello",
		0.5,
		true,
	}

	type O struct {
		A int
		B string
		C *float64
		D bool
	}

	var output O
	ec := 0.5
	expect := O{1, "hello", &ec, true}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestSliceToInterfaceTransfer ...
func TestSliceToInterfaceTransfer(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e"}
	var output interface{}
	expect := []interface{}{"a", "b", "c", "d", "e"}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestSliceTransfer ...
func TestSliceTransfer(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	var output []int
	expect := input

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestSliceTransferInvalidCase ...
func TestSliceTransferInvalidCase(t *testing.T) {
	input := []string{"a", "b", "c", "d", "e"}
	var output []int

	err := t2t.Transfer(input, &output)
	if err == nil {
		t.Error("invalid type case check fail")
		return
	}
	t.Log("Success: error message ->", err)
}

// TestSliceGreaterLengthTransfer ...
func TestSliceGreaterLengthTransfer(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	output := make([]int, 0, 10)
	expect := input

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestMapToInterface ...
func TestMapToInterface(t *testing.T) {
	input := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	var output interface{}

	expect := map[string]interface{}{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestMapToInterface2 ...
func TestMapToInterface2(t *testing.T) {
	ka := &MapKey{"a"}
	kb := &MapKey{"b"}
	kc := &MapKey{"c"}

	input := map[*MapKey]int{
		ka: 1,
		kb: 2,
		kc: 3,
	}

	var output interface{}
	expect := map[*MapKey]interface{}{
		ka: 1,
		kb: 2,
		kc: 3,
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestMapToMapTransfer ...
func TestMapToMapTransfer(t *testing.T) {
	input := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	var output map[string]int

	expect := input

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestMapToMapTransferInvalidCase ...
func TestMapToMapTransferInvalidCase(t *testing.T) {
	input := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	var output map[string]float32

	err := t2t.Transfer(input, &output)
	if err == nil {
		t.Error("invalid type case check fail")
		return
	}
	t.Log("Success: error message ->", err)
}

// TestMapToMapTransferInvalidCase2 ...
func TestMapToMapTransferInvalidCase2(t *testing.T) {
	input := map[int64]int{
		1: 1,
		2: 2,
		3: 3,
	}

	var output map[int]int

	err := t2t.Transfer(input, &output)
	if err == nil {
		t.Error("invalid type case check fail")
		return
	}
	t.Log("Success: error message ->", err)
}

// TestMapToStructTransfer ...
func TestMapToStructTransfer(t *testing.T) {
	input := map[string]int{
		"A": 1,
		"B": 2,
	}

	type O struct {
		A int
		B int
	}

	var output O
	expect := O{1, 2}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestMapToEmbeddedStructTransfer ...
func TestMapToEmbeddedStructTransfer(t *testing.T) {
	input := map[string]interface{}{
		"A": 1,
		"B": 2,
		"C": map[string]int{
			"AA": 3,
		},
		"D": 4,
	}

	type EmbAnonymous struct {
		D int
	}

	type Emb struct {
		AA int
	}

	type O struct {
		A            int
		B            int
		C            Emb
		EmbAnonymous `t2t:",inline"`
	}

	var output O
	expect := O{1, 2, Emb{3}, EmbAnonymous{4}}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestMapToEmbeddedStructTransfer2 ...
func TestMapToEmbeddedStructTransfer2(t *testing.T) {
	input := map[string]interface{}{
		"A": 1,
		"B": 2,
		"C": map[string]int{
			"AA": 3,
		},
		"EmbAnonymous": map[string]int{
			"D": 4,
		},
	}

	type EmbAnonymous struct {
		D int
	}

	type Emb struct {
		AA int
	}

	type O struct {
		A int
		B int
		C Emb
		EmbAnonymous
	}
	var output O
	expect := O{1, 2, Emb{3}, EmbAnonymous{4}}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestMapToEmbeddedPtrStructTransfer ...
func TestMapToEmbeddedPtrStructTransfer(t *testing.T) {
	input := map[string]interface{}{
		"A": 1,
		"B": 2,
		"C": map[string]int{
			"AA": 3,
		},
	}

	type Emb struct {
		AA int
	}

	type O struct {
		A int
		B *int
		C *Emb
	}

	var output O

	eb := 2
	expect := O{1, &eb, &Emb{3}}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestTextMarshallerTransfer ...
func TestTextMarshallerTransfer(t *testing.T) {
	ka := &MapKey{"A"}
	kb := &MapKey{"B"}
	kc := &MapKey{"C"}
	kaa := &MapKey{"AA"}

	input := map[*MapKey]interface{}{
		ka: 1,
		kb: 2,
		kc: map[*MapKey]int{
			kaa: 3,
		},
	}

	type Emb struct {
		AA int
	}

	type O struct {
		A int
		B int
		C Emb
	}

	var output O
	expect := O{1, 2, Emb{3}}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestTextMarshallerTransferInvalidCase ...
func TestTextMarshallerTransferInvalidCase(t *testing.T) {

	input := map[MapKey]interface{}{
		MapKey{"A"}: 1,
		MapKey{"B"}: 2,
		MapKey{"C"}: map[MapKey]int{
			MapKey{"AA"}: 3,
		},
	}

	type Emb struct {
		AA int
	}

	output := struct {
		A int
		B int
		C Emb
	}{}

	err := t2t.Transfer(input, &output)
	if err == nil {
		t.Error("invalid type case check fail")
		return
	}
	t.Log("Success: error message ->", err)
}

// StructToMapBase ...
type StructToMapBase struct {
	I int64
}

// TestStructToMapTransfer ...
func TestStructToMapTransfer(t *testing.T) {
	input := struct {
		A  int
		M  map[string]StructToMapBase
		MP map[string]*StructToMapBase
	}{A: 1,
		M: map[string]StructToMapBase{
			"a": {I: 2},
		},
		MP: map[string]*StructToMapBase{
			"b": {I: 999},
		}}

	output := map[string]interface{}{}
	expect := map[string]interface{}{
		"A": 1,
		"M": map[string]interface{}{
			"a": map[string]interface{}{
				"I": int64(2),
			},
		},
		"MP": map[string]interface{}{
			"b": map[string]interface{}{
				"I": int64(999),
			},
		},
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestStructToMapTransfer2 ...
func TestStructToMapTransfer2(t *testing.T) {
	input := struct {
		M  StructToMapBase
		MP *StructToMapBase
	}{
		M: StructToMapBase{
			I: 2,
		},
		MP: &StructToMapBase{
			I: 999,
		}}

	output := map[string]interface{}{}
	expect := map[string]interface{}{
		"M": map[string]interface{}{
			"I": int64(2),
		},
		"MP": map[string]interface{}{
			"I": int64(999),
		},
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestStructToMapTransfer3 ...
func TestStructToMapTransfer3(t *testing.T) {
	input := struct {
		M  StructToMapBase
		MP *StructToMapBase
	}{
		M: StructToMapBase{
			I: 2,
		},
		MP: &StructToMapBase{
			I: 999,
		}}

	output := map[string]StructToMapBase{}
	expect := map[string]StructToMapBase{
		"M": {
			I: 2,
		},
		"MP": {
			I: 999,
		},
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestStructToMapTransfer4 ...
func TestStructToMapTransfer4(t *testing.T) {
	input := struct {
		M  StructToMapBase
		MP *StructToMapBase
	}{
		M: StructToMapBase{
			I: 2,
		},
		MP: &StructToMapBase{
			I: 999,
		}}

	output := map[string]*StructToMapBase{}
	expect := map[string]*StructToMapBase{
		"M": {
			I: 2,
		},
		"MP": {
			I: 999,
		},
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestStructToInterfaceTransfer ...
func TestStructToInterfaceTransfer(t *testing.T) {
	type EmbAnonymous struct {
		C int
	}

	type EmbAnonymousInline struct {
		D int
	}

	type Emb struct {
		A int
	}

	type I struct {
		A int
		B Emb
		EmbAnonymous
		EmbAnonymousInline `t2t:",inline"`
	}

	input := I{1, Emb{2}, EmbAnonymous{3}, EmbAnonymousInline{4}}

	var output interface{}
	expect := map[string]interface{}{
		"A": 1,
		"B": map[string]interface{}{
			"A": 2,
		},
		"EmbAnonymous": map[string]interface{}{
			"C": 3,
		},
		"D": 4,
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestStructToStructTransfer ...
func TestStructToStructTransfer(t *testing.T) {
	input := struct {
		A int
		B string
		C float32
		D bool
	}{1, "hello", 0.5, true}

	type O struct {
		A int64
		B string
		C float64
		D bool
	}

	var output O
	expect := O{1, "hello", 0.5, true}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// StructToStructBase ...
type StructToStructBase struct {
	Int int64
}

// TestStructToStructTag ...
func TestStructToStructTag(t *testing.T) {
	input := struct {
		A  int    `t2t:"C"`
		B  string `t2t:"D"`
		a  int
		M  map[string]StructToStructBase
		MP map[string]*StructToStructBase
	}{A: 1,
		B: "hello",
		a: 2,
		M: map[string]StructToStructBase{
			"tt": {Int: 66},
		}, MP: map[string]*StructToStructBase{
			"tt": {Int: 66},
		}}

	var output struct {
		AA int    `t2t:"C"`
		BB string `t2t:"D"`
		a  int
		M  map[string]StructToStructBase
		MP map[string]*StructToStructBase
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if input.A != output.AA || input.B != output.BB {
		t.Errorf("unexpected value: %+v vs %+v", input, output)
		return
	}
}

// TestStructToSliceTransfer ...
func TestStructToSliceTransfer(t *testing.T) {
	input := struct {
		A int
		B string
		C float64
		D bool
	}{1, "hello", 0.5, true}

	output := make([]interface{}, 0, 4)
	expect := []interface{}{1, "hello", 0.5, true}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestStructToArrayTransfer ...
func TestStructToArrayTransfer(t *testing.T) {
	input := struct {
		A int
		B string
		C float64
		D bool
	}{1, "hello", 0.5, true}

	output := [3]interface{}{}
	expect := [3]interface{}{1, "hello", 0.5}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestNilPointerInputInvalidCase ...
func TestNilPointerInputInvalidCase(t *testing.T) {
	var input *int
	var output int

	err := t2t.Transfer(input, &output)
	if err == nil {
		t.Error("invalid type case check fail")
		return
	}
	t.Log("Success: error message ->", err)
}

// TestNilPointerOutputInvalidCase ...
func TestNilPointerOutputInvalidCase(t *testing.T) {
	input := 999
	var output *int

	err := t2t.Transfer(input, output)
	if err == nil {
		t.Error("invalid type case check fail")
		return
	}
	t.Log("Success: error message ->", err)
}

// TestInterfaceInvalidCase ...
func TestInterfaceInvalidCase(t *testing.T) {
	var input interface{}
	var output int

	err := t2t.Transfer(input, &output)
	if err == nil {
		t.Error("invalid type case check fail")
		return
	}
	t.Log("Success: error message ->", err)
}

// TestInterfaceInvalidCase2 ...
func TestInterfaceInvalidCase2(t *testing.T) {
	var input interface{}
	var output int

	err := t2t.Transfer(&input, &output)
	if err != nil {
		t.Error(err)
		return
	}
}

// UnmarshalerStruct ...
type UnmarshalerStruct struct {
	A int
	B string
}

// UnmarshalerT2T ...
func (m *UnmarshalerStruct) UnmarshalerT2T(input interface{}) error {
	switch input.(type) {
	case int:
		m.A = input.(int)
	case string:
		m.B = input.(string)
	default:
		return fmt.Errorf("invalid type of input: %v", reflect.ValueOf(input).Kind())
	}

	return nil
}

// TestNormalToUnmarshalerTransfer ...
func TestNormalToUnmarshalerTransfer(t *testing.T) {
	input := 999
	var output UnmarshalerStruct
	expect := UnmarshalerStruct{
		A: input,
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// MarshallerStruct ...
type MarshallerStruct struct {
	A string
}

// MarshallerT2T ...
func (m *MarshallerStruct) MarshallerT2T() (interface{}, error) {
	return m.A, nil
}

// TestMarshallerToNormalTransfer ...
func TestMarshallerToNormalTransfer(t *testing.T) {
	input := &MarshallerStruct{
		A: "hello",
	}
	var output string
	expect := input.A

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestMarshallerToUnmarshalerTransfer ...
func TestMarshallerToUnmarshalerTransfer(t *testing.T) {
	input := &MarshallerStruct{
		A: "hello",
	}
	var output UnmarshalerStruct
	expect := UnmarshalerStruct{
		B: input.A,
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestInlineTransfer ...
func TestInlineTransfer(t *testing.T) {
	type Emb struct {
		A int
	}
	type I struct {
		Data Emb `t2t:",inline"`
	}

	input := I{
		Data: Emb{
			A: 99,
		},
	}
	output := map[string]interface{}{}
	expect := map[string]interface{}{
		"A": 99,
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestInlineTransfer2 ...
func TestInlineTransfer2(t *testing.T) {
	type Emb struct {
		A int
	}
	type I struct {
		Data Emb
	}

	input := I{
		Data: Emb{
			A: 99,
		},
	}
	output := map[string]interface{}{}
	expect := map[string]interface{}{
		"Data": map[string]interface{}{
			"A": 99,
		},
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestInlineTransfer3 ...
func TestInlineTransfer3(t *testing.T) {
	type Emb struct {
		A int
	}
	type O struct {
		Data Emb `t2t:",inline"`
	}

	input := map[string]interface{}{
		"A": 99,
	}
	output := O{}
	expect := O{
		Data: Emb{
			A: 99,
		},
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestInlineTransfer4 ...
func TestInlineTransfer4(t *testing.T) {
	type Emb struct {
		A int
	}
	type O struct {
		Data Emb
	}

	input := map[string]interface{}{
		"Data": map[string]interface{}{
			"A": 99,
		},
	}
	output := O{}
	expect := O{
		Data: Emb{
			A: 99,
		},
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestOmitemptyTransfer ...
func TestOmitemptyTransfer(t *testing.T) {
	type I struct {
		A  int `t2t:",omitempty"`
		AA int `t2t:",omitempty"`

		B  string `t2t:",omitempty"`
		BB string `t2t:",omitempty"`

		C  float64 `t2t:",omitempty"`
		CC float64 `t2t:",omitempty"`

		D  map[string]string `t2t:",omitempty"`
		DD map[string]string `t2t:",omitempty"`

		E  *int `t2t:",omitempty"`
		EE *int `t2t:",omitempty"`

		F  []string `t2t:",omitempty"`
		FF []string `t2t:",omitempty"`
	}

	ee := 10
	input := I{
		AA: 999,
		BB: "Hello",
		CC: 5.5,
		DD: map[string]string{
			"key": "",
		},
		EE: &ee,
		FF: []string{
			"Good",
		},
	}
	output := map[string]interface{}{}
	expect := map[string]interface{}{
		"AA": 999,
		"BB": "Hello",
		"CC": 5.5,
		"DD": map[string]interface{}{
			"key": "",
		},
		"EE": 10,
		"FF": []interface{}{
			"Good",
		},
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestOmitemptyTransfer2 ...
func TestOmitemptyTransfer2(t *testing.T) {
	type O struct {
		A  int `t2t:",omitempty"`
		AA int `t2t:",omitempty"`

		B  string `t2t:",omitempty"`
		BB string `t2t:",omitempty"`

		C  float64 `t2t:",omitempty"`
		CC float64 `t2t:",omitempty"`

		D  map[string]string `t2t:",omitempty"`
		DD map[string]string `t2t:",omitempty"`

		E  *int `t2t:",omitempty"`
		EE *int `t2t:",omitempty"`

		F  []string `t2t:",omitempty"`
		FF []string `t2t:",omitempty"`
	}

	ee := 10
	input := map[string]interface{}{
		"AA": 999,
		"BB": "Hello",
		"CC": 5.5,
		"DD": map[string]interface{}{
			"key": "",
		},
		"EE": 10,
		"FF": []interface{}{
			"Good",
		},
	}
	output := O{}
	expect := O{
		AA: 999,
		BB: "Hello",
		CC: 5.5,
		DD: map[string]string{
			"key": "",
		},
		EE: &ee,
		FF: []string{
			"Good",
		},
	}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
	if output.E != nil {
		t.Errorf("unexpected output(%v): expect nil point", reflect.TypeOf(output).String())
		return
	}
}

// TestTypenameOfLiteral ...
func TestTypenameOfLiteral(t *testing.T) {
	type O string
	input := "hello"
	var output O
	expect := input

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if expect != string(output) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// JSONMarshaller ...
type JSONMarshaller struct {
	Hhh string `json:"a"`
}

// MarshalJSON ...
func (m *JSONMarshaller) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Hhh)
}

// TestJSONMarshaller ...
func TestJSONMarshaller(t *testing.T) {

	input := JSONMarshaller{
		Hhh: "hello",
	}
	var output string
	expect := "hello"

	err := t2t.TransferWithTagName(&input, &output, "json")
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// JSONMarshaller2 ...
type JSONMarshaller2 struct {
	Hhh string `json:"a"`
}

// MarshalJSON ...
func (m JSONMarshaller2) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.Hhh)
}

// JSONMarshallerParent ...
type JSONMarshallerParent struct {
	M JSONMarshaller2 `json:"m"`
}

// TestJSONMarshaller2 ...
func TestJSONMarshaller2(t *testing.T) {

	input := JSONMarshallerParent{
		M: JSONMarshaller2{"hello"},
	}
	var output interface{}
	expect := map[string]interface{}{
		"m": "hello",
	}

	err := t2t.TransferWithTagName(input, &output, "json")
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// JSONUnmarshaler ...
type JSONUnmarshaler struct {
	Hhh string `json:"a"`
}

// UnmarshalJSON ...
func (m *JSONUnmarshaler) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &m.Hhh)
}

// TestJSONUnmarshaler ...
func TestJSONUnmarshaler(t *testing.T) {
	input := "hello"
	var output JSONUnmarshaler
	expect := JSONUnmarshaler{
		Hhh: "hello",
	}

	err := t2t.TransferWithTagName(input, &output, "json")
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// JSONUnmarshalerParent ...
type JSONUnmarshalerParent struct {
	M JSONUnmarshaler `json:"m"`
}

// TestJSONUnmarshalerParent ...
func TestJSONUnmarshalerParent(t *testing.T) {
	input := map[string]interface{}{
		"m": "hello",
	}
	var output JSONUnmarshalerParent
	expect := JSONUnmarshalerParent{
		M: JSONUnmarshaler{"hello"},
	}

	err := t2t.TransferWithTagName(input, &output, "json")
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestJSONMarshallerAndUnmarshaler ...
func TestJSONMarshallerAndUnmarshaler(t *testing.T) {
	input := JSONMarshaller{
		Hhh: "hello",
	}
	var output JSONUnmarshaler
	expect := JSONUnmarshaler{
		Hhh: "hello",
	}

	err := t2t.TransferWithTagName(&input, &output, "json")
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// BsonGetter ...
type BsonGetter struct {
	hhh string `bson:"a"`
}

// GetBSON ...
func (m BsonGetter) GetBSON() (interface{}, error) {
	return map[string]interface{}{"bson": m.hhh}, nil
}

// TestBsonGetter ...
func TestBsonGetter(t *testing.T) {

	input := BsonGetter{
		hhh: "hello",
	}
	var output map[string]interface{}
	expect := map[string]interface{}{"bson": "hello"}

	err := t2t.TransferWithTagName(&input, &output, "bson")
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// BsonSetter ...
type BsonSetter struct {
	Hhh string `json:"a"`
}

// SetBSON ...
func (m *BsonSetter) SetBSON(raw bson.Raw) error {
	var v map[string]interface{}
	err := raw.Unmarshal(&v)
	if err != nil {
		fmt.Println("err: ", err)
		return err
	}

	s, exists := v["bson"]
	if !exists {
		return fmt.Errorf("setter field not exists")
	}

	str, ok := s.(string)
	if !ok {
		return fmt.Errorf("invalid typee of setter value")
	}

	m.Hhh = str
	return nil
}

// TestBsonSetter ...
func TestBsonSetter(t *testing.T) {
	input := map[string]interface{}{"bson": "hello"}
	var output BsonSetter
	expect := BsonSetter{
		Hhh: "hello",
	}

	err := t2t.TransferWithTagName(input, &output, "bson")
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestBsonGetterAndSetter ...
func TestBsonGetterAndSetter(t *testing.T) {
	input := BsonGetter{
		hhh: "hello",
	}
	var output BsonSetter
	expect := BsonSetter{
		Hhh: "hello",
	}

	err := t2t.TransferWithTagName(&input, &output, "bson")
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// TestWeakTypeTransfer ...
func TestWeakTypeTransfer(t *testing.T) {
	input := 999
	var output string
	expect := "999"

	tr := t2t.NewTransformer()
	tr.EnableWeakTypeTransfer()

	err := tr.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(output, expect) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
}

// PointerElement ...
type PointerElement struct {
	A []*int64 `t2t:"a"`
}

func TestPointerElement(t *testing.T) {
	v := int64(1)
	input := &PointerElement{[]*int64{
		&v,
	}}
	var output PointerElement
	expect := &PointerElement{[]*int64{
		&v,
	}}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	if len(expect.A) != len(output.A) {
		t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
			reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
		return
	}
	for i := 0; i < len(expect.A); i++ {
		if *expect.A[i] != *output.A[i] {
			t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
				reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
			return
		}
	}
}

func TestPointerElement2(t *testing.T) {
	ip1 := 1
	ip2 := 2
	input := []*int{&ip1, &ip2}
	var output [2]*int64
	ep1 := int64(1)
	ep2 := int64(2)
	expect := [2]*int64{&ep1, &ep2}

	err := t2t.Transfer(input, &output)
	if err != nil {
		t.Error(err)
		return
	}

	for i := 0; i < len(expect); i++ {
		if *expect[i] != *output[i] {
			t.Errorf("unexpected output(%v): %+v, expect(%v): %+v",
				reflect.TypeOf(output).String(), output, reflect.TypeOf(expect).String(), expect)
			return
		}
	}
}
