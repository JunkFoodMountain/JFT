package main

import (
	"JFT/internal/transport/rest"
	"JFT/internal/transport/rest/gen"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	handler := rest.JFMHandler{}

	gen.RegisterHandlers(e, gen.NewStrictHandler(&handler, nil))

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:8080")))
}

func setupDB() error {
	db, err := sql.Open("")
	killIf(err)

}

func killIf(err error) {
	if err != nil {
		panic(err)
	}
}
