package repository

import (
	"JFT/contract/models"
	"github.com/google/uuid"
)

type NFTRepository[T models.Nft] struct {
}

func (N NFTRepository[T]) GetOne(uuid uuid.UUID) (*T, error) {
	//TODO implement me
	return &T{}, nil
}

func (N NFTRepository[T]) GetAll() ([]*T, error) {
	//TODO implement me
	return []*T{}, nil
}
