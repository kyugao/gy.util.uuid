package ugener

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var srcLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ=")

func randomSecret() (secret []byte) {
	tempLetters := make([]byte, len(srcLetters))
	copy(tempLetters, srcLetters)
	for len(tempLetters) > 0 {
		pos, _ := rand.Int(rand.Reader, big.NewInt(int64(len(tempLetters))))
		secret = append(secret, tempLetters[pos.Int64()])
		tempLetters = append(tempLetters[:pos.Int64()], tempLetters[pos.Int64()+1:]...)
	}
	fmt.Println(string(secret))
	return
}

func Short63(id int64) (output string) {
	secret := randomSecret()
	secretLen := int64(len(secret))
	var outByte []byte
	for id != 0 {
		m := id % secretLen
		outByte = append(outByte, secret[m])
		id = (id - m) / secretLen
	}
	output = string(outByte)
	return
}
