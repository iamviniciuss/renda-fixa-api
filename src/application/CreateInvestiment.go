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
	exists, err := ci.repository.FindByCode(ativo.Code)

	if exists.ID != "" {
		return exists, err
	}

	return ci.repository.Create(ativo)
}
