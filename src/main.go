package main

import (
	infra "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/http"
	health "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/health"
	profit "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/profit"
)

func main() {

	http := infra.NewFiberHttp()
	profit.ProfitRouter(http)
	health.HealthRouter(http)

	err := http.ListenAndServe(":80")

	if err != nil {
		panic(err)
	}
}
