package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"nproxy/cmd/web"
	views "nproxy/cmd/web/views"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	s.RegisterApiRoutes(e)
	s.RegisterWebRoutes(e)

	return e
}

func (s *Server) RegisterApiRoutes(e *echo.Echo) http.Handler {
	return e
}

func (s Server) HandleIndex(c echo.Context) error {
	tiles, err := s.queries.GetAllTiles(s.ctx)
	if err != nil {
		return err
	}

	return views.Index(tiles).Render(c.Request().Context(), c.Response())
}

func (s *Server) RegisterWebRoutes(e *echo.Echo) http.Handler {
	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))

	e.GET("/", s.HandleIndex)

	return e
}
