package contract

import (
	"JFT/contract/models"
	"github.com/google/uuid"
)

type App interface {
	GetNFTByID(id uuid.UUID) (models.Nft, error)
}
