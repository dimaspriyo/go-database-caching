package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Response struct {
	Start            time.Time   `json:'start'`
	End              time.Time   `json:'end'`
	TotalMicrosecond int64       `json:'total_microsecond'`
	Cache            bool        `json:'cache'`
	Data             interface{} `json:'data'`
}

type DataMongo struct {
	ID          interface{} `json:"id" bson:"_id"`
	UUID        interface{} `json:"uuid" bson:"uuid"`
	Name        string      `json:"name" bson:"name"`
	Type        string      `json:"type" bson:"type"`
	Permalink   string      `json:"permalink" bson:"permalink"`
	CbUrl       string      `json:"cb_url" bson:"cb_url"`
	Rank        int64       `json:"rank" bson:"rank"`
	CreatedAt   interface{} `json:"created_at" bson:"created_at"`
	UpdatedAt   interface{} `json:"updated_at" bson:"updated_at"`
	Description string      `json:"description" bson:"description"`
}

func main() {
	e := echo.New()

	e.GET("", func(c echo.Context) error {
		var response Response
		var start = time.Now()

		rdb := redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@mongo:27017"))
		if err != nil {
			panic(err.Error())
		}

		var data []DataMongo
		val, err := rdb.Get("people_description").Result()
		if err != nil {
			collection := client.Database("sample").Collection("people_description")
			cur, err := collection.Find(ctx, bson.D{})
			if err != nil {
				panic(err.Error())
			}

			for cur.Next(ctx) {
				var tes bson.M
				err = cur.Decode(&tes)
				if err != nil {
					panic(err.Error())
				}
				log.Println(tes)
			}

			err = cur.All(ctx, &data)
			if err != nil {
				panic(err.Error())
			}

			var redisValue []byte
			redisValue, err = json.Marshal(data)
			if err != nil {
				panic(err.Error())
			}

			err = rdb.Set("people_description", string(redisValue), 0).Err()
			if err != nil {
				panic(err)
			}

			response.Data = data
			response.Cache = false
		} else {
			var res []DataMongo

			err = json.Unmarshal([]byte(val), &res)
			if err != nil {
				panic(err.Error())
			}
			response.Data = res
			response.Cache = true
		}

		var end = time.Now()
		response.Start = start
		response.End = end
		response.TotalMicrosecond = end.Sub(start).Microseconds()

		return c.JSON(http.StatusOK, response)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
