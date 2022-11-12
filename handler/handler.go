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
	GroupID     *string `db:"group_id" json:"GroupID"`
	Room        *string `db:"room" json:"room"`
	DateOfBirth *string `db:"date_of_birth" json:"dateOfBirth"`
}

type Room struct {
	RoomNumber string `db:"room_number" json:"room_umber"`
	NumOfDBeds int    `db:"number_of_beds" json:"number_of_beds"`
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

func (s *Server) GetIdStudentsHandler(context *gin.Context) {

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

	jsonInByte, err := json.Marshal(resultTable)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
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
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
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

func (s *Server) GetStudentsByIdHandler(context *gin.Context) {

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

	jsonInByte, err := json.Marshal(resultTable)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
}

func (s *Server) GetStudentsByNameHandler(context *gin.Context) {

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

	jsonInByte, err := json.Marshal(resultTable)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}
	//fmt.Println(string(jsonInByte))

	context.Writer.Write(jsonInByte)
}

func (s *Server) CreateRoomHandler(context *gin.Context) {

	var err error

	roomNum, ok := context.GetQuery("room_number")
	if roomNum == "" || !ok {
		context.Writer.WriteString("No room number")
		return
	}

	beds, ok := context.GetQuery("number_of_beds")
	if beds == "" || !ok {
		context.Writer.WriteString("No number of beds")
		return
	}

	_, err = s.DataBase.Exec("INSERT INTO rooms(room_number, number_of_beds) VALUES (?,?)", roomNum, beds)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something's not right. Try again")
		fmt.Println("!!!!!!!", err)
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) GetRoomStudentsHandler(context *gin.Context) {

	var err error

	roomNum, ok := context.GetQuery("room_number")
	if roomNum == "" || !ok {
		context.Writer.WriteString("Missing room number")
		return
	}

	var resultTable []Student

	err = s.DataBase.Select(&resultTable, "SELECT * FROM student WHERE room = ?", roomNum)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		fmt.Println("!!!!!!!!!!!! - ", err)
		return
	}

	if len(resultTable) == 0 {
		context.Status(404)
		context.Writer.WriteString("No data with this rooms")
		return
	}

	jsonInByte, err := json.Marshal(resultTable)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
}

func (s *Server) GetRoomHandler(context *gin.Context) {

	var err error

	roomNum, ok := context.GetQuery("room_number")
	if roomNum == "" || !ok {
		context.Writer.WriteString("Missing room number")
		return
	}

	var resultTable []Room

	err = s.DataBase.Select(&resultTable, "SELECT * FROM rooms WHERE room_number = ?", roomNum)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		fmt.Println("!!!!!!!!!!!! - ", err)
		return
	}

	if len(resultTable) == 0 {
		context.Status(404)
		context.Writer.WriteString("No data with this rooms")
		return
	}

	jsonInByte, err := json.Marshal(resultTable)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)

}

func (s *Server) DelRoomHandler(context *gin.Context) {

	var err error

	roomNum, ok := context.GetQuery("room_number")
	if roomNum == "" || !ok {
		context.Writer.WriteString("No room number")
		return
	}

	res, err := s.DataBase.Exec("DELETE FROM rooms WHERE room_number = ?", roomNum)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
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

func (s *Server) AddToRoomHandler(context *gin.Context) {

	var err error

	roomNum, ok := context.GetQuery("room_number")
	if roomNum == "" || !ok {
		context.Writer.WriteString("No room number")
		return
	}

	studId, ok := context.GetQuery("student_id")
	if studId == "" || !ok {
		context.Writer.WriteString("No student ID")
		return
	}

	res, err := s.DataBase.Exec("UPDATE student SET room = ? WHERE id = ?", roomNum, studId)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something's not right. Try again")
		fmt.Println("!!!!!!!", err)
		return
	}

	countOfDeletedRows, err := res.RowsAffected()
	if err != nil {
		context.Writer.WriteString("Something went wrong")
		context.Status(500)
		return
	}

	if countOfDeletedRows == 0 {
		context.Writer.WriteString("There is no student with this ID")
		context.Status(500)
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}
