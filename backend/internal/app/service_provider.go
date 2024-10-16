package app

import (
	"context"
	"log"

	DBclient "github.com/uxsnap/fresh_market_shop/backend/internal/client/database"
	"github.com/uxsnap/fresh_market_shop/backend/internal/client/database/pg"
	clientAuthService "github.com/uxsnap/fresh_market_shop/backend/internal/client/services/auth"
	"github.com/uxsnap/fresh_market_shop/backend/internal/config"
	deliveryHttp "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http"
	"github.com/uxsnap/fresh_market_shop/backend/internal/manager/transaction"
	repositoryCategories "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/categories"
	repositoryProducts "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/products"
	repositoryRecipes "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/recipes"
	repositoryUsers "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/users"
	ucProducts "github.com/uxsnap/fresh_market_shop/backend/internal/usecase/products"
	ucRecipes "github.com/uxsnap/fresh_market_shop/backend/internal/usecase/recipes"
	ucUsers "github.com/uxsnap/fresh_market_shop/backend/internal/usecase/users"
)

type serviceProvider struct {
	configPG          *config.ConfigPG
	configHTTP        *config.ConfigHTTP
	configExternalApi *config.ConfigExternalApi

	pgClient   DBclient.ClientDB
	authClient *clientAuthService.AuthClient

	productsRepository   *repositoryProducts.ProductsRepository
	categoriesRepository *repositoryCategories.CategoriesRepository
	usersRepository      *repositoryUsers.UsersRepository
	recipesRepository    *repositoryRecipes.RecipesRepository

	ucProducts *ucProducts.UseCaseProducts
	ucUsers    *ucUsers.UseCaseUsers
	ucRecipes  *ucRecipes.UseCaseRecipes

	txManager *transaction.Manager

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

func (sp *serviceProvider) TxManager(ctx context.Context) *transaction.Manager {
	if sp.txManager == nil {
		sp.txManager = transaction.NewTxManager(sp.PgClient(ctx))
	}
	return sp.txManager
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

func (sp *serviceProvider) UsersRepository(ctx context.Context) *repositoryUsers.UsersRepository {
	if sp.usersRepository == nil {
		sp.usersRepository = repositoryUsers.New(sp.PgClient(ctx))
	}
	return sp.usersRepository
}

func (sp *serviceProvider) RecipesRepository(ctx context.Context) *repositoryRecipes.RecipesRepository {
	if sp.recipesRepository == nil {
		sp.recipesRepository = repositoryRecipes.New(sp.PgClient(ctx))
	}
	return sp.recipesRepository
}

func (sp *serviceProvider) ProductsService(ctx context.Context) *ucProducts.UseCaseProducts {
	if sp.ucProducts == nil {
		sp.ucProducts = ucProducts.New(
			sp.ProductsRepository(ctx),
			sp.CategoriesRepository(ctx),
			sp.TxManager(ctx),
		)
	}
	return sp.ucProducts
}

func (sp *serviceProvider) UsersService(ctx context.Context) *ucUsers.UseCaseUsers {
	if sp.ucUsers == nil {
		sp.ucUsers = ucUsers.New(sp.UsersRepository(ctx))
	}
	return sp.ucUsers
}

func (sp *serviceProvider) RecipesService(ctx context.Context) *ucRecipes.UseCaseRecipes {
	if sp.ucRecipes == nil {
		sp.ucRecipes = ucRecipes.New(sp.RecipesRepository(ctx))
	}
	return sp.ucRecipes
}

func (sp *serviceProvider) HandlerHTTP(ctx context.Context) *deliveryHttp.Handler {
	if sp.handlerHTTP == nil {
		sp.handlerHTTP = deliveryHttp.New(
			nil,
			sp.AuthClient(ctx),
			sp.ProductsService(ctx),
			sp.UsersService(ctx),
			sp.RecipesService(ctx),
		)
	}
	return sp.handlerHTTP
}
