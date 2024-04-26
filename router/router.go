package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	authhandler "main.go/handlers/auth"
	devicehandler "main.go/handlers/devices"
	sensorhandler "main.go/handlers/sensors"
	userhandler "main.go/handlers/users"
	"main.go/middlewares"
	devicerepository "main.go/repositories/devices"
	sensorrepository "main.go/repositories/sensors"
	userrepository "main.go/repositories/users"
	authservice "main.go/services/auth"
	deviceservice "main.go/services/devices"
	sensorservice "main.go/services/sensors"
	userservice "main.go/services/user"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	var (
		//repo
		authRepository   userrepository.UserRepository     = userrepository.NewUserRepository(db)
		userRepository   userrepository.UserRepository     = userrepository.NewUserRepository(db)
		deviceRepository devicerepository.DeviceRepository = devicerepository.NewDeviceRepository(db)
		sensorRepository sensorrepository.SensorRepository = sensorrepository.NewSensorRepository(db)

		//service
		authService   authservice.AuthService     = authservice.NewAuthService(authRepository)
		userService   userservice.UserService     = userservice.NewUserService(userRepository)
		deviceService deviceservice.DeviceService = deviceservice.NewDeviceService(deviceRepository)
		sensorService sensorservice.SensorService = sensorservice.NewSensorService(sensorRepository)

		//handler
		authHandler   authhandler.AuthHandler     = authhandler.NewAuthHandler(authService)
		userHandler   userhandler.UserHandler     = userhandler.NewUserHandler(userService)
		deviceHandler devicehandler.DeviceHandler = devicehandler.NewDeviceHandler(deviceService)
		sensorHandler sensorhandler.SensorHandler = sensorhandler.NewSensorHandler(sensorService)

		//middleware
		authMiddleware = middlewares.AuthMiddleware()
	)

	apiRoutes := r.Group("api")
	{
		authRoutes := apiRoutes.Group("auth")
		{
			authRoutes.POST("/login", authHandler.Login)
		}

		userRoutes := apiRoutes.Group("users")
		{
			userRoutes.POST("", userHandler.CreateUser)
			userRoutes.PUT("/:id", authMiddleware, userHandler.UpdateUser)
			userRoutes.DELETE("/:id", authMiddleware, userHandler.DeleteUserByID)
			userRoutes.GET("/:id", authMiddleware, userHandler.GetUserByID)

		}

		deviceRoutes := apiRoutes.Group("devices")
		{
			deviceRoutes.POST("", authMiddleware, deviceHandler.CreateDevice)
			deviceRoutes.PUT("/:id", authMiddleware, deviceHandler.UpdateDevice)
			deviceRoutes.DELETE("/:id", authMiddleware, deviceHandler.DeleteDevice)
			deviceRoutes.GET("/:id", authMiddleware, deviceHandler.GetDeviceByID)

		}

		sensorRoutes := apiRoutes.Group("sensors")
		{
			sensorRoutes.POST("", authMiddleware, sensorHandler.CreateSensor)
			sensorRoutes.PUT("/:id", authMiddleware, sensorHandler.UpdateSensor)
			sensorRoutes.DELETE("/:id", authMiddleware, sensorHandler.DeleteSensor)
			sensorRoutes.GET("/:id", authMiddleware, sensorHandler.GetSensorByID)

		}

	}
}
