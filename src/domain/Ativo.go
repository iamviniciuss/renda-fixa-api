package domain

import (
	"math"

	helper "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/domain/helper"
)

const IPCA = 6.1

type Ativo struct {
	ID                            string  `json:"_id" bson:"_id"`
	Code                          string  `json:"code,omitempty"`
	Name                          string  `json:"nickName,omitempty"`
	MaturityDate                  string  `json:"maturityDate,omitempty"`
	Fee                           float64 `json:"fee,omitempty"`
	MinimumQuantityForApplication string  `json:"minimumQuantityForApplication,omitempty"`
	PuMinValue                    string  `json:"puMinValue,omitempty"`
	Product                       string  `json:"product,omitempty"`
	QualifiedInvestor             string  `json:"qualifiedInvestor,omitempty"`
	GuaranteeFGC                  bool    `json:"guaranteeFGC,omitempty"`
	GraceDate                     string  `json:"graceDate,omitempty"`
	RiskScore                     string  `json:"riskScore,omitempty"`
	Index                         string  `json:"indexers,omitempty"`
}

func (atv *Ativo) CalculateReturn(taxaJurosTotal float64, mesesDeAplicacao float64, capital float64) float64 {

	x := float64((1.0 / 12.0))
	y := float64(mesesDeAplicacao)

	tempoAplicacao := x * y

	taxaJuros := taxaJurosTotal / 100.0

	montante := capital * float64(math.Pow(float64(1.0+taxaJuros), float64(tempoAplicacao)))

	lucro := montante - capital
	return helper.Round(lucro)
}

func (atv *Ativo) GetAliquota(mesesDeAplicacao float64) float64 {
	aliquota := 22.5

	if mesesDeAplicacao >= 6 && mesesDeAplicacao <= 12 {
		aliquota = 20.0
	} else if mesesDeAplicacao > 12 && mesesDeAplicacao <= 24 {
		aliquota = 17.5
	} else if mesesDeAplicacao > 24 {
		aliquota = 15.0
	}

	return aliquota
}

func (atv *Ativo) CalculateProfit(mesesDeAplicacao float64, capital float64) float64 {

	taxaDeJurosR := IPCA + atv.Fee

	rendimento_liquido := atv.CalculateReturn(taxaDeJurosR, mesesDeAplicacao, capital)

	aliquota := atv.GetAliquota(mesesDeAplicacao)

	final := rendimento_liquido - ((aliquota / 100) * rendimento_liquido)

	return helper.Round(final)
}
func (atv *Ativo) GetRealProfitability(mesesDeAplicacao float64, capital float64) float64 {
	profit := atv.CalculateProfit(mesesDeAplicacao, capital)

	r := ((profit * 100.0) / capital) / mesesDeAplicacao

	return helper.Round(r)
}
