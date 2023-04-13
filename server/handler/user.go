package handler

import (
	"encoding/json"
	"fmt"
	models "gym/server/model"
	"gym/server/repository"
	"gym/server/request"
	"gym/server/response"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (handler UserHandler) CreateUserHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	var createUserRequest request.CreateUserRequest

	reqBody, _ := ioutil.ReadAll(context.Request.Body)

	err := json.Unmarshal(reqBody, &createUserRequest)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	var userCreated models.User
	userCreated.User_Name = createUserRequest.Name
	userCreated.Gender = createUserRequest.Gender
	fmt.Println("create user is", userCreated)
	userRepository := repository.UserRespository{
		DB: handler.DB}

	userRepository.Create(&userCreated)
	response.ShowResponse(
		"Success",
		200,
		"User created successfully",
		createUserRequest,
		context,
	)
}
