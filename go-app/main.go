package main

import (
	"github.com/caoquy2000/meeting-app/api/controller"
	"github.com/caoquy2000/meeting-app/api/repository"
	"github.com/caoquy2000/meeting-app/api/routes"
	"github.com/caoquy2000/meeting-app/api/service"
	"github.com/caoquy2000/meeting-app/infrastructure"
	"github.com/caoquy2000/meeting-app/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {
	router := infrastructure.NewGinRouter()
	db := infrastructure.NewDatabase()
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)
	postRoute := routes.NewPostRoute(postController, router)
	postRoute.Setup()

	db.DB.AutoMigrate(&models.Post{})
	router.Gin.Run(":8000")
}
