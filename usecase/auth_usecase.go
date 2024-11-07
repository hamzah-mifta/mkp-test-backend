package usecase

import (
	"context"
	"database/sql"
	"time"

	"github.com/hamzah-mifta/mkp-test-backend/domain"
	"github.com/hamzah-mifta/mkp-test-backend/transport/request"
	"github.com/hamzah-mifta/mkp-test-backend/utils"
	"github.com/hamzah-mifta/mkp-test-backend/utils/crypto"
	"github.com/hamzah-mifta/mkp-test-backend/utils/jwt"
)

type authUsecase struct {
	userRepo  domain.UserRepository
	cryptoSvc crypto.CryptoService
	jwtSvc    jwt.JWTService
}

func NewAuthUsecase(userRepo domain.UserRepository, cryptoSvc crypto.CryptoService, jwtSvc jwt.JWTService) *authUsecase {
	return &authUsecase{
		userRepo:  userRepo,
		cryptoSvc: cryptoSvc,
		jwtSvc:    jwtSvc,
	}
}

func (u *authUsecase) SignUp(ctx context.Context, request *request.SignUpReq) (err error) {
	user, err := u.userRepo.GetByEmail(ctx, request.Email)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	if user.ID != 0 {
		return utils.NewBadRequestError("email already registered")
	}

	passwordHash, err := u.cryptoSvc.CreatePasswordHash(ctx, request.Password)
	if err != nil {
		return
	}

	err = u.userRepo.Create(ctx, &domain.User{
		Email:     request.Email,
		Password:  passwordHash,
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return
}
