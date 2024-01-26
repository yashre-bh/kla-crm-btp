package middlewares

import (
	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"
	"github.com/yashre-bh/kla-crm-btp/pkg/types"
)

func CreateJWTClaims(EmployeeID int32, Role types.Role) (string, error) {
	var config types.Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"employeeID": EmployeeID,
		"role":       Role,
	})

	token, err := claims.SignedString([]byte(config.JWT.Secret))
	return token, err
}
