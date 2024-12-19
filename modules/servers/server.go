package servers

import (
	"log"

	"github.com/nguitarpb/7-solutions/configs"
	"github.com/nguitarpb/7-solutions/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	App        *fiber.App
	Cfg        *configs.Configs
}

func NewServer(cfg *configs.Configs) *Server {
	return &Server{
		App:        fiber.New(),
		Cfg:        cfg,
	}
}

func (s *Server) Start() {
	s.App.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,C_U,atr,rtk",
        AllowOrigins:     "http://localhost:3000",
        AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowCredentials: true,
	}))


	if err := s.MapHandlers(); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	fiberConnURL, err := utils.ConnectionUrlBuilder("fiber", s.Cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	host := s.Cfg.App.Host
	port := s.Cfg.App.Port
	log.Printf("server has been started on %s:%s ⚡", host, port)

	if err := s.App.Listen(fiberConnURL); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
}
