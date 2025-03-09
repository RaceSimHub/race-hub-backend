package track

import (
	"net/http"

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
	trackListTemplate       = "track/track_list"
	trackEditTemplate       = "track/track_edit"
	trackCreateTemplate     = "track/track_create"
	trackFormFieldsTemplate = "track/track_form_fields"
	tracksUrl               = "/tracks"
)

func NewTrack(serviceTrack serviceTrack.Track) *Track {
	return &Track{serviceTrack: serviceTrack, response: response.Response{}}
}

func (t Track) GetList(c *gin.Context) {
	search, offset, limit := request.Request{}.DefaultListParams(c)

	tracks, total, err := t.serviceTrack.GetList(search, offset, limit)
	if err != nil {
		t.response.ResponseWithNotification(c, response.NotificationTypeError, "Erro ao buscar lista de pistas. Erro: "+err.Error(), "")
		return
	}

	headers := []string{"ID", "Name", "Country"}

	headerTranslations := map[string]string{
		"ID":      "ID",
		"Name":    "Nome",
		"Country": "País",
	}

	data := list.ListTemplateData[sqlc.SelectListTracksRow]{
		Template:           "tracks",
		Headers:            headers,
		HeaderTranslations: headerTranslations,
		Data:               tracks,
		Total:              int(total),
		GinContext:         c,
	}

	template.Template{}.RenderPage(c, "Lista de Pistas", false, data, trackListTemplate)
}

func (t Track) Put(c *gin.Context) {
	id, err := request.Request{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	name := c.PostForm("name")
	country := c.PostForm("country")

	if name == "" || country == "" {
		t.response.ResponseWithNotification(c, response.NotificationTypeError, "Campos obrigatórios não preenchidos", "")
		return
	}

	err = t.serviceTrack.Update(id, name, country, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	t.response.ResponseWithNotification(c, response.NotificationTypeSuccess, "Pista atualizada com sucesso", tracksUrl)
}

func (t Track) Post(c *gin.Context) {
	name := c.PostForm("name")
	country := c.PostForm("country")

	_, err := t.serviceTrack.Create(name, country, 1)
	if err != nil {
		t.response.ResponseWithNotification(c, response.NotificationTypeError, "Erro ao criar pista. Erro: "+err.Error(), "")
		return
	}

	t.response.ResponseWithNotification(c, response.NotificationTypeSuccess, "Pista criada com sucesso", tracksUrl)
}

func (t Track) GetByID(c *gin.Context) {
	id, err := request.Request{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	track, err := t.serviceTrack.GetByID(id)
	if err != nil {
		t.response.ResponseWithNotification(c, response.NotificationTypeError, "Erro ao buscar pista. Erro: "+err.Error(), "")
		return
	}

	data := map[string]any{
		"Track": track,
	}

	template.Template{}.RenderPage(c, track.Name, false, data, trackEditTemplate, trackFormFieldsTemplate)
}

func (t Track) New(c *gin.Context) {
	template.Template{}.RenderPage(c, "Novo Piloto", false, nil, trackCreateTemplate, trackFormFieldsTemplate)
}

func (t Track) Delete(c *gin.Context) {
	id, err := request.Request{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	err = t.serviceTrack.Delete(id)
	if err != nil {
		t.response.ResponseWithNotification(c, response.NotificationTypeError, "Erro ao deletar pista. Erro: "+err.Error(), "")
		return
	}

	t.response.ResponseWithNotification(c, response.NotificationTypeSuccess, "Pista deletada com sucesso", tracksUrl)
}
