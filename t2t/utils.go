package t2t

import (
	"fmt"
	"strconv"
	"strings"
)

// StrToSlice 字符串转换成字符切片
func StrToSlice(str, sep string) []string {
	if len(str) == 0 {
		return nil
	}
	strArray := strings.Split(str, sep)
	if len(strArray) == 1 && len(strArray[0]) == 0 {
		return nil
	}
	result := make([]string, 0, len(strArray))
	for i := 0; i < len(strArray); i++ {
		strItem := strings.TrimSpace(strArray[i])
		result = append(result, strItem)
	}
	return result
}

// SliceToString 字符切片转换成字符串
func SliceToString(strS []string, sep string) string {
	return strings.Join(strS, sep)
}

// StrToInt64Slice 字符串转换成整型切片
func StrToInt64Slice(str, sep string) []int64 {
	if len(str) == 0 {
		return nil
	}
	strArray := strings.Split(str, sep)
	if len(strArray) == 1 && len(strArray[0]) == 0 {
		return nil
	}
	result := make([]int64, 0, len(strArray))
	for i := 0; i < len(strArray); i++ {
		result = append(result, StrToInt64(strArray[i]))
	}
	return result
}

// StrToFloat64Slice 字符串转换成浮点型切片
func StrToFloat64Slice(str, sep string) []float64 {
	if len(str) == 0 {
		return nil
	}
	strArray := strings.Split(str, sep)
	if len(strArray) == 1 && len(strArray[0]) == 0 {
		return nil
	}
	result := make([]float64, 0, len(strArray))
	for i := 0; i < len(strArray); i++ {
		result = append(result, StrToFloat64(strArray[i]))
	}
	return result
}

// StrToInt64 字符串转换成64位整型
func StrToInt64(str string) int64 {
	if str == "" {
		return 0
	}
	str = strings.TrimSpace(str)
	intVal, _ := strconv.ParseInt(str, 10, 64)
	return intVal
}

// Int64ToStr 64位整型转换为字符串
func Int64ToStr(input int64) string {
	return strconv.FormatInt(input, 10)
}

// StrToFloat32 字符串转换为32位浮点型
func StrToFloat32(str string) float32 {
	if str == "" {
		return 0.0
	}
	intVal, _ := strconv.ParseFloat(str, 32)
	return float32(intVal)
}

// Float32ToStr 32位浮点型转换为字符串
func Float32ToStr(input float32) string {
	return fmt.Sprint(input)
}

// StrToFloat64 字符串转换为64位浮点型
func StrToFloat64(str string) float64 {
	if str == "" {
		return 0.0
	}
	str = strings.TrimSpace(str)
	intVal, _ := strconv.ParseFloat(str, 64)
	return intVal
}

// Float64ToStr 64位浮点型转换为字符串
func Float64ToStr(input float64) string {
	return fmt.Sprint(input)
}

// Float64ToPerStr 2个64位浮点型得到一个百分比字符串
func Float64ToPerStr(molecule, denominator float64) string {
	if molecule == 0 || denominator == 0 {
		return "0"
	}
	rate := molecule / denominator * 100
	rate, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", rate), 64)
	rateStr := fmt.Sprintf("%v", rate) // "1.23"
	rateStr += "%"
	return rateStr
}

// Float32ToString 32位浮点型 + 期望保留位数转换为字符串
func Float32ToString(input float32, precision int) string {
	return Float64ToString(float64(input), precision)
}

// Float64ToString 64位浮点型 + 期望保留位数转换为字符串
func Float64ToString(input float64, precision int) string {
	return strconv.FormatFloat(input, 'f', precision, 64)
}
