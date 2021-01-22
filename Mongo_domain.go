package main

import "go.mongodb.org/mongo-driver/mongo"

type MongoConfig struct {
	Host string
	Database string
	Port string
}

type CounterCollectionInterface interface {
	GetCounterDB() (*Counter, error)
	NilCounterDB() error
}


type CounterCollectionClass struct{
	dbcon *mongo.Database
}
