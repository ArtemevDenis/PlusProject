package quiz

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (q *Quiz) Status(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "ok")
	if err != nil {
		return err
	}

	return nil
}
