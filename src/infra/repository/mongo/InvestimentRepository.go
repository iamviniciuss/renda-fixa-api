package repository

import (
	"context"
	"os"
	"time"

	"github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/domain"
	"github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/database"
	"github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/infra/database/mongodb"
	"github.com/Vinicius-Santos-da-Silva/renda-fixa-api/src/util/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

type InvestimentRepositoryMongo[T mongodb.MongoInteface] struct {
	connection database.Connection[T]
}

func NewInvestimentRepositoryMongo(connection database.Connection[mongodb.MongoInteface]) *InvestimentRepositoryMongo[mongodb.MongoInteface] {
	return &InvestimentRepositoryMongo[mongodb.MongoInteface]{
		connection: connection,
	}
}

func (erm *InvestimentRepositoryMongo[T]) Create(ativo *domain.Ativo) (*domain.Ativo, error) {
	var id primitive.ObjectID

	coll := erm.connection.
		Client().
		Mongo().
		Database(os.Getenv("DATABASE")).
		Collection("investiment")

	if ativo.ID == "" {
		id = primitive.NewObjectID()
	} else {
		id = mongo.GetObjectIDFromString(ativo.ID)
	}

	data := bson.M{
		"_id":                           id,
		"code":                          ativo.Code,
		"nickName":                      ativo.Name,
		"maturityDate":                  ativo.MaturityDate,
		"fee":                           ativo.FeeString,
		"minimumQuantityForApplication": ativo.MinimumQuantityForApplication,
		"puMinValue":                    ativo.PuMinValue,
		"product":                       ativo.Product,
		"qualifiedInvestor":             ativo.QualifiedInvestor,
		"guaranteeFGC":                  ativo.GuaranteeFGC,
		"graceDate":                     ativo.GraceDate,
		"riskScore":                     ativo.RiskScore,
		"indexers":                      ativo.Index,
		"created_at":                    time.Now().UTC(),
		"professionalInvestor":          ativo.ProfessionalInvestor,
		"generalInvestor":               ativo.GeneralInvestor,
		"incentive":                     ativo.Incentive,
		"ratingName":                    ativo.RatingName,
		"agencyName":                    ativo.AgencyName,
		"preRegistration":               ativo.PreRegistration,
		"redemptionType":                ativo.RedemptionType,
		"continuedOffering":             ativo.ContinuedOffering,
		"quantityAvailable":             ativo.QuantityAvailable,
		"descriptionAmortization":       ativo.DescriptionAmortization,
		"descriptionInterestrates":      ativo.DescriptionInterestrates,
		"originCode":                    ativo.OriginCode,
		"isCampaign":                    ativo.IsCampaign,
		"suitability":                   ativo.Suitability,
		"targetInvestor":                ativo.TargetInvestor,
	}

	res, err1 := coll.InsertOne(context.TODO(), data)

	if err1 != nil {
		return nil, err1
	}

	ativo.ID = res.InsertedID.(primitive.ObjectID).Hex()

	return ativo, nil
}

func (erm *InvestimentRepositoryMongo[T]) FindByCode(code string) (*domain.Ativo, error) {
	ativo := &domain.Ativo{}

	err := erm.connection.
		Client().
		Mongo().
		Database(os.Getenv("DATABASE")).
		Collection("investiment").
		FindOne(context.TODO(), bson.M{"code": code}).
		Decode(ativo)

	return ativo, err
}

func (erm *InvestimentRepositoryMongo[T]) Count() int64 {
	total, err := erm.connection.
		Client().
		Mongo().
		Database(os.Getenv("DATABASE")).
		Collection("investiment").
		CountDocuments(context.TODO(), bson.M{})

	if err != nil {
		return 0
	}

	return total
}
