package server

import (
	"gym/server/handler"
)

func ConfigureRoutes(server *Server) {

	//User routes
	server.engine.POST("/createUser", handler.CreateUserHandler)

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

}
