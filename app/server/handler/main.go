package handler

import (
	"github.com/amiame/mock-go-backend/app/server/usecase"
)

type Handler interface {
	MathHandler
}

type handler struct {
	MathHandler
}

func New(u usecase.Usecase) Handler {
	return &handler{
		MathHandler: newMathHandler(u),
	}
}
