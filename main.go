package main

import (
	"github.com/joho/godotenv"
	"intikom-test-be/config"
	"intikom-test-be/handler"
	"intikom-test-be/helper"
	"intikom-test-be/service"
	"intikom-test-be/usecase"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	db := config.DbConnect(dbUser, dbPass, dbHost, dbPort, dbName)
	googleConfig := config.GoogleConfig()

	v1 := router.Group("/intikom-test/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.SetCookie("Authorization", "", 0, "", "", false, true)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	us := service.NewUserService(db)
	uu := usecase.NewUserUsecase(us)
	uh := handler.NewUserHandler(uu)
	user := v1.Group("user")
	{
		user.GET("", uh.ReadAll)
		user.GET("/:id", uh.ReadById)
		user.POST("", uh.Create)
		user.PUT("/:id", uh.Update)
		user.DELETE("/:id", uh.Delete)
	}

	ts := service.NewTaskService(db)
	tu := usecase.NewTaskUsecase(ts)
	th := handler.NewTaskHandler(tu)
	task := v1.Group("task")
	task.Use(helper.ValidateToken)
	{
		task.GET("", th.ReadAll)
		task.GET("/:id", th.ReadById)
		task.POST("", th.Create)
		task.PUT("/:id", th.Update)
		task.DELETE("/:id", th.Delete)
	}

	ah := handler.NewAuthHandler(googleConfig, uu)
	auth := v1.Group("auth")
	{
		auth.GET("/google_login", ah.Googlelogin)
		auth.GET("/google_callback", ah.GoogleCallback)
	}

	err = router.Run()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
