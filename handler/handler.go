package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	DataBase *sqlx.DB
}

func (s *Server) CreateStudentHandler(context *gin.Context) {

	var err error

	firstN, ok := context.GetQuery("first_name")
	if firstN == "" || !ok {
		context.Writer.WriteString("No first name")
		return
	}

	lastN, ok := context.GetQuery("last_name")
	if lastN == "" || !ok {
		context.Writer.WriteString("No last name")
		return
	}

	_, err = s.DataBase.Exec("INSERT INTO student(first_name, last_name) VALUES (?,?)", firstN, lastN)
	if err != nil {
		context.Writer.WriteString("Something's not right. Try again")
		fmt.Println("!!!!!!!", err)
		context.Status(500)
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) DelStudentHandler(context *gin.Context) {
	var err error

	firstN, ok := context.GetQuery("first_name")
	if firstN == "" || !ok {
		context.Writer.WriteString("No first name")
		return
	}

	lastN, ok := context.GetQuery("last_name")
	if lastN == "" || !ok {
		context.Writer.WriteString("No last name")
		return
	}

	res, err := s.DataBase.Exec("DELETE FROM student WHERE first_name = ? AND last_name = ?", firstN, lastN)
	if err != nil {
		context.Writer.WriteString("Something's not right. Try again")
		context.Status(500)
		return
	}

	countOfDeletedRows, err := res.RowsAffected()
	if err != nil {
		context.Writer.WriteString("Something went wrong")
		context.Status(500)
		return
	}

	if countOfDeletedRows == 0 {
		context.Writer.WriteString("Wrong first name or last name. Try again")
		context.Status(500)
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}
