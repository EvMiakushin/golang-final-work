package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	ginEngine *gin.Engine
}

func NewServer() *Server {
	return &Server{
		ginEngine: gin.Default(),
	}
}

func (s *Server) Start(port string) {
	service := NewService()
	handler := NewHandler(service)
	handler.Register(s.ginEngine)

	go func() {
		s.Router(port)
	}()
	log.Println("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit
}

func (s *Server) Router(port string) {
	srv := &http.Server{
		Addr:           port,
		Handler:        s.ginEngine,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", srv.Addr, err)
	}
}
