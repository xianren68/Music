package test

import (
	"Music/utils"
	"fmt"
	"testing"
)

func TestScrypt(t *testing.T) {
	r1 := utils.Scrypt("123456")
	r2 := utils.Scrypt("123456")
	fmt.Println(r1)
	if r1 == r2 {
		t.Logf("%s", "ok")
	} else {
		t.Logf("%s", "fail")
	}
}
