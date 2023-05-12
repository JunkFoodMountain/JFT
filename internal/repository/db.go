package repository

import (
	"JFT/contract/models"
	"embed"
	"github.com/google/uuid"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

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
