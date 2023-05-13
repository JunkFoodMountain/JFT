package rest

import (
	"JFT/contract"
	"JFT/transport/rest/gen"
	"context"
)

//go:generate oapi-codegen --config server.config.yaml openapi.yaml

type JFTHandler struct {
	app contract.App
}

func (j *JFTHandler) GetNftByID(ctx context.Context, request gen.GetNftByIDRequestObject) (gen.GetNftByIDResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func New(app contract.App) {

}

// ensure server implements the server interface
var _ gen.StrictServerInterface = (*JFTHandler)(nil)
