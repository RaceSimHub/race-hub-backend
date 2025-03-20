package template

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/RaceSimHub/race-hub-backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Template struct {
	db sqlc.Querier
}

type PageData struct {
	Title       string
	MinimalPage bool
	Data        any
	CurrentPath string
	Role        string
}

func NewTemplate(db sqlc.Querier) *Template {
	return &Template{db: db}
}

func (t Template) Home(c *gin.Context) {
	t.RenderPage(c, "Home - Race Hub", nil, "base/index")
}

func (t Template) RenderPageMinimal(c *gin.Context, title string, content any, templates ...string) {
	t.renderPage(c, title, true, content, templates...)
}

func (t Template) RenderPage(c *gin.Context, title string, content any, templates ...string) {
	t.renderPage(c, title, false, content, templates...)
}

func (t Template) renderPage(c *gin.Context, title string, minimal bool, content any, templates ...string) {
	currentPath := c.Request.URL.Path

	if currentPath == "/" {
		currentPath = "/"
	} else {
		if strings.Contains(currentPath, "/") {
			currentPath = "/" + strings.Split(currentPath[1:], "/")[0]
		}
	}

	var role string
	if !minimal {
		claims := middleware.RetrieveJwtClaims(c)

		role, _ = claims["role"].(string)
	}

	data := PageData{
		Title:       title,
		MinimalPage: minimal,
		Data:        content,
		CurrentPath: currentPath,
		Role:        role,
	}

	t.render(c, data, templates...)
}

func (Template) render(c *gin.Context, data any, templates ...string) {
	basePath, err := os.Getwd()
	if err != nil {
		c.String(500, "Internal Server Error")
		return
	}
	parentPath := filepath.Dir(basePath)

	baseTemplate := template.New("base").Funcs(template.FuncMap{
		"sub": func(a, b int) int { return a - b },
		"add": func(a, b int) int { return a + b },
		"div": func(a, b int) int {
			if b == 0 {
				return 0
			}
			return a / b
		},
		"mod": func(a, b int) int {
			if b == 0 {
				return 0
			}
			return a % b
		},
		"dict": func(values ...any) map[string]any {
			m := make(map[string]any)
			for i := 0; i < len(values); i += 2 {
				m[values[i].(string)] = values[i+1]
			}
			return m
		},
	})

	var templatesPaths []string
	for _, templateName := range templates {
		templatesPaths = append(templatesPaths, filepath.Join(parentPath, "internal", "template", templateName+".html"))
	}

	templatesPaths = append(templatesPaths, filepath.Join(parentPath, "internal", "template", "base", "base.html"))
	templatesPaths = append(templatesPaths, filepath.Join(parentPath, "internal", "template", "base", "header.html"))
	templatesPaths = append(templatesPaths, filepath.Join(parentPath, "internal", "template", "base", "footer.html"))
	templatesPaths = append(templatesPaths, filepath.Join(parentPath, "internal", "template", "base", "sidebar.html"))
	templatesPaths = append(templatesPaths, filepath.Join(parentPath, "internal", "template", "list", "list.html"))

	baseTemplate, err = baseTemplate.ParseFiles(
		templatesPaths...,
	)

	if err != nil {
		log.Printf("Erro ao carregar templates: %v", err)
		c.String(500, "Erro ao carregar templates")
		return
	}

	// Garantir que o template base seja renderizado corretamente
	err = baseTemplate.ExecuteTemplate(c.Writer, "base.html", data)
	if err != nil {
		log.Printf("Erro ao renderizar template: %v", err)
		c.String(500, "Erro ao renderizar template")
	}
}
