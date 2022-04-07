package main

import (
	"backend/config"
	"backend/controller/logistic"
	"backend/controller/user"
	_ "backend/docs"
	"backend/middleware"
	"os"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var (
	//DB CONNECTION
	db *gorm.DB = config.SetupDBConnection()

	//USER ROUTES
	userRepository user.UsersRepository = user.NewUsersRepository(db)
	userService    user.UserService     = user.NewUserService(userRepository)
	userController user.UserController  = user.NewUserController(userService, jwtService)

	//AUTH ROUTES
	authService    user.AuthService    = user.NewAuthService(userRepository)
	jwtService     user.JWTService     = user.NewJWTService()
	authController user.AuthController = user.NewAuthController(authService, jwtService)

	//LOGISTIC ROUTES
	logisticRepository logistic.LogisticRepository = logistic.NewLogisticRepository(db)
	logisticService    logistic.LogisticService    = logistic.NewLogisticService(logisticRepository)
	logisticController logistic.LogisticController = logistic.NewLogisticController(logisticService, jwtService)
)

// @title API DOCUMENTATION - API FOR LOLIPAD 2022
// @version 1.0
// @description This is a Api Documentation for LOLIPAD 2022 with Golang Backend, Gin Gonic Framework, GORM MySQL, and JWT Authentication
// @termsOfService http://swagger.io/terms/
// contact.name API Support
// contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// routes.GlobalRoutes(db)
	// routes.UserRoutes(db)
	r.Use(middleware.SetupCorsMiddleware())
	r.StaticFS("/web", http.Dir("web"))

	authRoute := r.Group("api/auth")
	{
		authRoute.POST("/login", authController.Login)
		authRoute.POST("/register", authController.Register)
	}

	usersRoute := r.Group("api/users")
	{
		usersRoute.GET("/", userController.Profile)
		usersRoute.PUT("/:id", userController.Update)
	}

	logisticRoute := r.Group("api/logistic")
	{
		logisticRoute.GET("/", logisticController.GetAllLogistic)
		logisticRoute.POST("/", logisticController.InsertLogistic)
		logisticRoute.GET("/bydestination", logisticController.FindLogisticByID)
		logisticRoute.PUT("/:id", logisticController.UpdateLogistic)
		logisticRoute.DELETE("/:id", logisticController.DeleteLogistic)
	}

	r.Run(": " + os.Getenv("SERVER_PORT"))
}
