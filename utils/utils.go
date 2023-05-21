package utils

import (
	"encoding/json"
	"time"

	"golang_example/model"
)

func WriteResponse(r any) model.BaseResponse {
	re := model.BaseResponse{
		Timestamp: int(time.Now().Unix()),
		Data:      r,
	}

	return re
}

func WriteError(e error, infos ...any) model.ErrorResponse {
	cod := "Unknow Error"

	if len(infos) > 0 {
		cod = infos[0].(string)
	}

	err := model.ErrorResponse{
		Code:    cod,
		Message: e.Error(),
	}

	if len(infos) > 1 {
		if v, ok := infos[1].(string); ok {
			err.Description = v
		}
	}

	return err
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
