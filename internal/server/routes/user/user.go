package login

import (
	"net/url"

	"github.com/RaceSimHub/race-hub-backend/internal/server/routes/template"
	"github.com/RaceSimHub/race-hub-backend/internal/service/user"
	"github.com/RaceSimHub/race-hub-backend/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type User struct {
	serviceUser user.User
}

func NewUser(serviceUser user.User) *User {

	return &User{
		serviceUser: serviceUser,
	}
}

func (u *User) PostLogin(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	token, err := u.serviceUser.GenerateToken(email, password)
	if err != nil {
		if err.Error() == "error.user.pending" {
			response.Response{}.NewNotification(response.NotificationTypeWarning, "Usuário não confirmado. Verifique seu e-mail").
				WithRedirect("/email-confirm?email=" + url.QueryEscape(email)).
				Show(c)
			return
		}

		response.Response{}.NewNotification(response.NotificationTypeError, "Usuário não encontrado ou senha inválida").
			Show(c)
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "", false, true)
	response.Response{}.NewNotification(response.NotificationTypeSuccess, "Login efetuado com sucesso!").
		WithRedirect("/").
		Show(c)
}

func (u *User) PostLogout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", false, true)
	c.Header("HX-Location", "/login")
}

func (u *User) GetLogin(c *gin.Context) {
	template.Template{}.RenderPageMinimal(c, "Login", nil, "login/login")
}

func (u *User) GetSignUp(c *gin.Context) {
	template.Template{}.RenderPageMinimal(c, "Crie sua conta", nil, "login/sign_up")
}

func (u *User) PostUser(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	if len(password) < 6 {
		response.Response{}.NewNotification(response.NotificationTypeWarning, "A senha deve ter no mínimo 6 caracteres").
			Show(c)
		return
	}

	if password != confirmPassword {
		response.Response{}.NewNotification(response.NotificationTypeError, "As senhas não conferem").
			Show(c)
		return
	}

	_, err := u.serviceUser.Create(email, name, password)
	if err != nil {
		response.Response{}.NewNotification(response.NotificationTypeError, "Erro ao criar conta. Erro: "+err.Error()).
			Show(c)
		return
	}

	c.Header("HX-Location", "/email-confirm?email="+url.QueryEscape(email))
}

func (u *User) GetEmailConfirm(c *gin.Context) {
	data := map[string]string{
		"Email": c.Query("email"),
	}

	template.Template{}.RenderPageMinimal(c, "Confirme seu E-mail", data, "login/email_confirm")

}

func (u *User) PostEmailVerify(c *gin.Context) {
	email := c.PostForm("email")
	token := c.PostForm("token")

	err := u.serviceUser.VerifyEmail(email, token)
	if err != nil {
		logrus.Error(err)
		response.Response{}.NewNotification(response.NotificationTypeError, "Erro ao confirmar e-mail").
			Show(c)
		return
	}

	response.Response{}.NewNotification(response.NotificationTypeSuccess, "E-mail confirmado com sucesso").
		WithRedirect("/login").
		Show(c)
}

func (u *User) GetForgotPassword(c *gin.Context) {
	template.Template{}.RenderPageMinimal(c, "Esqueci minha senha", nil, "login/forgot_password")
}

func (u *User) PostForgotPassword(c *gin.Context) {
	email := c.PostForm("email")

	err := u.serviceUser.ForgotPassword(email)
	if err != nil {
		logrus.Error(err)
		response.Response{}.NewNotification(response.NotificationTypeError, "Erro ao solicitar redefinição de senha").
			Show(c)
		return
	}

	response.Response{}.NewNotification(response.NotificationTypeSuccess, "E-mail de redefinição de senha enviado").
		WithRedirect("/").
		Show(c)
}

func (u *User) GetResetPassword(c *gin.Context) {
	data := map[string]string{
		"Email": c.Query("email"),
		"Token": c.Query("token"),
	}

	template.Template{}.RenderPageMinimal(c, "Redefina sua senha", data, "login/reset_password")
}

func (u *User) PostResetPassword(c *gin.Context) {
	email := c.PostForm("email")
	token := c.PostForm("token")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	if len(password) < 6 {
		response.Response{}.NewNotification(response.NotificationTypeWarning, "A senha deve ter no mínimo 6 caracteres").
			Show(c)
		return
	}

	if password != confirmPassword {
		response.Response{}.NewNotification(response.NotificationTypeError, "As senhas não conferem").
			Show(c)
		return
	}

	err := u.serviceUser.UpdatePassword(email, token, password)
	if err != nil {
		logrus.Error(err)
		response.Response{}.NewNotification(response.NotificationTypeError, "Erro ao redefinir senha").
			Show(c)
		return
	}

	response.Response{}.NewNotification(response.NotificationTypeSuccess, "Senha redefinida com sucesso").
		WithRedirect("/login").
		Show(c)
}

func (u *User) PostResendEmailConfirmation(c *gin.Context) {
	email := c.Query("email")

	err := u.serviceUser.ResendEmailConfirmation(email)
	if err != nil {
		logrus.Error(err)
		response.Response{}.NewNotification(response.NotificationTypeError, "Erro ao reenviar e-mail de confirmação").
			Show(c)
		return
	}

	response.Response{}.NewNotification(response.NotificationTypeSuccess, "E-mail de confirmação reenviado").
		Show(c)
}
