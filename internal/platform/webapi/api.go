package webapi

import "github.com/labstack/echo/v4"

func NewWebApi() (*echo.Echo, error) {
	e := echo.New()
	e.HideBanner = true

	return e, nil
}
