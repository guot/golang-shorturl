package shorturl

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"hash/crc32"

	"compress/zlib"
	"fmt"
	"io"

	"testing"
	"time"
)

type Buffer struct {
	bytes.Buffer
	io.ReaderFrom // conflicts with and hides bytes.Buffer's ReaderFrom.
	io.WriterTo   // conflicts with and hides bytes.Buffer's WriterTo.
}

func TestMain(t *testing.T) {
	var input = []byte("data:image/png;base64,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")
	buf := new(Buffer)
	compressor, _ := zlib.NewWriterLevel(buf, zlib.BestCompression)

	compressor.Write(input)
	compressor.Close()
	fmt.Println(buf.Bytes())
	fmt.Println(len(buf.Bytes()))
	fmt.Println(len(input))
	//	b := bytes.NewReader(buf.Bytes())
	// 	b := bytes.NewReader(buf)
	//	tt := []byte{
	//		0x78, 0x9c, 0x4b, 0xcf, 0xcf, 0x4f, 0x49, 0xaa,
	//			0x4c, 0xd5, 0x51, 0x28, 0xcf, 0x2f, 0xca, 0x49,
	//			0x01, 0x00, 0x28, 0xa5, 0x05, 0x5e,
	//	}
	tt := []byte(buf.Bytes())
	b := new(bytes.Buffer)
	//	in := bytes.NewReader(tt)
	zlib, err := zlib.NewReaderDict(buf, nil)
	if err != nil {

		t.Errorf("%s: NewReader: %s", tt, err)

	}
	//	defer zlib.Close()
	b.Reset()
	io.Copy(b, zlib)
	if err != nil {

		t.Errorf("%s: io.Copy: %v want %v", err)

	}
	s := b.String()
	fmt.Println("11:", s)

	//var b bytes.Buffer
	//	w := zlib.NewWriter(&b)
	//	w.Write([]byte("hello, world\n"))
	//	w.Close()
	//	r, _ := zlib.NewReader(&b)
	//	io.Copy(os.Stdout, r)
	//	r.Close()
}
func ExampleEncoding_EncodeToString() {
	data := []byte("any + old & datadsfasfasdfasdfaslkjweiojasdfkjwe rfiowjalsdjkf weijlajsdf isajfewl")
	str := base64.StdEncoding.EncodeToString(data)
	md := md5.New()
	md.Write([]byte("com_ceo1")) // 需要加密的字符串为 123456
	time.Now().String()
	crc := crc32.NewIEEE()
	crc.Write([]byte("com_ceo1"))
	fmt.Println("asdfsadf:", hex.EncodeToString(crc.Sum(nil)))
	fmt.Println("asdfsadf:", hex.EncodeToString(md.Sum(nil)), ":", str)
	// Output:
	// YW55ICsgb2xkICYgZGF0YQ==
}
