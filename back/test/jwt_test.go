package test

import (
	_ "Music/config"
	"Music/middleware"
	"testing"
)

// 注意：需要修改config路径
func TestParseToken(t *testing.T) {
	token, _ := middleware.CreateToken("xianren", 1)
	claim, _ := middleware.ParseToken(token)
	if claim.Role == 1 && claim.Username == "xianren" {
		t.Log("success")
	} else {
		t.Error("fail")
	}

}
