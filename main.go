package main

import (
	"JFM-NFT/internal/transport/rest"
	"JFM-NFT/internal/transport/rest/gen"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	gen.RegisterHandlers(e, rest.New())

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:8080")))
}
