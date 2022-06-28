package controller

import (
	"encoding/json"
	"time"

	profit "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/application"
	domain "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/domain"
	http "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/http"
)

type ProfitCtrl struct {
	Now time.Time
}

func NewProfitCtrl(now time.Time) *ProfitCtrl {
	ProfitCtrl := new(ProfitCtrl)
	ProfitCtrl.Now = now

	return ProfitCtrl
}

type ProfitInput struct {
	Capital float64      `json:"capital,omitempty"`
	Ativo   domain.Ativo `json:"ativo,omitempty"`
}

func (gs *ProfitCtrl) Profit(params map[string]string, body []byte, queryArgs http.QueryParams) (interface{}, error) {
	var inputJSON ProfitInput
	err := json.Unmarshal(body, &inputJSON)

	useCase := profit.NewEstimateInvestment(gs.Now)

	output, err := useCase.Execute(&inputJSON.Ativo, inputJSON.Capital)

	return output, err
}
