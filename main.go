package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	router := gin.New()

	_, err := sqlx.Open("mysql", "root:040498usa_wot@tcp(127.0.0.1:3306)/userdata")
	if err == nil {
		panic(err)
		return
	}
}
