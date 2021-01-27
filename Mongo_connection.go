package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)


var (
	collection *mongo.Collection
)


func CounterTicker()   {
	ticker := time.NewTicker(time.Hour)
	for {
		counter:=&Counter{}
		filter:=bson.D{}
		err:=collection.FindOne(context.TODO(),filter).Decode(&counter)
		if err!=nil{
			break
		}

		t := time.Now()
		if t.Hour() == 1 && t.Day() != counter.CurrentDate.Day()  {

			update := bson.D{{"$set", bson.D{
				{"days", counter.Days + 1},
			}}}

			_, err = collection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				break
			}
		}

		<-ticker.C
	}

}



func InitCounterCollection(config MongoConfig) (CounterCollectionInterface, error) {

	clientOptions:=options.Client().ApplyURI("mongodb://"+config.Host+":"+config.Port)
	client,err := mongo.Connect(context.TODO(),clientOptions)
	if err!=nil{
		return nil, err
	}
	err = client.Ping(context.TODO(),nil)
	if err!=nil{
		return nil, err
	}
	db:=client.Database(config.Database)
	collection=db.Collection("Counter")


	counter := Counter{
		Days:        0,
		CurrentDate: time.Now(),
	}


	_, err = collection.InsertOne(context.TODO(), counter)
	if err!=nil{
		return nil, err
	}

	fmt.Println("Counter start")

	return &CounterCollectionClass{dbcon:db,}, nil

}

func (mg *CounterCollectionClass) GetCounterDB() (*Counter, error) {

	filter:=bson.D{}
	counter:=&Counter{}
	err:=collection.FindOne(context.TODO(),filter).Decode(&counter)
	if err!=nil{
		return nil, err
	}
	return counter, nil

}

func (mg *CounterCollectionClass) NilCounterDB() error{
	filter:=bson.D{}
	update:=bson.D{{"$set",bson.D{
		{"days",0},
		{"currentdate",time.Now()},

	}}}
	_, err:=collection.UpdateOne(context.TODO(),filter,update)
	if err!=nil {
		return err
	}
	return nil
}