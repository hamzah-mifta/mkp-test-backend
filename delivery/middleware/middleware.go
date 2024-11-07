package middleware

import "github.com/hamzah-mifta/mkp-test-backend/utils/jwt"

type Middleware struct {
	jwtSvc jwt.JWTService
}

func NewMiddleware(jwtSvc jwt.JWTService) *Middleware {
	return &Middleware{
		jwtSvc: jwtSvc,
	}
}
