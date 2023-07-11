package test

import (
	"Music/utils"
	"testing"
)

func TestScrypt(t *testing.T) {
	r1 := utils.Scrypt("123456")
	r2 := utils.Scrypt("123456")
	if r1 == r2 {
		t.Logf("%s", "sucess")
	} else {
		t.Logf("%s", "fail")
	}
}
