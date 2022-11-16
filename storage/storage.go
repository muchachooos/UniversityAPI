package storage

import "github.com/jmoiron/sqlx"

type UniversityStorage struct {
	DataBase *sqlx.DB
}

