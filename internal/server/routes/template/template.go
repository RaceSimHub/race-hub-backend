package template

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/gin-gonic/gin"
)

type Template struct {
	db sqlc.Querier
}

func NewTemplate(db sqlc.Querier) *Template {
	return &Template{db: db}
}

func (t Template) Home(c *gin.Context) {
	data := map[string]any{
		"Title": "Home - Race Hub",
	}
	t.Render(c, "index", data)
}

func (Template) Render(c *gin.Context, templateName string, data any) {
	basePath, err := os.Getwd()
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	parentPath := filepath.Dir(basePath)

	templates := template.Must(template.ParseFiles(
		filepath.Join(parentPath, "internal", "template", "base.html"),
		filepath.Join(parentPath, "internal", "template", templateName+".html"),
		filepath.Join(parentPath, "internal", "template", "partial", "header.html"),
		filepath.Join(parentPath, "internal", "template", "partial", "footer.html"),
		filepath.Join(parentPath, "internal", "template", "partial", "sidebar.html"),
	))
	templates.Execute(c.Writer, data)
}
