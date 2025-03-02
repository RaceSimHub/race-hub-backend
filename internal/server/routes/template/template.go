package template

import (
	"log"
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

	tmpl := template.New("base").Funcs(template.FuncMap{
		"sub": func(a, b int) int { return a - b },
		"add": func(a, b int) int { return a + b },
		"div": func(a, b int) int {
			if b == 0 {
				return 0 // Evita divisão por zero
			}
			return a / b
		},
		"mod": func(a, b int) int {
			if b == 0 {
				return 0 // Evita divisão por zero
			}
			return a % b
		},
	})

	// Registrar todos os arquivos de template corretamente
	tmpl, err = tmpl.ParseFiles(
		filepath.Join(parentPath, "internal", "template", "base.html"),
		filepath.Join(parentPath, "internal", "template", templateName+".html"),
		filepath.Join(parentPath, "internal", "template", "partial", "header.html"),
		filepath.Join(parentPath, "internal", "template", "partial", "footer.html"),
		filepath.Join(parentPath, "internal", "template", "partial", "sidebar.html"),
		filepath.Join(parentPath, "internal", "template", "list.html"),
	)

	if err != nil {
		log.Printf("Erro ao carregar templates: %v", err)
		c.String(500, "Erro ao carregar templates")
		return
	}

	// Garantir que o template base seja renderizado corretamente
	err = tmpl.ExecuteTemplate(c.Writer, "base.html", data)
	if err != nil {
		log.Printf("Erro ao renderizar template: %v", err)
		c.String(500, "Erro ao renderizar template")
	}
}
