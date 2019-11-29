package utils

//RequestParameter struct model
type RequestParameter struct {
	Skip int64 `json:"skip" bson:"skip"`
	Take int64 `json:"take" bson:"take"`
}
