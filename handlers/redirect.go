package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"url-shortener-go/models"
)

func ResolveURL(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		shortID := c.Params("shortId")
		var url models.URL
		result := db.First(&url, "short_id = ?", shortID)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).SendString("URL not found")
		}
		url.Clicks++
		db.Save(&url)
		return c.Redirect(url.LongURL, fiber.StatusTemporaryRedirect)
	}
}