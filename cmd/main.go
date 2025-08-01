package main

import (
	"log"

	"github.com/Chandra5468/go-chat-app-dh/db"
	"github.com/Chandra5468/go-chat-app-dh/internal/users"
	"github.com/Chandra5468/go-chat-app-dh/router"
)

type x struct {
}

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}
	// db := dbConn.GetDb()

	// xy := &x{}
	userRepo := users.NewRepository(dbConn.GetDb())
	userSrv := users.NewService(userRepo)
	userHandler := users.NewHandler(userSrv)

	router.InitRouter(userHandler)
	router.Start("localhost:8000")
}
