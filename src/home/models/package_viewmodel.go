package models

import "golang_example/model"

type PackageViewModel struct {
	model.Package         `xorm:"extends"`
	model.LearningProgram `xorm:"extends"`
	model.SubjectType     `xorm:"extends"`
}
