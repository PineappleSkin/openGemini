package compress

import (
	"github.com/openGemini/openGemini/lib/compress/mlf"
	"github.com/openGemini/openGemini/lib/util"
)

var comp = &mlf.Compressor{}
var decomp = &mlf.Decompressor{}

func MLFCompress(dst []byte, src []byte) ([]byte, error) {
	floats := util.Bytes2Float64Slice(src)
	comp.Prepare(floats)
	return comp.Encode(dst[:0], floats), nil
}

func MLFDecompress(dst []byte, src []byte) ([]byte, error) {
	floats := decomp.Decode(src)
	return util.Float64Slice2byte(floats), nil
}
