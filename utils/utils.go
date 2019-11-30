package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

import "net/http"

import "encoding/json"

import "fmt"

//RequestParameter struct model
type RequestParameter struct {
	Skip int64 `json:"skip" bson:"skip"`
	Take int64 `json:"take" bson:"take"`
}

//UpdateCollectionParameter struct
type UpdateCollectionParameter struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
}

//Send to api
func Send(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("err")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}
