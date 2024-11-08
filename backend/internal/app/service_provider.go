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
	repositoryOrderProducts "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/order_products"
	repositoryOrders "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/orders"
	repositoryProducts "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/products"
	repositoryProductsCount "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/products_count"
	repositoryRecipes "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/recipes"
	repositoryUsers "github.com/uxsnap/fresh_market_shop/backend/internal/repository/postgres/users"
	ucOrders "github.com/uxsnap/fresh_market_shop/backend/internal/usecase/orders"
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

	productsRepository      *repositoryProducts.ProductsRepository
	categoriesRepository    *repositoryCategories.CategoriesRepository
	usersRepository         *repositoryUsers.UsersRepository
	recipesRepository       *repositoryRecipes.RecipesRepository
	ordersRepository        *repositoryOrders.OrdersRepository
	productsCountRepository *repositoryProductsCount.ProductsCountRepository
	orderProductsRepository *repositoryOrderProducts.OrderProductsRepository

	ucProducts *ucProducts.UseCaseProducts
	ucUsers    *ucUsers.UseCaseUsers
	ucRecipes  *ucRecipes.UseCaseRecipes
	ucOrders   *ucOrders.UseCaseOrders

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
		client, err := pg.NewClient(ctx, sp.ConfigPG().DSN())
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

func (sp *serviceProvider) OrdersRepository(ctx context.Context) *repositoryOrders.OrdersRepository {
	if sp.ordersRepository == nil {
		sp.ordersRepository = repositoryOrders.New(sp.PgClient(ctx))
	}
	return sp.ordersRepository
}

func (sp *serviceProvider) ProductsCountRepository(ctx context.Context) *repositoryProductsCount.ProductsCountRepository {
	if sp.productsCountRepository == nil {
		sp.productsCountRepository = repositoryProductsCount.New(sp.PgClient(ctx))
	}
	return sp.productsCountRepository
}

func (sp *serviceProvider) OrderProductsRepository(ctx context.Context) *repositoryOrderProducts.OrderProductsRepository {
	if sp.orderProductsRepository == nil {
		sp.orderProductsRepository = repositoryOrderProducts.New(sp.PgClient(ctx))
	}
	return sp.orderProductsRepository
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
		sp.ucRecipes = ucRecipes.New(sp.RecipesRepository(ctx), sp.ProductsRepository(ctx))
	}
	return sp.ucRecipes
}

func (sp *serviceProvider) OrdersService(ctx context.Context) *ucOrders.UseCaseOrders {
	if sp.ucOrders == nil {
		sp.ucOrders = ucOrders.New(
			sp.OrdersRepository(ctx),
			sp.OrderProductsRepository(ctx),
			sp.ProductsCountRepository(ctx),
			sp.ProductsRepository(ctx),
			sp.TxManager(ctx),
		)
	}
	return sp.ucOrders
}

func (sp *serviceProvider) HandlerHTTP(ctx context.Context) *deliveryHttp.Handler {
	if sp.handlerHTTP == nil {
		sp.handlerHTTP = deliveryHttp.New(
			nil,
			sp.AuthClient(ctx),
			sp.ProductsService(ctx),
			sp.UsersService(ctx),
			sp.RecipesService(ctx),
			sp.OrdersService(ctx),
		)
	}
	return sp.handlerHTTP
}
