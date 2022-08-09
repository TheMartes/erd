package environments

import (
	"context"
	"fmt"
	"log"

	"github.com/themartes/erd/env"
	"github.com/themartes/erd/persistance"
	mongoserviceprovider "github.com/themartes/erd/persistance/mongodb"
	"github.com/themartes/erd/replication"
)

// Local :)
type Local struct{}

// InitLocalEnv :))
func (l Local) InitLocalEnv() {
	mongoClient := persistance.GetMongoClient()
	arr := replication.GenerateFakeData(10000)

	dropErr := mongoClient.Database(env.Params.MongoDB).
		Collection(env.Params.MongoCollection).
		Drop(context.TODO())

	if dropErr != nil {
		log.Print(dropErr)
	}

	createErr := mongoClient.Database(env.Params.MongoDB).CreateCollection(
		context.TODO(),
		env.Params.MongoCollection,
	)

	if createErr != nil {
		log.Print(createErr)
	}

	instance := mongoserviceprovider.GetCollectionFromDB(
		mongoClient,
		env.Params.MongoDB,
		env.Params.MongoCollection,
	)

	models := replication.ConvertToMongoCompatible(arr)

	res, err := instance.BulkWrite(context.TODO(), models)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(
		"inserted %v and deleted %v documents\n",
		res.InsertedCount,
		res.DeletedCount)
}
