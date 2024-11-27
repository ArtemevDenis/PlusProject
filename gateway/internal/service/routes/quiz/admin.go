package quiz

import (
	"gateway/internal/service/routes/quiz/dto"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (q *Quiz) AdminGenerateQuiz(ctx echo.Context) error {
	//id := ctx.Param("id")

	err := ctx.String(http.StatusOK, "ok")
	if err != nil {
		return err
	}

	return nil
}

func (q *Quiz) AdminGetQuiz(ctx echo.Context) error {
	//id := ctx.Param("id")

	err := ctx.String(http.StatusOK, "ok")
	if err != nil {
		return err
	}

	return nil
}

func (q *Quiz) AdminUpdateQuiz(ctx echo.Context) error {
	var quiz dto.Quiz
	if err := ctx.Bind(&quiz); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Здесь можно добавить логику для сохранения опросника в базе данных
	// Например, сохранение в MongoDB

	return ctx.JSON(http.StatusCreated, quiz)

}

func (q *Quiz) AdminDeleteQuiz(ctx echo.Context) error {
	//id := ctx.Param("id")

	err := ctx.String(http.StatusOK, "ok")
	if err != nil {
		return err
	}

	return nil
}
