package controller

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/domain"
	"github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/database/mongodb"
	repository "github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/repository/mongo"
	"github.com/tryvium-travels/memongo"

	"github.com/stretchr/testify/suite"
)

type CreateInvestimentTestSuite struct {
	MongoConnection *memongo.Server
	suite.Suite
}

func TestCreateInvestimentTestSuite(t *testing.T) {
	suite.Run(t, &CreateInvestimentTestSuite{
		MongoConnection: mongodb.CreateMongoTemp(),
	})
}

func (s *CreateInvestimentTestSuite) AfterTest(suiteName, testName string) {
	fmt.Println("After test")
	mongodb.CreateMongoTemp()
}

func (suite *CreateInvestimentTestSuite) TestCreateAInvestiment() {
	mongo := mongodb.NewMongoConnection()
	investimentRepository := repository.NewInvestimentRepositoryMongo(mongo)

	body := []byte(`{
		"ativo": {
			"nickName" : "CDB PAGBANK - MAI/2023",
			"maturityDate" : "2023-05-11T00:00:00",
			"fee" : 7,
			"minimumQuantityForApplication" : "1",
			"puMinValue" : "998.18747",
			"product" : "CDB",
			"indexers" : "Inflação",
			"qualifiedInvestor" : "N",
			"guaranteeFGC" : true,
			"code" : "10850585",
			"graceDate" : "2023-06-11T10:00:00Z",
			"originCode" : "1",
			"riskScore" : "3"
		}
	}`)

	var params = map[string]string{}

	profitCtrl := NewInvestimentCtrl(investimentRepository)
	response, err := profitCtrl.Create(params, body, nil)
	suite.Equal(nil, err)

	byteData, err := json.Marshal(response)
	suite.Nil(err)
	ativo := domain.Ativo{}
	json.Unmarshal(byteData, &ativo)

	suite.Equal("CDB PAGBANK - MAI/2023", ativo.Name)
	suite.Equal(int64(1), investimentRepository.Count())

}

func (suite *CreateInvestimentTestSuite) TestShouldNotDuplicateInvestiment() {
	mongo := mongodb.NewMongoConnection()
	investimentRepository := repository.NewInvestimentRepositoryMongo(mongo)

	body := []byte(`{
		"ativo": {
			"nickName" : "CDB PAGBANK - MAI/2023",
			"maturityDate" : "2023-05-11T00:00:00",
			"fee" : 7,
			"minimumQuantityForApplication" : "1",
			"puMinValue" : "998.18747",
			"product" : "CDB",
			"indexers" : "Inflação",
			"qualifiedInvestor" : "N",
			"guaranteeFGC" : true,
			"code" : "10850585",
			"graceDate" : "2023-06-11T10:00:00Z",
			"originCode" : "1",
			"riskScore" : "3"
		}
	}`)

	var params = map[string]string{}

	profitCtrl := NewInvestimentCtrl(investimentRepository)
	response, err := profitCtrl.Create(params, body, nil)
	suite.Equal(nil, err)

	byteData, err := json.Marshal(response)
	suite.Nil(err)
	ativo := domain.Ativo{}
	json.Unmarshal(byteData, &ativo)

	suite.Equal("CDB PAGBANK - MAI/2023", ativo.Name)
	suite.Equal(int64(1), investimentRepository.Count())

	_, err1 := profitCtrl.Create(params, body, nil)
	suite.Equal(nil, err1)
	suite.Equal(int64(1), investimentRepository.Count())
}
