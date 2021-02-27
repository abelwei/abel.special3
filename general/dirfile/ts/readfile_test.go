package ts

import (
	"github.com/abelwei/abel.special3/general/dirfile"
	"testing"
)

func TestReadfileDft(t *testing.T) {
	txtling := dirfile.ReadfileLine("t1.txt")
	if len(txtling)>0 {
		for i, txt := range txtling {
			t.Log(i, ":", txt)
		}
	}else{
		t.Error("eq 0?")
	}
}