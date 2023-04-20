package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/yellowpuki/simple-bank/db/sqlc"
)

// Server serves HTTP routes for out banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponce(err error) gin.H {
	return gin.H{"error": err.Error()}
}
