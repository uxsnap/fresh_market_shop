package app

import (
	"context"
	"log"

	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	"github.com/uxsnap/fresh_market_shop/backend/internal/client/database/pg"
	clientAuthService "github.com/uxsnap/fresh_market_shop/backend/internal/client/services/auth"
	"github.com/uxsnap/fresh_market_shop/backend/internal/config"
	deliveryHttp "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http"
	repositoryCategories "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/categories"
	repositoryProducts "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/products"
	ucCategories "github.com/uxsnap/fresh_market_shop/backend/internal/usecase/categories"
	ucProducts "github.com/uxsnap/fresh_market_shop/backend/internal/usecase/products"
)

type serviceProvider struct {
	configPG          *config.ConfigPG
	configHTTP        *config.ConfigHTTP
	configExternalApi *config.ConfigExternalApi

	pgClient   DBclient.ClientDB
	authClient *clientAuthService.AuthClient

	productsRepository   *repositoryProducts.ProductsRepository
	categoriesRepository *repositoryCategories.CategoriesRepository

	ucProducts        *ucProducts.UseCaseProducts
	ucCategories *ucCategories.UseCaseCategories

	handlerHTTP *deliveryHttp.Handler
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (sp *serviceProvider) ConfigPG() *config.ConfigPG {
	if sp.configPG == nil {
		sp.configPG = config.NewConfigPG()
	}
	return sp.configPG
}

func (sp *serviceProvider) ConfigHTTP() *config.ConfigHTTP {
	if sp.configHTTP == nil {
		sp.configHTTP = config.NewConfigHTTP()
	}
	return sp.configHTTP
}

func (sp *serviceProvider) ConfigExternalApi() *config.ConfigExternalApi {
	if sp.configExternalApi == nil {
		sp.configExternalApi = config.NewConfigExternalApi()
	}
	return sp.configExternalApi
}

func (sp *serviceProvider) PgClient(ctx context.Context) DBclient.ClientDB {
	if sp.pgClient == nil {
		client, err := pg.NewClient(ctx, sp.configPG.DSN())
		if err != nil {
			log.Fatalf("failed to connect to postgres: %v", err)
		}
		sp.pgClient = client
	}
	return sp.pgClient
}

func (sp *serviceProvider) AuthClient(ctx context.Context) *clientAuthService.AuthClient {
	if sp.authClient == nil {
		client, err := clientAuthService.New(ctx, sp.ConfigExternalApi())
		if err != nil {
			log.Fatalf("failed to create auth service client: %v", err)
		}
		sp.authClient = client
	}
	return sp.authClient
}

func (sp *serviceProvider) ProductsRepository(ctx context.Context) *repositoryProducts.ProductsRepository {
	if sp.productsRepository == nil {
		sp.productsRepository = repositoryProducts.New(sp.PgClient(ctx))
	}
	return sp.productsRepository
}

func (sp *serviceProvider) CategoriesRepository(ctx context.Context) *repositoryCategories.CategoriesRepository {
	if sp.categoriesRepository == nil {
		sp.categoriesRepository = repositoryCategories.New(sp.PgClient(ctx))
	}
	return sp.categoriesRepository
}

func (sp *serviceProvider) ProductsService(ctx context.Context) *ucProducts.UseCaseProducts {
	if sp.ucProducts == nil {
		sp.ucProducts = ucProducts.New(
			sp.ProductsRepository(ctx),
		)
	}
	return sp.ucProducts
}

func (sp *serviceProvider) CategoriesService(ctx context.Context) *ucCategories.UseCaseCategories {
	if sp.ucCategories == nil {
		sp.ucCategories = ucCategories.New(
			sp.CategoriesRepository(ctx),
		)
	}
	return sp.ucCategories
}

func (sp *serviceProvider) HandlerHTTP(ctx context.Context) *deliveryHttp.Handler {
	if sp.handlerHTTP == nil {
		sp.handlerHTTP = deliveryHttp.New(
			nil,
			sp.AuthClient(ctx),
			sp.ProductsService(ctx),
			sp.CategoriesService(ctx),
		)
	}
	return sp.handlerHTTP
}
