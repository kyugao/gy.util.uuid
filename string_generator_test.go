package ugener

import "testing"

func TestGenString(t *testing.T) {
	t.Log(StrGen.RandHex(10))
	t.Log(StrGen.RandLow(10))
	t.Log(StrGen.RandUp(10))
	t.Log(StrGen.RandNum(10))
}
