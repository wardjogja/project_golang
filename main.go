package main

import (
	"github.com/julienschmidt/httprouter"
	"fmt"
	"net/http"
	TasksController "project_golang/controllers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"project_golang/models"
)



func main() {
	db, err := gorm.Open(sqlite.Open("database.db"),&gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Model.Task_model{})


	router := httprouter.New()

	router.ServeFiles("/static/*filepath",http.Dir("assets"))
	router.GET("/", TasksController.Index)
	router.GET("/create", TasksController.Create)
	router.POST("/create", TasksController.Create)
	router.GET("/update/:id", TasksController.Update)
	router.POST("/update/:id", TasksController.Update)
	router.GET("/delete/:id", TasksController.DeleteTask)
	
	
	fmt.Println("localhost://localhost:8080")
	http.ListenAndServe(":8080",router)

	
}