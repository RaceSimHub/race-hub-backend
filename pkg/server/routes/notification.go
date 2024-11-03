package routes

import (
	"github.com/RaceSimHub/race-hub-backend/pkg/server/model/request"
	"github.com/RaceSimHub/race-hub-backend/pkg/service"
	"github.com/RaceSimHub/race-hub-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type Notification struct{}

func (Notification) Post(c *gin.Context) {
	bodyRequest := request.PostNotification{}
	err := utils.Utils{}.BindJson(c, &bodyRequest)
	if err != nil {
		return
	}

	id, err := service.Notification{}.Create(bodyRequest.Message, bodyRequest.FirstDriver, bodyRequest.SecondDriver, bodyRequest.ThirdDriver, bodyRequest.LicensePoints)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseCreated(c, int(id))
}

func (Notification) Put(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	bodyRequest := request.PutNotification{}
	err = utils.Utils{}.BindJson(c, &bodyRequest)
	if err != nil {
		return
	}

	err = service.Notification{}.Update(id, bodyRequest.Message, bodyRequest.FirstDriver, bodyRequest.SecondDriver, bodyRequest.ThirdDriver, bodyRequest.LicensePoints)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseNoContent(c)
}

func (Notification) Delete(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	err = service.Notification{}.Delete(id)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseNoContent(c)
}

func (Notification) GetLastMessage(c *gin.Context) {
	message, err := service.Notification{}.GetLastMessage()
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseOK(c, message)
}

func (Notification) GetList(c *gin.Context) {
	offset, limit, err := utils.Utils{}.GetListParams(c)
	if err != nil {
		return
	}

	list, err := service.Notification{}.GetList(offset, limit)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseOK(c, list)
}
