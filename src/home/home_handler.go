package home

import (
	"net/http"

	"golang_example/src/home/controller"
	"golang_example/utils"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	result, e := controller.Handle(r)

	if e != nil {
		utils.WriteError(w, e)
		return
	}

	utils.WriteResponse(w, result)
}