package db

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	source   *xorm.Engine
	activity *xorm.Engine
)

func Open() {
	if source == nil {
		s, e := xorm.NewEngine("sqlite3", "source.db")
		if e != nil {
			panic(e)
		}

		source = s
	}

	if activity == nil {
		a, e := xorm.NewEngine("sqlite3", "activity.db")
		if e != nil {
			panic(e)
		}

		activity = a
	}
}

func Source() *xorm.Engine {
	return source
}

func Activity() *xorm.Engine {
	return activity
}
