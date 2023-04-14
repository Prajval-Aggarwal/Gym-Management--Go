package handler

import (
	"encoding/json"
	"gym/server/request"
	"gym/server/response"
	slot "gym/server/services/slots"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func SlotUpdateHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	var userId request.UpdateSlotRequest

	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &userId)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	slot.SlotServices(context, userId)

}
