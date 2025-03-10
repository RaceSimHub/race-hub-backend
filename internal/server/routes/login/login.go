package login

import (
	"github.com/RaceSimHub/race-hub-backend/internal/server/routes/template"
	"github.com/RaceSimHub/race-hub-backend/internal/service/user"
	"github.com/RaceSimHub/race-hub-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type Login struct {
	serviceUser user.User
}

func NewLogin(serviceUser user.User) *Login {
	return &Login{serviceUser: serviceUser}
}

func (u *Login) PostLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	token, err := u.serviceUser.GenerateToken(email, password)
	if err != nil {
		response.Response{}.ResponseError(c, err)
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "", false, true)
	c.Header("HX-Location", "/")
}

func (u *Login) PostLogout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", false, true)
	c.Header("HX-Location", "/login")
}

func (u *Login) GetLogin(c *gin.Context) {
	template.Template{}.RenderPage(c, "Login", true, nil, "base/login")
}
