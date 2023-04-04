package main

import (
	"github.com/amiosamu/chatapp/db"
	"github.com/amiosamu/chatapp/internal/user"
	"github.com/amiosamu/chatapp/internal/wsocket"
	"github.com/amiosamu/chatapp/router"
	"log"
	"os"
)

func main() {
	dbConn, err := db.NewDB(db.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Username: os.Getenv("USERNAME"),
		DBName:   os.Getenv("DBNAME"),
		SSLMode:  os.Getenv("SSLMODE"),
		Password: os.Getenv("PASSWORD"),
	})
	if err != nil {
		log.Println("could not init db...")
	}

	userRepo := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	hub := wsocket.NewHub()
	wsHandler := wsocket.NewHandler(hub)

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:8000")
}
