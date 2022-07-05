package application

import (
	"testing"
	"time"

	"github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/domain"
	"github.com/stretchr/testify/suite"
)

type EstimeInvestimentTestSuite struct {
	suite.Suite
}

func (suite *EstimeInvestimentTestSuite) TestCalcReturn() {
	now := time.Date(2023, 03, 11, 10, 0, 0, 0, time.UTC)

	ativo := &domain.Ativo{
		Fee:       7,
		GraceDate: "2023-06-11T10:00:00Z",
		Index:     "Inflação",
		Product:   "CDB",
	}

	useCase := NewEstimateInvestment(now)
	output, err := useCase.Execute(ativo, 1000)

	suite.Equal(nil, err)
	suite.Equal(3.07, output.Months)
	suite.Equal(0.81, output.PercentageProfit)
	suite.Equal(24.79, output.Profit)
}

func (suite *EstimeInvestimentTestSuite) TestWithInvalidDate() {
	now := time.Date(2023, 03, 11, 10, 0, 0, 0, time.UTC)

	ativo := &domain.Ativo{
		Fee:       7,
		GraceDate: "invalid",
		Index:     "Inflação",
		Product:   "CDB",
	}

	useCase := NewEstimateInvestment(now)
	output, err := useCase.Execute(ativo, 1000)

	suite.NotNil(err)
	suite.Nil(output)
}

func (suite *EstimeInvestimentTestSuite) TestWithInvalidProduct() {
	now := time.Date(2023, 03, 11, 10, 0, 0, 0, time.UTC)

	ativo := &domain.Ativo{
		Fee:       7,
		GraceDate: "2023-06-11T10:00:00Z",
		Index:     "Inflação",
		Product:   "",
	}

	useCase := NewEstimateInvestment(now)
	output, err := useCase.Execute(ativo, 1000)

	suite.Nil(err)
	suite.NotNil(output)
	suite.Equal(0.0, output.Investiment)
	suite.Equal(0.0, output.MonthProfit)
	suite.Equal(0.0, output.Months)
	suite.Equal(0.0, output.PercentageProfit)
}

func (suite *EstimeInvestimentTestSuite) TestWithInvalidIndex() {
	now := time.Date(2023, 03, 11, 10, 0, 0, 0, time.UTC)

	ativo := &domain.Ativo{
		Fee:       7,
		GraceDate: "2023-06-11T10:00:00Z",
		Index:     "",
		Product:   "CDB",
	}

	useCase := NewEstimateInvestment(now)
	output, err := useCase.Execute(ativo, 1000)

	suite.Nil(err)
	suite.NotNil(output)
	suite.Equal(0.0, output.Investiment)
	suite.Equal(0.0, output.MonthProfit)
	suite.Equal(0.0, output.Months)
	suite.Equal(0.0, output.PercentageProfit)
}

func TestEstimeInvestimentTestSuite(t *testing.T) {
	suite.Run(t, &EstimeInvestimentTestSuite{})
}
