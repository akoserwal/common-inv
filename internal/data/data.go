package data

import (
	"common-inv/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewClusterRepo)

// Data .
type Data struct {
	DbFactory *ConnectionFactory
	Config    *conf.Data
}

// NewData .
func NewData(config *conf.Data, logger log.Logger) (*Data, func(), error) {
	dbFactory, cleanup := NewConnectionFactory(config)
	return &Data{DbFactory: dbFactory}, cleanup, nil
}
