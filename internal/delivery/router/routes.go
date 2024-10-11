package router

import (
	"github.com/Arasy41/go-gin-quiz-api/internal/delivery/http"
	"github.com/Arasy41/go-gin-quiz-api/internal/delivery/middleware"
	"github.com/Arasy41/go-gin-quiz-api/internal/domain/repositories"
	"github.com/Arasy41/go-gin-quiz-api/internal/domain/usecases"
	"github.com/Arasy41/go-gin-quiz-api/pkg/constant"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initializes the main router
func InitRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"}

	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	r.Use(cors.New(corsConfig))

	// Middleware
	r.Use(middleware.RequestLogger())

	// Swagger
	r.GET("/api/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Initialize usecases
	userUsecase := usecases.NewUserUsecase(repositories.NewUserRepository(db))
	roleUsecase := usecases.NewRoleUsecase(repositories.NewRoleRepository(db))
	categoryUc := usecases.NewCategoryUsecase(repositories.NewCategoryRepository(db))

	// Initialize handlers
	userHandler := http.NewUserHandler(userUsecase)
	authHandler := http.NewAuthHandler(userUsecase)
	roleHandler := http.NewRoleHandler(roleUsecase)
	categoryHandler := http.NewCategoryHandler(categoryUc)

	// Routes for Admin
	adminRoute := r.Group("/cms", middleware.JWTAuthMiddleware(db, constant.RoleAdmin))
	{
		// User Admin Routes
		adminRoute.GET("/users", userHandler.GetAllUsers)
		adminRoute.GET("/user/:id", userHandler.GetUserByID)
		adminRoute.POST("/user", userHandler.CreateUser)
		adminRoute.PUT("/user/:id", userHandler.UpdateUser)
		adminRoute.DELETE("/user/:id", userHandler.DeleteUser)

		// Role Admin Routes
		adminRoute.GET("/roles", roleHandler.GetAllRoles)
		adminRoute.GET("/role/:id", roleHandler.GetRoleByID)
		adminRoute.POST("/role", roleHandler.CreateRole)
		adminRoute.PUT("/role/:id", roleHandler.UpdateRole)
		adminRoute.DELETE("/role/:id", roleHandler.DeleteRole)

		// Category Admin Routes
		adminRoute.GET("/categories", categoryHandler.GetAllCategories)
		adminRoute.GET("/category/:id", categoryHandler.GetCategoryByID)
		adminRoute.GET("/category/:name", categoryHandler.GetCategoryByName)
		adminRoute.POST("/category", categoryHandler.CreateCategory)
		adminRoute.PUT("/category/:id", categoryHandler.UpdateCategory)
		adminRoute.DELETE("/category/:id", categoryHandler.DeleteCategory)
	}

	// Auth Routes
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/login", authHandler.Login)
		authRoute.POST("/register", authHandler.Register)
		authRoute.PUT("/change-password", middleware.JWTAuthMiddleware(db, constant.AllRoles...), authHandler.ChangePassword)
		authRoute.GET("/user", middleware.JWTAuthMiddleware(db, constant.AllRoles...), authHandler.GetCurrentUser)
	}

	return r
}
