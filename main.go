package main

import (
	"log"

	"users/database"
	"users/usersHandlers"
	"users/usersRepository"
	"users/usersServices"
	"github.com/gin-gonic/gin"
)

func main() {
	// เปิดการเชื่อมต่อกับฐานข้อมูล
	db := database.Postgresql()
	defer db.Close()

	// ตรวจสอบการเชื่อมต่อกับฐานข้อมูล
	err := db.Ping()
	if err != nil {
		log.Fatal("Database connection error: ", err)
	}

	// สร้าง instances ของ repositories, services, และ handlers
	r := usersRepository.NewRepository(db)
	s := usersServices.NewServices(r)
	h := usersHandlers.NewHandlers(s)

	
	router := gin.Default()
	
	router.POST("/api/register", h.PostRegisterHandler)
	
	
	err = router.Run(":8062")
	if err != nil {
		log.Fatal(err.Error())
	}
}
