package notification

import (
	"github.com/RaceSimHub/race-hub-backend/internal/server/model/request"
	"github.com/RaceSimHub/race-hub-backend/internal/service/notification"
	"github.com/RaceSimHub/race-hub-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type Notification struct {
	serviceNotification notification.Notification
}

func NewNotification(serviceNotification notification.Notification) *Notification {
	return &Notification{serviceNotification: serviceNotification}
}

func (n *Notification) Post(c *gin.Context) {
	bodyRequest := request.PostNotification{}
	err := utils.Utils{}.BindJson(c, &bodyRequest)
	if err != nil {
		return
	}

	id, err := n.serviceNotification.Create(bodyRequest.Message, bodyRequest.FirstDriver, bodyRequest.SecondDriver, bodyRequest.ThirdDriver, bodyRequest.LicensePoints)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseCreated(c, int(id))
}

func (n *Notification) Put(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	bodyRequest := request.PutNotification{}
	err = utils.Utils{}.BindJson(c, &bodyRequest)
	if err != nil {
		return
	}

	err = n.serviceNotification.Update(id, bodyRequest.Message, bodyRequest.FirstDriver, bodyRequest.SecondDriver, bodyRequest.ThirdDriver, bodyRequest.LicensePoints)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseNoContent(c)
}

func (n *Notification) Delete(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	err = n.serviceNotification.Delete(id)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseNoContent(c)
}

func (n *Notification) GetLastMessage(c *gin.Context) {
	message, err := n.serviceNotification.GetLastMessage()
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseOK(c, message)
}

func (n *Notification) GetList(c *gin.Context) {
	offset, limit, err := utils.Utils{}.GetListParams(c)
	if err != nil {
		return
	}

	list, err := n.serviceNotification.GetList(offset, limit)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseOK(c, list)
}
