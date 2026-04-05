package models

import (
	"time"
	"html/template"
	"io"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Template struct{
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error{
	return t.templates.ExecuteTemplate(w, name, data) 
}

type JWT struct{
	UserID uuid.UUID
	Exp time.Time
	Iat time.Time
	Iss string
}

type User struct {
	ID uuid.UUID
	Name string
	Login string
	Password string
	Email string
	CreatedAt time.Time
}