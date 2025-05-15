package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kirildevops/weather-api/db/sqlc"
)

type CreateSubscriptionRequest struct {
	Email     string `json:"email" binding:"required"`
	City      string `json:"city" binding:"required"`
	Frequency string `json:"frequency" binding:"required,oneof=hourly daily"`
}

func (server *Server) subscribe(ctx *gin.Context) {
	var req CreateSubscriptionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.InsertSubscriptionParams{
		Email:     req.Email,
		City:      req.City,
		Frequency: db.FrequencyEnum(req.Frequency),
	}

	subscription, err := server.store.InsertSubscription(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, subscription)

}
