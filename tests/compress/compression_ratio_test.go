package compress

import (
	"bytes"
	"fmt"
	"testing"
)

type CompressRatioReport struct {
	Algorithm string
	Ratios    [3]float64
}

func (r *CompressRatioReport) String() string {
	return fmt.Sprintf("%s,%.02f,%.02f,%.02f", r.Algorithm, r.Ratios[0], r.Ratios[1], r.Ratios[2])
}

// 生成混合测试数据（JSON + 二进制 + 时序数据）
func generateTestData() [][]byte {
	var data [][]byte
	// 生成三种类型的浮点数组
	floats := GenerateFloats(1024)
	data = append(data, Float64Slice2byte(floats))

	unitFloats := GenerateUnitFloats(1024)
	data = append(data, Float64Slice2byte(unitFloats))

	incrFloats := GenerateIncrementFloats(1024)
	data = append(data, Float64Slice2byte(incrFloats))
	return data
}

func BenchmarkCompressionRatios(b *testing.B) {
	testData := generateTestData()
	b.ResetTimer()

	cases := []struct {
		name       string
		compressFn func(dst, src []byte) ([]byte, error)
	}{
		{"LZ4", CompressLZ4},
		{"MLF", MLFCompress},
		{"Snappy", CompressSnappy},
		{"ZSTD", CompressZSTD},
	}

	var dst []byte
	dataTypes := []string{"floats", "unit_floats", "incr_floats"}
	var reports []*CompressRatioReport
	defer func() {
		for _, report := range reports {
			fmt.Println(report)
		}
	}()

	for _, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			report := &CompressRatioReport{
				Algorithm: c.name,
				Ratios:    [3]float64{},
			}
			reports = append(reports, report)

			for i, data := range testData {
				b.Run(dataTypes[i], func(b *testing.B) {
					var totalSize, compressedSize int
					for i := 0; i < b.N; i++ {
						dst, _ = c.compressFn(dst[:0], data)
						totalSize += len(data)
						compressedSize += len(dst)
					}

					report.Ratios[i] = float64(totalSize) / float64(compressedSize)
					b.ReportMetric(
						float64(totalSize)/float64(compressedSize),
						"compression-ratio",
					)
				})
			}
		})
	}
}

func TestSnappyCompression(t *testing.T) {
	testData := generateTestData()

	t.Run("valid_data", func(t *testing.T) {
		for i, data := range testData {
			compressed, err := CompressSnappy(nil, data)
			if err != nil {
				t.Fatalf("压缩失败: %v", err)
			}

			decompressed, err := DecompressSnappy(nil, compressed)
			if err != nil {
				t.Fatalf("解压失败: %v", err)
			}

			if !bytes.Equal(decompressed, data) {
				t.Fatalf("数据类型%d解压后数据不一致", i)
			}
		}
	})
}

func TestLZ4Compression(t *testing.T) {
	testData := generateTestData()

	t.Run("valid_data", func(t *testing.T) {
		for i, data := range testData {
			compressed, err := CompressLZ4(nil, data)
			if err != nil {
				t.Fatalf("LZ4压缩失败: %v", err)
			}

			decompressed, err := DecompressLZ4(nil, compressed)
			if err != nil {
				t.Fatalf("LZ4解压失败: %v", err)
			}

			if !bytes.Equal(decompressed, data) {
				t.Fatalf("LZ4数据类型%d解压后数据不一致", i)
			}
		}
	})
}

func TestZSTDCompression(t *testing.T) {
	testData := generateTestData()

	t.Run("valid_data", func(t *testing.T) {
		for i, data := range testData {
			compressed, err := CompressZSTD(nil, data)
			if err != nil {
				t.Fatalf("ZSTD压缩失败: %v", err)
			}

			decompressed, err := DecompressZSTD(nil, compressed)
			if err != nil {
				t.Fatalf("ZSTD解压失败: %v", err)
			}

			if !bytes.Equal(decompressed, data) {
				t.Fatalf("ZSTD数据类型%d解压后数据不一致", i)
			}
		}
	})
}

func TestMLFCompression(t *testing.T) {
	testData := generateTestData()

	t.Run("valid_data", func(t *testing.T) {
		for i, data := range testData {
			compressed, err := MLFCompress(nil, data)
			if err != nil {
				t.Fatalf("MLF压缩失败: %v", err)
			}

			decompressed, err := MLFDecompress(nil, compressed)
			if err != nil {
				t.Fatalf("MLF解压失败: %v", err)
			}

			if !bytes.Equal(decompressed, data) {
				t.Fatalf("MLF数据类型%d解压后数据不一致", i)
			}
		}
	})
}
