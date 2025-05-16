package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Weather struct {
	Current struct {
		Temp      float32 `json:"temp_c"`
		Hum       int     `json:"humidity"`
		Condition struct {
			Description string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

type WeatherError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func getWeather(ctx *gin.Context) {
	godotenv.Load(".env")
	city := ctx.Query("city")
	if city == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("Provide the 'city' query param")))
		return
	}
	apiKey := os.Getenv("WEATHER_API_KEY")
	urlWithKey := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, city)
	w := Weather{}

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlWithKey, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("error 1")))
		return
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("error 2")))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println(resp.Body)
		we := WeatherError{}
		if err := json.NewDecoder(resp.Body).Decode(&we); err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(errors.New("error 3")))
			return
		}
		// fmt.Println(we.Error.Code)
		if we.Error.Code == 1006 {
			ctx.JSON(http.StatusNotFound, errorResponse(errors.New("City not found")))
			return
		}

	}

	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// fmt.Printf("Temp %.1f, Hum %d Descr %s\n", w.Current.Temp, w.Current.Hum, w.Current.Condition.Description)

	ctx.JSON(http.StatusOK, gin.H{"temperature": w.Current.Temp, "humidity": w.Current.Hum, "description": w.Current.Condition.Description})
}
