package contract

import "JFT/contract/models"

type App interface {
	Getter[models.Nft]
}
