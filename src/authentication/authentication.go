package authentication

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang_example/pkg/jwt"
	"golang_example/structs"
	"golang_example/utils"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var request structs.AuthorizationRequest

	e := json.NewDecoder(r.Body).Decode(&request)

	if e != nil || request.Empty() {
		utils.WriteError(w, http.ErrBodyNotAllowed)
	}

	result, e := Authorizating(request)

	if e != nil {
		fmt.Println(e)
		return
	}

	utils.WriteResponse(w, result)
}

func Authorizating(request structs.AuthorizationRequest) (*structs.AuthoriztionReponse, error) {
	token, e := jwt.GenerateAuthorizationToken(request)

	if e != nil {
		return nil, e
	}

	return &structs.AuthoriztionReponse{
		Token: token,
	}, nil
}
