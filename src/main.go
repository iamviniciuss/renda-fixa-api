package main

import (
	infra "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/http"
	profit "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/profit"
)

func main() {

	http := infra.NewFiberHttp()
	profit.ProfitRouter(http)

	err := http.ListenAndServe(":9002")

	if err != nil {
		panic(err)
	}
}
