package infra

import (
	"github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/domain"
	infra "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/http"
	investiment "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/investiment/controller"
)

func InvestimentRouter(http infra.HttpService, repository domain.InvestimentRepository) {
	http.Post("/investiment", investiment.NewInvestimentCtrl(repository).Create)
}
