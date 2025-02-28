package driver

import (
	"net/http"

	"github.com/RaceSimHub/race-hub-backend/internal/server/routes/template"
	serviceDriver "github.com/RaceSimHub/race-hub-backend/internal/service/driver"
	"github.com/RaceSimHub/race-hub-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type Driver struct {
	serviceDriver serviceDriver.Driver
}

func NewDriver(serviceDriver serviceDriver.Driver) *Driver {
	return &Driver{serviceDriver: serviceDriver}
}

func (d Driver) GetList(c *gin.Context) {
	drivers, err := d.serviceDriver.GetList(0, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	data := map[string]interface{}{
		"Drivers": drivers,
	}
	template.Template{}.Render(c, "driver/driver_list", data)
}

func (d Driver) Put(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	name := c.PostForm("name")
	raceName := c.PostForm("race_name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")

	err = d.serviceDriver.Update(id, name, raceName, email, phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	d.GetList(c)
}

func (d Driver) Post(c *gin.Context) {
	name := c.PostForm("name")
	raceName := c.PostForm("race_name")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	_, err := d.serviceDriver.Create(name, raceName, email, phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	d.GetList(c)
}

func (d Driver) GetByID(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	driver, err := d.serviceDriver.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	data := map[string]interface{}{
		"Title":  driver.Name,
		"Driver": driver,
	}

	template.Template{}.Render(c, "driver/driver_edit", data)
}
