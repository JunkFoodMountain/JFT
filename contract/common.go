package contract

import "github.com/google/uuid"

type Getter[T any] interface {
	GetOne(uuid uuid.UUID) (*T, error)
	GetAll() ([]*T, error)
}
