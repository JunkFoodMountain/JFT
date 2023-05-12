package app

import "JFT/contract/models"

type NFT[T models.Nft] struct {
}

func (n *NFT[T]) Get(id string) models.Nft {
	return models.Nft{
		Id:   nil,
		Name: nil,
	}
}

func (n *NFT[T]) GetAll() []models.Nft {
	return []models.Nft{}
}
