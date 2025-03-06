package driver

import (
	"database/sql"
	"net/http"

	"github.com/RaceSimHub/race-hub-backend/internal/database/sqlc"
	"github.com/RaceSimHub/race-hub-backend/internal/server/routes/list"
	"github.com/RaceSimHub/race-hub-backend/internal/server/routes/template"
	serviceDriver "github.com/RaceSimHub/race-hub-backend/internal/service/driver"
	"github.com/RaceSimHub/race-hub-backend/pkg/conv"
	"github.com/RaceSimHub/race-hub-backend/pkg/request"
	"github.com/gin-gonic/gin"
)

type Driver struct {
	serviceDriver serviceDriver.Driver
}

const (
	driverListTemplate   = "driver/driver_list"
	driverEditTemplate   = "driver/driver_edit"
	driverCreateTemplate = "driver/driver_create"
	driversUrl           = "/drivers"
)

func NewDriver(serviceDriver serviceDriver.Driver) *Driver {
	return &Driver{serviceDriver: serviceDriver}
}

func (d Driver) GetList(c *gin.Context) {
	search, offset, limit := request.Request{}.DefaultListParams(c)

	drivers, total, err := d.serviceDriver.GetList(search, offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	headers := []string{"ID", "Name"}

	headerTranslations := map[string]string{
		"ID":   "ID",
		"Name": "Nome",
	}

	data := list.ListTemplateData[sqlc.SelectListDriversRow]{
		Title:              "Lista de Pilotos",
		Template:           "drivers",
		Headers:            headers,
		HeaderTranslations: headerTranslations,
		Data:               drivers,
		Total:              int(total),
		GinContext:         c,
	}

	template.Template{}.RenderPage(c, "Lista de Pilotos", false, data, driverListTemplate)
}

func (d Driver) Put(c *gin.Context) {
	id, err := request.Request{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	name := c.PostForm("name")
	email := c.PostForm("email")
	secondaryEmail := c.PostForm("secondary_email")
	phone := c.PostForm("phone")
	secondaryPhone := c.PostForm("secondary_phone")
	license := c.PostForm("license")
	neighborhood := c.PostForm("neighborhood")
	state := c.PostForm("state")
	city := c.PostForm("city")
	cep := c.PostForm("cep")
	address := c.PostForm("address")
	addressNumber := c.PostForm("address_number")
	country := c.PostForm("country")
	team := c.PostForm("team")
	idIracing := c.PostForm("id_iracing")
	idSteam := c.PostForm("id_steam")
	instagram := c.PostForm("instagram")
	facebook := c.PostForm("facebook")
	twitch := c.PostForm("twitch")
	photoURL := c.PostForm("photo_url")
	numberStr := c.PostForm("number")
	secondaryNumberStr := c.PostForm("secondary_number")

	number := conv.StringConv{}.StringToNullInt(numberStr)
	secondaryNumber := conv.StringConv{}.StringToNullInt(secondaryNumberStr)

	if name == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obrigatórios não preenchidos"})
		return
	}

	driver := sqlc.UpdateDriverParams{
		Name:            name,
		Email:           email,
		SecondaryEmail:  sql.NullString{String: secondaryEmail, Valid: secondaryEmail != ""},
		Phone:           sql.NullString{String: phone, Valid: phone != ""},
		SecondaryPhone:  sql.NullString{String: secondaryPhone, Valid: secondaryPhone != ""},
		License:         sql.NullString{String: license, Valid: license != ""},
		State:           sql.NullString{String: state, Valid: state != ""},
		Neighborhood:    sql.NullString{String: neighborhood, Valid: neighborhood != ""},
		City:            sql.NullString{String: city, Valid: city != ""},
		Cep:             sql.NullString{String: cep, Valid: cep != ""},
		Address:         sql.NullString{String: address, Valid: address != ""},
		AddressNumber:   sql.NullString{String: addressNumber, Valid: addressNumber != ""},
		Country:         sql.NullString{String: country, Valid: country != ""},
		Team:            sql.NullString{String: team, Valid: team != ""},
		IDIracing:       sql.NullString{String: idIracing, Valid: idIracing != ""},
		IDSteam:         sql.NullString{String: idSteam, Valid: idSteam != ""},
		Instagram:       sql.NullString{String: instagram, Valid: instagram != ""},
		Facebook:        sql.NullString{String: facebook, Valid: facebook != ""},
		Twitch:          sql.NullString{String: twitch, Valid: twitch != ""},
		PhotoUrl:        sql.NullString{String: photoURL, Valid: photoURL != ""},
		Number:          number,
		SecondaryNumber: secondaryNumber,
	}

	err = d.serviceDriver.Update(id, driver, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.Header("HX-Location", driversUrl)
	c.Status(200)
}

func (d Driver) Post(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	secondaryEmail := c.PostForm("secondary_email")
	phone := c.PostForm("phone")
	secondaryPhone := c.PostForm("secondary_phone")
	license := c.PostForm("license")
	neighborhood := c.PostForm("neighborhood")
	state := c.PostForm("state")
	city := c.PostForm("city")
	cep := c.PostForm("cep")
	address := c.PostForm("address")
	addressNumber := c.PostForm("address_number")
	country := c.PostForm("country")
	team := c.PostForm("team")
	idIracing := c.PostForm("id_iracing")
	idSteam := c.PostForm("id_steam")
	instagram := c.PostForm("instagram")
	facebook := c.PostForm("facebook")
	twitch := c.PostForm("twitch")
	photoURL := c.PostForm("photo_url")
	numberStr := c.PostForm("number")
	secondaryNumberStr := c.PostForm("secondary_number")

	number := conv.StringConv{}.StringToNullInt(numberStr)
	secondaryNumber := conv.StringConv{}.StringToNullInt(secondaryNumberStr)

	if name == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obrigatórios não preenchidos"})
		return
	}

	driver := sqlc.InsertDriverParams{
		Name:            name,
		Email:           email,
		SecondaryEmail:  sql.NullString{String: secondaryEmail, Valid: secondaryEmail != ""},
		Phone:           sql.NullString{String: phone, Valid: phone != ""},
		SecondaryPhone:  sql.NullString{String: secondaryPhone, Valid: secondaryPhone != ""},
		License:         sql.NullString{String: license, Valid: license != ""},
		State:           sql.NullString{String: state, Valid: state != ""},
		City:            sql.NullString{String: city, Valid: city != ""},
		Neighborhood:    sql.NullString{String: neighborhood, Valid: neighborhood != ""},
		Cep:             sql.NullString{String: cep, Valid: cep != ""},
		Address:         sql.NullString{String: address, Valid: address != ""},
		AddressNumber:   sql.NullString{String: addressNumber, Valid: addressNumber != ""},
		Country:         sql.NullString{String: country, Valid: country != ""},
		Team:            sql.NullString{String: team, Valid: team != ""},
		IDIracing:       sql.NullString{String: idIracing, Valid: idIracing != ""},
		IDSteam:         sql.NullString{String: idSteam, Valid: idSteam != ""},
		Instagram:       sql.NullString{String: instagram, Valid: instagram != ""},
		Facebook:        sql.NullString{String: facebook, Valid: facebook != ""},
		Twitch:          sql.NullString{String: twitch, Valid: twitch != ""},
		PhotoUrl:        sql.NullString{String: photoURL, Valid: photoURL != ""},
		Number:          number,
		SecondaryNumber: secondaryNumber,
	}

	_, err := d.serviceDriver.Create(driver, 1)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastrar piloto"})
		return
	}

	c.Header("HX-Location", driversUrl)
	c.Status(200)
}

func (d Driver) GetByID(c *gin.Context) {
	id, err := request.Request{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	driver, err := d.serviceDriver.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	data := map[string]any{
		"Driver": driver,
	}

	template.Template{}.RenderPage(c, driver.Name, false, data, driverEditTemplate)
}

func (d Driver) New(c *gin.Context) {
	template.Template{}.RenderPage(c, "Novo Piloto", false, nil, driverCreateTemplate)
}

func (d Driver) Delete(c *gin.Context) {
	id, err := request.Request{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	err = d.serviceDriver.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	c.Header("HX-Location", driversUrl)
	c.Status(200)
}
