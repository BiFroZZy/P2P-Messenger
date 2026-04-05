package server

import (
	"context"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"messenger/internal/auth"
	"messenger/internal/config"
	// "messenger/internal/logger"
	l "messenger/internal/logger"
	"messenger/web/handlers"
)
var log = l.InitLogger()

type Server struct{
	echo *echo.Echo
	cfg *config.Config
	ctx context.Context
}

type Template struct{
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error{
	return t.templates.ExecuteTemplate(w, name, data) 
}

func New(cfg *config.Config) *Server{
	e := echo.New()
	ctx := context.Background()
	srv := &Server{
		echo: e,
		cfg: cfg,
		ctx: ctx,
	}
	srv.echo.HideBanner = true
	srv.echo.HidePort = true
	return srv
}

func (s *Server) Middleware(){
	s.echo.Use(middleware.RequestLogger())
	s.echo.Use(middleware.Recover())
}

func (s *Server) SetTemplates(){
	tmpl, err := template.ParseGlob(
		"web/templates/*.html",
	)
	if err != nil{
		log.Err(err).Msg("Error occured while getting HTML templates")
	}
	s.echo.Renderer = &Template{templates: tmpl}
}
func (s *Server) SetRoutes(){
	s.echo.GET("/", func(c echo.Context) error{
		return c.Redirect(http.StatusSeeOther, "/public/auth")
	})
	public := s.echo.Group("public")
	{
		public.POST("/auth/post", auth.Auth)
		public.GET("/auth", handlers.ShowAuth)
	}
	users := s.echo.Group("/users")
	{
		users.GET("/main", handlers.MainHandler)
		users.GET("/messages", handlers.HandlerMessages)
	}
}
func (s *Server) Start() error{
	s.Middleware()
	s.SetRoutes()
	s.SetTemplates()
	// TODO: добавить graceful shutdown
	err := s.echo.Start(s.cfg.ServerPort)
	if err != nil{
		log.Error().Err(err).Msg("Error occured while starting server")
		return err
	}
	s.echo.Logger.Fatal(err)
	return nil
}