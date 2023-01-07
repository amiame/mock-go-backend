package server

import (
	"os"

	"github.com/amiame/mock-go-backend/app/server/handler"
	"github.com/labstack/echo/v4"
)

type Server struct {
	echo *echo.Echo
}

func New(e *echo.Echo) *Server {
	return &Server{
		echo: e,
	}
}

func (s *Server) RegisterHandler(h handler.Handler) {
	s.echo.POST("/twoSum", h.TwoSum)
}

func (s *Server) Start() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	return s.echo.Start(":" + port)
}
