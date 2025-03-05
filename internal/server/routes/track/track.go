package track

import (
	"net/http"

	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/RaceSimHub/race-hub-backend/internal/server/routes/list"
	"github.com/RaceSimHub/race-hub-backend/internal/server/routes/template"
	serviceTrack "github.com/RaceSimHub/race-hub-backend/internal/service/track"
	"github.com/RaceSimHub/race-hub-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type Track struct {
	serviceTrack serviceTrack.Track
}

const (
	trackListTemplate       = "track/track_list"
	trackEditTemplate       = "track/track_edit"
	trackCreateTemplate     = "track/track_create"
	trackFormFieldsTemplate = "track/track_form_fields"
	tracksUrl               = "/tracks"
)

func NewTrack(serviceTrack serviceTrack.Track) *Track {
	return &Track{serviceTrack: serviceTrack}
}

func (d Track) GetList(c *gin.Context) {
	search, offset, limit := utils.Utils{}.DefaultListParams(c)

	tracks, total, err := d.serviceTrack.GetList(search, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	mapFields := map[string]string{
		"ID":      "ID",
		"Name":    "Nome",
		"Country": "País",
	}

	data := list.ListTemplateData[sqlc.SelectListTracksRow]{
		Template:   "tracks",
		MapFields:  mapFields,
		Data:       tracks,
		Total:      int(total),
		GinContext: c,
	}

	template.Template{}.RenderPage(c, "Lista de Pistas", false, data, trackListTemplate)
}

func (d Track) Put(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	name := c.PostForm("name")
	country := c.PostForm("country")

	if name == "" || country == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obrigatórios não preenchidos"})
		return
	}

	err = d.serviceTrack.Update(id, name, country, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.Header("HX-Location", tracksUrl)
	c.Status(200)
}

func (d Track) Post(c *gin.Context) {
	name := c.PostForm("name")
	country := c.PostForm("country")

	_, err := d.serviceTrack.Create(name, country, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastrar pista"})
		return
	}

	c.Header("HX-Location", tracksUrl)
	c.Status(200)
}

func (d Track) GetByID(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	track, err := d.serviceTrack.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	data := map[string]any{
		"Track": track,
	}

	template.Template{}.RenderPage(c, track.Name, false, data, trackEditTemplate, trackFormFieldsTemplate)
}

func (d Track) New(c *gin.Context) {
	template.Template{}.RenderPage(c, "Novo Piloto", false, nil, trackCreateTemplate, trackFormFieldsTemplate)
}

func (d Track) Delete(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	err = d.serviceTrack.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.Header("HX-Location", tracksUrl)
	c.Status(200)
}
