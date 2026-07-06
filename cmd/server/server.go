package server

import (
	"context"
	"io"
	"errors"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"messenger/internal/auth"
	"messenger/internal/config"

	l "messenger/internal/logger"
	"messenger/web/handlers"
)
var logger = l.InitLogger()

type Server struct{
	echo *echo.Echo
	cfg *config.Config
	ctx context.Context
	templates *template.Template
}

func (s *Server) Render(w io.Writer, name string, data interface{}, c echo.Context) error{
	return s.templates.ExecuteTemplate(w, name, data) 
}

func New(cfg *config.Config) (*Server, error){
	e := echo.New()
	ctx := context.Background()
	tmpl, err := template.ParseGlob("web/templates/*.html",)
	if err != nil{
		logger.Err(err).Msg("Error occured while getting HTML templates")
		return nil, err
	}
	e.Renderer=&Server{templates: tmpl}
	srv := &Server{
		echo: e,
		cfg: cfg,
		ctx: ctx,
		templates: tmpl,
	}
	srv.echo.HideBanner = true
	srv.echo.HidePort = true
	return srv, nil
}

func (s *Server) Middleware(){
	s.echo.Use(middleware.RequestLogger())
	s.echo.Use(middleware.Recover())
}

func (s *Server) SetRoutes(){
	s.echo.GET("/", func(c echo.Context) error{
		return c.Redirect(http.StatusSeeOther, "/public/auth")
	})
	public := s.echo.Group("public")
	{
		public.POST("/auth/post", auth.AuthUser)
		public.GET("/auth", handlers.ShowAuth)
		public.POST("/register/post", auth.RegUser)
		public.GET("/register", handlers.ShowReg)
	}
	users := s.echo.Group("/users")
	users.Use(auth.AuthMiddleware) // Здесь че-то не то
	{	
		users.GET("/main", handlers.MainHandler)
		users.GET("/messages", handlers.ShowMessages)
	}
}
func (s *Server) Start() error{
	s.Middleware()
	s.SetRoutes()
	serverErr := make(chan error, 1)
	go func(){
		if err := s.echo.Start(s.cfg.ServerPort); !errors.Is(err, http.ErrServerClosed){
			logger.Error().Err(err).Msg("Error occured while starting server")
			serverErr<-err
		}
		close(serverErr)
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case err := <-serverErr:
		return err
	case <-quit:
		logger.Info().Msg("Shutdown signal recieved")
	case <- s.ctx.Done():
		logger.Info().Msg("Context canceled")
	}
	shutdownCtx, cancel := context.WithTimeout(context.Background(), s.cfg.ShutdownTimeout)
	defer cancel()
	if err := s.echo.Shutdown(shutdownCtx); err != nil{
		if serverClose := s.echo.Close(); serverClose != nil{
			return errors.Join(err, serverClose)
		}
	}
	return nil
}