package db

import (
	"context"

	"github.com/hungdoo/bot/src/packages/dotenv"
	"github.com/hungdoo/bot/src/packages/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _db *MongoDB

type MongoDB struct {
	Client *mongo.Client
}

func newMongoDB() (*MongoDB, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://" + dotenv.GetEnv("DBHost") + ":" + dotenv.GetEnv("DBPort"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	log.GeneralLogger.Printf("MongoDB Connected to: %s\n", clientOptions.GetURI())
	return &MongoDB{Client: client}, nil
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
