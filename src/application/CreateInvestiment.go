package application

import "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/domain"

type CreateInvestment struct {
	repository domain.InvestimentRepository
}

func NewCreateInvestment(repository domain.InvestimentRepository) *CreateInvestment {
	return &CreateInvestment{
		repository: repository,
	}
}

func (ci *CreateInvestment) Execute(ativo *domain.Ativo) (*domain.Ativo, error) {
	return ci.repository.Create(ativo)
}
