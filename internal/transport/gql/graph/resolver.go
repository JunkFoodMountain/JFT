package graph

import "JFT/contract"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	app *contract.App
}

func New(app *contract.App) Resolver {
	return Resolver{app: app}
}
