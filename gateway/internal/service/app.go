package service

import (
	"gateway/internal/config"
	"gateway/internal/service/routes"
	"gateway/internal/service/routes/quiz"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"log"
	"log/slog"
)

type App struct {
	eQuiz    *quiz.Quiz
	eGateway *routes.Gateway
	echo     *echo.Echo
}

func New(cfg *config.Config, log *slog.Logger) (*App, error) {
	a := &App{}

	a.eQuiz = quiz.New(cfg, log)
	a.eGateway = routes.New(cfg)
	a.echo = echo.New()

	a.echo.Use(echoMiddleware.Logger())
	a.echo.Use(echoMiddleware.Recover())
	a.echo.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	api := a.echo.Group("/api")
	api.GET("/status", a.eGateway.Status)
	api.GET("/version", a.eGateway.Version)

	//api.POST("/login", a.e.Login)
	//api.POST("/refresh", a.e.Refresh)

	quiz.Routes(api, a.eQuiz)

	return a, nil
}

func (a *App) Run(cfg *config.Config) error {
	log.Println("server running")

	err := a.echo.Start(":" + cfg.Gateway.Port)
	if err != nil {
		log.Fatalf("failed to start http server: %w", err)
		return err
	}

	return nil
}
