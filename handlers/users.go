package handlers

import (
	"log/slog"

	"github.com/KristianElde/butte/models"
	"github.com/KristianElde/butte/services"
	"github.com/KristianElde/butte/views"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userSrv services.IUserService
}

func NewUserHandler(userSrv services.IUserService) *UserHandler {
	return &UserHandler{userSrv: userSrv}
}

func (h *UserHandler) HandleCreate(c *fiber.Ctx) error {
	req := new(models.CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return views.Render(c, "auth/register", fiber.Map{"Error": "Invalid request"})
	}

	v := validator.New()
	if err := v.Struct(req); err != nil {
		slog.Error("validation error", "err", err)
		return views.Render(c, "auth/register", fiber.Map{"Error": "Invalid request"})
	}

	if err := h.userSrv.Create(req.Email, req.Password, req.Name); err != nil {
		slog.Error("failed to create user", "err", err)
		return views.Render(c, "auth/register", fiber.Map{"Error": "Failed to create user"})
	}

	return c.Redirect("/auth/login")
}

func (h *UserHandler) HandleGetById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userSrv.FindByID(id)
	if err != nil {
		slog.Error("failed to find user", "err", err)
		return views.Render(c, "error", fiber.Map{"Error": "Failed to find user"})
	}

	return views.Render(c, "user", fiber.Map{"User": user})
}
