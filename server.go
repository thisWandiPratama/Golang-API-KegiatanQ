package main

import (
	"golang_api_kegiatanQ/config"
	"golang_api_kegiatanQ/controllers"
	"golang_api_kegiatanQ/middleware"
	"golang_api_kegiatanQ/repository"
	"golang_api_kegiatanQ/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db                        *gorm.DB                             = config.SetupDatabaseConnection()
	userRepository            repository.UserRepository            = repository.NewUserRepository(db)
	kegiatanqRepository       repository.KegiatanRepository        = repository.NewKegiatanRepository(db)
	kegiatanqfinishRepository repository.KegiatanQFinishRepository = repository.NewKegiatanQFinishRepository(db)
	jwtService                service.JWTService                   = service.NewJWTService()
	userService               service.UserService                  = service.NewUserService(userRepository)
	kegiatanqService          service.KegiatanQService             = service.NewKegiatanQService(kegiatanqRepository)
	isfinishService           service.KegiatanQFinishService       = service.NewKegiatanQFinishService(kegiatanqfinishRepository)
	authService               service.AuthService                  = service.NewAuthService(userRepository)
	authController            controllers.AuthController           = controllers.NewAuthController(authService, jwtService)
	userController            controllers.UserController           = controllers.NewUserController(userService, jwtService)
	kegiatanqController       controllers.KegiatanQController      = controllers.NewKegiatanQController(kegiatanqService, jwtService)
	isfinishController        controllers.FinishController         = controllers.NewIsFinishController(isfinishService, jwtService)
)

func main() {

	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/register", authController.Register)
		}

		userRoutes := v1.Group("/user", middleware.AuthorizeJWT(jwtService))
		{
			userRoutes.GET("/profile", userController.Profile)
			userRoutes.PUT("/profile", userController.Update)
		}

		hutangRoutes := v1.Group("/kegiatan", middleware.AuthorizeJWT(jwtService))
		{
			hutangRoutes.GET("/", kegiatanqController.All)
			hutangRoutes.POST("/", kegiatanqController.Insert)
			hutangRoutes.GET("/:id", kegiatanqController.FindByID)
			hutangRoutes.PUT("/:id", kegiatanqController.Update)
			hutangRoutes.DELETE("/:id", kegiatanqController.Delete)
			hutangRoutes.PUT("/finish/:id", isfinishController.UpdateIsFinishHutang)
		}
	}
	r.GET("/api/v1/alluser", kegiatanqController.AllUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
