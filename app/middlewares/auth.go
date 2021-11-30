package middlewares

import (
	"currency-exchange/controllers"
	"currency-exchange/helpers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	UserID  int `json:"id"`
	IsAdmin int `json:"is_admin"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		SuccessHandler: func(context echo.Context) {
		},
		ErrorHandlerWithContext: func(e error, c echo.Context) error {
			return controllers.NewErrorResponse(c, http.StatusUnauthorized, helpers.ErrInvalidTokenCredential)
		},
	}
}

func (configJwt *ConfigJWT) GenererateToken(userID int, isAdmin int) (string, error) {
	claims := JwtCustomClaims{
		userID,
		isAdmin,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(configJwt.ExpiresDuration))).Unix(),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, _ := t.SignedString([]byte(configJwt.SecretJWT))

	return token, nil
}

// func GetUser(c echo.Context) (res *JwtCustomClaims) {
// 	user := c.Get("user")
// 	if user != nil {
// 		res = user.(*JwtCustomClaims)
// 	}
// 	return res
// }

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtCustomClaims)
		isAdmin := claims.IsAdmin
		if isAdmin != 0 {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func IsUserId(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		userId := c.Param("userId")
		convUserId, err := helpers.StringToInt(userId)
		if err != nil {
			return echo.ErrBadRequest
		}
		claims := user.Claims.(*JwtCustomClaims)
		claimsUserId := claims.UserID
		if claimsUserId != convUserId {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
