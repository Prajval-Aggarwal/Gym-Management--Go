package server

import (
	"gym/server/handler"
)

func ConfigureRoutes(server *Server) {

	//users
	server.engine.POST("/create", handler.CreateUserHandler)
	server.engine.POST("/getUser", handler.GetUserByIdHandler)
	server.engine.POST("/userAttendence", handler.UserAttendenceHandler)

	//slots
	server.engine.POST("/slotUpdate", handler.SlotUpdateHandler)

	//subscriptions 
	server.engine.POST("/createSubscription", handler.CreateSubscriptionHandler)
	server.engine.DELETE("/endSubscription", handler.EndSubscriptionHandler)






}
