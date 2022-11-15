package handler

import "github.com/jmoiron/sqlx"

type Server struct {
	DataBase *sqlx.DB
}
