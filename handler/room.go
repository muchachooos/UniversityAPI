package handler

import (
	"UniversityAPI/storage"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

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
