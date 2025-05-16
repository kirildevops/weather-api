package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (server *Server) confirmSubscription(ctx *gin.Context) {
	token := ctx.Param("token")
	if token == "" {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("Confirmation Token Not Found. Use the link from the Email.")))
		return
	}
	token = strings.Trim(token, "/")
	uuid_token, err := uuid.Parse(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	_, err = server.store.GetSubscriptionByToken(ctx, uuid_token)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("Token not found")))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err = server.store.ConfirmSubscription(ctx, uuid_token); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
}
