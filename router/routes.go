package router

import (
	"github.com/Chandra5468/go-chat-app-dh/internal/users"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *users.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.CreateUser)
	r.POST("/logout", userHandler.CreateUser)
}

func Start(addr string) error {
	return r.Run(addr)
}
