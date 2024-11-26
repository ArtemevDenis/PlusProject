package quiz

import (
	"gateway/internal/service/middleware/is-admin"
	"github.com/labstack/echo/v4"
)

func Routes(api *echo.Group, quiz *Quiz) {

	group := api.Group("/quiz")
	group.GET("/status", quiz.Status)

	//admin
	adminGroup := group.Group("/admin")
	adminGroup.Use(is_admin.New(quiz.log, quiz.secret, quiz.address))

	adminGroup.GET("/status", quiz.AdminStatus)
}
