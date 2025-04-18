package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/iandanarko/concert/config"
	"github.com/iandanarko/concert/internal/entity"
	cErr "github.com/iandanarko/concert/internal/error"
	"github.com/iandanarko/concert/internal/serializer"
	"github.com/labstack/echo/v4"
)

const (
	Actor         = "actor"
	authBearerKey = "Bearer"
)

func Auth(cfg config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCtx echo.Context) error {

			requestHeader := echoCtx.Request().Header
			authHeader := requestHeader.Get(echo.HeaderAuthorization)
			token := strings.Split(authHeader, " ")
			if len(token) != 2 || token[0] != authBearerKey {
				echoCtx.JSON(http.StatusUnauthorized, serializer.NewErrorResponse(cErr.ErrUnauthorized))
				return cErr.ErrUnauthorized
			}

			claims := jwt.MapClaims{}
			_, err := jwt.ParseWithClaims(token[1], claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(cfg.Jwt.Secret), nil
			})
			if err != nil {
				echoCtx.JSON(http.StatusUnauthorized, serializer.NewErrorResponse(cErr.ErrUnauthorized))
				return cErr.ErrUnauthorized
			}

			userID := uint64(claims["user_id"].(float64))
			echoCtx.Set(Actor, &entity.Actor{UserID: userID})
			return next(echoCtx)
		}
	}
}
