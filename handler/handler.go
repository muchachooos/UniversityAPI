package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type Server struct {
	DataBase *sqlx.DB
}

type Student struct {
	ID          int     `db:"id"`
	FirstName   string  `db:"first_name"`
	LastName    string  `db:"last_name"`
	ClassID     *string `db:"class_id"`
	Room        *string `db:"room"`
	DateOfBirth *string `db:"date_of_birth"`
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

	birth, ok := context.GetQuery("date_of_birth")
	if birth == "" || !ok {
		context.Writer.WriteString("No date of birth")
		return
	}

	_, err = s.DataBase.Exec("INSERT INTO student(first_name, last_name, date_of_birth) VALUES (?,?,?)", firstN, lastN, birth)
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

	id, ok := context.GetQuery("id")
	if id == "" || !ok {
		context.Writer.WriteString("No ID")
		return
	}

	res, err := s.DataBase.Exec("DELETE FROM student WHERE id = ?", id)
	if err != nil {
		context.Writer.WriteString("Something went wrong. Try again")
		context.Status(500)
		return
	}

	countOfDeletedRows, err := res.RowsAffected()
	if err != nil {
		context.Writer.WriteString("Something went wrong. Try again")
		context.Status(500)
		return
	}

	if countOfDeletedRows == 0 {
		context.Writer.WriteString("Wrong ID. Try again")
		context.Status(500)
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) GetIdStudents(context *gin.Context) {

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

	birth, ok := context.GetQuery("date_of_birth")
	if birth == "" || !ok {
		context.Writer.WriteString("No date of birth")
		return
	}

	var resultTable []Student

	err = s.DataBase.Select(&resultTable, "SELECT * FROM student WHERE first_name = ? AND last_name = ? AND date_of_birth = ?", firstN, lastN, birth)
	if err != nil {
		context.Writer.WriteString("Something went wrong. Try again")
		fmt.Println("!!!!!!!!!!!! - ", err)
		context.Status(500)
		return
	}

	if len(resultTable) == 0 {
		context.Writer.WriteString("Wrong login or password. Try again")
		return
	}

	var i Student
	i = resultTable[0]
	var j int
	j = i.ID
	r := strconv.Itoa(j)
	context.Writer.WriteString("Welcome to the club Body! ID = ")
	context.Writer.WriteString(r)
}

//func (s *Server) GetStudentsByName(context *gin.Context) {
//
//}
