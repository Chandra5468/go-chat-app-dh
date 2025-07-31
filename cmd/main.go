package main

import (
	"log"

	"github.com/Chandra5468/go-chat-app-dh/db"
	"github.com/Chandra5468/go-chat-app-dh/internal/users"
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
	users.NewRepository(dbConn.GetDb())
}
