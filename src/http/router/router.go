package router

import (
	"github.com/gin-gonic/gin"
	"user-service/http/middleware"
	"user-service/interactor"
)

func NewRouter(handler interactor.AppHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	router.GET("/getById", handler.GetById)
	router.GET("/getProfileInfoByUsername", handler.GetProfileInfoByUsername)
	router.POST("/isPrivate", handler.IsPrivatePostMethod)
	router.GET("/getProfileUsernameImageById", handler.GetProfileUsernameImageById)
	router.GET("/getUserById", handler.GetUserById)
	router.GET("/getUserProfileById", handler.GetUserProfileById)
	router.POST("/saveNewUser", handler.SaveNewUser)
	router.GET("/getAllPublicUsers", handler.GetAllPublicProfiles)
	router.POST("/editUser", handler.EditUser)
	router.GET("/getProfileInfo", handler.GetProfileInfoById)
	router.GET("/searchUser", handler.SearchUser)
	router.GET("/searchPublicUsers", handler.SearchPublicUser)


	return router
}
