package server

import (
	"gym/server/handler"
)

func ConfigureRoutes(server *Server) {

	userHandler := handler.UserHandler{
		DB: server.db,
	}
	server.engine.GET("/create", userHandler.CreateUserHandler)
}
