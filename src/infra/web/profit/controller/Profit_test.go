package controller

import (
	"encoding/json"
	"testing"
	"time"

	application "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/application"

	"github.com/stretchr/testify/assert"
)

func TestGetScore(t *testing.T) {
	t.Run("Return NPS Score equal to 33.33 with ES Repository", func(t *testing.T) {
		now := time.Date(2023, 03, 11, 10, 0, 0, 0, time.UTC)

		body := []byte(`{
			"capital": 1000,
			"ativo": {
				"nickName" : "CDB PAGBANK - MAI/2023",
				"maturityDate" : "2023-05-11T00:00:00",
				"fee" : 7,
				"minimumQuantityForApplication" : 1,
				"puMinValue" : 998.18747,
				"product" : "CDB",
				"indexers" : "Inflação",
				"qualifiedInvestor" : "N",
				"guaranteeFGC" : true,
				"code" : 10850585,
				"graceDate" : "2023-06-11T10:00:00Z",
				"originCode" : 1,
				"riskScore" : 3
			}
		}`)

		var params = map[string]string{}

		profitCtrl := NewProfitCtrl(now)
		response, err := profitCtrl.Profit(params, body, nil)
		assert.Equal(t, nil, err)

		byteData, _ := json.Marshal(response)
		outputNPSScore := application.EstimateInvestmentOutput{}
		json.Unmarshal(byteData, &outputNPSScore)

		assert.Equal(t, 3.07, outputNPSScore.Months)
		assert.Equal(t, 0.81, outputNPSScore.PercentageProfit)
		assert.Equal(t, 24.79, outputNPSScore.Profit)
	})
}
