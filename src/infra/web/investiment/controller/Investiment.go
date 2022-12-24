package controller

import (
	"encoding/json"
	"fmt"

	application "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/application"
	domain "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/domain"
	http "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/http"
)

type InvestimentCtrl struct {
	repository domain.InvestimentRepository
}

func NewInvestimentCtrl(repository domain.InvestimentRepository) *InvestimentCtrl {
	InvestimentCtrl := new(InvestimentCtrl)
	InvestimentCtrl.repository = repository

	return InvestimentCtrl
}

type InvestimentInput struct {
	Ativo *domain.Ativo `json:"ativo"`
}

func (gs *InvestimentCtrl) Create(params map[string]string, body []byte, queryArgs http.QueryParams) (interface{}, error) {
	var inputJSON InvestimentInput
	err := json.Unmarshal(body, &inputJSON)
	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	output, err := application.
		NewCreateInvestment(gs.repository).
		Execute(inputJSON.Ativo)

	fmt.Println(err)

	return output, err
}
