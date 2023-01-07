package main

import (
	"github.com/amiame/mock-go-backend/app/server"
	"github.com/amiame/mock-go-backend/app/server/domain"
	"github.com/amiame/mock-go-backend/app/server/handler"
	"github.com/amiame/mock-go-backend/app/server/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	d := domain.New()
	u := usecase.New(d)
	h := handler.New(u)
	s := server.New(e)
	s.RegisterHandler(h)

	log.Fatal(s.Start())
}
