package main

import (
	"fmt"
	"intikom-test-be/config"
	"intikom-test-be/handler"
	"intikom-test-be/service"
	"intikom-test-be/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	dbUser := "root"
	dbPass := ""
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "intikom-test"
	db := config.DbConnect(dbUser, dbPass, dbHost, dbPort, dbName)
	googleConfig := config.GoogleConfig()

	v1 := router.Group("/intikom-test/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
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
	{
		task.GET("", th.ReadAll)
		task.GET("/:id", th.ReadById)
		task.POST("", th.Create)
		task.PUT("/:id", th.Update)
		task.DELETE("/:id", th.Delete)
	}

	ah := handler.NewAuthHandler(googleConfig)
	auth := v1.Group("auth")
	{
		auth.GET("/google_login", ah.Googlelogin)
		auth.GET("/google_callback", ah.GoogleCallback)
	}

	err := router.Run()
	if err != nil {
		fmt.Errorf("error: %v", err)
	}
}
