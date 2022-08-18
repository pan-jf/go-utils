package t2t

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrToSlice(t *testing.T) {
	input := "hello,world,t2t"
	output := StrToSlice(input, ",")
	expected := []string{"hello", "world", "t2t"}
	assert.Equal(t, expected, output)
}

func TestSliceToStr(t *testing.T) {
	input := []string{"hello", "world", "t2t"}
	output := SliceToString(input, ",")
	expected := "hello,world,t2t"
	assert.Equal(t, expected, output)
}

func TestStrToInt64Slice(t *testing.T) {
	input := "55,96,32,100,5,1.2,word"
	output := StrToInt64Slice(input, ",")
	expected := []int64{55, 96, 32, 100, 5, 0, 0}
	assert.Equal(t, expected, output)
}

func TestStrToFloat64Slice(t *testing.T) {
	input := "55,96,32,100,5,1.2,word"
	output := StrToFloat64Slice(input, ",")
	expected := []float64{55, 96, 32, 100, 5, 1.2, 0}
	assert.Equal(t, expected, output)
}

func TestStrToInt64(t *testing.T) {
	input := "55"
	output := StrToInt64(input)
	expected := int64(55)
	assert.Equal(t, expected, output)

	input = "55.22"
	output = StrToInt64(input)
	expected = int64(0)
	assert.Equal(t, expected, output)

	input = "word"
	output = StrToInt64(input)
	expected = int64(0)
	assert.Equal(t, expected, output)
}

func TestInt64ToStr(t *testing.T) {
	input := int64(55)
	output := Int64ToStr(input)
	expected := "55"
	assert.Equal(t, expected, output)
}

func TestStrToFloat32(t *testing.T) {
	input := "55"
	output := StrToFloat32(input)
	expected := float32(55)
	assert.Equal(t, expected, output)

	input = "55.22"
	output = StrToFloat32(input)
	expected = float32(55.22)
	assert.Equal(t, expected, output)

	input = "word"
	output = StrToFloat32(input)
	expected = float32(0)
	assert.Equal(t, expected, output)
}

func TestFloat32ToStr(t *testing.T) {
	input := float32(55.22)
	output := Float32ToStr(input)
	expected := "55.22"
	assert.Equal(t, expected, output)
}

func TestStrToFloat64(t *testing.T) {
	input := "55"
	output := StrToFloat64(input)
	expected := float64(55)
	assert.Equal(t, expected, output)

	input = "55.22"
	output = StrToFloat64(input)
	expected = 55.22
	assert.Equal(t, expected, output)

	input = "word"
	output = StrToFloat64(input)
	expected = float64(0)
	assert.Equal(t, expected, output)
}

func TestFloat64ToStr(t *testing.T) {
	input := 55.22
	output := Float64ToStr(input)
	expected := "55.22"
	assert.Equal(t, expected, output)
}

func TestFloat64ToPerStr(t *testing.T) {
	input1 := 55.22333
	input2 := 100.00
	output := Float64ToPerStr(input1, input2)
	expected := "55.22%"
	assert.Equal(t, expected, output)
}

func TestFloat32ToString(t *testing.T) {
	input1 := float32(55.22333)
	input2 := 2
	output := Float32ToString(input1, input2)
	expected := "55.22"
	assert.Equal(t, expected, output)
}

func TestFloat64ToString(t *testing.T) {
	input1 := 55.22333
	input2 := 2
	output := Float64ToString(input1, input2)
	expected := "55.22"
	assert.Equal(t, expected, output)
}
