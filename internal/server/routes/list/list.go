package list

import (
	"reflect"

	"github.com/RaceSimHub/race-hub-backend/internal/localization"
	"github.com/RaceSimHub/race-hub-backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type ListTemplateData[T any] struct {
	Title            string
	Template         string
	Data             []T
	Columns          []string
	Total            int
	LastColumnIcon   string
	GinContext       *gin.Context
	ShowPostAction   bool
	ShowPutAction    bool
	ShowDeleteAction bool
	CreateIcon       string
	EditIcon         string
	DeleteIcon       string
}

func (l ListTemplateData[T]) ColumnsCount() int {
	return len(l.Columns) + 1
}

func (l ListTemplateData[T]) ExtraPage() int {
	if int(l.Total)%l.Limit() != 0 {
		return 1
	}

	return 0
}

func (l ListTemplateData[T]) DefaultURL() string {
	return "/" + l.Template
}

func (l ListTemplateData[T]) NewURL() string {
	return "/" + l.Template + "/new"
}

func (l ListTemplateData[T]) Items() (rows [][]any) {
	listAny := make([]any, len(l.Data))
	for i, v := range l.Data {
		listAny[i] = v
	}

	for _, item := range listAny {
		val := reflect.ValueOf(item)

		if val.Kind() != reflect.Struct {
			continue
		}

		var row []any
		for _, field := range l.Columns {
			fieldValue := val.FieldByName(field)
			if fieldValue.IsValid() {
				row = append(row, fieldValue.Interface())
			} else {
				row = append(row, nil)
			}
		}

		rows = append(rows, row)
	}

	return rows
}

func (l ListTemplateData[T]) Search() string {
	search, _, _ := request.Request{}.DefaultListParams(l.GinContext)

	return search
}

func (l ListTemplateData[T]) Offset() int {
	_, offset, _ := request.Request{}.DefaultListParams(l.GinContext)

	return offset
}

func (l ListTemplateData[T]) Limit() int {
	_, _, limit := request.Request{}.DefaultListParams(l.GinContext)

	return limit
}

func (l ListTemplateData[T]) ActionIcon() string {
	if l.LastColumnIcon != "" {
		return l.LastColumnIcon
	}

	return "fa-solid fa-pen-to-square"
}

func (l ListTemplateData[T]) PutIcon() string {
	if l.EditIcon != "" {
		return l.EditIcon
	}

	return "fa-solid fa-pen-to-square"
}

func (l ListTemplateData[T]) DelIcon() string {
	if l.DeleteIcon != "" {
		return l.DeleteIcon
	}

	return "fa-solid fa-trash"
}

func (l ListTemplateData[T]) PostIcon() string {
	if l.CreateIcon != "" {
		return l.CreateIcon
	}

	return "fa-solid fa-plus"
}

func (l ListTemplateData[T]) ShowGet() bool {
	return !l.ShowPostAction && !l.ShowPutAction && !l.ShowDeleteAction
}

func (l ListTemplateData[T]) HeadersTranslated() []string {
	translations, _ := localization.LoadTranslations("pt")

	headers := make([]string, len(l.Columns))
	for i, header := range l.Columns {
		if translations[header] == "" {
			headers[i] = "[*]" + header
			continue
		}

		headers[i] = translations[header]
	}

	return headers
}
