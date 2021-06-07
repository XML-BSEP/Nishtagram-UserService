package router

import (
	"github.com/gin-gonic/gin"
	"user-service/http/middleware"
	"user-service/interactor"
)

func NewRouter(handler interactor.AppHandler) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(middleware.CORSMiddleware())

	router.GET("/getById", handler.GetById)
	router.GET("/getProfileInfoByUsername", handler.GetProfileInfoByUsername)
	router.GET("/isPrivate", handler.IsPrivate)
	router.GET("/getProfileUsernameImageById", handler.GetProfileUsernameImageById)
	router.GET("/getUserById", handler.GetUserById)
	router.GET("/getUserProfileById", handler.GetUserProfileById)
	router.POST("/saveNewUser", handler.SaveNewUser)
	router.GET("/getAllPublicUsers", handler.GetAllPublicProfiles)
	router.POST("/editUser", handler.EditUser)


	return router
}
