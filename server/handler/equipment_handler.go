package handler

import (
	"encoding/json"
	"gym/server/model"
	"gym/server/response"
	"gym/server/services/equipment"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func CreateEquipmentHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	var createEquipment model.Equipment
	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &createEquipment)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	//Error handling
	validationErr := Validate.Struct(createEquipment)
	if validationErr != nil {
		response.ErrorResponse(context, 400, validationErr.Error())
		return
	}
	equipment.CreateEquipmentService(context, createEquipment)

}
func GetEquipmentHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")
	equipment.GetEquipmentService(context)
}
