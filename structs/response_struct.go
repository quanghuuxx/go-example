package structs

type BaseResponse struct {
	Timestamp int `json:"timestamp" bson:"timestamp"`
	Data      any `json:"data" bson:"data"`
}
