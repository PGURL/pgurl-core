package gen

import (
	"testing"
)

func TestGenStr(t *testing.T) {
	for gslen := 1; gslen <= 10; gslen++ {
		gs := GenRandomStr(gslen)
		if len(gs) == gslen {
			t.Logf("%s len is %d", gs, gslen)
		} else {
			t.Error("len error")
		}
	}
}
