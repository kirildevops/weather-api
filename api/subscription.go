package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	db "github.com/kirildevops/weather-api/db/sqlc"
	"github.com/lib/pq"
)

type CreateSubscriptionRequest struct {
	Email     string `json:"email" binding:"required"`
	City      string `json:"city" binding:"required"`
	Frequency string `json:"frequency" binding:"required,oneof=hourly daily"`
}

func (server *Server) subscribe(ctx *gin.Context) {
	var req CreateSubscriptionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Invalid input")))
			fmt.Println(validationErrors)
			return
		}

	}

	arg := db.InsertSubscriptionParams{
		Email:     req.Email,
		City:      req.City,
		Frequency: db.FrequencyEnum(req.Frequency),
	}

	subscription, err := server.store.InsertSubscription(ctx, arg)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			// duplicate key value violates unique constraint
			if err.Code == "23505" {
				ctx.JSON(http.StatusConflict, errorResponse(errors.New("Email already subscribed")))
				return
			} else if err.Code == "22P02" {
				ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Provide all of ['email', 'city', 'frequency'] Frequency is 'daily' OR 'hourly'")))
				return
			}
			// fmt.Println(err.Code)
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// ctx.JSON(http.StatusOK, subscription)
	fmt.Println(subscription)
	ctx.JSON(http.StatusOK, normalResponse("Subscription successful. Confirmation email sent."))

}

func (server *Server) confirmSubscription(ctx *gin.Context) {
	token := ctx.Param("token")
	if token == "" {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("Invalid token")))
		return
	}
	token = strings.Trim(token, "/")
	uuid_token, err := uuid.Parse(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, normalResponse("Invalid token"))
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

	ctx.JSON(http.StatusOK, normalResponse("Subscription confirmed successfully"))
}

func (server *Server) unsubscribe(ctx *gin.Context) {
	token := ctx.Param("token")
	if token == "" {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("Invalid token")))
		return
	}
	token = strings.Trim(token, "/")
	uuid_token, err := uuid.Parse(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, normalResponse("Invalid token"))
		return
	}

	sub, err := server.store.GetSubscriptionByToken(ctx, uuid_token)
	if err == sql.ErrNoRows {
		ctx.JSON(http.StatusNotFound, errorResponse(errors.New("Token not found")))
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.DeleteSubscriptionParams{
		Email: sub.Email,
		Token: uuid_token,
	}

	if err = server.store.DeleteSubscription(ctx, arg); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, normalResponse("Unsubscribed successfully"))
}
