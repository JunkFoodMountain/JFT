package contract

import (
	"JFT/contract/models"
	"github.com/google/uuid"
)

type NFTRepository interface {
	Getter[models.Nft]
}

type Getter[T any] interface {
	GetOne(uuid uuid.UUID) (*T, error)
	GetAll() ([]*T, error)
}
