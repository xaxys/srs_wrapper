package util

import (
	"srs_wrapper/config"
	"time"

	"github.com/iris-contrib/middleware/jwt"
)

var jwtkey = []byte(config.AppConfig.GetString("app.jwtkey"))

func GetJwtString(id uint) (string, error) {
	token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": id,

		"iss": config.AppConfig.GetString("app.name"),
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(24 * time.Hour * time.Duration(1)).Unix(),
	})

	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
