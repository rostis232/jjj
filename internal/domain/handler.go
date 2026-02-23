package domain

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type Site struct {
	Domain     string
	CustomText string
}

type Handler struct {
	DB *pgxpool.Pool
}

func (h *Handler) HandleRequest(c echo.Context) error {
	host := c.Request().Host
	// In case host contains port, strip it
	domain := stripPort(host)

	var customText string
	err := h.DB.QueryRow(c.Request().Context(), "SELECT custom_text FROM sites WHERE domain = $1", domain).Scan(&customText)

	if err != nil {
		if err == pgx.ErrNoRows {
			return h.renderPage(c, http.StatusNotFound, "404 Not Found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	return h.renderPage(c, http.StatusOK, customText)
}

func (h *Handler) renderPage(c echo.Context, status int, text string) error {
	tmpl := `<html><body>{{.}}</body></html>`
	t, err := template.New("page").Parse(tmpl)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, text); err != nil {
		return err
	}
	return c.HTML(status, buf.String())
}

func stripPort(host string) string {
	for i := len(host) - 1; i >= 0; i-- {
		if host[i] == ':' {
			return host[:i]
		}
		if host[i] == ']' {
			break
		}
	}
	return host
}
