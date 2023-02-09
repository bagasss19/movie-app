package main

import (
	"github.com/movie-api/config"
	movieController "github.com/movie-api/controller"
	"github.com/movie-api/docs"

	movieService "github.com/movie-api/service"

	"log"
	"os"

	movierepository "github.com/movie-api/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
	err    error
)

func readEnvironmentFile() {
	//Environment file Load --------------------------------
	err := godotenv.Load(".secret.env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	mode := os.Getenv("MODE")

	switch mode {
	case "development":
		docs.SwaggerInfo.Host = "localhost:8080"
		docs.SwaggerInfo.Schemes = []string{"http"}
	case "staging":
		docs.SwaggerInfo.Host = "talenthub-dev.kemnaker.go.id"
		docs.SwaggerInfo.Schemes = []string{"https"}
	case "production":
		docs.SwaggerInfo.Host = "talenthub.kemnaker.go.id"
		docs.SwaggerInfo.Schemes = []string{"https"}
	}
}

// @title           Movie API
// @version         1.0.0
// @description     This is collection of movie endpoint.
// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey UserAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey AdminAuth
// @in header
// @name Authorization

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	readEnvironmentFile()
	DBConn, err = config.DBConnect()
	if err != nil {
		log.Fatalf("Database connection error: %s", err)
	}

	// docs.SwaggerInfo.BasePath = "/"

	// Init Repository
	movieRepo := movierepository.NewMovieRepository(DBConn)

	// Init Usecase
	movieService := movieService.NewMovieUsecase(movieRepo)

	// Init Controller
	movieController := movieController.NewMovieController(movieService)

	r := gin.Default()
	r.Use(cors.Default())

	// Route
	api := r.Group("/api/v1")
	api.GET("/ping", ping)

	{
		// USER
		// api.GET("/login", userController.Login)
		// user := api.Group("/user")
		// user.Use(middleware.UserAuth())

		{
			// user.GET("/profile", userController.GetUserByID)
		}

		// ADMIN
		// api.POST("/login-admin", adminController.AdminLogin)
		admin := api.Group("/admin")
		// admin.Use(middleware.AdminAuth())

		// Admin
		{
			admin.POST("/movie", movieController.CreateMovie)
		}
	}

	// Swagger
	swaggerGroup := r.Group("/swagger-internal", gin.BasicAuth(gin.Accounts{
		"root": "root",
	}))
	swaggerGroup.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run() // listen and serve on 0.0.0.0:8080
}
