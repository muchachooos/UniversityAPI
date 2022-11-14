package main

import (
	"UniversityAPI/handler"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	router := gin.New()

	dataBase, err := sqlx.Open("mysql", "root:040498usa_wot@tcp(127.0.0.1:3306)/university")
	if err != nil {
		panic(err)
		return
	}

	server := handler.Server{
		dataBase,
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
	router.POST("/add_to_group", server.AddToGroupHandler)            // Not done
	router.POST("/remove_from_group", server.RemoveFromGroupHandler)  // Not done
	router.GET("/get_students_group", server.GetGroupStudentsHandler) // Not done
	router.GET("/get_group_info", server.GetGroupHandler)             // Not done

	router.Run("localhost:8080")
}
