package model

import "time"

import "go.mongodb.org/mongo-driver/bson/primitive"

//UserModel type
type UserModel struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname    string             `json:"first_name,omitempty" bson:"first_name,omitempty"`
	Lastname     string             `json:"last_name,omitempty" bson:"last_name,omitempty"`
	IsActive     bool               `json:"is_active,omitempty" bson:"is_active,omitempty"`
	CreatedBy    string             `json:"created_by,omitempty" bson:"created_by,omitempty"`
	ModifiedBy   string             `json:"modified_by,omitempty" bson:"modified_by,omitempty"`
	CreatedDate  time.Time          `json:"created_date,omitempty" bson:"created_date,omitempty"`
	ModifiedDate time.Time          `json:"modified_date,omitempty" bson:"modified_date,omitempty"`
}
