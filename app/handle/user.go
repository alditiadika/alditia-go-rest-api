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
	"github.com/gorilla/mux"
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
	mongoCursor, _ := collection.Find(context.TODO(), bson.M{"is_active": true}, options.Find().SetLimit(param.Take).SetSkip(param.Skip))
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

//GetOneUser func
func GetOneUser(res http.ResponseWriter, req *http.Request, client *mongo.Client) {
	//define variable
	strID := mux.Vars(req)["id"]
	id, _ := primitive.ObjectIDFromHex(strID)
	collection := client.Database(dbName).Collection("users")

	var ret model.UserModel
	err := collection.FindOne(context.TODO(), bson.M{"_id": bson.M{"$eq": id}}).Decode(&ret)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	collection.FindOne(context.TODO(), bson.M{"_id": bson.M{"$eq": id}}).Decode(&ret)
	//send data
	fmt.Printf("%s %s\n", "[GET]", req.URL)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ret)
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

	fmt.Printf("%s %s\n", "[POST]", req.URL)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(result)
}

// UpdateUser scheme
func UpdateUser(res http.ResponseWriter, req *http.Request, client *mongo.Client) {
	//define variable
	strID := mux.Vars(req)["id"]
	id, _ := primitive.ObjectIDFromHex(strID)
	collection := client.Database(dbName).Collection("users")
	var param model.UserModel

	//search data before update
	var ret model.UserModel
	err := collection.FindOne(context.TODO(), bson.M{"_id": bson.M{"$eq": id}}).Decode(&ret)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	//convert req.body to struct
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &param)

	//query run
	data := bson.M{
		"$set": bson.M{
			"first_name":    param.Firstname,
			"last_name":     param.Lastname,
			"is_active":     param.IsActive,
			"created_by":    param.CreatedBy,
			"modified_by":   param.ModifiedBy,
			"modified_date": time.Now(),
		},
	}
	collection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": id}}, data)
	collection.FindOne(context.TODO(), bson.M{"_id": bson.M{"$eq": id}}).Decode(&ret)
	//send data
	fmt.Printf("%s %s\n", "[PUT]", req.URL)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ret)

}

// DeleteUser scheme
func DeleteUser(res http.ResponseWriter, req *http.Request, client *mongo.Client) {
	//define variable
	strID := mux.Vars(req)["id"]
	id, _ := primitive.ObjectIDFromHex(strID)
	collection := client.Database(dbName).Collection("users")
	var param model.UserModel

	//search data before update
	var ret model.UserModel
	err := collection.FindOne(context.TODO(), bson.M{"_id": bson.M{"$eq": id}}).Decode(&ret)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	//convert req.body to struct
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &param)

	//query run
	data := bson.M{"$set": bson.M{"is_active": false}}
	collection.UpdateOne(context.TODO(), bson.M{"_id": bson.M{"$eq": id}}, data)
	collection.FindOne(context.TODO(), bson.M{"_id": bson.M{"$eq": id}}).Decode(&ret)
	//send data
	fmt.Printf("%s %s", "[DEL]", req.URL)
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(ret)

}
