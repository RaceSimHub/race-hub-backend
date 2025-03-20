package driverlink

import (
	"database/sql"

	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/RaceSimHub/race-hub-backend/internal/middleware"
	"github.com/RaceSimHub/race-hub-backend/internal/model"
	"github.com/RaceSimHub/race-hub-backend/internal/server/routes/list"
	"github.com/RaceSimHub/race-hub-backend/internal/server/routes/template"
	serviceDriver "github.com/RaceSimHub/race-hub-backend/internal/service/driver"
	serviceDriverLink "github.com/RaceSimHub/race-hub-backend/internal/service/driver_link"
	"github.com/RaceSimHub/race-hub-backend/pkg/request"
	"github.com/RaceSimHub/race-hub-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type DriverLink struct {
	serviceDriverLink serviceDriverLink.DriverLink
	serviceDriver     serviceDriver.Driver
	response          response.Response
}

const (
	driverLinkListTemplate    = "list/list_content"
	driverLinkPendingTemplate = "driver_link/driver_link_pending"
	driverLinksUrl            = "/driver/link"
	driverLinksPendingUrl     = "/driver/link/pending"
	driverListTemplate        = "driver_link/driver_link_list"
)

func NewDriverLink(serviceDriverLink serviceDriverLink.DriverLink, serviceDriver serviceDriver.Driver) *DriverLink {
	return &DriverLink{serviceDriverLink: serviceDriverLink, serviceDriver: serviceDriver, response: response.Response{}}
}

func (d DriverLink) Pending(c *gin.Context) {
	template.Template{}.RenderPage(c, "Solicitação de Piloto Pendente", nil, driverLinkPendingTemplate)
}

func (d DriverLink) GetDriverLink(c *gin.Context) {
	claims := middleware.RetrieveJwtClaims(c)

	userID := claims["userID"].(float64)

	status, err := d.serviceDriverLink.GetStatusByUserID(int64(userID))
	if err != nil && err != sql.ErrNoRows {
		d.response.NewNotification(response.NotificationTypeError, "Erro ao buscar status do piloto: "+err.Error()).
			Show(c)
		return
	}

	if status == string(model.DriverLinkStatusPending) {
		template.Template{}.RenderPage(c, "Solicitação de Piloto Pendente", nil, driverLinkPendingTemplate)

		return
	}

	search, offset, limit := request.Request{}.DefaultListParams(c)

	drivers, total, err := d.serviceDriver.GetList(search, offset, limit)
	if err != nil {
		d.response.NewNotification(response.NotificationTypeError, "Erro ao buscar lista de pilotos: "+err.Error()).
			Show(c)
		return
	}

	headers := []string{"ID", "Name"}

	headerTranslations := map[string]string{
		"ID":   "ID",
		"Name": "Nome",
	}

	data := list.ListTemplateData[sqlc.SelectListDriversRow]{
		GinContext:         c,
		Title:              "Relacionar Usuário e Piloto",
		Template:           "driver/link",
		Headers:            headers,
		HeaderTranslations: headerTranslations,
		Data:               drivers,
		Total:              int(total),
		ShowPostAction:     true,
		CreateIcon:         "fas fa-link",
	}

	template.Template{}.RenderPage(c, "Relacionar Usuário e Piloto", data, driverListTemplate)
}

func (d DriverLink) GetList(c *gin.Context) {
	search, offset, limit := request.Request{}.DefaultListParams(c)

	driverLinks, total, err := d.serviceDriverLink.GetList(search, offset, limit)
	if err != nil {
		d.response.NewNotification(response.NotificationTypeError, "Erro ao buscar lista de pilotos: "+err.Error()).
			Show(c)
		return
	}

	headers := []string{"ID", "UserName", "DriverName", "Status"}

	headerTranslations := map[string]string{
		"ID":               "ID",
		"UserName":         "Usuário",
		"DriverName":       "Piloto",
		"DriverLinkStatus": "Status",
	}

	data := list.ListTemplateData[sqlc.SelectDriverLinksRow]{
		GinContext:         c,
		Title:              "Relacionar Usuário e Piloto",
		Template:           "admin/driver/link",
		Headers:            headers,
		HeaderTranslations: headerTranslations,
		Data:               driverLinks,
		Total:              int(total),
	}

	template.Template{}.RenderPage(c, "Relacionar Usuário e Piloto", data, driverLinkListTemplate)
}

func (d DriverLink) Post(c *gin.Context) {
	driverID, err := request.Request{}.BindParamInt(c, "driverID", true)
	if err != nil {
		return
	}

	claims := middleware.RetrieveJwtClaims(c)

	userID := claims["userID"].(float64)

	_, err = d.serviceDriverLink.Create(int64(driverID), int64(userID))
	if err != nil {
		d.response.NewNotification(response.NotificationTypeError, "Erro ao relacionar o piloto: "+err.Error()).
			Show(c)
		return
	}

	d.response.NewNotification(response.NotificationTypeSuccess, "Requisição feita com sucesso!").
		WithRedirect(driverLinksPendingUrl).
		Show(c)
}
