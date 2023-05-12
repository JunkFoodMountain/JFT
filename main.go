package main

import (
	"JFT/internal/transport/rest"
	"JFT/internal/transport/rest/gen"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	handler := rest.JFMHandler{}

	gen.RegisterHandlers(e, gen.NewStrictHandler(&handler, nil))

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:8080")))
}
