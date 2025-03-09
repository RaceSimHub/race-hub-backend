package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct{}

func (r Response) ResponseWithNotification(ctx *gin.Context, notificationType NotificationType, message string, redirect string) {
	notification := Notification{
		Message:  message,
		Type:     notificationType,
		Redirect: redirect,
	}

	ctx.JSON(200, notification)
}

func (r Response) ResponseForbidden(ctx *gin.Context) {
	err := errors.New("error.request.forbidden")

	bodyResponse := Exception{}.Make(err.Error())

	r.Response(ctx, http.StatusForbidden, bodyResponse)
}

func (r Response) ResponseError(ctx *gin.Context, err error) {
	bodyResponse := Exception{}.Make(err.Error())

	statusCode := http.StatusBadRequest

	if strings.HasSuffix(bodyResponse.Key, "not.found") {
		statusCode = http.StatusBadRequest
	}

	r.Response(ctx, statusCode, bodyResponse)
}
func (r Response) ResponseNotFound(ctx *gin.Context, err error) {
	bodyResponse := Exception{}.Make(err.Error())

	r.Response(ctx, http.StatusNotFound, bodyResponse)
}

func (r Response) ResponseCreated(ctx *gin.Context, id int) {
	var bodyResponse = Id{Id: id}

	r.Response(ctx, http.StatusCreated, bodyResponse)
}

func (r Response) ResponseCreatedObj(ctx *gin.Context, bodyResponse any) {
	r.Response(ctx, http.StatusCreated, bodyResponse)
}

func (r Response) ResponseOK(ctx *gin.Context, bodyResponse any) {
	r.Response(ctx, http.StatusOK, bodyResponse)
}

func (Response) ResponseNoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (Response) ResponseBadRequest(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
}

func (Response) ResponseListOk(ctx *gin.Context, bodyResponse any, total, limit, offset int) {
	var list List

	list.Pagination = Pagination{Total: total, Limit: limit, Offset: offset}
	list.Data = bodyResponse

	ctx.JSON(http.StatusOK, list)
}

func (r Response) Response(ctx *gin.Context, statusCode int, bodyResponse any) {
	jsonResponse, err := json.Marshal(bodyResponse)

	if err != nil {
		r.ResponseError(ctx, errors.New("error.system.json: "+err.Error()))
		return
	}

	ctx.Header("Content-Type", "application/json")
	ctx.Status(statusCode)

	bodyResponse = string(jsonResponse)
	if bodyResponse == "null" {
		return
	}

	_, _ = fmt.Fprint(ctx.Writer, bodyResponse)
}
