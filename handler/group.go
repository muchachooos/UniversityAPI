package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

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

	err := s.Storage.CreateGroupInDB(id, course, places, spec)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something's not right. Try again")
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

	res, err := s.Storage.DeleteGroupFromDB(id)
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

	res, err := s.Storage.AddToGroupInDB(groupId, studId)
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

	res, err := s.Storage.RemoveFromGroupInnDB(groupId, studId)
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

	res, err := s.Storage.GetGroupStudentFromDB(groupId)
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

	res, err := s.Storage.GetGroupFromDB(groupId)
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
