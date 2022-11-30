package handler

import (
	"chatbotbasic/internal/domain/template"
	"chatbotbasic/internal/platform/webapi"
	"chatbotbasic/internal/platform/webapi/request"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

type TemplateHandler struct {
	useCase template.UseCase
}

func NewTemplateHandler(h template.UseCase) TemplateHandler {
	return TemplateHandler{useCase: h}
}

func (h *TemplateHandler) SaveTemplateHandler(c echo.Context) error {
	var r request.TemplateParamRequest

	if err := c.Bind(&r); err != nil {
		log.Error().Str("handler", "SaveNewTemplate").Msg(err.Error())
		return c.JSON(http.StatusInternalServerError, webapi.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  webapi.Err,
		})
	}

	//for _, i := range r.Stories[0].Blocks[0].Contents {
	//	res, _ := json.Marshal(i.Payload)
	//	fmt.Printf("payload id:%s :%s\n", i.ID, string(res))
	//}

	// generated id
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	err = h.useCase.SaveTemplate(id.String(), r.ToEntity())

	if err != nil {
		log.Error().Str("handler", "SaveTemplateHandler").Msg(err.Error())

		return c.JSON(http.StatusInternalServerError, webapi.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  webapi.Err,
		})
	}

	return c.JSON(http.StatusOK, webapi.Response{
		Status:  webapi.Success,
		Code:    http.StatusOK,
		Message: "success created",
		Data: map[string]string{
			"id":    id.String(),
			"title": r.Title,
		},
	})
}

func (h *TemplateHandler) GetTemplateByIdHandler(c echo.Context) error {
	id := c.Param("id")

	result, err := h.useCase.GetTemplateById(id)
	if err != nil {
		log.Error().Str("handler", "GetTemplateByIdHandler").Msg(err.Error())

		return c.JSON(http.StatusNotFound, webapi.Response{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("id %s not found", id),
		})
	}

	return c.JSON(http.StatusOK, webapi.Response{
		Status: webapi.Success,
		Code:   http.StatusOK,
		Data:   result,
	})
}

func (h *TemplateHandler) GetAllTemplateHandler(c echo.Context) error {
	qPage, err := strconv.Atoi(c.QueryParam("page"))
	qOwnerID := c.QueryParam("owner_id")

	if qPage < 1 {
		qPage = 1
	}

	templates, err := h.useCase.GetAllTemplate(qPage, qOwnerID)
	if err != nil {
		log.Error().Str("handler", "GetAllTemplateHandler").Msg(err.Error())
		return c.JSON(http.StatusInternalServerError, webapi.Response{
			Status:  webapi.Err,
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, webapi.Response{
		Status:  webapi.Success,
		Code:    http.StatusOK,
		Message: fmt.Sprintf("page %d", qPage),
		Data:    templates,
	})
}

func (h *TemplateHandler) UpdateTemplateHandler(c echo.Context) error {
	var r request.TemplateParamRequest

	if err := c.Bind(&r); err != nil {
		log.Error().Str("handler", "UpdateTemplateHandler").Msg(err.Error())
		return c.JSON(http.StatusInternalServerError, webapi.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  webapi.Err,
		})
	}

	if r.ID == "" {
		return c.JSON(http.StatusBadRequest, webapi.Response{
			Code:    http.StatusBadRequest,
			Message: "ID is empty",
			Status:  webapi.Err,
		})
	}

	err := h.useCase.UpdateTemplate(r.ToEntity())

	if err != nil {
		log.Error().Str("handler", "UpdateTemplateHandler").Msg(err.Error())

		return c.JSON(http.StatusInternalServerError, webapi.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Status:  webapi.Err,
		})
	}

	return c.JSON(http.StatusOK, webapi.Response{
		Status:  webapi.Success,
		Code:    http.StatusOK,
		Message: "success updated",
		Data: map[string]string{
			"id":    r.ID,
			"title": r.Title,
		},
	})
}

func (h *TemplateHandler) DeleteTemplateHandler(c echo.Context) error {
	id := c.Param("id")

	err := h.useCase.RemoveTemplate(id)
	if err != nil {
		log.Error().Str("handler", "DeleteTemplateHandler").Msg(err.Error())

		return c.JSON(http.StatusNotFound, webapi.Response{
			Code:    http.StatusNotFound,
			Message: fmt.Sprintf("id %s not found", id),
		})
	}

	return c.JSON(http.StatusOK, webapi.Response{
		Status:  webapi.Success,
		Code:    http.StatusOK,
		Message: "success deleted",
	})
}

//func (h *TemplateHandler) GetTemplateCosterByIdHandler(c echo.Context) error {
//	id := c.Param("id")
//
//	result, err := h.useCase.GetTemplateForCoster(id)
//	//result, err := h.useCase.GetTemplateForDev(id)
//	if err != nil {
//		log.Error().Str("handler", "GetTemplateCosterByIdHandler").Msg(err.Error())
//
//		return c.JSON(http.StatusNotFound, webapi.Response{
//			Code:    http.StatusNotFound,
//			Message: fmt.Sprintf("id %s not found", id),
//		})
//	}
//
//	return c.JSON(http.StatusOK, webapi.Response{
//		Status: webapi.Success,
//		Code:   http.StatusOK,
//		Data:   result,
//	})
//}
