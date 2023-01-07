package usecase

import (
	"github.com/amiame/mock-go-backend/app/server/domain"
	"github.com/amiame/mock-go-backend/app/server/handler/request"
	"github.com/amiame/mock-go-backend/app/server/handler/response"
)

type MathUsecase interface {
	TwoSum(args request.TwoSumArguments) response.TwoSumResponse
}

type mathUsecase struct {
	domain domain.Domain
}

func newMathUsecase(d domain.Domain) MathUsecase {
	return &mathUsecase{
		domain: d,
	}
}

func (m *mathUsecase) TwoSum(args request.TwoSumArguments) response.TwoSumResponse {
	return response.TwoSumResponse{
		TwoSum: m.domain.TwoSum(args.Numbers, args.Target),
	}
}
