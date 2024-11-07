package http

import (
	"net/http"

	"github.com/hamzah-mifta/mkp-test-backend/domain"
	"github.com/hamzah-mifta/mkp-test-backend/transport/request"
	"github.com/hamzah-mifta/mkp-test-backend/utils"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

func NewAuthHandler(e *echo.Echo, authUsecase domain.AuthUsecase) {
	handler := &AuthHandler{
		AuthUsecase: authUsecase,
	}

	apiV1 := e.Group("/api/v1")
	apiV1.POST("/auth/signup", handler.SignUp)
}

func (h *AuthHandler) SignUp(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.SignUpReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewUnprocessableEntityError(err.Error()))
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestError(err.Error()))
	}

	if err := h.AuthUsecase.SignUp(ctx, &req); err != nil {
		return c.JSON(utils.ParseHttpError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "sigend up succesfully",
	})
}
