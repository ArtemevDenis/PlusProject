package quiz

import (
	"github.com/labstack/echo/v4"
)

//
//type QuizHandler interface {
//	Status(ctx echo.Context) error
//	AdminGenerateQuiz(ctx echo.Context) error
//	AdminGetQuiz(ctx echo.Context) error
//	AdminUpdateQuiz(ctx echo.Context) error
//	AdminDeleteQuiz(ctx echo.Context) error
//}
//
//type QuizHandler struct {
//	log        *slog.Logger
//	appSecret  string
//	ssoAddress string
//}

func Routes(api *echo.Group, quiz *Quiz) {
	group := api.Group("/quiz")
	group.GET("/status", quiz.Status)

	//admin
	adminGroup := group.Group("/admin")
	adminGroup.Use(quiz.isAdminMiddleware)

	adminGroup.POST("/", quiz.AdminGenerateQuiz)
	adminGroup.GET("/:id", quiz.AdminGetQuiz)
	adminGroup.PUT("/:id", quiz.AdminUpdateQuiz)
	adminGroup.DELETE("/:id", quiz.AdminDeleteQuiz)
}
