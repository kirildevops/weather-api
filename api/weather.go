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

// type Weather struct {
// 	Current []struct {
// 		Temp string `json:"temp_c"`
// 		Hum  string `json:"humidity"`
// 	} `json:"current"`
// }

func getWeather(ctx *gin.Context) {
	godotenv.Load(".env")
	apiKey := os.Getenv("WEATHER_API_KEY")
	urlWithKey := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=Kyiv", apiKey)
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

	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Printf("Temp %.1f, Hum %d Descr %s\n", w.Current.Temp, w.Current.Hum, w.Current.Condition.Description)
	ctx.JSON(http.StatusOK, w)
}
