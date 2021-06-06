package router

import (
	"github.com/gin-gonic/gin"
	"user-service/interactor"
)

func NewRouter(handler interactor.AppHandler) *gin.Engine {
	router := gin.Default()

	g := router.Group("/profile")
	//g.Use(middleware.AuthMiddleware())

	//router.GET("/getById", handler.GetById)
	g.GET("/getProfileInfoByUsername", handler.GetProfileInfoByUsername)
	g.GET("/isPrivate", handler.IsPrivate)
	g.GET("/getProfileUsernameImageById", handler.GetProfileUsernameImageById)
	//router.GET("/getUserById", handler.GetUserById)
	g.GET("/getUserProfileById", handler.GetUserProfileById)
	g.POST("/saveNewUser", handler.SaveNewUser)



	return router
}
