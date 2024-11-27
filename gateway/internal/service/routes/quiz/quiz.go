package quiz

import (
	"gateway/internal/config"
	is_admin "gateway/internal/service/middleware/is-admin"
	"github.com/labstack/echo/v4"
	"log/slog"
)

type Quiz struct {
	address string
	secret  string

	log *slog.Logger

	isAdminMiddleware echo.MiddlewareFunc
}

func New(config *config.Config, log *slog.Logger) *Quiz {
	return &Quiz{
		address:           config.Quiz.Address,
		secret:            config.Quiz.Secret,
		log:               log,
		isAdminMiddleware: is_admin.New(log, config.Quiz.Secret, config.Quiz.Address),
	}
}
