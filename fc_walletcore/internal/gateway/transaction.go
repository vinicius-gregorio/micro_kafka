package gateway

import "github.com/vinicius-gregorio/fc_walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
