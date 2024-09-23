package app

import (
	"github.com/balobas/dbClient"
	clientAuthService "github.com/uxsnap/fresh_market_shop/backend/internal/client/services/auth"
	"github.com/uxsnap/fresh_market_shop/backend/internal/config"
)

type serviceProvider struct {
	configPG   *config.ConfigPG
	configHTTP *config.ConfigHTTP

	pgClient   *dbClient.ClientDB
	authClient *clientAuthService.AuthClient
	
}
