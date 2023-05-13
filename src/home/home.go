package home

import (
	"fmt"
	"net/http"

	"golang_example/utils"
	"golang_example/pkg/db"
	"golang_example/pkg/jwt"
	"golang_example/structs"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	e := jwt.ValidateRequest(r)

	if e != nil {
		utils.WriteError(w, e)
		return
	}

	data, e := GetHomeData()

	if e != nil {
		fmt.Println(e)
		utils.WriteError(w, e)
		return
	}

	utils.WriteResponse(w, data)
}

func GetHomeData() ([]any, error) {
	var dafts []struct {
		structs.Package         `xorm:"extends"`
		structs.LearningProgram `xorm:"extends"`
		structs.SubjectType     `xorm:"extends"`
	}

	db.Source().Table("package").
		Join("inner", "learning_program", "package.learning_program_id = learning_program.id").
		Join("inner", "subject_type", "package.subject_type_id = subject_type.id").
		Find(&dafts)

	len := len(dafts)

	if len > 0 {

		collects := make([]any, 0, len)
		for _, v := range dafts {
			collects = append(collects, struct {
				structs.Package
				LearningProgram structs.LearningProgram `json:"learning_program" bson:"learning_program"`
				SubjectType     structs.SubjectType     `json:"subject_type" bson:"subject_type"`
			}{
				Package:         v.Package,
				SubjectType:     v.SubjectType,
				LearningProgram: v.LearningProgram,
			})
		}

		return collects, nil
	}

	return nil, fmt.Errorf("Not found any Package")
}
