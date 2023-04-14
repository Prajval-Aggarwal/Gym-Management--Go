package user

import (
	"fmt"
	"gym/server/db"
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"

	"github.com/gin-gonic/gin"
)

func CreateUserService(context *gin.Context, decodedData request.CreateUserRequest) {

	var userCreated model.User
	userCreated.User_Name = decodedData.Name
	userCreated.Gender = decodedData.Gender

	err := db.CreateRecord(&userCreated)
	if err != nil {
		fmt.Println("FAT GYA")
		response.ErrorResponse(context, 500, err.Error())
	}

}
