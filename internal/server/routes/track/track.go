package track

import (
	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/RaceSimHub/race-hub-backend/internal/server/routes/list"
	"github.com/RaceSimHub/race-hub-backend/internal/server/routes/template"
	serviceTrack "github.com/RaceSimHub/race-hub-backend/internal/service/track"
	"github.com/RaceSimHub/race-hub-backend/pkg/request"
	"github.com/RaceSimHub/race-hub-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type Track struct {
	serviceTrack serviceTrack.Track
	response     response.Response
}

const (
	trackListTemplate       = "list/list_content"
	trackEditTemplate       = "track/track_edit"
	trackCreateTemplate     = "track/track_create"
	trackFormFieldsTemplate = "track/track_form_fields"
	tracksUrl               = "/admin/tracks"
)

func NewTrack(serviceTrack serviceTrack.Track) *Track {
	return &Track{serviceTrack: serviceTrack, response: response.Response{}}
}

func (t Track) GetList(c *gin.Context) {
	search, offset, limit := request.Request{}.DefaultListParams(c)

	tracks, total, err := t.serviceTrack.GetList(search, offset, limit)
	if err != nil {
		t.response.NewNotification(response.NotificationTypeError, "Erro ao buscar lista de pistas. Erro: "+err.Error()).
			Show(c)
		return
	}

	headers := []string{"ID", "Name", "Country"}

	headerTranslations := map[string]string{
		"ID":      "ID",
		"Name":    "Nome",
		"Country": "País",
	}

	data := list.ListTemplateData[sqlc.SelectListTracksRow]{
		Template:           "admin/tracks",
		Headers:            headers,
		HeaderTranslations: headerTranslations,
		Data:               tracks,
		Total:              int(total),
		GinContext:         c,
	}

	template.Template{}.RenderPage(c, "Lista de Pistas", data, trackListTemplate)
}

func (t Track) Put(c *gin.Context) {
	id, err := request.Request{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	name := c.PostForm("name")
	country := c.PostForm("country")

	if name == "" || country == "" {
		t.response.NewNotification(response.NotificationTypeError, "Campos obrigatórios não preenchidos").
			Show(c)
		return
	}

	err = t.serviceTrack.Update(id, name, country, 1)
	if err != nil {
		t.response.NewNotification(response.NotificationTypeError, "Erro ao fazer a atualização. Erro: "+err.Error()).
			Show(c)
		return
	}

	t.response.NewNotification(response.NotificationTypeSuccess, "Pista atualizada com sucesso").
		WithRedirect(tracksUrl).
		Show(c)
}

func (t Track) Post(c *gin.Context) {
	name := c.PostForm("name")
	country := c.PostForm("country")

	_, err := t.serviceTrack.Create(name, country, 1)
	if err != nil {
		t.response.NewNotification(response.NotificationTypeError, "Erro ao criar pista. Erro: "+err.Error()).
			Show(c)
		return
	}

	t.response.NewNotification(response.NotificationTypeSuccess, "Pista criada com sucesso").
		WithRedirect(tracksUrl).
		Show(c)

}

func (t Track) GetByID(c *gin.Context) {
	id, err := request.Request{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	track, err := t.serviceTrack.GetByID(id)
	if err != nil {
		t.response.NewNotification(response.NotificationTypeError, "Erro ao buscar pista. Erro: "+err.Error()).
			Show(c)
		return
	}

	data := map[string]any{
		"Track": track,
	}

	template.Template{}.RenderPage(c, track.Name, data, trackEditTemplate, trackFormFieldsTemplate)
}

func (t Track) New(c *gin.Context) {
	template.Template{}.RenderPage(c, "Novo Piloto", nil, trackCreateTemplate, trackFormFieldsTemplate)
}

func (t Track) Delete(c *gin.Context) {
	id, err := request.Request{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	err = t.serviceTrack.Delete(id)
	if err != nil {
		t.response.NewNotification(response.NotificationTypeError, "Erro ao deletar pista. Erro: "+err.Error()).
			Show(c)
		return
	}

	t.response.NewNotification(response.NotificationTypeSuccess, "Pista deletada com sucesso").
		WithRedirect(tracksUrl).
		Show(c)
}
