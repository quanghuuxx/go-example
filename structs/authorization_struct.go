package structs

type AuthorizationRequest struct {
	Device   string `json:"device" bson:"device"`
	Platform string `json:"platform" bson:"platform"`
	Language string `json:"language,omitempty" bson:"language,omitempty"`
}

func (r AuthorizationRequest) Empty() bool {
	return r.Device == "" || r.Platform == "" || r.Language == ""
}

type AuthoriztionReponse struct {
	Token          string `json:"token" bson:"token"`
	SuccessMessage string `json:"success_message,omitempty" bson:"success_message,omitempty"`
}
