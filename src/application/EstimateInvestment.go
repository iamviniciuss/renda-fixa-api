package application

import (
	"math"
	"time"

	"github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/domain"
	helper "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/domain/helper"
)

type EstimateInvestmentOutput struct {
	Months           float64
	Profit           float64
	PercentageProfit float64
	Investiment      float64
	MonthProfit      float64
}

type EstimateInvestment struct {
	PurchaseDate time.Time
}

func NewEstimateInvestment(purchaseDate time.Time) *EstimateInvestment {
	return &EstimateInvestment{
		PurchaseDate: purchaseDate,
	}
}

func (ei *EstimateInvestment) Execute(ativo *domain.Ativo, investiment float64) (*EstimateInvestmentOutput, error) {

	if ativo.Product != "CDB" {
		return &EstimateInvestmentOutput{}, nil
	}

	if ativo.Index != "Inflação" {
		return &EstimateInvestmentOutput{}, nil
	}

	GraceDate, err := time.Parse(time.RFC3339, ativo.GraceDate)

	if err != nil {
		return nil, err
	}

	difference := ei.PurchaseDate.Sub(GraceDate)
	months := math.Abs(helper.Round(difference.Hours() / 24 / 30))

	percentageProfit := ativo.GetRealProfitability(months, investiment)
	profit := ativo.CalculateProfit(months, investiment)

	return &EstimateInvestmentOutput{
		Months:           months,
		Profit:           profit,
		PercentageProfit: percentageProfit,
		Investiment:      investiment,
		MonthProfit:      profit / months,
	}, nil
}
