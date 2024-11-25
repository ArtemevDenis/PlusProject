package quiz

import "github.com/labstack/echo/v4"

func Routes(api *echo.Group, quiz *Quiz) {
	group := api.Group("/quiz")
	//clients endpoints
	group.GET("/status", quiz.Status)

	//admin
	adminGroup := group.Group("/admin")
	adminGroup.GET("/status", quiz.AdminStatus)
}
