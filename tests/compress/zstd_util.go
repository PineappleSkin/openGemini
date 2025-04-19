package compress

import (
	"github.com/DataDog/zstd"
)

// CompressZSTD 使用Zstandard算法压缩数据
func CompressZSTD(dst, src []byte) ([]byte, error) {
	return zstd.CompressLevel(dst[:0], src, zstd.BestSpeed)
}

// DecompressZSTD 使用Zstandard算法解压数据
func DecompressZSTD(dst, compressed []byte) ([]byte, error) {
	return zstd.Decompress(dst[:0], compressed)
}
