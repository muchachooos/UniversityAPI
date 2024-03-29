package handler

import (
	"encoding/json"
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

	err := s.Storage.CreateStudentInDB(firstN, lastN, birth)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something's not right. Try again")
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

	res, err := s.Storage.GetIdStudentFromDB(firstN, lastN, birth)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
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

	isDeleted, err := s.Storage.DeleteStudentFromDB(id)
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

	res, err := s.Storage.GetStudentFromDB(id)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
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

	res, err := s.Storage.GetStudentByNameFromDB(firstN, lastN)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
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
