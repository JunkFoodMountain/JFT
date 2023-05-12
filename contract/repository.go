package contract

import "JFT/contract/models"

type NFTRepository interface {
	Getter[models.Nft]
}
