package utils

import (
	"encoding/json"
	"errors"

	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Utils struct{}

func (su Utils) ResponseForbidden(ctx *gin.Context) {
	err := errors.New("error.request.forbidden")

	bodyResponse := Exception{}.Make(err.Error())

	su.Response(ctx, http.StatusForbidden, bodyResponse)
}

func (su Utils) ResponseError(ctx *gin.Context, err error) {
	bodyResponse := Exception{}.Make(err.Error())

	statusCode := http.StatusBadRequest

	if strings.HasSuffix(bodyResponse.Key, "not.found") {
		statusCode = http.StatusBadRequest
	}

	su.Response(ctx, statusCode, bodyResponse)
}
func (su Utils) ResponseNotFound(ctx *gin.Context, err error) {
	bodyResponse := Exception{}.Make(err.Error())

	su.Response(ctx, http.StatusNotFound, bodyResponse)
}

func (su Utils) ResponseCreated(ctx *gin.Context, id int) {
	var bodyResponse = Id{Id: id}

	su.Response(ctx, http.StatusCreated, bodyResponse)
}

func (su Utils) ResponseCreatedObj(ctx *gin.Context, bodyResponse interface{}) {
	su.Response(ctx, http.StatusCreated, bodyResponse)
}

func (su Utils) ResponseOK(ctx *gin.Context, bodyResponse interface{}) {
	su.Response(ctx, http.StatusOK, bodyResponse)
}

func (Utils) ResponseNoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (Utils) ResponseBadRequest(ctx *gin.Context, err error) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
}

func (Utils) ResponseListOk(ctx *gin.Context, bodyResponse any, total, limit, offset int) {
	var list List

	list.Pagination = Pagination{Total: total, Limit: limit, Offset: offset}
	list.Data = bodyResponse

	ctx.JSON(http.StatusOK, list)
}

func (su Utils) Response(ctx *gin.Context, statusCode int, bodyResponse interface{}) {
	jsonResponse, err := json.Marshal(bodyResponse)

	if err != nil {
		su.ResponseError(ctx, errors.New("error.system.json: "+err.Error()))
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

func (su Utils) ParseBody(ctx *gin.Context, b io.Reader, dto interface{}) error {
	err := json.NewDecoder(b).Decode(dto)

	if err != nil {
		su.ResponseError(ctx, errors.New("error.system.json: "+err.Error()))
		return err
	}

	return nil
}

func (su Utils) BindParam(ctx *gin.Context, key string, required bool) (value string, err error) {
	value = ctx.Param(key)

	if required && len(strings.TrimSpace(value)) == 0 {
		err = errors.New("error.request.parameter.invalid: Param " + key + " has value invalid")
		su.ResponseBadRequest(ctx, err)
	}

	return
}

func (su Utils) GetPathVariableString(ctx *gin.Context, name string) (value string, err error) {
	value, err = su.BindParam(ctx, name, true)
	if err != nil {
		return
	}

	if value == "" {
		err := errors.New("error.request.variable.path_invalid")
		su.ResponseError(ctx, err)

		return "", err
	}

	return value, nil
}

func (su Utils) GetPathParamInt(ctx *gin.Context, values url.Values, name string, required bool) (int, error) {
	for parameter, value := range values {
		if strings.EqualFold(parameter, name) {
			v, err := strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				su.ResponseError(ctx, errors.New("error.request.parameter.invalid_value: "+name))
				return 0, err
			}

			return int(v), nil
		}
	}

	if required {
		err := errors.New("error.request.parameter.invalid: " + name)
		su.ResponseError(ctx, err)
		return 0, err
	}

	return 0, nil
}

func (su Utils) GetPathParamBoolean(ctx *gin.Context, values url.Values, name string, required bool) (bool, error) {
	for parameter, value := range values {
		if strings.EqualFold(parameter, name) {
			v, err := strconv.ParseBool(value[0])
			if err != nil {
				su.ResponseError(ctx, errors.New("error.request.parameter.invalid_value: "+name))
				return false, err
			}

			return v, nil
		}
	}

	if required {
		err := errors.New("error.request.parameter.invalid: " + name)
		su.ResponseError(ctx, err)
		return false, err
	}

	return false, nil
}

func (su Utils) BindJson(ctx *gin.Context, obj any) error {
	err := ctx.ShouldBindJSON(obj)
	if err != nil {
		err := errors.New("error.request.invalid: " + err.Error())
		su.ResponseBadRequest(ctx, err)
	}
	return err
}

func (su Utils) GetListParams(ctx *gin.Context) (offset int, limit int, err error) {
	limit, err = su.bindQueryParamInt(ctx, "limit", false)
	if err != nil {
		return
	}

	if limit == 0 {
		limit = 20
	}

	offset, err = su.bindQueryParamInt(ctx, "offset", false)
	if err != nil {
		return
	}

	return
}

func (su Utils) bindQueryParamInt(ctx *gin.Context, key string, required bool) (int, error) {
	param, err := su.BindQueryParam(ctx, key, required)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(param)
	if err != nil && required {
		su.ResponseBadRequest(ctx, err)
		return 0, err
	}

	return value, nil
}

func (su Utils) BindQueryParam(ctx *gin.Context, key string, required bool) (value string, err error) {
	value = ctx.Query(key)

	if required && len(strings.TrimSpace(value)) == 0 {
		err = errors.New("error.request.query.param.invalid: " + key)
		su.ResponseBadRequest(ctx, err)
	}

	return
}

func (su Utils) BindParamInt(ctx *gin.Context, key string, required bool) (int, error) {
	param, err := su.BindParam(ctx, key, required)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(param)
	if err != nil && required {
		su.ResponseBadRequest(ctx, errors.New("error.param.invalid: "+err.Error()))
		return 0, err
	}

	return value, nil
}
