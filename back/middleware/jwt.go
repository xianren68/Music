package middleware

import (
	"net/http"
	"strings"
	"time"

	"Music/config"
	"Music/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// 定义jwt结构体
type MyClaims struct {
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.RegisteredClaims
}

// 创建token
func CreateToken(username string, role int) (string, int) {
	// 初始化一个声明结构体
	info := MyClaims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			// 设置过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			// 设置颁发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 主题
			Subject: "music",
		},
	}
	// 生成token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, info).SignedString([]byte(config.JwtKey))
	if err != nil {
		return "", utils.ERROR
	}
	return token, utils.SUCCESS
}

// 解析token
func ParseToken(token string) (*MyClaims, int) {
	// 解析token
	tokenstr, err := jwt.ParseWithClaims(token, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JwtKey), nil
	})
	if err != nil {
		return nil, utils.ERROR
	}
	// token是否为空
	if claims, ok := tokenstr.Claims.(*MyClaims); ok && tokenstr.Valid {
		return claims, utils.SUCCESS
	} else {
		return nil, utils.ERROR
	}

}

// token中间件
func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		code := utils.SUCCESS
		// 不存在
		if tokenHeader == "" {
			code = utils.ERROR_TOKEN_NOT_EXIST
			returnError(c, code)
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		// token格式出错
		if len(checkToken) != 2 || checkToken[0] != "Bearer" {
			code = utils.ERROR_TOKEN_FORMAT
			returnError(c, code)
			return
		}
		token, err := ParseToken(checkToken[1])
		if err == utils.ERROR {
			code = utils.ERROR_TOKEN_WRONG
			returnError(c, code)
			return
		}
		// token过期
		if token.VerifyExpiresAt(time.Now(), false) {
			returnError(c, utils.ERROR_TOKEN_EXPIRE)
			return
		}
		// 将token中的username和role设置到context中
		c.Set("username", token.Username)
		c.Set("role", token.Role)
	}
}

// 返回错误信息
func returnError(c *gin.Context, code int) {
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    utils.CodeMsg[code],
	})
	c.Abort()
}
