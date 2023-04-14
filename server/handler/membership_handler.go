package handler

import (
	"encoding/json"
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/membership"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func CreateMembershipHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	var createMembership model.Membership
	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &createMembership)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	//Error handling
	validationErr := Validate.Struct(createMembership)
	if validationErr != nil {
		response.ErrorResponse(context, 400, validationErr.Error())
		return
	}
	membership.CreateMembershipService(context, createMembership)
}

func GetMembershipHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	membership.GetMembershipsService(context)
}

func UpdateMembershipHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")

	var updateMembership model.Membership
	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &updateMembership)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	//Error handling
	validationErr := Validate.Struct(updateMembership)
	if validationErr != nil {
		response.ErrorResponse(context, 400, validationErr.Error())
		return
	}
	membership.UpdateMembershipService(context, updateMembership)

}

func DeleteMembershipHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	var deleteMembersip request.DeleteMembershipRequest
	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &deleteMembersip)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	//Error handling
	validationErr := Validate.Struct(deleteMembersip)
	if validationErr != nil {
		response.ErrorResponse(context, 400, validationErr.Error())
		return
	}
	membership.DeleteMembershipService(context, deleteMembersip)
}
