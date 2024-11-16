package track

import (
	"github.com/RaceSimHub/race-hub-backend/internal/server/model/request"
	"github.com/RaceSimHub/race-hub-backend/internal/service/track"
	"github.com/RaceSimHub/race-hub-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type Track struct {
	serviceTrack track.Track
}

func NewTrack(serviceTrack track.Track) *Track {
	return &Track{serviceTrack: serviceTrack}
}

func (n *Track) Post(c *gin.Context) {
	bodyRequest := request.PostTrack{}
	err := utils.Utils{}.BindJson(c, &bodyRequest)
	if err != nil {
		return
	}

	id, err := n.serviceTrack.Create(bodyRequest.Name, bodyRequest.Country)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseCreated(c, int(id))
}

func (n *Track) Put(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	bodyRequest := request.PutTrack{}
	err = utils.Utils{}.BindJson(c, &bodyRequest)
	if err != nil {
		return
	}

	err = n.serviceTrack.Update(id, bodyRequest.Name, bodyRequest.Country)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseNoContent(c)
}

func (n *Track) Delete(c *gin.Context) {
	id, err := utils.Utils{}.BindParamInt(c, "id", true)
	if err != nil {
		return
	}

	err = n.serviceTrack.Delete(id)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseNoContent(c)
}

func (n *Track) GetList(c *gin.Context) {
	offset, limit, err := utils.Utils{}.GetListParams(c)
	if err != nil {
		return
	}

	list, err := n.serviceTrack.GetList(offset, limit)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseOK(c, list)
}
