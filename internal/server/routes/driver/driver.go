package driver

import (
	"github.com/RaceSimHub/race-hub-backend/internal/server/model/request"
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

func (d *Driver) Post(c *gin.Context) {
	bodyRequest := request.PostDriver{}
	err := utils.Utils{}.BindJson(c, &bodyRequest)
	if err != nil {
		return
	}

	id, err := d.serviceDriver.Create(bodyRequest.Name, bodyRequest.RaceName, bodyRequest.Email, bodyRequest.Phone)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseCreated(c, int(id))
}

func (d *Driver) Put(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	bodyRequest := request.PutDriver{}
	err = utils.Utils{}.BindJson(c, &bodyRequest)
	if err != nil {
		return
	}

	err = d.serviceDriver.Update(id, bodyRequest.Name, bodyRequest.RaceName, bodyRequest.Email, bodyRequest.Phone)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseNoContent(c)
}

func (d *Driver) Delete(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	err = d.serviceDriver.Delete(id)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseNoContent(c)
}

func (d *Driver) GetList(c *gin.Context) {
	offset, limit, err := utils.Utils{}.GetListParams(c)
	if err != nil {
		return
	}

	drivers, err := d.serviceDriver.GetList(offset, limit)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseOK(c, drivers)
}

func (d *Driver) GetByID(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	driver, err := d.serviceDriver.GetByID(id)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseOK(c, driver)
}
