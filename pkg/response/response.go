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

func (r Response) NewNotification(notificationType NotificationType, message string) *notification {
	return &notification{
		Message: message,
		Type:    notificationType,
	}
}

func (n *notification) WithRedirect(redirect string) *notification {
	n.Redirect = redirect

	return n
}

func (n *notification) Show(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, n)
}

func (r Response) Forbidden(ctx *gin.Context) {
	err := errors.New("error.request.forbidden")

	bodyResponse := Exception{}.Make(err.Error())

	r.response(ctx, http.StatusForbidden, bodyResponse)
}

func (r Response) Error(ctx *gin.Context, err error) {
	bodyResponse := Exception{}.Make(err.Error())

	statusCode := http.StatusBadRequest

	if strings.HasSuffix(bodyResponse.Key, "not.found") {
		statusCode = http.StatusBadRequest
	}

	r.response(ctx, statusCode, bodyResponse)
}
func (r Response) NotFound(ctx *gin.Context, err error) {
	bodyResponse := Exception{}.Make(err.Error())

	r.response(ctx, http.StatusNotFound, bodyResponse)
}

func (r Response) Created(ctx *gin.Context, id int) {
	var bodyResponse = Id{Id: id}

	r.response(ctx, http.StatusCreated, bodyResponse)
}

func (r Response) CreatedObj(ctx *gin.Context, bodyResponse any) {
	r.response(ctx, http.StatusCreated, bodyResponse)
}

func (r Response) OK(ctx *gin.Context, bodyResponse any) {
	r.response(ctx, http.StatusOK, bodyResponse)
}

func (Response) NoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (Response) BadRequest(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
}

func (Response) ListOk(ctx *gin.Context, bodyResponse any, total, limit, offset int) {
	var list List

	list.Pagination = Pagination{Total: total, Limit: limit, Offset: offset}
	list.Data = bodyResponse

	ctx.JSON(http.StatusOK, list)
}

func (r Response) response(ctx *gin.Context, statusCode int, bodyResponse any) {
	jsonResponse, err := json.Marshal(bodyResponse)

	if err != nil {
		r.Error(ctx, errors.New("error.system.json: "+err.Error()))
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
