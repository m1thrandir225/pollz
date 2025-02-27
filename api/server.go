package api

import (
	"github.com/gin-gonic/gin"
	db "m1thrandir225/cicd2025/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer() (*Server, error) {
	server := &Server{}
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
