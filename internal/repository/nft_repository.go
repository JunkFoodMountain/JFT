package repository

import (
	"JFT/contract"
	"github.com/google/uuid"
)

type NFTRepository[T contract.Nft] struct {
}

func (N NFTRepository[T]) GetOne(uuid uuid.UUID) (*T, error) {
	//TODO implement me
	return &T{}, nil
}

func (N NFTRepository[T]) GetAll() ([]*T, error) {
	//TODO implement me
	return []*T{}, nil
}
