package shorturl

import (
	"bytes"
	"compress/lzw"
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
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
