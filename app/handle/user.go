package handle

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/alditiadika/alditia-go-rest-api/app/model"
	"github.com/alditiadika/alditia-go-rest-api/config"
	"github.com/alditiadika/alditia-go-rest-api/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbName string = config.GetConf().DBName

//GetUser func
func GetUser(res http.ResponseWriter, req *http.Request, client *mongo.Client) {
	//define variable
	collection := client.Database(dbName).Collection("users")
	var results []*model.UserModel
	var param utils.RequestParameter
	//convert req.body to struct
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &param)
	//call query get
	mongoCursor, _ := collection.Find(context.TODO(), bson.D{{}}, options.Find().SetLimit(param.Take).SetSkip(param.Skip))
	for mongoCursor.Next(context.TODO()) {
		var element model.UserModel
		err := mongoCursor.Decode(&element)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &element)
	}
	fmt.Printf("%s %s\n", "[GET]", req.URL)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(results)
}

//Insertuser to user collection
func Insertuser(res http.ResponseWriter, req *http.Request, client *mongo.Client) {
	//define variable
	collection := client.Database(dbName).Collection("users")
	id := primitive.NewObjectID()
	var param model.UserModel
	//convert req.body to struct
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &param)
	param.ID = id
	param.CreatedDate = time.Now().Local()
	param.ModifiedDate = time.Now().Local()
	collection.InsertOne(context.TODO(), param)
	result := model.UserModel{
		ID:           id,
		Firstname:    param.Firstname,
		Lastname:     param.Lastname,
		IsActive:     param.IsActive,
		CreatedBy:    param.CreatedBy,
		CreatedDate:  param.CreatedDate,
		ModifiedBy:   param.ModifiedBy,
		ModifiedDate: param.ModifiedDate,
	}

	json.NewEncoder(res).Encode(result)
}
