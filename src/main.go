package main

import (
	"os"

	"github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/database/mongodb"
	infra "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/http"
	repository "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/repository/mongo"
	health "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/health"
	investiment "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/investiment"
	profit "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/profit"
)

func main() {

	http := infra.NewFiberHttp()
	profit.ProfitRouter(http)
	health.HealthRouter(http)
	mongo := mongodb.NewMongoConnection()

	investimentRepository := repository.NewInvestimentRepositoryMongo(mongo)
	investiment.InvestimentRouter(http, investimentRepository)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":" + port)

	if err != nil {
		panic(err)
	}

}
