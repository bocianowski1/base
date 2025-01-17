package handlers

import (
	"log/slog"

	"github.com/KristianElde/butte/models"
	"github.com/KristianElde/butte/services"
	"github.com/KristianElde/butte/views"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authSrv services.IAuthService
	userSrv services.IUserService
}

func NewAuthHandler(userSrv services.IUserService, authSrv services.IAuthService) *AuthHandler {
	return &AuthHandler{authSrv: authSrv, userSrv: userSrv}
}

func (h *AuthHandler) HandleLogin(c *fiber.Ctx) error {
	req := new(models.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		slog.Error("failed to parse request", "err", err)
		return views.Render(c, "auth/login", fiber.Map{"Error": "Invalid request"})
	}

	token, err := h.authSrv.Login(req.Email, req.Password)
	if err != nil {
		slog.Error("failed to login", "err", err)
		return views.Render(c, "auth/login", fiber.Map{"Error": "Invalid credentials"})
	}

	slog.Error("login success", "token", token)
	// c.Cookie(&fiber.Cookie{
	// 	Name:  "token",
	// 	Value: token,
	// })

	// store in cache or session

	return c.Redirect("/")
}
