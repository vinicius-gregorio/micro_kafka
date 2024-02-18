package gateway

import "github.com/vinicius-gregorio/fc_walletcore/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
