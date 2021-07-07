package router

import (
	"github.com/gin-gonic/gin"
	logger "github.com/jelena-vlajkov/logger/logger"
	"user-service/http/middleware"
	"user-service/interactor"
)

func NewRouter(handler interactor.AppHandler, logger *logger.Logger) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.AuthMiddleware(logger))

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
	router.GET("/getAllRequestVerificationsForWaiting", handler.GetAllRequestVerificationForWaiting)
	router.GET("/getAllRequestVerifications", handler.GetAllRequestVerification)
	router.POST("/saveNewRequestVerification", handler.SaveNewVerificationRequest)
	router.POST("/approveRequestVerification", handler.ApproveRequestVerification)
	router.POST("/rejectRequestVerification", handler.RejectRequestVerification)
	router.POST("/changePrivacyAndTagging", handler.ChangePrivacyAndTaggin)
	router.GET("/getPrivacyAndTagging", handler.GetPrivacyAndTagging)
	router.POST("/banProfile", handler.BanProfile)
	router.POST("/IsInfluencerAndPrivate", handler.IsInfluencerAndPrivate)
	router.POST("/getSearchInfo", handler.GetProfileInfo)


	return router
}
