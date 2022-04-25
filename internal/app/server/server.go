package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/paul-ss/pgram-backend/internal/pkg/config"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	srv *http.Server
	log *os.File
}

func NewServer() *Server {
	conf := config.C().Server

	switch conf.Mode {
	case gin.ReleaseMode:
		gin.SetMode(gin.ReleaseMode)
	case gin.DebugMode:
	default:
		log.Error("nit existing gin mode provided")
	}

	var file *os.File
	if !conf.StdoutLog {
		file, err := os.OpenFile(conf.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err.Error())
		}

		gin.DefaultWriter = file
	}

	r := gin.Default()
	addRoutes(r)

	return &Server{
		log: file,
		srv: &http.Server{
			Addr:    conf.Address,
			Handler: r,
		},
	}
}

func addRoutes(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}

func (s *Server) Run() {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown:", err)
	}

	log.Info("Server exiting")

	if s.log != nil {
		if err := s.log.Close(); err != nil {
			log.Error(err.Error())
		}
	}
	log.Info("Server cleanup done")
}
