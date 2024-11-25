package quiz

import (
	"gateway/internal/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Quiz struct {
	address string
	port    string
}

func New(config *config.Config) *Quiz {
	return &Quiz{
		address: config.Quiz.Address,
		port:    config.Quiz.Port,
	}
}

func (q *Quiz) Status(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "ok")
	if err != nil {
		return err
	}

	return nil
}
