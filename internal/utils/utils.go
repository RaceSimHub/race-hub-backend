package utils

import (
	"encoding/json"
	"errors"
	"github.com/RaceSimHub/race-hub-backend/internal/middleware"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"

	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Utils struct{}

func (u Utils) ResponseForbidden(ctx *gin.Context) {
	err := errors.New("error.request.forbidden")

	bodyResponse := Exception{}.Make(err.Error())

	u.Response(ctx, http.StatusForbidden, bodyResponse)
}

func (u Utils) ResponseError(ctx *gin.Context, err error) {
	bodyResponse := Exception{}.Make(err.Error())

	statusCode := http.StatusBadRequest

	if strings.HasSuffix(bodyResponse.Key, "not.found") {
		statusCode = http.StatusBadRequest
	}

	u.Response(ctx, statusCode, bodyResponse)
}
func (u Utils) ResponseNotFound(ctx *gin.Context, err error) {
	bodyResponse := Exception{}.Make(err.Error())

	u.Response(ctx, http.StatusNotFound, bodyResponse)
}

func (u Utils) ResponseCreated(ctx *gin.Context, id int) {
	var bodyResponse = Id{Id: id}

	u.Response(ctx, http.StatusCreated, bodyResponse)
}

func (u Utils) ResponseCreatedObj(ctx *gin.Context, bodyResponse interface{}) {
	u.Response(ctx, http.StatusCreated, bodyResponse)
}

func (u Utils) ResponseOK(ctx *gin.Context, bodyResponse interface{}) {
	u.Response(ctx, http.StatusOK, bodyResponse)
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

func (u Utils) Response(ctx *gin.Context, statusCode int, bodyResponse interface{}) {
	jsonResponse, err := json.Marshal(bodyResponse)

	if err != nil {
		u.ResponseError(ctx, errors.New("error.system.json: "+err.Error()))
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

func (u Utils) ParseBody(ctx *gin.Context, b io.Reader, dto interface{}) error {
	err := json.NewDecoder(b).Decode(dto)

	if err != nil {
		u.ResponseError(ctx, errors.New("error.system.json: "+err.Error()))
		return err
	}

	return nil
}

func (u Utils) BindParam(ctx *gin.Context, key string, required bool) (value string, err error) {
	value = ctx.Param(key)

	if required && len(strings.TrimSpace(value)) == 0 {
		err = errors.New("error.request.parameter.invalid: Param " + key + " has value invalid")
		u.ResponseBadRequest(ctx, err)
	}

	return
}

func (u Utils) GetPathVariableString(ctx *gin.Context, name string) (value string, err error) {
	value, err = u.BindParam(ctx, name, true)
	if err != nil {
		return
	}

	if value == "" {
		err := errors.New("error.request.variable.path_invalid")
		u.ResponseError(ctx, err)

		return "", err
	}

	return value, nil
}

func (u Utils) GetPathParamInt(ctx *gin.Context, values url.Values, name string, required bool) (int, error) {
	for parameter, value := range values {
		if strings.EqualFold(parameter, name) {
			v, err := strconv.ParseInt(value[0], 10, 64)
			if err != nil {
				u.ResponseError(ctx, errors.New("error.request.parameter.invalid_value: "+name))
				return 0, err
			}

			return int(v), nil
		}
	}

	if required {
		err := errors.New("error.request.parameter.invalid: " + name)
		u.ResponseError(ctx, err)
		return 0, err
	}

	return 0, nil
}

func (u Utils) BindJson(ctx *gin.Context, obj any) error {
	err := ctx.ShouldBindJSON(obj)
	if err != nil {
		err := errors.New("error.request.invalid: " + err.Error())
		u.ResponseBadRequest(ctx, err)
	}
	return err
}

func (u Utils) GetListParams(ctx *gin.Context) (offset int, limit int, err error) {
	limit, err = u.bindQueryParamInt(ctx, "limit", false)
	if err != nil {
		return
	}

	if limit == 0 {
		limit = 20
	}

	offset, err = u.bindQueryParamInt(ctx, "offset", false)
	if err != nil {
		return
	}

	return
}

func (u Utils) bindQueryParamInt(ctx *gin.Context, key string, required bool) (int, error) {
	param, err := u.BindQueryParam(ctx, key, required)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(param)
	if err != nil && required {
		u.ResponseBadRequest(ctx, err)
		return 0, err
	}

	return value, nil
}

func (u Utils) BindQueryParam(ctx *gin.Context, key string, required bool) (value string, err error) {
	value = ctx.Query(key)

	if required && len(strings.TrimSpace(value)) == 0 {
		err = errors.New("error.request.query.param.invalid: " + key)
		u.ResponseBadRequest(ctx, err)
	}

	return
}

func (u Utils) BindParamInt(ctx *gin.Context, key string, required bool) (int, error) {
	param, err := u.BindParam(ctx, key, required)
	if err != nil {
		return 0, err
	}

	value, err := strconv.Atoi(param)
	if err != nil && required {
		u.ResponseBadRequest(ctx, errors.New("error.param.invalid: "+err.Error()))
		return 0, err
	}

	return value, nil
}

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(middleware.JwtSecret)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
