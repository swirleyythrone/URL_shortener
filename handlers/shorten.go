package handlers

import (
	"url-shortener-go/models"
	"url-shortener-go/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type shortenRequest struct {
	LongURL string `json:"long_url"`
}

func ShortenURL(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req shortenRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		shortID := utils.GenerateShortID()

		url := models.URL{LongURL: req.LongURL, ShortID: shortID}
		db.Create(&url)

		return c.JSON(fiber.Map{"short_url": c.BaseURL() + "/" + shortID})
	}
}