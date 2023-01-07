package handler

import (
	"net/http"

	"github.com/amiame/mock-go-backend/app/server/handler/request"
	"github.com/amiame/mock-go-backend/app/server/usecase"
	"github.com/labstack/echo/v4"
)

type MathHandler interface {
	TwoSum(c echo.Context) error
}

type mathHandler struct {
	usecase usecase.Usecase
}

func newMathHandler(u usecase.Usecase) MathHandler {
	return &mathHandler{
		usecase: u,
	}
}

func (m *mathHandler) TwoSum(c echo.Context) error {
	parameters := request.TwoSumArguments{}
	if err := c.Bind(&parameters); err != nil {
		return err
	}

	resp := m.usecase.TwoSum(parameters)

	return c.JSON(http.StatusOK, resp)
}
