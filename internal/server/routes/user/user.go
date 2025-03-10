package user

import (
	modelRequest "github.com/RaceSimHub/race-hub-backend/internal/server/model/request"
	"github.com/RaceSimHub/race-hub-backend/internal/service/user"
	"github.com/RaceSimHub/race-hub-backend/pkg/request"
	"github.com/RaceSimHub/race-hub-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type User struct {
	serviceUser user.User
}

func NewUser(serviceUser user.User) *User {
	return &User{serviceUser: serviceUser}
}

func (u *User) Post(c *gin.Context) {
	bodyRequest := modelRequest.PostUser{}
	err := request.Request{}.BindJson(c, &bodyRequest)
	if err != nil {
		return
	}

	id, err := u.serviceUser.Create(bodyRequest.Email, bodyRequest.Name, bodyRequest.Password)
	if err != nil {
		response.Response{}.Error(c, err)
		return
	}

	response.Response{}.Created(c, int(id))
}
