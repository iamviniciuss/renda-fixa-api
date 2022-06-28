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
	}

	useCase := NewEstimateInvestment(now)
	output, err := useCase.Execute(ativo, 1000)

	suite.Equal(nil, err)
	suite.Equal(3.07, output.Months)
	suite.Equal(0.81, output.PercentageProfit)
	suite.Equal(24.79, output.Profit)
}

func TestEstimeInvestimentTestSuite(t *testing.T) {
	suite.Run(t, &EstimeInvestimentTestSuite{})
}
