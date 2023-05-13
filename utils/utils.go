package utils

import (
	"encoding/json"
	"net/http"
	"time"

	"golang_example/structs"
)

func WriteResponse(w http.ResponseWriter, r any) {
	re := structs.BaseResponse{
		Timestamp: int(time.Now().Unix()),
		Data:      r,
	}

	w.WriteHeader(200)

	e := json.NewEncoder(w).Encode(re)

	if e != nil {
		WriteError(w, e)
	}
}

func WriteError(w http.ResponseWriter, e error, infos ...any) {
	cod := "Unknow Error"

	if len(infos) > 0 {
		cod = infos[0].(string)
	}

	err := structs.ErrorResponse{
		Code:    cod,
		Message: e.Error(),
	}

	if len(infos) > 1 {
		w.WriteHeader(infos[1].(int))
	}

	if len(infos) > 2 {
		if v, ok := infos[2].(string); ok {
			err.Description = v
		}
	}

	WriteResponse(w, err)
}

// using json to convert v to m
func Convert(v any, m any) error {
	b, e := json.Marshal(v)

	if e != nil {
		return e
	}

	e = json.Unmarshal(b, m)
	if e != nil {
		return e
	}
	return nil
}

func RemoveAt[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}
