package routes

import (
	"gateway/internal/common"
	"gateway/internal/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

//type Service interface {
//	Ping()
//}

type Gateway struct {
	//s Service
	version string
}

func New(config *config.Config) *Gateway {
	return &Gateway{
		//s: s,
		version: config.Gateway.Version,
	}
}

func (e *Gateway) Status(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "ok")
	if err != nil {
		return err
	}

	return nil
}

func (e *Gateway) Version(ctx echo.Context) error {
	version := common.Version{
		Version: e.version,
		Name:    "Gateway",
	}

	err := ctx.JSON(http.StatusOK, version)
	if err != nil {
		return err
	}

	return nil
}
