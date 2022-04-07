package logistic

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/controller/user"
	"backend/helper"
	"backend/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//LogisticController interface is a contract what this controller can do
type LogisticController interface {
	GetAllLogistic(context *gin.Context)
	FindLogisticByID(context *gin.Context)
	InsertLogistic(context *gin.Context)
	UpdateLogistic(context *gin.Context)
	DeleteLogistic(context *gin.Context)
}

type logisticController struct {
	logisticService LogisticService
	jwtService      user.JWTService
}

//NewLogisticController create a new instances of LogisticController
func NewLogisticController(logisticServ LogisticService, jwtServ user.JWTService) LogisticController {
	return &logisticController{
		logisticService: logisticServ,
		jwtService:      jwtServ,
	}
}

// CreateLogistic
// @Security ApiKeyAuth
// @Description API untuk membuat logistic baru.
// @Summary Membuat logistic baru.
// @Tags Logistic
// @Accept json
// @Produce json
// @Param logistic body models.LOGISTIC true "Logistic Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /logistic [post]
func (c *logisticController) InsertLogistic(context *gin.Context) {
	var logisticCreateDTO models.LOGISTIC
	errDTO := context.ShouldBind(&logisticCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")
		userID := c.getUserIDByToken(authHeader)
		convertedUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			logisticCreateDTO.CREATE_USER = convertedUserID
			logisticCreateDTO.UPDATE_USER = convertedUserID
		}
		result := c.logisticService.InsertLogistic(logisticCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}
}

// GetLogistic Get All Logistic
// @Security ApiKeyAuth
// @Description API untuk mengambil semua logistic yang terdapat dalam database.
// @Summary Mengambil Semua Logistic
// @Tags Logistic
// @Accept json
// @Produce json
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /logistic [get]
func (c *logisticController) GetAllLogistic(context *gin.Context) {
	var logistic []models.LOGISTIC = c.logisticService.GetAllLogistic()
	res := helper.BuildResponse(true, "OK", logistic)
	context.JSON(http.StatusOK, res)
}

// GetLogistic by ID Logistic
// @Security ApiKeyAuth
// @Description API untuk mencari logistic by ID yang terdapat dalam database.
// @Summary Mengambil Logistic by ID
// @Tags Logistic
// @Accept json
// @Produce json
// @Param Origin_Name path string true "Origin Name"
// @Param Destination_Name path string true "Destination Name"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /logistic/bydestination [get]
func (c *logisticController) FindLogisticByID(context *gin.Context) {
	// id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	// OriginName := context.Param("origin_name")
	// DestinationName := context.Param("destination_name")
	// if err != nil {
	// 	res := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObj{})
	// 	context.AbortWithStatusJSON(http.StatusBadRequest, res)
	// 	return
	// }

	// var logistic models.LOGISTIC = c.logisticService.FindLogisticByOriginName(OriginName, DestinationName)
	// if (logistic == models.LOGISTIC{}) {
	// 	res := helper.BuildErrorResponse("Data not found", "No data with given id", helper.EmptyObj{})
	// 	context.JSON(http.StatusNotFound, res)
	// } else {
	// 	res := helper.BuildResponse(true, "OK", logistic)
	// 	context.JSON(http.StatusOK, res)
	// }

	OriginName := context.Param("origin_name")
	DestinationName := context.Param("destination_name")
	var tickets models.LOGISTIC = c.logisticService.FindLogisticByOriginName(OriginName, DestinationName)
	res := helper.BuildResponse(true, "OK", tickets)
	context.JSON(http.StatusOK, res)
}

// UpdateLogistic
// @Security ApiKeyAuth
// @Description API untuk update logistic.
// @Summary Update logistic.
// @Tags Logistic
// @Accept json
// @Produce json
// @Param logistic body models.LOGISTIC true "Logistic Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /logistic/{id} [put]
func (c *logisticController) UpdateLogistic(context *gin.Context) {
	var logisticUpdateDTO models.LOGISTIC
	id, errDTO := strconv.ParseUint(context.Param("id"), 0, 0)
	if errDTO != nil {
		res := helper.BuildErrorResponse("No param id was found", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	errDTO = context.ShouldBind(&logisticUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	idUser, errID := strconv.ParseUint(userID, 10, 64)
	logisticUpdateDTO.LOGISTIC_ID = id
	logisticUpdateDTO.UPDATE_USER = idUser
	if errID == nil {
		response := helper.BuildErrorResponse("User Id Not Found", "User Id not found", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
	result := c.logisticService.UpdateLogistic(logisticUpdateDTO)
	response := helper.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

// DeleteLogisticId
// @Security ApiKeyAuth
// @Description API untuk delete logistic.
// @Summary Delete logistic.
// @Tags Logistic
// @Accept json
// @Produce json
// @Param id path string true "Logistic ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /logistic/{id} [delete]
func (c *logisticController) DeleteLogistic(context *gin.Context) {
	var logistic models.LOGISTIC
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed tou get id", "No param id were found", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	logistic.LOGISTIC_ID = id
	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.logisticService.IsAllowedToEdit(userID, logistic.LOGISTIC_ID) {
		c.logisticService.DeleteLogistic(logistic)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		response := helper.BuildErrorResponse("You dont have permission", "You are not the owner", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}
}

func (c *logisticController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
