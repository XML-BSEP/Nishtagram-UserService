package router

import (
	"github.com/gin-gonic/gin"
	"user-service/interactor"
)

func NewRouter(handler interactor.AppHandler) *gin.Engine {
	router := gin.Default()

	router.GET("/getById", handler.GetById)
	router.GET("/getProfileInfoByUsername", handler.GetProfileInfoByUsername)

	return router
}
