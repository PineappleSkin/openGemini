package compress

import (
	"fmt"
	"testing"
	"time"
)

type CompressPerformanceReport struct {
	Algorithm string
	Speed     [3]float64
}

func (r *CompressPerformanceReport) String() string {
	return fmt.Sprintf("%s,%.02f,%.02f,%.02f", r.Algorithm, r.Speed[0], r.Speed[1], r.Speed[2])
}

func TestCompressionPerformance(t *testing.T) {
	testCompressionPerformance()
}

func testCompressionPerformance() {
	testData := generateTestData()

	cases := []struct {
		name         string
		compressFn   func(dst, src []byte) ([]byte, error)
		decompressFn func(dst, src []byte) ([]byte, error)
	}{
		{"LZ4", CompressLZ4, DecompressLZ4},
		{"MLF", MLFCompress, MLFDecompress},
		{"Snappy", CompressSnappy, DecompressSnappy},
		{"ZSTD", CompressZSTD, DecompressZSTD},
	}

	var compressReports []*CompressPerformanceReport
	var decompressReports []*CompressPerformanceReport
	defer func() {
		for _, report := range compressReports {
			fmt.Println(report)
		}
		fmt.Println("--------")
		for _, report := range decompressReports {
			fmt.Println(report)
		}
	}()

	for _, c := range cases {
		report := &CompressPerformanceReport{
			Algorithm: c.name,
			Speed:     [3]float64{},
		}
		compressReports = append(compressReports, report)

		for idx, data := range testData {
			fmt.Println("compress", c.name, idx)
			report.Speed[idx] = runCompressionPerformance(c.compressFn, data)
		}
	}

	for _, c := range cases {
		report := &CompressPerformanceReport{
			Algorithm: c.name,
			Speed:     [3]float64{},
		}
		decompressReports = append(decompressReports, report)

		for idx, data := range testData {
			fmt.Println("decompress", c.name, idx)
			compressed, _ := c.compressFn(nil, data)
			report.Speed[idx] = runCompressionPerformance(c.decompressFn, compressed)
		}
	}
}

func runCompressionPerformance(fn func(dst, src []byte) ([]byte, error), data []byte) float64 {
	begin := time.Now()
	var buf []byte
	var err error
	const N = 100000

	for range N {
		buf, err = fn(buf[:0], data)
		assertError(err)
	}

	totalSize := len(data) * N
	return float64(totalSize) / (time.Since(begin).Seconds() * 1e6)
}

func assertError(err error) {
	if err != nil {
		panic(err)
	}
}
