package jwt

import (
	"errors"
	jwtgo "github.com/dgrijalva/jwt-go"
	"time"
)

const (
	// SignKey ：加密盐,使用时，要换成自己的密钥
	SignKey     = "liam4e3052a9fd6450077a3287003a-mb1907@jm@!$$v1"
	TokenIssuer = "Liam"
)

var Server *JWT

func init() {
	Server = NewJWT("")
}

var (
	// ErrorTokenExpired : 过期了
	ErrorTokenExpired = errors.New("Token is expired")
	// ErrorTokenNotValidYet ：还没有生效
	ErrorTokenNotValidYet = errors.New("Token not active yet")
	// ErrorTokenMalformed  ：That's not even a token
	ErrorTokenMalformed = errors.New("That's not even a token")
	// ErrorTokenInvalid ：Couldn't handle this token
	ErrorTokenInvalid = errors.New("Couldn't handle this token")
)

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserID int64  `json:"id"`
	Name   string `json:"name"`
	jwtgo.StandardClaims
}

// NewJWT 新建一个jwt实例
func NewJWT(key string) *JWT {
	if key == "" {
		key = SignKey
	}
	return &JWT{
		[]byte(key),
	}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims *CustomClaims) (string, error) {
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwtgo.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwtgo.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwtgo.ValidationError); ok {
			if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
				return nil, ErrorTokenMalformed
			} else if ve.Errors&jwtgo.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrorTokenExpired
			} else if ve.Errors&jwtgo.ValidationErrorNotValidYet != 0 {
				return nil, ErrorTokenNotValidYet
			} else {
				return nil, ErrorTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrorTokenInvalid
}

// RefreshToken 更新token,认加 x 分钟
func (j *JWT) RefreshToken(tokenString string, second int) (string, error) {
	info, err2 := j.ParseToken(tokenString)
	if err2 != nil {
		return "", err2
	}
	info.ExpiresAt = time.Now().Add(time.Second * time.Duration(second)).Unix()
	return j.CreateToken(info)
}
