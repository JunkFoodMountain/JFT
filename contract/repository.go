package contract

import (
	"github.com/google/uuid"
)

type NFTRepository interface {
	Getter[Nft]
}

type Getter[T any] interface {
	GetOne(uuid uuid.UUID) (*T, error)
	GetAll() ([]*T, error)
}
