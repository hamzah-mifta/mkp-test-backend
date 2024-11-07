package domain

import (
	"context"

	"github.com/hamzah-mifta/mkp-test-backend/transport/request"
)

type AuthUsecase interface {
	SignUp(ctx context.Context, req *request.SignUpReq) error
	SignIn(ctx context.Context, req *request.SignInReq) (string, error)
}
