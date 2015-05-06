package shorturl

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"strings"
)

//	"bufio"
//	"bytes"

const (
	//	ALPHABET = "0123456789abcdfghijkmlnopqrstuvwxyzABCDFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+{}:\\-=[];'<>?,./"
	ALPHABET = "0123456789abcdfghijkmlnopqrstuvwxyzABCDFGHIJKLMNOPQRSTUVWXYZ~!@#$%^&*()_+{}:\\-=[];'<>?,./"
	BASE     = len(ALPHABET)
)

//生成短URL
func Encoder(str string) {

	bb := new(bytes.Buffer)
	//	var buf = make([]byte, 100)
	//	for num > 0 {
	//		buf = append(buf, ALPHABET[num])
	//	}
	//	src := "http://guotao98@.163.com"
	//	var num int64

	//	var sbbbb = []byte("1222")
	var sbbbb = []byte("data:image/png;base64,+ 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")
	fmt.Println("len src size", len(sbbbb))

	w, _ := zlib.NewWriterLevelDict(bb, zlib.BestCompression, sbbbb)

	// 	w.Write(sbbbb)
	w.Close()
	fmt.Println("len zip size", len(bb.Bytes()))
	//78f900320032010000ffff000000012029584434
	//78f901f400c8026140000000ffff01f400c82029584884
	//78f984fdeb1e010000ffff000000012029618429
	//78f9be55c81a010000ffff000000012029633109
	fmt.Printf("len zip size \n%x", bb.Bytes())
	var val uint32

	err := binary.Read(bb, binary.BigEndian, &val)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println(val)
	var strbuf bytes.Buffer
	for val > 0 {
		mod := val % uint32(BASE)
		//fmt.Printf("%c",ALPHABET[mod])

		fmt.Printf("val:%c_%d\n", ALPHABET[mod], val)
		strbuf.WriteByte(ALPHABET[mod])
		val = val / uint32(BASE)
	}
	result := strbuf.String()
	fmt.Println(result)
	var sum uint32
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Printf("sum:%c_%d\n", result[i], sum)
		sum = sum*uint32(BASE) + uint32(strings.IndexByte(ALPHABET, result[i]))
	}
	bb.Reset()
	fmt.Println("sum:", sum)
	binary.Write(bb, binary.BigEndian, sum)
	fmt.Print(len(bb.Bytes()))

	buf := new(bytes.Buffer)

	err1 := binary.Write(buf, binary.BigEndian, sum)
	if err1 != nil {
		fmt.Println("binary.Write failed:", err1)
	}
	fmt.Printf("\n% x", buf.Bytes())

	//	w1:=zlib.NewWriterLevelDict();

	//	fmt.Println(hex.Dump([]byte("hhasdfsd")))

	//	for _, index := range src {
	//		fmt.Println("hh11:%s", string(index))
	//		pos := strings.IndexRune(ALPHABET, index)
	//		num = num*int64(BASE) + int64(pos)
	//		fmt.Println("hh11:%d", num)
	//	}
	//	str1 := "中国"
	//	str2 := "中"
	//	str3 := "人"
	//	str4 := "民"
	//	dst := make([]byte, hex.DecodedLen(len(str1)))
	//	hex.Decode(dst, []byte(str1))
	//	sst, _ := hex.DecodeString(str1)

	//	fmt.Println("htllo1:" + hex.Dump(sst))

	//	fmt.Println("htllo2:" + hex.EncodeToString([]byte(str1)))
	//	fmt.Println("htllo3:" + hex.EncodeToString([]byte(str2)))
	//	fmt.Println("htllo4:" + hex.EncodeToString([]byte(str3)))
	//	fmt.Println("htllo5:" + hex.EncodeToString([]byte(str4)))
	//	fmt.Println("base len%d", BASE)
}

//对短URL进行还原
func Decoder() {

}

//public class ShortURL {

//    public static final String ALPHABET = "23456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ-_";
//    public static final int BASE = ALPHABET.length();

//    public static String encode(int num) {
//        StringBuilder str = new StringBuilder();
//        while (num > 0) {
//            str.insert(0, ALPHABET.charAt(num % BASE));
//            num = num / BASE;
//        }
//        return str.toString();
//    }

//    public static int decode(String str) {
//        int num = 0;
//        for (int i = 0; i < str.length(); i++) {
//            num = num * BASE + ALPHABET.indexOf(str.charAt(i));
//        }
//        return num;
//    }

//}
