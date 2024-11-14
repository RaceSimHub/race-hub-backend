package user

import (
	"github.com/RaceSimHub/race-hub-backend/internal/server/model/request"
	"github.com/RaceSimHub/race-hub-backend/internal/server/model/response"
	"github.com/RaceSimHub/race-hub-backend/internal/service/user"
	"github.com/RaceSimHub/race-hub-backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type User struct {
	serviceUser user.User
}

func NewUser(serviceUser user.User) *User {
	return &User{serviceUser: serviceUser}
}

func (u *User) PostLogin(c *gin.Context) {
	bodyRequest := request.PostUserLogin{}
	err := utils.Utils{}.BindJson(c, &bodyRequest)
	if err != nil {
		return
	}

	token, err := u.serviceUser.GenerateToken(bodyRequest.Email, bodyRequest.Password)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseOK(c, response.PostUserLogin{Token: token})
}

func (u *User) Post(c *gin.Context) {
	bodyRequest := request.PostUser{}
	err := utils.Utils{}.BindJson(c, &bodyRequest)
	if err != nil {
		return
	}

	id, err := u.serviceUser.Create(bodyRequest.Email, bodyRequest.Name, bodyRequest.Password)
	if err != nil {
		utils.Utils{}.ResponseError(c, err)
		return
	}

	utils.Utils{}.ResponseCreated(c, int(id))
}
