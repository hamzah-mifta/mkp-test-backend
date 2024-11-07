package http

import (
	"net/http"
	"strconv"

	"github.com/hamzah-mifta/mkp-test-backend/delivery/middleware"
	"github.com/hamzah-mifta/mkp-test-backend/domain"
	"github.com/hamzah-mifta/mkp-test-backend/transport/request"
	"github.com/hamzah-mifta/mkp-test-backend/utils"
	"github.com/labstack/echo/v4"
)

type ScheduleHandler struct {
	ScheduleUsecase domain.ScheduleUsecase
}

func NewScheduleHandler(e *echo.Echo, middleware *middleware.Middleware, scheduleUsecase domain.ScheduleUsecase) {
	handler := &ScheduleHandler{
		ScheduleUsecase: scheduleUsecase,
	}

	apiV1 := e.Group("/api/v1")

	apiV1.POST("/schedules", handler.Create, middleware.JWTAuth())
	apiV1.GET("/schedules", handler.Fetch, middleware.JWTAuth())
	apiV1.GET("/schedules/:id", handler.GetByID, middleware.JWTAuth())
	apiV1.PUT("/schedules/:id", handler.Update, middleware.JWTAuth())
}

func (h *ScheduleHandler) Create(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.CreateScheduleReq

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewUnprocessableEntityError(err.Error()))
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestError(err.Error()))
	}

	if err := h.ScheduleUsecase.Create(ctx, &req); err != nil {
		return c.JSON(utils.ParseHttpError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "schedule created successfully"})
}

func (h *ScheduleHandler) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	schedules, err := h.ScheduleUsecase.Fetch(ctx)
	if err != nil {
		return c.JSON(utils.ParseHttpError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success fetch schedules",
		"data":    schedules,
	})
}

func (h *ScheduleHandler) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNotFoundError("schedule not found"))
	}

	schedule, err := h.ScheduleUsecase.GetByID(ctx, int64(id))
	if err != nil {
		return c.JSON(utils.ParseHttpError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get schedule",
		"data":    schedule,
	})
}

func (h *ScheduleHandler) Update(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.UpdateScheduleReq

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.NewNotFoundError("schedule not found"))
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewUnprocessableEntityError(err.Error()))
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewBadRequestError(err.Error()))
	}

	if err := h.ScheduleUsecase.Update(ctx, int64(id), &req); err != nil {
		return c.JSON(utils.ParseHttpError(err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "schedule updated successfully"})
}
