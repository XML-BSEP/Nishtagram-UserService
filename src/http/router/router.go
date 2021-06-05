package router

import (
	"github.com/gin-gonic/gin"
	"user-service/interactor"
)

func NewRouter(handler interactor.AppHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/getById", handler.GetById)
	router.GET("/getProfileInfoByUsername", handler.GetProfileInfoByUsername)
	router.GET("/isPrivate", handler.IsPrivate)
	router.GET("/getProfileUsernameImageById", handler.GetProfileUsernameImageById)
	router.GET("/getUserById", handler.GetUserById)
	router.GET("/getUserProfileById", handler.GetUserProfileById)



	return router
}
