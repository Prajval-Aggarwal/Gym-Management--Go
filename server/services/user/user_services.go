package user

import (
	"gym/server/db"
	"gym/server/model"
	"gym/server/request"
	"gym/server/response"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUserService(context *gin.Context, decodedData request.CreateUserRequest) {

	var userCreated model.User
	userCreated.User_Name = decodedData.Name
	userCreated.Gender = decodedData.Gender

	err := db.CreateRecord(&userCreated)
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
	}
	response.ShowResponse(
		"Success",
		200,
		"User Created successfully",
		userCreated,
		context,
	)

}

func GetUserByIdService(context *gin.Context, decodedData request.UserRequest) {
	var userGetter model.User
	err := db.FindById(&userGetter, decodedData.UserId, "user_id")
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	response.ShowResponse(
		"Success",
		200,
		"User retrieved",
		userGetter,
		context,
	)
}

func UserAttendenceService(context *gin.Context, userId request.UserRequest) {
	var userAttendence model.UAttendence
	now := time.Now()
	userAttendence.User_Id = userId.UserId
	userAttendence.Present = "Present"
	userAttendence.Date = now.Format("02 Jan 2006")

	err := db.CreateRecord(&userAttendence)
	if err != nil {
		response.ErrorResponse(context, 500, err.Error())
		return
	}
	response.ShowResponse(
		"Success",
		200,
		"User Attendence logged successfully",
		userAttendence,
		context,
	)

}

func GetAllUserServices(context *gin.Context){
	var displayAllUsers []model.Display

	query := "SELECT users.user_id,users.user_name,users.gender, payments.amount, payments.offer_amount,payments.offer,payments.payment_type, payments.payment_id, subscriptions.subs_name, subscriptions.start_date, subscriptions.deleted_at,subscriptions.end_date, subscriptions.duration, subscriptions.emp_id FROM users JOIN payments ON users.user_id = payments.user_id JOIN subscriptions ON payments.payment_id = subscriptions.payment_id;"

	err := db.QueryExecutor(query , &displayAllUsers)
	if err!=nil{
		response.ErrorResponse(context, 500, err.Error())
		return
	}

	response.ShowResponse(
		"Success",
		200,
		"Fetched all the resgistered users",
		displayAllUsers,
		context,
	)
}
