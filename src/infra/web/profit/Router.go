package infra

import (
	"time"

	infra "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/http"
	profit "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/profit/controller"
)

func ProfitRouter(http infra.HttpService) {
	http.Post("/profit", profit.NewProfitCtrl(time.Now()).Profit)
}
