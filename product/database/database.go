package database

import (
	"context"
	"log"
	"time"
	"os"

	"nx-microray-api/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &DB{
		client: client,
	}
}

func (db *DB) FindByID(ID string) *model.Product {
	collection := db.client.Database("nx-microray").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"id": ID})
	product := model.Product{}
	res.Decode(&product)
	return &product
}

func (db *DB) All() []*model.Product {
	collection := db.client.Database("nx-microray").Collection("products")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var products []*model.Product
	for cur.Next(ctx) {
		var product *model.Product
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	return products
}

func (db *DB) GetFeaturedProducts(count int) ([]*model.Product, error) {
    collection := db.client.Database("nx-microray").Collection("products")
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // Using aggregate to get a random sample of products
    pipeline := mongo.Pipeline{
        {{"$sample", bson.M{"size": count}}},
    }

    cur, err := collection.Aggregate(ctx, pipeline)
    if err != nil {
        return nil, err
    }

    var products []*model.Product
    if err = cur.All(ctx, &products); err != nil {
        return nil, err
    }

    return products, nil
}