package notification

import (
	modelRequest "github.com/RaceSimHub/race-hub-backend/internal/server/model/request"
	"github.com/RaceSimHub/race-hub-backend/internal/service/notification"
	"github.com/RaceSimHub/race-hub-backend/pkg/request"
	"github.com/RaceSimHub/race-hub-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type Notification struct {
	serviceNotification notification.Notification
}

func NewNotification(serviceNotification notification.Notification) *Notification {
	return &Notification{serviceNotification: serviceNotification}
}

func (n *Notification) Post(c *gin.Context) {
	bodyRequest := modelRequest.PostNotification{}
	err := request.Request{}.BindJson(c, &bodyRequest)
	if err != nil {
		return
	}

	id, err := n.serviceNotification.Create(bodyRequest.Message, bodyRequest.FirstDriver, bodyRequest.SecondDriver, bodyRequest.ThirdDriver, bodyRequest.LicensePoints)
	if err != nil {
		response.Response{}.Error(c, err)
		return
	}

	response.Response{}.Created(c, int(id))
}

func (n *Notification) Put(c *gin.Context) {
	id, err := request.Request{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	bodyRequest := modelRequest.PutNotification{}
	err = request.Request{}.BindJson(c, &bodyRequest)
	if err != nil {
		return
	}

	err = n.serviceNotification.Update(id, bodyRequest.Message, bodyRequest.FirstDriver, bodyRequest.SecondDriver, bodyRequest.ThirdDriver, bodyRequest.LicensePoints)
	if err != nil {
		response.Response{}.Error(c, err)
		return
	}

	response.Response{}.NoContent(c)
}

func (n *Notification) Delete(c *gin.Context) {
	id, err := request.Request{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	err = n.serviceNotification.Delete(id)
	if err != nil {
		response.Response{}.Error(c, err)
		return
	}

	response.Response{}.NoContent(c)
}

func (n *Notification) GetLastMessage(c *gin.Context) {
	message, err := n.serviceNotification.GetLastMessage()
	if err != nil {
		response.Response{}.Error(c, err)
		return
	}

	response.Response{}.OK(c, message)
}

func (n *Notification) GetList(c *gin.Context) {
	offset, limit, err := request.Request{}.GetListParams(c)
	if err != nil {
		return
	}

	list, err := n.serviceNotification.GetList(offset, limit)
	if err != nil {
		response.Response{}.Error(c, err)
		return
	}

	response.Response{}.OK(c, list)
}
