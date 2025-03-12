package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	db "m1thrandir225/cicd2025/db/sqlc"
	"m1thrandir225/cicd2025/token"
	"m1thrandir225/cicd2025/util"
)

type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.TokenMaker
}

func NewServer(store db.Store, config util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	if server.config.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
