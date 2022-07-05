package infra

import (
	"time"

	infra "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/http"
	"github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/web/health/controller"
)

func HealthRouter(http infra.HttpService) {
	http.Get("/", controller.NewHealthCtrl(time.Now()).Index)
}
