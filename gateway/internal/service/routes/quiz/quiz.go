package quiz

import (
	"gateway/internal/config"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type Quiz struct {
	address string
	secret  string

	//ssoClient grpc.Client
	log *slog.Logger
}

func New(config *config.Config, log *slog.Logger) *Quiz {
	return &Quiz{
		address: config.Quiz.Address,
		secret:  config.Quiz.Secret,
		log:     log,
	}
}

func (q *Quiz) Status(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "ok")
	if err != nil {
		return err
	}

	return nil
}
