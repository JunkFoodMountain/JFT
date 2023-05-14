package app

import (
	"JFT/contract"
	"github.com/google/uuid"
)

type App struct {
	nftRepo contract.NFTRepository
}

func New(nftRepo contract.NFTRepository) App {
	return App{
		nftRepo: nftRepo,
	}
}

func (a *App) GetNFTByID(id uuid.UUID) (*contract.Nft, error) {
	return nil, nil
}
