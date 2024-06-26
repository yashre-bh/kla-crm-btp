package middlewares

import (
	"errors"
	"fmt"
	"strings"

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

func CreateJWTClaims(EmployeeID int32, Role types.Role, Checkpoint []interface{}) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"employeeID": EmployeeID,
		"role":       Role,
		"checkpoint": Checkpoint,
	})

	token, err := claims.SignedString([]byte(GetTOMLCongfig().JWT.Secret))
	return token, err
}

func ExtractJWTClaims(c *gin.Context) (jwt.MapClaims, error) {
	var claims jwt.MapClaims
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return claims, errors.New("missing authorization header")
	}

	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
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
