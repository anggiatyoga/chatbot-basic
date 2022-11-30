package router

import (
	"chatbotbasic/internal/domain/template"
	"chatbotbasic/internal/platform/webapi"
	"chatbotbasic/internal/platform/webapi/handler"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

func RunRouter(modules webapi.AppModule) (*echo.Echo, error) {
	e, err := webapi.NewWebApi()
	if err != nil {
		log.Error().Str("router", "run init").Msg(err.Error())
		return nil, err
	}

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"name":    "chatbot-studio basic api",
			"version": "1.0.0",
		})
	})

	api := e.Group("/api/studio/template")

	templateRoute(api, &modules.TemplateModule)
	return e, nil
}

func templateRoute(eg *echo.Group, module template.UseCase) {
	templateHandler := handler.NewTemplateHandler(module)

	eg.POST("/save", templateHandler.SaveTemplateHandler)
	eg.GET("/get/:id", templateHandler.GetTemplateByIdHandler)
	eg.GET("/get", templateHandler.GetAllTemplateHandler)
	eg.POST("/update", templateHandler.UpdateTemplateHandler)
	eg.GET("/delete/:id", templateHandler.DeleteTemplateHandler)
	//eg.GET("/coster/:id", templateHandler.GetTemplateCosterByIdHandler)
}
