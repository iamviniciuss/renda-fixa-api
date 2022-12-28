package controller

import (
	"fmt"
	"time"

	http "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/http"
)

type HealthCtrl struct {
	Now time.Time
}

func NewHealthCtrl(now time.Time) *HealthCtrl {
	HealthCtrl := new(HealthCtrl)
	HealthCtrl.Now = now

	return HealthCtrl
}

func (gs *HealthCtrl) Index(params map[string]string, body []byte, queryArgs http.QueryParams) (interface{}, error) {
	fmt.Println("** Health Check **")
	return "OK 1.5", nil
}
