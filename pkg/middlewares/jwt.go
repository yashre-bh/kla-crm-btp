package middlewares

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func GetTOMLCongfig() types.Config {
	var config types.Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}
	return config
}

func CreateJWTClaims(EmployeeID int32, Role types.Role) (string, error) {

	claims := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"employeeID": EmployeeID,
		"role":       Role,
	})

	token, err := claims.SignedString([]byte(GetTOMLCongfig().JWT.Secret))
	return token, err
}

func ExtractJWTClaims(c *gin.Context) (jwt.MapClaims, error) {
	var claims jwt.MapClaims
	cookie, err := c.Cookie("auth")
	if err != nil {
		return claims, err
	}

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetTOMLCongfig().JWT.Secret), nil
	})

	if err != nil {
		return claims, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return claims, errors.New("error retrieving token")
	}

	return claims, err
}
