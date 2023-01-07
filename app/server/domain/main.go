package domain

type Domain interface {
	MathDomain
}

type domain struct {
	MathDomain
}

func New() Domain {
	return &domain{
		MathDomain: newMathDomain(),
	}
}
