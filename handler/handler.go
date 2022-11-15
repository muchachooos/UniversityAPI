package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Student struct {
	ID          int     `db:"id" json:"id"`
	FirstName   string  `db:"first_name" json:"firstName"`
	LastName    string  `db:"last_name" json:"lastName"`
	GroupID     *string `db:"group_id" json:"GroupID"`
	Room        *string `db:"room" json:"room"`
	DateOfBirth *string `db:"date_of_birth" json:"date_of_birth"`
}

type Room struct {
	RoomNumber string `db:"room_number" json:"room_umber"`
	NumOfDBeds int    `db:"number_of_beds" json:"number_of_beds"`
}

type ID struct {
	ID int `db:"id" json:"id"`
}

func (s *Server) CreateStudentHandler(context *gin.Context) {

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

	err := createStudentInDB(s.DataBase, firstN, lastN, birth)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something's not right. Try again")
		fmt.Println("!!!!!!!", err)
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) GetIdStudentsHandler(context *gin.Context) {

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

	res, err := getIdStudentFromDB(s.DataBase, firstN, lastN, birth)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		fmt.Println("!!!!!!!!!!!! - ", err)
		return
	}

	if len(res) == 0 {
		context.Status(404)
		context.Writer.WriteString("No students with this data")
		return
	}

	jsonInByte, err := json.Marshal(res)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
}

func (s *Server) DelStudentHandler(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if id == "" || !ok {
		context.Writer.WriteString("No ID")
		return
	}

	isDeleted, err := deleteStudentFromDB(s.DataBase, id)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if isDeleted == false {
		context.Writer.WriteString("Something went wrong")
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) GetStudentsByIdHandler(context *gin.Context) {

	id, ok := context.GetQuery("id")
	if id == "" || !ok {
		context.Writer.WriteString("No ID")
		return
	}

	res, err := getStudentFromDB(s.DataBase, id)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		fmt.Println("!!!!!!!!!!!! - ", err)
		return
	}

	if len(res) == 0 {
		context.Status(404)
		context.Writer.WriteString("No students with this ID")
		return
	}

	jsonInByte, err := json.Marshal(res)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
}

func (s *Server) GetStudentsByNameHandler(context *gin.Context) {

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

	res, err := getStudentByNameFromDB(s.DataBase, firstN, lastN)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		fmt.Println("!!!!!!!!!!!! - ", err)
		return
	}

	if len(res) == 0 {
		context.Status(404)
		context.Writer.WriteString("No students with this name")
		return
	}

	jsonInByte, err := json.Marshal(res)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
}

func (s *Server) CreateRoomHandler(context *gin.Context) {

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

	err := createRoomInDB(s.DataBase, roomNum, beds)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something's not right. Try again")
		fmt.Println("!!!!!!!", err)
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) DelRoomHandler(context *gin.Context) {

	roomNum, ok := context.GetQuery("room_number")
	if roomNum == "" || !ok {
		context.Writer.WriteString("No room number")
		return
	}

	isDeleted, err := deleteRoomFromDB(s.DataBase, roomNum)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if isDeleted == false {
		context.Writer.WriteString("Something went wrong")
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) AddToRoomHandler(context *gin.Context) {

	studId, ok := context.GetQuery("student_id")
	if studId == "" || !ok {
		context.Writer.WriteString("No student ID")
		return
	}

	roomNum, ok := context.GetQuery("room_number")
	if roomNum == "" || !ok {
		context.Writer.WriteString("No room number")
		return
	}

	res, err := addToRoomInDB(s.DataBase, roomNum, studId)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something's not right. Try again")
		fmt.Println("!!!!!!!", err)
		return
	}

	if res == false {
		context.Writer.WriteString("Something went wrong")
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) RemoveFromRoomHandler(context *gin.Context) {

	studId, ok := context.GetQuery("student_id")
	if studId == "" || !ok {
		context.Writer.WriteString("Missing student ID")
		return
	}

	roomNum, ok := context.GetQuery("room_number")
	if roomNum == "" || !ok {
		context.Writer.WriteString("Missing student ID")
		return
	}

	res, err := removeFromRoomInDB(s.DataBase, studId, roomNum)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something's not right. Try again")
		fmt.Println("!!!!!!!", err)
		return
	}

	if res == false {
		context.Writer.WriteString("Something went wrong")
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

//func getRoomStudentFromDB(db *sqlx.DB, roomNum string) ([]Student, error) {
//
//}

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

func (s *Server) CreateGroupHandler(context *gin.Context) {

	var err error

	id, ok := context.GetQuery("group_id")
	if id == "" || !ok {
		context.Writer.WriteString("Missing group ID")
		return
	}

	course, ok := context.GetQuery("course")
	if course == "" || !ok {
		context.Writer.WriteString("Missing course")
		return
	}

	places, ok := context.GetQuery("number_of_places")
	if places == "" || !ok {
		context.Writer.WriteString("Missing number of places")
		return
	}

	spec, ok := context.GetQuery("specialization")
	if spec == "" || !ok {
		context.Writer.WriteString("Missing specialization")
		return
	}

	_, err = s.DataBase.Exec("INSERT INTO `group`(id, course, number_of_places, specialization) VALUES (?,?,?,?)", id, course, places, spec)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something's not right. Try again")
		fmt.Println("!!!!!!!", err)
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) DeleteGroupHandler(context *gin.Context) {

	var err error

	id, ok := context.GetQuery("group_id")
	if id == "" || !ok {
		context.Writer.WriteString("Missing group ID")
		return
	}

	res, err := s.DataBase.Exec("DELETE FROM `group` WHERE id = ?", id)
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
		context.Writer.WriteString("Wrong group ID. Try again")
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) AddToGroupHandler(context *gin.Context) {

}

func (s *Server) RemoveFromGroupHandler(context *gin.Context) {

}

func (s *Server) GetGroupStudentsHandler(context *gin.Context) {

}

func (s *Server) GetGroupHandler(context *gin.Context) {

}
