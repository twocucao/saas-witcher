package data

import (
	"github.com/google/wire"
	"github.com/twocucao/saas-witcher/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO warpped database client
}

// NewData .
func NewData(c *conf.Data) (*Data, error) {
	return &Data{}, nil
}
