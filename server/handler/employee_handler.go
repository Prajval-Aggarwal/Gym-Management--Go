package handler

import (
	"encoding/json"
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/employee"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func CreateEmployeeHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	var createEmployee request.CreateEmployeeRequest
	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &createEmployee)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	//Error handling
	validationErr := Validate.Struct(createEmployee)
	if validationErr != nil {
		response.ErrorResponse(context, 400, validationErr.Error())
		return
	}

	employee.CreateEmployeeService(context, createEmployee)
}

func GetEmployeeHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	employee.GetEmployeeService(context)
}

func GetEmployeeRoleHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	employee.GetEmployeeRoleService(context)
}

func GetUsersWithEmployeesHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	employee.GetUsersWithEmployeService(context)
}

func EmployeeAttendenceHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")

	var empId request.EmployeeRequest

	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &empId)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	//Error handling
	validationErr := Validate.Struct(empId)
	if validationErr != nil {
		response.ErrorResponse(context, 400, validationErr.Error())
		return
	}

	employee.EmployeeAttendenceService(context, empId)
}

func EmployeeRoleHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	var createEmpRole model.EmpTypes
	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &createEmpRole)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	//Error handling
	validationErr := Validate.Struct(createEmpRole)
	if validationErr != nil {
		response.ErrorResponse(context, 400, validationErr.Error())
		return
	}
	employee.EmployeeRoleService(context, createEmpRole)
}

func GetEmployeeByIdHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")

	var empId request.EmployeeRequest

	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &empId)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	employee.GetEmployeeByIdService(context, empId)
}
