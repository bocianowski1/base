package views

import (
	"encoding/json"

	"log"
	"log/slog"
	"strings"

	"github.com/bocianowski1/base/models"
	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx, view string, data fiber.Map) error {
	// add default data
	data["Title"] = makeDocumentTitleFrom(c.Path())

	// add user data if available
	// user, err := unmarshalUser(c)
	// if err == nil && user != nil {
	// 	data["User"] = user
	// }

	slog.Info("Rendering", "path", c.Path(), "view", view, "with data", data)

	return c.Render(view, data, "layouts/main")
}

func makeDocumentTitleFrom(path string) string {
	// remove leading slash
	path = strings.TrimLeft(path, "/")
	// replace all slashes with dashes
	return strings.ReplaceAll(path, "/", "-")
}

func unmarshalUser(c *fiber.Ctx) (*models.User, error) {
	userStr, ok := c.Locals("user").(string)
	if !ok {
		return nil, nil
	}

	var user models.User
	err := json.Unmarshal([]byte(userStr), &user)
	if err != nil {
		log.Println("Failed to unmarshal user", err)
		return nil, err
	}

	return &user, nil
}
