package main

import (
	"fmt"
	"os"

	"github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/database/mongodb"
	infra "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/http"
	repository "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/repository/mongo"
	health "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/health"
	investiment "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/investiment"
	profit "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/profit"
	// keys "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/keys"
)

func main() {

	fmt.Println("** Starting app **")
	fmt.Println(os.Getenv("MONGO_URL"))

	// keys.SetSecret("rendafixa/mongo", "MONGO_URL")

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

	fmt.Println("** Using port: ", port)

	err := http.ListenAndServe(":" + port)
	// err := http.ListenAndServe(":80")

	if err != nil {
		panic(err)
	}

}
