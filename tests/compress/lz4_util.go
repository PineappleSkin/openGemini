package compress

import (
	"fmt"

	"github.com/pierrec/lz4/v4"
)

// CompressLZ4 使用LZ4算法压缩数据
func CompressLZ4(dst, src []byte) ([]byte, error) {
	// 预分配内存校验
	ml := lz4.CompressBlockBound(len(src))
	if cap(dst) < ml {
		dst = make([]byte, ml)
	} else {
		dst = dst[:ml]
	}

	n, err := lz4.CompressBlock(src, dst, nil)
	if err != nil {
		return nil, fmt.Errorf("lz4 compress failed: %w", err)
	}
	if n <= 0 {
		return nil, fmt.Errorf("lz4 compress invalid output size: %d", n)
	}
	return dst[:n], nil
}

// DecompressLZ4 使用LZ4算法解压数据
func DecompressLZ4(dst, compressed []byte) ([]byte, error) {
	// 自动扩容逻辑
	dl := len(compressed) * 10 // 保守估计解压尺寸
	if cap(dst) < dl {
		dst = make([]byte, dl)
	}

	n, err := lz4.UncompressBlock(compressed, dst[:dl])
	if err != nil {
		return nil, fmt.Errorf("lz4 decompress failed: %w", err)
	}
	if n <= 0 {
		return nil, fmt.Errorf("lz4 decompress invalid output size: %d", n)
	}
	return dst[:n], nil
}
