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

	mainRouter.Serve()
}

func initApp() {
	// Init environment provider
	en.Init()
	// Connect database
	db.Connect(en)
	// Init service
	service.NewUserService()
}

func initRoutes() {
	mainRouter.Get("/health", controller.Health)

	userRouter := mainRouter.RegisterSubRoute("/user")
	userRouter.Post("/signup", controller.Signup)
	userRouter.Post("/login", controller.Login)
	userRouter.Get("/{id}", controller.GetUserById)

	orgRouter := mainRouter.RegisterSubRoute("/org")
	orgRouter.Get("/{id}", controller.GetOrgById)
}
