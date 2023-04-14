package handler

import (
	"encoding/json"
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/payment"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func MakePaymentHandler(context *gin.Context) {
	var createPayment request.CreatePaymentRequest

	//Decoding request body
	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &createPayment)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	//Error handling
	validationErr := Validate.Struct(createPayment)
	if validationErr != nil {
		response.ErrorResponse(context, 400, validationErr.Error())
		return
	}

	payment.MakePaymentService(context, createPayment)
}
