package contract

import "JFM-NFT/contract/models"

type NFTRepository interface {
	Getter[models.Nft]
}
