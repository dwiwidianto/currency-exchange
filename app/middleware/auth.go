package middleware

import (
	"currency-exchange/controllers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	UserID int `json:"id"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return controllers.NewErrorResponse(c, http.StatusForbidden, e)
		}),
	}
}

func (configJwt ConfigJWT) GenererateToken(userID int) string {
	claims := JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(configJwt.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, _ := t.SignedString([]byte(configJwt.SecretJWT))
	return token
}

func GetUser(c echo.Context) (res *JwtCustomClaims) {
	user := c.Get("user")
	if user != nil {
		res = user.(*JwtCustomClaims)
	}
	return res
}
