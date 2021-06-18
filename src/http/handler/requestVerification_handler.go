package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/dto"
	"user-service/usecase"
)

type requestVerificationHandler struct {
	RequestVerificationUseCase usecase.RequestVerificationUseCase
}



type RequestVerificationHandler interface {
	GetAllRequestVerificationForWaiting(ctx *gin.Context)
	GetAllRequestVerification(ctx *gin.Context)
	SaveNewVerificationRequest(ctx *gin.Context)
	ApproveRequestVerification(ctx *gin.Context)
	RejectRequestVerification(ctx *gin.Context)

}

func (r *requestVerificationHandler) GetAllRequestVerificationForWaiting(ctx *gin.Context) {
	verifications := r.RequestVerificationUseCase.GetAllRequestVerificationForWaiting(ctx)

	ctx.JSON(200, verifications)
}

func (r *requestVerificationHandler) GetAllRequestVerification(ctx *gin.Context) {
	verifications := r.RequestVerificationUseCase.GetAllRequestVerification(ctx)

	ctx.JSON(200, verifications)
}

func (r *requestVerificationHandler) SaveNewVerificationRequest(ctx *gin.Context) {

	var newRequestVerification dto.RequestVerificationDTO

	err := json.NewDecoder(ctx.Request.Body).Decode(&newRequestVerification)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Decoding error")
		ctx.Abort()
		return
	}

	exists, err := r.RequestVerificationUseCase.ExistsRequestForProfile(newRequestVerification.ProfileId, ctx)
	if exists {
		ctx.JSON(400, gin.H{"message" : "Request already sent for this profile"})
		return
	}

	message, err := r.RequestVerificationUseCase.SaveNewRequestVerification(newRequestVerification, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"message" : message})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message" : message})
}

func (r *requestVerificationHandler) ApproveRequestVerification(ctx *gin.Context) {
	var requestToApprove dto.RequestVerificationToChangeStateDTO

	err := json.NewDecoder(ctx.Request.Body).Decode(&requestToApprove)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Decoding error")
		ctx.Abort()
		return
	}

	_, err = r.RequestVerificationUseCase.ApproveRequestVerification(requestToApprove.ID, requestToApprove.ProfileId, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Failed to approve")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message" : "Request is approved"})

}

func (r *requestVerificationHandler) RejectRequestVerification(ctx *gin.Context) {
	var requestToApprove dto.RequestVerificationToChangeStateDTO

	err := json.NewDecoder(ctx.Request.Body).Decode(&requestToApprove)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Decoding error")
		ctx.Abort()
		return
	}

	_, err = r.RequestVerificationUseCase.RejectRequestVerification(requestToApprove.ID, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Failed to approve")
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message" : "Request is rejected"})

}

func NewRequestVerificationHandler(useCase usecase.RequestVerificationUseCase) RequestVerificationHandler {
	return &requestVerificationHandler{RequestVerificationUseCase: useCase}
}