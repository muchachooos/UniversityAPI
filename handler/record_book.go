package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreateRecordBookHandler(context *gin.Context) {

	studId, ok := context.GetQuery("student_id")
	if studId == "" || !ok {
		context.Writer.WriteString("Missing student ID")
		return
	}

	err := s.Storage.CreateRecordBookInDB(studId)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something's not right. Try again")
		return
	}

	context.Writer.WriteString("Welcome to the club Body")
}

func (s *Server) GetIdRecordBookHandler(context *gin.Context) {

	studId, ok := context.GetQuery("student_id")
	if studId == "" || !ok {
		context.Writer.WriteString("No student ID")
		return
	}

	res, err := s.Storage.GetIDRecordBookFromDB(studId)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if len(res) == 0 {
		context.Status(404)
		context.Writer.WriteString("No record book with this data")
		return
	}

	jsonInByte, err := json.Marshal(res)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
}

func (s *Server) DelRecordBookHandler(context *gin.Context) {

	bookId, ok := context.GetQuery("id_record_book")
	if bookId == "" || !ok {
		context.Writer.WriteString("Missing record book ID")
		return
	}

	isDeleted, err := s.Storage.DeleteRecordBookFromDB(bookId)
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

func (s *Server) GetRecordBookHandler(context *gin.Context) {

	bookId, ok := context.GetQuery("id_record_book")
	if bookId == "" || !ok {
		context.Writer.WriteString("Missing record book ID ")
		return
	}

	res, err := s.Storage.GetRecordBookFromDB(bookId)
	if err != nil {
		context.Status(500)
		context.Writer.WriteString("Something went wrong. Try again")
		return
	}

	if len(res) == 0 {
		context.Status(404)
		context.Writer.WriteString("No record book with this data")
		return
	}

	jsonInByte, err := json.Marshal(res)
	if err != nil {
		context.Writer.WriteString("json creating error")
		return
	}

	context.Writer.Write(jsonInByte)
}
