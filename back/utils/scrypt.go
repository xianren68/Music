package utils

import (
	"crypto/sha1"
	"fmt"
)

// 加密函数
func Scrypt(str string) string {
	hashValue := sha1.New()
	hashValue.Write([]byte(str))
	res := hashValue.Sum(nil)
	// 转化为16进制
	return fmt.Sprintf("%x", res)
}
