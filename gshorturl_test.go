package shorturl

import (
	"bytes"
	"compress/lzw"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"testing"
	"time"
)

func aTestMain(t *testing.T) {
	//	Encoder("")
	litw := 8

	var data = []byte("hello world")

	fmt.Printf("input: %#v\n", string(data))

	var buf bytes.Buffer

	com := lzw.NewWriter(&buf, lzw.LSB, litw)

	w, err := com.Write(data)

	if err != nil {
		fmt.Println("write error:", err)
	}

	fmt.Println("wrote", w, "bytes")
	fmt.Printf("D;ewr:%x\n", buf.Len())

	com.Close()

	var output = make([]byte, len(data))

	dec := lzw.NewReader(&buf, lzw.LSB, litw)

	r, err := dec.Read(output)

	if err != nil {
		fmt.Println("read error:", err)
	}

	fmt.Println("read", r, "bytes")
	fmt.Printf("::out:%s", output)

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
