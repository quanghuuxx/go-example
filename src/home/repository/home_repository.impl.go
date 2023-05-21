package repository

import (
	"golang_example/model"
	"golang_example/src/home/storage"
)

type _HomeRepositoryImpl struct {
	storage storage.HomeStorage
}

func (r _HomeRepositoryImpl) FetchPackages() ([]any, error) {
	dafts, e := r.storage.FetchPackages()

	len := len(dafts)

	if len > 0 {

		collects := make([]any, 0, len)
		for _, v := range dafts {
			collects = append(collects, struct {
				model.Package
				LearningProgram model.LearningProgram `json:"learning_program" bson:"learning_program"`
				SubjectType     model.SubjectType     `json:"subject_type" bson:"subject_type"`
			}{
				Package:         v.Package,
				SubjectType:     v.SubjectType,
				LearningProgram: v.LearningProgram,
			})
		}

		return collects, e
	}

	return nil, e
}