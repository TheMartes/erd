package replication

import (
	"github.com/bxcodec/faker/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FakeData []string

func GenerateFakeData(cycles int) FakeData {
	var fakeTitles FakeData

	for i := 0; i < cycles; i++ {
		fakeTitles = append(fakeTitles, faker.Word())
	}

	return fakeTitles
}

func ConvertToMongoCompatible(arr []string) []mongo.WriteModel {
	models := []mongo.WriteModel{}

	for i, title := range arr {
		update := bson.D{
			bson.E{Key: "$set", Value: bson.D{
				bson.E{Key: "title", Value: title},
			}},
		}

		wm := mongo.NewUpdateOneModel().
			SetFilter(bson.D{
				primitive.E{
					Key:   "_id",
					Value: i,
				},
			},
			).SetUpdate(update).SetUpsert(true)

		models = append(models, wm)
	}

	return models
}
