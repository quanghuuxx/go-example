package storage

import (
	"golang_example/pkg/db"
	"golang_example/src/home/models"

	"github.com/go-xorm/xorm"
)

type _HomeStorageImpl struct {
	db *xorm.Engine
}

func (s _HomeStorageImpl) FetchPackages() ([]models.PackageViewModel, error) {
	var dafts []models.PackageViewModel

	e := db.Source().Table("package").
		Join("inner", "learning_program", "package.learning_program_id = learning_program.id").
		Join("inner", "subject_type", "package.subject_type_id = subject_type.id").
		Find(&dafts)

	return dafts, e
}
