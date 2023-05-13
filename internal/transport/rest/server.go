package rest

import (
	"JFT/internal/transport/rest/gen"
	"context"
)

//go:generate oapi-codegen --config server.config.yaml openapi.yaml

type JFMHandler struct {
}

func (j *JFMHandler) GetNftByID(ctx context.Context, request gen.GetNftByIdRequestObject) (gen.GetNftByIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

// ensure server implements the server interface
var _ gen.StrictServerInterface = (*JFMHandler)(nil)
