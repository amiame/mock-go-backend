package usecase

import (
	"github.com/amiame/mock-go-backend/app/server/domain"
)

type Usecase interface {
	MathUsecase
}

type usecase struct {
	MathUsecase
}

func New(d domain.Domain) Usecase {
	return &usecase{
		MathUsecase: newMathUsecase(d),
	}
}
