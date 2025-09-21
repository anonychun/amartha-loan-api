package loan

import (
	"github.com/anonychun/amartha-loan-api/internal/api"
	"github.com/labstack/echo/v4"
)

func (h *Handler) FindAll(c echo.Context) error {
	res, err := h.usecase.FindAll(c.Request().Context())
	if err != nil {
		return err
	}

	return api.NewResponse(c).SetData(res).Send()
}

func (h *Handler) Create(c echo.Context) error {
	var req CreateRequest
	err := c.Bind(&req)
	if err != nil {
		return err
	}

	res, err := h.usecase.Create(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return api.NewResponse(c).SetData(res).Send()
}
