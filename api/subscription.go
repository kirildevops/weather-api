package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
			}
			// fmt.Println("Severity:", err.Severity)
			// fmt.Println("Code:", err.Code)
			// fmt.Println("Message:", err.Message)
			// fmt.Println("Detail:", err.Detail)
			// fmt.Println("Schema:", err.Schema)
			// fmt.Println("Table:", err.Table)
			// fmt.Println("Constraint:", err.Constraint)
			// fmt.Println("File:", err.File)
			// fmt.Println("Line:", err.Line)
			// fmt.Println("Routine:", err.Routine)
			//
			// Produces
			//
			// Severity: ERROR
			// Code: 23505
			// Message: duplicate key value violates unique constraint "subscriptions_email_key"
			// Detail: Key (email)=(your@email.com) already exists.
			// Schema: public
			// Table: subscriptions
			// Constraint: subscriptions_email_key
			// File: nbtinsert.c
			// Line: 563
			// Routine: _bt_check_unique
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// ctx.JSON(http.StatusOK, subscription)
	fmt.Println(subscription)
	ctx.JSON(http.StatusOK, "Subscription successful. Confirmation email sent.")

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
