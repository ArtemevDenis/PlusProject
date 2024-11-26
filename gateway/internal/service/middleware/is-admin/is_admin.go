package is_admin

import (
	"context"
	"errors"
	"gateway/internal/clients/sso/grpc"
	climes "gateway/internal/lib/jwt"
	"gateway/internal/lib/logger"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strings"
	"time"
)

// // extractBearerToken extracts is-admin token from Authorization header.
func extractBearerToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}

	return splitToken[1]
}

var (
	ErrInvalidToken       = errors.New("invalid token")
	ErrFailedIsAdminCheck = errors.New("failed to check if user is admin")
)

// New creates new is-admin middleware.
func New(log *slog.Logger, appSecret string, ssoAddress string) echo.MiddlewareFunc {
	const op = "middleware.is-admin.New"

	log = log.With(slog.String("op", op))

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Получаем JWT-токен из запроса
			tokenStr := extractBearerToken(c.Request())
			if tokenStr == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Provide auth token",
				})
			}

			token, err := jwt.ParseWithClaims(tokenStr, &climes.Climes{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(appSecret), nil
			})
			if err != nil {
				log.Warn("failed to parse token", sl.Err(err))

				// But if token is invalid, we shouldn't handle request
				ctx := context.WithValue(c.Request().Context(), "errorKey", ErrInvalidToken)
				c.SetRequest(c.Request().WithContext(ctx))

				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Bad token format",
				})
			}

			claims := token.Claims.(*climes.Climes)

			log.Info("user authorized", slog.Any("claims", claims))

			// Отправляем запрос для проверки, является ли пользователь админов
			permProvider, err := grpc.New(c.Request().Context(), log, ssoAddress, time.Duration(1000*1000*1000*60*60), 5)
			isAdmin, err := permProvider.IsAdmin(c.Request().Context(), claims.UID)
			if err != nil {
				log.Error("failed to check if user is admin", sl.Err(err))

				ctx := context.WithValue(c.Request().Context(), "errorKey", ErrFailedIsAdminCheck)
				c.SetRequest(c.Request().WithContext(ctx))

				return c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "Bad token format",
				})
			}

			// Полученны данные сохраняем в контекст,
			// откуда его смогут получить следующие хэндлеры.
			ctx := context.WithValue(c.Request().Context(), "uidKey", claims.UID)
			ctx = context.WithValue(ctx, "isAdminKey", isAdmin)
			c.SetRequest(c.Request().WithContext(ctx))

			if !isAdmin {
				return c.JSON(http.StatusMethodNotAllowed, map[string]string{
					"message": "Bad auth token",
				})
			}

			return next(c)
		}
	}
}
func UIDFromContext(ctx context.Context) (int64, bool) {
	uid, ok := ctx.Value("uidKey").(int64)
	return uid, ok
}

func ErrorFromContext(ctx context.Context) (error, bool) {
	err, ok := ctx.Value("errorKey").(error)
	return err, ok
}
