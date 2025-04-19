package compress

import (
	"fmt"

	"github.com/golang/snappy"
)

// CompressSnappy 使用Snappy算法压缩数据
func CompressSnappy(dst, src []byte) ([]byte, error) {
	ml := snappy.MaxEncodedLen(len(src))
	if cap(dst) < ml {
		dst = make([]byte, ml)
	}
	return snappy.Encode(dst[:ml], src), nil
}

// DecompressSnappy 使用Snappy算法解压数据
func DecompressSnappy(dst, compressed []byte) ([]byte, error) {
	dl, _ := snappy.DecodedLen(compressed)
	if cap(dst) < dl {
		dst = make([]byte, dl)
	}

	dst, err := snappy.Decode(dst[:dl], compressed)
	if err != nil {
		return nil, fmt.Errorf("snappy decompress failed: %w", err)
	}
	return dst, nil
}
