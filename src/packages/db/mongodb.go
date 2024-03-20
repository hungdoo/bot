package db

import (
	"context"
	"time"

	"github.com/hungdoo/bot/src/packages/dotenv"
	"github.com/hungdoo/bot/src/packages/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _db *MongoDB

type MongoDB struct {
	IsTestEnv bool
	Client    *mongo.Client
}

func newMongoDB() (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://" + dotenv.GetEnv("DBHost") + ":" + dotenv.GetEnv("DBPort"))

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.GeneralLogger.Printf("MongoDB Connected to: %s\n", clientOptions.GetURI())

	var isTestEnv bool
	env := dotenv.GetEnv("DBEnv")
	if env == "test" {
		isTestEnv = true
	}

	return &MongoDB{Client: client, IsTestEnv: isTestEnv}, nil
}

func (db *MongoDB) Close() {
	log.GeneralLogger.Println("MongoDB Closed")
	db.Client.Disconnect(context.TODO())
}
func (db *MongoDB) GetCollection(coll string) *mongo.Collection {
	return db.Client.Database(dotenv.GetEnv("DBName")).Collection(coll)
}

// Manipulate

func (db *MongoDB) Insert(coll string, data interface{}) error {
	collection := db.GetCollection(coll)
	_, err := collection.InsertOne(context.TODO(), data)
	return err
}

func (db *MongoDB) Update(coll string, filter interface{}, update interface{}) error {
	collection := db.GetCollection(coll)
	_, err := collection.UpdateMany(context.TODO(), filter, update)
	return err
}

func (db *MongoDB) Delete(coll string, query interface{}) error {
	collection := db.GetCollection(coll)
	_, err := collection.DeleteMany(context.TODO(), query)
	return err
}

// Query

func (db *MongoDB) Find(coll string, query interface{}) (*mongo.Cursor, error) {
	collection := db.GetCollection(coll)
	return collection.Find(context.TODO(), query)
}

func GetDb() *MongoDB {
	if _db != nil {
		return _db
	}
	var err error
	_db, err = newMongoDB()
	if err != nil {
		log.ErrorLogger.Fatal(err)
	}

	return _db
}
