package contract

import (
	"github.com/google/uuid"
)

type App interface {
	GetNFTByID(id uuid.UUID) (*Nft, error)
}
