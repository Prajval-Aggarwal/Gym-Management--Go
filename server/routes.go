package server

import (
	"gym/server/handler"
)

func ConfigureRoutes(server *Server) {


	server.engine.POST("/create", handler.CreateUserHandler)
}
