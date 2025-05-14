// package gapi

// import (
// 	"fmt"

// 	"github.com/kirildevops/weather-api/pb"
// 	db "github.com/techschool/simplebank/db/sqlc"
// 	"github.com/techschool/simplebank/token"
// 	"github.com/techschool/simplebank/util"
// )

// // Server serves gRPC requests for our weather service.
// type Server struct {
// 	pb.UnimplementedWeatherAppServer
// 	config     util.Config
// 	store      db.Store
// 	tokenMaker token.Maker
// }

// // NewServer creates a new gRPC server.
// func NewServer(config util.Config, store db.Store) (*Server, error) {
// 	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
// 	if err != nil {
// 		return nil, fmt.Errorf("cannot create token maker: %w", err)
// 	}

// 	server := &Server{
// 		config:     config,
// 		store:      store,
// 		tokenMaker: tokenMaker,
// 	}

// 	return server, nil
// }
