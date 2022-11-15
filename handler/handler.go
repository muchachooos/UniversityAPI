package handler

import (
	"UniversityAPI/storage"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

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

	err := storage.CreateStudentInDB(s.DataBase, firstN, lastN, birth)
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

	res, err := storage.GetIdStudentFromDB(s.DataBase, firstN, lastN, birth)
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

	isDeleted, err := storage.DeleteStudentFromDB(s.DataBase, id)
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

	res, err := storage.GetStudentFromDB(s.DataBase, id)
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

	res, err := storage.GetStudentByNameFromDB(s.DataBase, firstN, lastN)
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

	err := storage.CreateRoomInDB(s.DataBase, roomNum, beds)
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

	isDeleted, err := storage.DeleteRoomFromDB(s.DataBase, roomNum)
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

	res, err := storage.AddToRoomInDB(s.DataBase, roomNum, studId)
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

	res, err := storage.RemoveFromRoomInDB(s.DataBase, studId, roomNum)
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

	roomNum, ok := context.GetQuery("room_number")
	if roomNum == "" || !ok {
		context.Writer.WriteString("Missing room number")
		return
	}

	res, err := storage.GetRoomStudentFromDB(s.DataBase, roomNum)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		fmt.Println("!!!!!!!!!!!! - ", err)
		return
	}

	if len(res) == 0 {
		context.Status(404)
		context.Writer.WriteString("No data with this rooms")
		return
	}

	jsonInByte, err := json.Marshal(res)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
}

func (s *Server) GetRoomHandler(context *gin.Context) {

	roomNum, ok := context.GetQuery("room_number")
	if roomNum == "" || !ok {
		context.Writer.WriteString("Missing room number")
		return
	}

	res, err := storage.GetRoomFromDB(s.DataBase, roomNum)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		fmt.Println("!!!!!!!!!!!! - ", err)
		return
	}

	if len(res) == 0 {
		context.Status(404)
		context.Writer.WriteString("No data with this rooms")
		return
	}

	jsonInByte, err := json.Marshal(res)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
}

func (s *Server) CreateGroupHandler(context *gin.Context) {

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

	err := storage.CreateGroupInDB(s.DataBase, id, course, places, spec)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something's not right. Try again")
		fmt.Println("!!!!!!!", err)
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) DeleteGroupHandler(context *gin.Context) {

	id, ok := context.GetQuery("group_id")
	if id == "" || !ok {
		context.Writer.WriteString("Missing group ID")
		return
	}

	res, err := storage.DeleteGroupFromDB(s.DataBase, id)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if res == false {
		context.Writer.WriteString("Something's not right")
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) AddToGroupHandler(context *gin.Context) {

	groupId, ok := context.GetQuery("group_id")
	if groupId == "" || !ok {
		context.Writer.WriteString("Missing group ID")
		return
	}

	studId, ok := context.GetQuery("student_id")
	if studId == "" || !ok {
		context.Writer.WriteString("Missing student ID")
		return
	}

	res, err := storage.AddToGroupInDB(s.DataBase, groupId, studId)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if res == false {
		context.Writer.WriteString("Something's not right")
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) RemoveFromGroupHandler(context *gin.Context) {

	groupId, ok := context.GetQuery("group_id")
	if groupId == "" || !ok {
		context.Writer.WriteString("Missing group ID")
		return
	}

	studId, ok := context.GetQuery("student_id")
	if studId == "" || !ok {
		context.Writer.WriteString("Missing student ID")
		return
	}

	res, err := storage.RemoveFromGroupInnDB(s.DataBase, groupId, studId)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if res == false {
		context.Writer.WriteString("Something's not right")
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) GetGroupStudentsHandler(context *gin.Context) {

	groupId, ok := context.GetQuery("group_id")
	if groupId == "" || !ok {
		context.Writer.WriteString("Missing group ID")
		return
	}

	res, err := storage.GetGroupStudentFromDB(s.DataBase, groupId)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		fmt.Println("!!!!!!!!!!!! - ", err)
		return
	}

	if len(res) == 0 {
		context.Status(404)
		context.Writer.WriteString("No data with this groups")
		return
	}

	jsonInByte, err := json.Marshal(res)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
}

func (s *Server) GetGroupHandler(context *gin.Context) {

	groupId, ok := context.GetQuery("group_id")
	if groupId == "" || !ok {
		context.Writer.WriteString("Missing group ID")
		return
	}

	res, err := storage.GetGroupFromDB(s.DataBase, groupId)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		fmt.Println("!!!!!!!!!!!! - ", err)
		return
	}

	if len(res) == 0 {
		context.Status(404)
		context.Writer.WriteString("No data with this groups")
		return
	}

	jsonInByte, err := json.Marshal(res)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
}
