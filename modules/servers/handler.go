package servers

import (
	_searchlistHttp "github.com/nguitarpb/7-solutions/modules/searchlist/controllers"
	_searchlistRepository "github.com/nguitarpb/7-solutions/modules/searchlist/repositories"
	_searchlistUsecase "github.com/nguitarpb/7-solutions/modules/searchlist/usecases"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	searchlistGroup := s.App.Group("/beef")
	searchlistRepository := _searchlistRepository.NewSearchlistRepository()
	searchlistUsecase := _searchlistUsecase.NewSearchlistUsecase(searchlistRepository)
	_searchlistHttp.NewSearchlistController(searchlistGroup, searchlistUsecase)

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}
