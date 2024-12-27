package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang/glog"
	_ "github.com/joho/godotenv/autoload"

	"task-tracker/internal/database"
)

type Server struct {
	port int

	db database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,

		db: database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
    glog.Info("HTTP Server Listening on: ",server.Addr)
	return server
}
