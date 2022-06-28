package domain

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AtivoTestSuite struct {
	suite.Suite
}

func (suite *AtivoTestSuite) TestCalcReturn() {
	ativo := &Ativo{
		Fee: 7,
	}

	returned := ativo.CalculateReturn(13.1, 3.07, 1000)

	suite.Equal(31.99, returned)
}

func (suite *AtivoTestSuite) TestCalcProfit() {
	ativo := &Ativo{
		Fee: 7,
	}

	profit := ativo.CalculateProfit(3.07, 1000)

	suite.Equal(24.79, profit)
}

func (suite *AtivoTestSuite) TestCalcRealProfitability() {
	ativo := &Ativo{
		Fee: 7,
	}

	profit := ativo.GetRealProfitability(3.07, 1000)

	suite.Equal(0.81, profit)
}

func (suite *AtivoTestSuite) TestGetAliquota1() {
	ativo := &Ativo{}

	aliquota1 := ativo.GetAliquota(1)
	aliquota2 := ativo.GetAliquota(2)
	aliquota3 := ativo.GetAliquota(3)
	aliquota4 := ativo.GetAliquota(4)
	aliquota5 := ativo.GetAliquota(5)

	suite.Equal(22.5, aliquota1)
	suite.Equal(22.5, aliquota2)
	suite.Equal(22.5, aliquota3)
	suite.Equal(22.5, aliquota4)
	suite.Equal(22.5, aliquota5)
}

func (suite *AtivoTestSuite) TestGetAliquota2() {
	ativo := &Ativo{}

	aliquota6 := ativo.GetAliquota(6)
	aliquota7 := ativo.GetAliquota(7)
	aliquota8 := ativo.GetAliquota(8)
	aliquota9 := ativo.GetAliquota(9)
	aliquota10 := ativo.GetAliquota(10)
	aliquota11 := ativo.GetAliquota(11)
	aliquota12 := ativo.GetAliquota(12)

	suite.Equal(20.0, aliquota6)
	suite.Equal(20.0, aliquota7)
	suite.Equal(20.0, aliquota8)
	suite.Equal(20.0, aliquota9)
	suite.Equal(20.0, aliquota10)
	suite.Equal(20.0, aliquota11)
	suite.Equal(20.0, aliquota12)
}

func (suite *AtivoTestSuite) TestGetAliquota3() {
	ativo := &Ativo{}

	aliquota13 := ativo.GetAliquota(13)
	aliquota18 := ativo.GetAliquota(18)
	aliquota24 := ativo.GetAliquota(24)

	suite.Equal(17.5, aliquota13)
	suite.Equal(17.5, aliquota18)
	suite.Equal(17.5, aliquota24)
}

func (suite *AtivoTestSuite) TestGetAliquota4() {
	ativo := &Ativo{}

	aliquota13 := ativo.GetAliquota(25)
	aliquota24 := ativo.GetAliquota(30)

	suite.Equal(15.0, aliquota13)
	suite.Equal(15.0, aliquota24)
}

func TestAtivoTestSuite(t *testing.T) {
	suite.Run(t, &AtivoTestSuite{})
}
