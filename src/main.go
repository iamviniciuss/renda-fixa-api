package main

import (
	"os"

	infra "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/http"
	health "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/health"
	profit "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/profit"
)

func main() {

	http := infra.NewFiberHttp()
	profit.ProfitRouter(http)
	health.HealthRouter(http)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":" + port)

	if err != nil {
		panic(err)
	}

}
