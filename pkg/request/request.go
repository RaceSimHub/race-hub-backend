package request

import (
	"encoding/json"
	"errors"
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/RaceSimHub/race-hub-backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type Request struct{}

func (r Request) ParseBody(ctx *gin.Context, b io.Reader, dto any) error {
	err := json.NewDecoder(b).Decode(dto)

	if err != nil {
		response.Response{}.Error(ctx, errors.New("error.system.json: "+err.Error()))
		return err
	}

	return nil
}

func (r Request) BindParam(ctx *gin.Context, key string, required bool) (value string, err error) {
	value = ctx.Param(key)

	if required && len(strings.TrimSpace(value)) == 0 {
		err = errors.New("error.request.parameter.invalid: Param " + key + " has value invalid")
		response.Response{}.BadRequest(ctx, err)
	}

	return
}

func (r Request) GetPathVariableString(ctx *gin.Context, name string) (value string, err error) {
	value, err = r.BindParam(ctx, name, true)
	if err != nil {
		return
	}

	if value == "" {
		err := errors.New("error.request.variable.path_invalid")
		response.Response{}.Error(ctx, err)

		return "", err
	}

	return value, nil
}

func (r Request) GetPathParamInt(ctx *gin.Context, values url.Values, name string, required bool) (int, error) {
	for parameter, value := range values {
		if strings.EqualFold(parameter, name) {
			v, err := strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				response.Response{}.Error(ctx, errors.New("error.request.parameter.invalid_value: "+name))
				return 0, err
			}

			return int(v), nil
		}
	}

	if required {
		err := errors.New("error.request.parameter.invalid: " + name)
		response.Response{}.Error(ctx, err)
		return 0, err
	}

	return 0, nil
}

func (r Request) BindJson(ctx *gin.Context, obj any) error {
	err := ctx.ShouldBindJSON(obj)
	if err != nil {
		err := errors.New("error.request.invalid: " + err.Error())
		response.Response{}.BadRequest(ctx, err)
	}
	return err
}

func (r Request) GetListParams(ctx *gin.Context) (offset int, limit int, err error) {
	limit, err = r.bindQueryParamInt(ctx, "limit", false)
	if err != nil {
		return
	}

	if limit == 0 {
		limit = 20
	}

	offset, err = r.bindQueryParamInt(ctx, "offset", false)
	if err != nil {
		return
	}

	return
}

func (r Request) bindQueryParamInt(ctx *gin.Context, key string, required bool) (int, error) {
	param, err := r.BindQueryParam(ctx, key, required)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(param)
	if err != nil && required {
		response.Response{}.BadRequest(ctx, err)
		return 0, err
	}

	return value, nil
}

func (r Request) BindQueryParam(ctx *gin.Context, key string, required bool) (value string, err error) {
	value = ctx.Query(key)

	if required && len(strings.TrimSpace(value)) == 0 {
		err = errors.New("error.request.query.param.invalid: " + key)
		response.Response{}.BadRequest(ctx, err)
	}

	return
}

func (r Request) BindParamInt(ctx *gin.Context, key string, required bool) (int, error) {
	param, err := r.BindParam(ctx, key, required)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(param)
	if err != nil && required {
		response.Response{}.BadRequest(ctx, errors.New("error.param.invalid: "+err.Error()))
		return 0, err
	}

	return value, nil
}

func (r Request) DefaultListParams(ctx *gin.Context) (search string, offset, limit int) {
	limitStr := ctx.DefaultQuery("limit", "10")
	limit, _ = strconv.Atoi(limitStr)

	offsetStr := ctx.DefaultQuery("offset", "0")
	offset, _ = strconv.Atoi(offsetStr)

	search = ctx.DefaultQuery("search", "")

	return
}
