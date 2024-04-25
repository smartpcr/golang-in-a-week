//go:build wireinject
// +build wireinject

package inject

import (
	"github.com/google/wire"
	"webapi/pkg/config"
	"webapi/pkg/store"
)

func Initialize(dbConfig *config.DbConfig) (*store.DbStorage, error) {
	wire.Build(store.NewDbStorage)
	return &store.DbStorage{}, nil
}
