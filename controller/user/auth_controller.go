package user

import (
	"net/http"
	"strconv"

	"backend/helper"
	"backend/models"

	"github.com/gin-gonic/gin"
)

//AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
}

type authController struct {
	authService AuthService
	jwtService  JWTService
}

//NewAuthController creates a new instance of AuthController
func NewAuthController(authService AuthService, jwtService JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

// Login User
// @Description Login User
// @Summary Login User
// @ID Authentication
// @Consume application/x-www-form-urlencoded
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.CredentialsLogin true "Input username & password"
// @Success 200 {object}  object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /auth/login [post]
func (c *authController) Login(context *gin.Context) {
	var credUser models.CredentialsLogin
	errDTO := context.ShouldBind(&credUser)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(credUser.Username, credUser.Password)
	if v, ok := authResult.(models.USER); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.UUID), 10))
		v.TOKEN = generatedToken
		response := helper.BuildResponse(true, "OK!", v)
		context.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your Username or Password", "Invalid Username or Password", helper.EmptyObj{})
	context.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

// Register User
// @Description Register User
// @Summary Register User
// @Consume application/x-www-form-urlencoded
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.RegisterUser true "Register User Data"
// @Success 200 {object} object
// @Header 200 {string} Token "qwerty"
// @Failure 400,500 {object} object
// @Router /auth/register [post]
func (c *authController) Register(context *gin.Context) {
	var registerDTO models.USER
	errDTO := context.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateUserName(registerDTO.USERNAME) {
		response := helper.BuildErrorResponse("Failed to process request", "Username Sudah Digunakan", helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
	} else if !c.authService.IsDuplicateMSISDN(int(registerDTO.MSISDN)) {
		response := helper.BuildErrorResponse("Failed to process request", "MSISDN Sudah Digunakan", helper.EmptyObj{})
		context.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		token := c.jwtService.GenerateToken(strconv.FormatUint(createdUser.MSISDN, 10))
		createdUser.USERNAME = token
		response := helper.BuildResponse(true, "OK!", createdUser)
		context.JSON(http.StatusCreated, response)
	}
}
