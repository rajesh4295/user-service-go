package main

import (
	"github.com/rajesh4295/user-service-go/controller"
	"github.com/rajesh4295/user-service-go/database"
	"github.com/rajesh4295/user-service-go/env"
	"github.com/rajesh4295/user-service-go/http"
	"github.com/rajesh4295/user-service-go/service"
)

var (
	en         env.Provider      = env.NewEnv()
	db         database.Provider = database.NewPG()
	mainRouter http.Router       = http.NewMuxRouter()
)

func main() {
	initApp()
	initRoutes()

	mainRouter.Serve(en)
}

func initApp() {
	en.Init()
	db.Connect(en)

	service.NewUserService()
}

func initRoutes() {
	mainRouter.Get("/health", controller.Health)

	userRouter := mainRouter.RegisterSubRoute("/user")
	userRouter.Post("/signup", controller.Signup)
	userRouter.Get("/{id}", controller.GetUserById)
}
