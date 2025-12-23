package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) ScrapCards(c echo.Context) error {
	go func() {
		// create a context background or similar if needed, or just run
		// ideally we should have logging
		_ = h.Card.ScrapCards()
	}()
	return responseJSON(c, "Scraping started in background")
}
