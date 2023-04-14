package handler

import (
	"encoding/json"
	"gym/server/request"
	"gym/server/response"
	"gym/server/services/subscriptions"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func CreateSubscriptionHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")

	var subscriptionCreate request.CreateSubRequest

	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &subscriptionCreate)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	subscriptions.CreateSubscriptionService(context, subscriptionCreate)
}

func EndSubscriptionHandler(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")

	var subscriptionEnd request.EndSubRequest

	reqBody, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}
	err = json.Unmarshal(reqBody, &subscriptionEnd)
	if err != nil {
		response.ErrorResponse(context, 400, err.Error())
		return
	}

	subscriptions.EndSubscriptionService(context, subscriptionEnd)
}

// func UpdateSubscriptionHandler(context *gin.Context){
// 	context.Writer.Header().Set("Content-Type", "application/json")

// 	var subscriptionEnd request.UpdateSubRequest

// 	reqBody, _ := ioutil.ReadAll(context.Request.Body)

// 	err := json.Unmarshal(reqBody, &subscriptionEnd)
// 	if err != nil {
// 		response.ErrorResponse(context, 400, err.Error())
// 		return
// 	}

// 	subscriptions.UpdateSubscriptionService(context, subscriptionEnd)
// }
