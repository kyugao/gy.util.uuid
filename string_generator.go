package ugen

import (
	"encoding/hex"
	"math/rand"
	"time"
)

var StrGen strGen

type strGen struct{}

var numbers = []byte("0123456789")
var letters = []byte("abcdefghjkmnpqrstuvwxyz123456789")
var longLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ=_")

func init() {
	rand.Seed(time.Now().Unix())
}

// RandNum 随机字符串，包含 0~9
func (strGen) RandNum(n int) string {
	if n <= 0 {
		return string([]byte{})
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return string([]byte{})
	}
	for i, _ := range b {
		arc = uint8(rand.Intn(10))
		b[i] = numbers[arc]
	}
	return string(b)
}

// RandLow 随机字符串，包含 1~9 和 a~z - [i,l,o]
func (strGen) RandLow(n int) string {
	if n <= 0 {
		return string([]byte{})
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return string([]byte{})
	}
	for i, x := range b {
		arc = x & 31
		b[i] = letters[arc]
	}
	return string(b)
}

// RandUp 随机字符串，包含 英文字母和数字附加=_两个符号
func (strGen) RandUp(n int) string {
	if n <= 0 {
		return string([]byte{})
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return string([]byte{})
	}
	for i, x := range b {
		arc = x & 63
		b[i] = longLetters[arc]
	}
	return string(b)
}

// RandHex 生成16进制格式的随机字符串
func (strGen) RandHex(n int) string {
	if n <= 0 {
		return string([]byte{})
	}
	var need int
	if n&1 == 0 { // even
		need = n
	} else { // odd
		need = n + 1
	}
	size := need / 2
	dst := make([]byte, need)
	src := dst[size:]
	if _, err := rand.Read(src[:]); err != nil {
		return string([]byte{})
	}
	hex.Encode(dst, src)
	return string(dst[:n])
}
