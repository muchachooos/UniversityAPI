package main

import (
	"UniversityAPI/handler"
	"UniversityAPI/model"
	"UniversityAPI/storage"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"
)

func main() {
	router := gin.Default()

	var conf model.Config

	byte, err := os.ReadFile("../configuration.json")
	if err != nil {
		fmt.Println("Error Read File:", err)
		return
	}

	err = json.Unmarshal(byte, &conf)
	if err != nil {
		fmt.Println("Error Unmarshal:", err)
		return
	}

	dataBase, err := sqlx.Open("mysql", conf.DataSourceName)
	if err != nil {
		panic(err)
		return
	}

	if dataBase == nil {
		fmt.Println("dB nil")
		panic(err)
		return
	}

	server := handler.Server{
		Storage: &storage.UniversityStorage{
			DataBase: dataBase,
		},
	}

	//Student
	router.POST("/create_student", server.CreateStudentHandler)
	router.GET("/get_id_student", server.GetIdStudentsHandler)
	router.DELETE("/delete_student", server.DelStudentHandler)
	router.GET("/get_students_info_by_id", server.GetStudentsByIdHandler)
	router.GET("/get_students_info_by_name", server.GetStudentsByNameHandler)

	//Room
	router.POST("/create_room", server.CreateRoomHandler)
	router.DELETE("/delete_room", server.DelRoomHandler)
	router.POST("/add_to_room", server.AddToRoomHandler)
	router.POST("/remove_from_room", server.RemoveFromRoomHandler)
	router.GET("/get_students_room", server.GetRoomStudentsHandler)
	router.GET("/get_room_info", server.GetRoomHandler)

	//Group
	router.POST("/create_group", server.CreateGroupHandler)
	router.DELETE("/delete_group", server.DeleteGroupHandler)
	router.POST("/add_to_group", server.AddToGroupHandler)
	router.POST("/remove_from_group", server.RemoveFromGroupHandler)
	router.GET("/get_students_group", server.GetGroupStudentsHandler)
	router.GET("/get_group_info", server.GetGroupHandler)

	//Record_book
	router.POST("/create_record_book", server.CreateRecordBookHandler)
	router.GET("/get_id_record_book", server.GetIdRecordBookHandler)
	router.DELETE("/delete_record_book", server.DelRecordBookHandler)
	router.GET("/get_record_book", server.GetRecordBookHandler)

	port := ":" + strconv.Itoa(conf.Port)

	err = router.Run(port)
	if err != nil {
		panic(err)
		return
	}
}
