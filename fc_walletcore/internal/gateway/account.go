package gateway

import "github.com/vinicius-gregorio/fc_walletcore/internal/entity"

type AccountGateway interface {
	FindByID(id string) (*entity.Account, error)
	Save(client *entity.Account) error
	UpdateBalance(account *entity.Account) error
}
