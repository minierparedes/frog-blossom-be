package api

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for CMS service
type Server struct {
	router *gin.Engine
	addr   string
	db     *sql.DB
}

// NewServer creates a new HTTP server and sets up routing
func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		addr: addr,
		db:   db,
	}
}

// Start runs the HTTP server on specific address
func (s *Server) Start() error {
	router := gin.Default()
	subrouter := router.Group("/api/v1")

	subrouter.POST("users", s.createUsers)

	s.router = router

	log.Println("Listening on", s.addr)

	return s.router.Run(s.addr)
}

func errorRosponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
