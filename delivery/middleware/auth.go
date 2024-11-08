package middleware

import (
	"net/http"
	"strings"

	jwtLib "github.com/golang-jwt/jwt"
	"github.com/hamzah-mifta/mkp-test-backend/utils"
	"github.com/labstack/echo/v4"
)

func (m *Middleware) JWTAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			authorizationHeader := c.Request().Header.Get("Authorization")
			bearerToken := strings.Split(authorizationHeader, " ")

			if len(bearerToken) != 2 {
				return c.JSON(http.StatusUnauthorized, utils.NewUnauthorizedError("invalid token"))
			}

			tokenString := bearerToken[1]
			token, err := m.jwtSvc.ValidateToken(ctx, tokenString)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, utils.NewUnauthorizedError("invalid token"))
			}

			if !token.Valid {
				return c.JSON(http.StatusUnauthorized, utils.NewUnauthorizedError("invalid token"))
			}

			claims := token.Claims.(jwtLib.MapClaims)
			c.Set("user_id", int64(claims["user_id"].(float64)))

			return next(c)
		}
	}
}
