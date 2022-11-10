package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	DataBase *sqlx.DB
}

type Student struct {
	ID          int     `db:"id" json:"id"`
	FirstName   string  `db:"first_name" json:"firstName"`
	LastName    string  `db:"last_name" json:"lastName"`
	ClassID     *string `db:"class_id" json:"classID"`
	Room        *string `db:"room" json:"room"`
	DateOfBirth *string `db:"date_of_birth" json:"dateOfBirth"`
}

type ID struct {
	ID int `db:"id" json:"id"`
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
		context.Status(500)
		context.Writer.WriteString("Something's not right. Try again")
		fmt.Println("!!!!!!!", err)
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
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if countOfDeletedRows == 0 {
		context.Status(500)
		context.Writer.WriteString("Wrong ID. Try again")
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

	var resultTable []ID

	err = s.DataBase.Select(&resultTable, "SELECT id FROM student WHERE first_name = ? AND last_name = ? AND date_of_birth = ?", firstN, lastN, birth)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		fmt.Println("!!!!!!!!!!!! - ", err)
		return
	}

	if len(resultTable) == 0 {
		context.Status(404)
		context.Writer.WriteString("No students with this data")
		return
	}

	for i := range resultTable {
		idStudents := resultTable[i]

		jsonInByte, err := json.Marshal(idStudents)
		if err != nil {
			context.Writer.WriteString("json creating error")
			return
		}
		context.Writer.Write(jsonInByte)
	}
}

func (s *Server) GetStudentsInfoByName(context *gin.Context) {

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

	var resultTable []Student

	err = s.DataBase.Select(&resultTable, "SELECT * FROM student WHERE first_name = ? AND last_name = ?", firstN, lastN)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		fmt.Println("!!!!!!!!!!!! - ", err)
		return
	}

	if len(resultTable) == 0 {
		context.Status(404)
		context.Writer.WriteString("No students with this name")
		return
	}

	for i := range resultTable {
		studentInfo := resultTable[i]

		jsonInByte, err := json.Marshal(studentInfo)
		if err != nil {
			context.Writer.WriteString("json creating error")
			return
		}
		context.Writer.Write(jsonInByte)
	}
}

func (s *Server) GetStudentsInfoByID(context *gin.Context) {

	var err error

	id, ok := context.GetQuery("id")
	if id == "" || !ok {
		context.Writer.WriteString("No ID")
		return
	}

	var resultTable []Student

	err = s.DataBase.Select(&resultTable, "SELECT * FROM student WHERE id = ?", id)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		fmt.Println("!!!!!!!!!!!! - ", err)
		return
	}

	if len(resultTable) == 0 {
		context.Status(404)
		context.Writer.WriteString("No students with this ID")
		return
	}

	studentInfo := resultTable[0]

	jsonInByte, err := json.Marshal(studentInfo)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
}
