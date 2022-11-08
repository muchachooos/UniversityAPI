package main

import (
	"UniversityAPI/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	router := gin.New()

	dataBase, err := sqlx.Open("mysql", "root:040498usa_wot@tcp(127.0.0.1:3306)/university")
	if err != nil {
		panic(err)
		return
	}

	server := handler.Server{
		dataBase,
	}

	router.POST("/create_student", server.CreateStudentHandler)
	router.POST("/delete_student", server.DelStudentHandler)

	router.Run("localhost:8080")
}
