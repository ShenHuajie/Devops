package middleware

import (
	"Devops/utils"
	"Devops/utils/errMsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	jwt.StandardClaims
}


// SetToken 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10 * time.Hour)
	SetClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "go-devops",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errMsg.ERROR
	}
	return token, errMsg.SUCCESS
}

// CheckToken CheckTOken 验证token
func CheckToken(token string) (*MyClaims, int) {

	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errMsg.SUCCESS
	} else {
		return nil, errMsg.ERROR
	}

}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header.Get("Authorization")
		code := errMsg.SUCCESS
		if tokenHeader == "" {
			code = errMsg.ErrorTokenExist
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg": errMsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errMsg.ErrorTokenWrong
			ctx.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode == errMsg.ERROR {
			code = errMsg.ErrorTokenWrong
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg": errMsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errMsg.ErrorTokenRuntime
			ctx.Abort()
			return
		}

		ctx.Set("username", key.Username)
		ctx.Next()
	}
}