package user

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/helper"
	"backend/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//UserController is a ....
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
}

type userController struct {
	userService UserService
	jwtService  JWTService
}

//NewUserController is creating anew instance of UserControlller
func NewUserController(userService UserService, jwtService JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

// Update User
// @Description Update User
// @Summary Update User
// @Security ApiKeyAuth
// @Consume application/x-www-form-urlencoded
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /user/update/{id} [put]
func (c *userController) Update(context *gin.Context) {
	var userUpdate models.USER
	errDTO := context.ShouldBind(&userUpdate)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdate.UUID = id
	u := c.userService.Update(userUpdate)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

// Get User Profile
// @Description Get User Profile
// @Summary Get User Profile
// @Security ApiKeyAuth
// @Consume application/x-www-form-urlencoded
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /user/profile/{id} [get]
func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}
