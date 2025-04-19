package compress

import (
	"math"
	"math/rand"
	"unsafe"
)

// GenerateFloats 生成指定数量的50-60之间的浮点数（保留两位小数）
func GenerateFloats(n int) []float64 {
	result := make([]float64, n)
	for i := range result {
		// 生成50.00-60.00范围内的随机数
		val := 50.0 + math.Floor(rand.Float64()*1000.0)/100
		result[i] = val
	}
	return result
}

// GenerateUnitFloats 生成指定数量的0-1之间的浮点数（保留六位小数）
func GenerateUnitFloats(n int) []float64 {
	result := make([]float64, n)
	for i := range result {
		// 生成0.000000-1.000000范围内的随机数
		val := math.Floor(rand.Float64()*1e6) / 1e6
		result[i] = val
	}
	return result
}

// GenerateIncrementFloats 生成指定数量的递增浮点数，最小值为minValue，增量为10-100之间的整数
func GenerateIncrementFloats(n int) []float64 {
	if n <= 0 {
		return nil
	}
	result := make([]float64, n)
	result[0] = 200
	for i := 1; i < n; i++ {
		step := float64(rand.Intn(91) + 10) // 生成100-1000的随机整数
		result[i] = result[i-1] + step
	}
	return result
}

func Bytes2Float64Slice(b []byte) []float64 {
	if len(b) == 0 {
		return nil
	}
	return unsafe.Slice((*float64)(unsafe.Pointer(&b[0])), len(b)/8)
}

func Float64Slice2byte(b []float64) []byte {
	if len(b) == 0 {
		return nil
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(&b[0])), len(b)*8)
}
