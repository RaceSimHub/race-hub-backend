package list

import (
	"reflect"

	"github.com/RaceSimHub/race-hub-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type ListTemplateData[T any] struct {
	Title      string
	Template   string
	Data       []T
	MapFields  map[string]string
	Total      int
	GinContext *gin.Context
}

func (l ListTemplateData[T]) ColumnsCount() int {
	return len(l.MapFields) + 1
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

func (l ListTemplateData[T]) DeleteURL() string {
	return "/" + l.Template + "/delete"
}

func (l ListTemplateData[T]) Items() (rows []map[string]any) {
	listAny := make([]any, len(l.Data))
	for i, v := range l.Data {
		listAny[i] = v
	}

	for _, item := range listAny {
		val := reflect.ValueOf(item)

		if val.Kind() != reflect.Struct {
			continue
		}

		row := make(map[string]any)

		for field, column := range l.MapFields {
			fieldValue := val.FieldByName(field)
			if fieldValue.IsValid() {
				row[column] = fieldValue.Interface()
			} else {
				row[column] = nil
			}
		}

		idField := val.FieldByName("ID")
		if idField.IsValid() {
			row["ID"] = idField.Interface()
		}

		rows = append(rows, row)
	}

	return rows
}

func (l ListTemplateData[T]) Headers() (headers []string) {
	for _, columnName := range l.MapFields {
		headers = append(headers, columnName)
	}

	return
}

    func (l ListTemplateData[T]) Search() string {
	search, _, _ := utils.Utils{}.DefaultListParams(l.GinContext)

	return search
}

func (l ListTemplateData[T]) Offset() int {
	_, offset, _ := utils.Utils{}.DefaultListParams(l.GinContext)

	return offset
}

func (l ListTemplateData[T]) Limit() int {
	_, _, limit := utils.Utils{}.DefaultListParams(l.GinContext)

	return limit
}
