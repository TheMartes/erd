package environments

import (
	"context"
	"fmt"
	"log"

	"github.com/themartes/erd/config"
	"github.com/themartes/erd/config/envparams"
	"github.com/themartes/erd/persistance"
	mongoserviceprovider "github.com/themartes/erd/persistance/mongodb"
	"github.com/themartes/erd/replication"
)

// Local :)
type Local struct{}

// InitLocalEnv :))
func (l Local) InitLocalEnv() {
	mongoClient := persistance.GetMongoClient()
	db := config.GetEnvValue(envparams.MongoDB)
	collection := config.GetEnvValue(envparams.MongoCollection)
	arr := replication.GenerateFakeData(10000)

	mongoClient.Database(db).Collection(collection).Drop(context.TODO())
	mongoClient.Database(db).CreateCollection(context.TODO(), collection)

	instance := mongoserviceprovider.GetCollectionFromDB(mongoClient, db, collection)

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
