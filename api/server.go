package api

import (
	db "github.com/AI1411/golang-postgres-k8s/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/accounts/:id", server.getAccount)
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func (server *Server) Start() error {
	return server.router.Run(":8082")
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}