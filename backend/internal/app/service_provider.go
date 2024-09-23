package app

import (
	"github.com/balobas/dbClient"
	"github.com/uxsnap/fresh_market_shop/backend/internal/config"
)

type serviceProvider struct {
	configPG   *config.ConfigPG
	configHTTP *config.ConfigHTTP

	pgClient *dbClient.ClientDB
	
}
