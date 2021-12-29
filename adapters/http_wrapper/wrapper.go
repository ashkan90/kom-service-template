package http_wrapper

import "github.com/labstack/echo/v4"

const (
	RegisterMethod = "GET"
)

type Registrar struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

func RegisterWithEcho(e *echo.Echo, registrar Registrar) {
	if registrar.Method == RegisterMethod {
		e.GET(registrar.Path, registrar.Handler)
	}
}
