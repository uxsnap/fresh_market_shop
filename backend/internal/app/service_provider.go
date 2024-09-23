package app

import (
	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	clientAuthService "github.com/uxsnap/fresh_market_shop/backend/internal/client/services/auth"
	"github.com/uxsnap/fresh_market_shop/backend/internal/config"
)

type serviceProvider struct {
	configPG   *config.ConfigPG
	configHTTP *config.ConfigHTTP

	pgClient   *DBclient.ClientDB
	authClient *clientAuthService.AuthClient
}
