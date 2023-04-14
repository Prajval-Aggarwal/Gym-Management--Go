package server

import (
	"gym/server/handler"
)

func ConfigureRoutes(server *Server) {

	//Membership routes
	server.engine.GET("getMembership", handler.GetMembershipHandler)
	server.engine.POST("/createMembership", handler.CreateMembershipHandler)
	server.engine.PUT("/updateMembership", handler.UpdateMembershipHandler)
	server.engine.PUT("/deleteMembership", handler.DeleteMembershipHandler)

	//Equipments routes
	server.engine.POST("/createEquipment", handler.CreateEquipmentHandler)
	server.engine.GET("/getEquipments", handler.GetEquipmentHandler)

	//employee routes
	server.engine.GET("/createEmp", handler.CreateEmployeeHandler)
	server.engine.GET("/getEmp", handler.GetEmployeeHandler)
	server.engine.GET("/getEmpRole", handler.GetEmployeeRoleHandler)
	server.engine.POST("/createEmpRole", handler.EmployeeRoleHandler)
	server.engine.GET("/empWithuser", handler.GetUsersWithEmployeesHandler)
	server.engine.POST("/empAttendence", handler.EmployeeAttendenceHandler)

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
