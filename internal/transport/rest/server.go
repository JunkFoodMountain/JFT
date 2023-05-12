package rest

import (
	"JFM-NFT/transport/rest/gen"
	"context"
)

type JFMHandler struct {
}

func (J JFMHandler) GetNftById(ctx context.Context, request gen.GetNftByIdRequestObject) (gen.GetNftByIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

// ensure server implements the server interface
var _ gen.StrictServerInterface = (*JFMHandler)(nil)
