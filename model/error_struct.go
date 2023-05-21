package model

type ErrorResponse struct {
	Code        string `json:"code" bson:"code"`
	Message     string `json:"message" bson:"message"`
	Description string `json:"description,omitempty" bson:"description"`
}
