package app

import (
	"JFT/contract/models"
	"github.com/google/uuid"
)

type App struct {
}

func (a *App) GetNFTByID(id uuid.UUID) (models.Nft, error) {

}
