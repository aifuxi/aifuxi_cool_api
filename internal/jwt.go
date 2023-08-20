package internal

import (
	"time"

	"api.aifuxi.cool/myerror"
	"api.aifuxi.cool/settings"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	Email string `json:"email"`

	jwt.RegisteredClaims
}

func GenToken(email string) (string, error) {
	temp := time.Now().Add(time.Hour * time.Duration(settings.AppConfig.JwtExpiredHour))

	claims := MyClaims{
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(temp),
			Issuer:    "aifuxi",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(settings.AppConfig.JwtSecret))
}

// TODO: 需要处理 token 过期的情况，单独抛出过期错误
func ParseToken(tokenStr string) (*MyClaims, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(settings.AppConfig.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, myerror.ErrorParseTokenFailed
}
