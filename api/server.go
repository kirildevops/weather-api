package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/kirildevops/weather-api/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"content": "You reached the album management site...",
	// 	})
	// })

	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.HTML(http.StatusOK, "templates/form.html", gin.H{
	// 		"action": "/subscribe",
	// 	})
	// })
	apiRouterGroup := router.Group("/api")
	apiRouterGroup.GET("/weather", getWeather)
	apiRouterGroup.POST("/subscribe", server.subscribe)
	apiRouterGroup.GET("/confirm/:token", server.confirmSubscription)
	apiRouterGroup.GET("/unsubscribe/:token", server.unsubscribe)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
