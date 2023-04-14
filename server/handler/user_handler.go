package handler

import (
	"encoding/json"
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/user"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	var createUserRequest request.CreateUserRequest

	reqBody, _ := ioutil.ReadAll(context.Request.Body)

	err := json.Unmarshal(reqBody, &createUserRequest)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	user.CreateUserService(context, createUserRequest)

}
