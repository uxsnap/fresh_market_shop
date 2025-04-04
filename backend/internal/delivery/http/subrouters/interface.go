package subrouters

import (
	"context"
	"mime/multipart"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type ProductsService interface {
	CreateProduct(ctx context.Context, product entity.Product) (uuid.UUID, error)
	UpdateProduct(ctx context.Context, product entity.Product) error
	GetProductByUid(ctx context.Context, uid uuid.UUID) (entity.Product, bool, error)
	GetProducts(ctx context.Context, qFilters entity.QueryFilters) ([]entity.Product, error)
	GetProductsWithExtra(ctx context.Context, qFilters entity.QueryFilters) (entity.ProductsWithExtra, error)
	GetProductsByNameLike(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.Product, error)
	GetProductsByNameLikeWithExtra(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error)
	GetProductsLikeNamesWithLimitOnEach(ctx context.Context, names []string, qFilters entity.QueryFilters) ([]entity.Product, error)
	GetProductsLikeNamesWithLimitOnEachWithExtra(ctx context.Context, names []string, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error)
	DeleteProduct(ctx context.Context, uid uuid.UUID) error
	ReviveProduct(ctx context.Context, uid uuid.UUID) error

	UploadProductPhotos(ctx context.Context, uid uuid.UUID, form *multipart.Form) error
	DeleteProductPhotos(ctx context.Context, productUid uuid.UUID, photosUids ...uuid.UUID) error

	UpdateProductCount(ctx context.Context, productUid uuid.UUID, stockQuantity int64) error
	IncrementProductCount(ctx context.Context, productUid uuid.UUID, incValue int64) error
	DecrementProductCount(ctx context.Context, productUid uuid.UUID, decValue int64) error
	GetProductCount(ctx context.Context, productUid uuid.UUID) (int64, bool, error)

	CreateCategory(ctx context.Context, category entity.Category) (uuid.UUID, error)
	GetCategoriesByNameLike(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.Category, error)
	GetCategoryByUid(ctx context.Context, uid uuid.UUID) (entity.Category, error)
	GetCategoriesByUserOrders(ctx context.Context, userUid uuid.UUID) ([]uuid.UUID, error)
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
	UpdateCategory(ctx context.Context, category entity.Category) error
	DeleteCategory(ctx context.Context, uid uuid.UUID) error
}

type AuthService interface {
	Register(ctx context.Context, email string, password string) (uuid.UUID, error)
	CreateAdmin(ctx context.Context, email string, password string, token string) (uuid.UUID, error)
	GetAdmins(ctx context.Context, token string) ([]entity.AuthUser, error)
	UpdateAuthUser(ctx context.Context, accessToken string, uid uuid.UUID, email string, password string) (accessJwt string, refreshJwt string, err error)
	GetAuthUser(ctx context.Context, accessJwt string, uid uuid.UUID, email string) (entity.AuthUser, error)
	DeleteAuthUser(ctx context.Context, accessJwt string, uid uuid.UUID) error
	Login(ctx context.Context, email string, password string) (accessJwt string, refreshJwt string, err error)
	Logout(ctx context.Context, accessJwt string, uid uuid.UUID) error
	Refresh(ctx context.Context, refreshToken string) (accessJwt string, refreshJwt string, err error)
	VerifyJwt(ctx context.Context, token string) error
	VerifyEmail(ctx context.Context, token string) error
	HealthCheck(ctx context.Context) entity.AuthServiceHealthCheck
}

type UsersService interface {
	CreateUser(ctx context.Context, user entity.User) (uuid.UUID, error)
	UpdateUser(ctx context.Context, user entity.User) error
	GetUser(ctx context.Context, uid uuid.UUID) (entity.User, bool, error)
	DeleteUser(ctx context.Context, uid uuid.UUID) error

	AddDeliveryAddress(ctx context.Context, address entity.DeliveryAddress) (uuid.UUID, error)
	UpdateDeliveryAddress(ctx context.Context, address entity.DeliveryAddress) error
	GetDeliveryAddress(ctx context.Context, uid uuid.UUID) (entity.DeliveryAddress, bool, error)
	GetUserDeliveryAddresses(ctx context.Context, userUid uuid.UUID) ([]entity.DeliveryAddress, error)
	DeleteDeliveryAddress(ctx context.Context, uid uuid.UUID) error
	DeleteUserDeliveryAddresses(ctx context.Context, userUid uuid.UUID) error
}

type RecipesService interface {
	CreateRecipe(ctx context.Context, recipe entity.Recipe) (uuid.UUID, error)
	GetRecipeByUid(ctx context.Context, uid uuid.UUID) (entity.Recipe, bool, error)
	GetRecipesByNameLike(ctx context.Context, name string, qFilters entity.QueryFilters) ([]entity.Recipe, error)
	GetRecipes(ctx context.Context, qFilters entity.QueryFilters) (entity.RecipesWithTotal, error)
	GetRecipesProducts(ctx context.Context, recipe_uid uuid.UUID) ([]entity.ProductWithExtra, error)
	GetRecipeSteps(ctx context.Context, recipe_uid uuid.UUID) ([]entity.RecipeStep, error)
	UpdateRecipe(ctx context.Context, recipe entity.Recipe) (uuid.UUID, error)
	DeleteRecipe(ctx context.Context, uid uuid.UUID) error

	UploadRecipePhotos(ctx context.Context, uid uuid.UUID, form *multipart.Form) error
	DeleteRecipePhotos(ctx context.Context, uid uuid.UUID, photoNames ...string) error

	AddRecipeSteps(ctx context.Context, uid uuid.UUID, rSteps []entity.RecipeStep) error
	DeleteRecipeStep(ctx context.Context, uid uuid.UUID, step int) error
}

type OrdersService interface {
	CreateOrder(ctx context.Context, userUid uuid.UUID, productsCounts entity.ProductsCounts) (uuid.UUID, error)
	GetOrder(ctx context.Context, qFilters entity.QueryFilters) (entity.Order, bool, error)
	PayOrder(ctx context.Context, userUid uuid.UUID, orderUid uuid.UUID, cardUid uuid.UUID, deliveryUid uuid.UUID) (uuid.UUID, error)
	GetOrderedProducts(ctx context.Context, qFilters entity.QueryFilters) ([]entity.ProductWithExtra, error)
	GetOrderHistory(ctx context.Context, userUid uuid.UUID, qFilters entity.QueryFilters) ([]entity.OrderWithProducts, error)
}

type DeliveryService interface {
	UpdateDelivery(ctx context.Context, delivery entity.Delivery) error
	CalculateDelivery(
		ctx context.Context,
		userUid uuid.UUID,
		orderUid uuid.UUID,
		orderPrice int64,
		deliveryAddressUid uuid.UUID,
	) (deliveryUid uuid.UUID, deliveryPrice int64, deliveryTime time.Duration, err error)
	GetDeliveryTimeAndPriceForOrder(ctx context.Context, orderUid uuid.UUID) (deliveryTime int64, deliveryPrice int64, err error)
	GetDeliveryByOrderUid(ctx context.Context, orderUid uuid.UUID) (entity.Delivery, bool, error)
	GetDeliveryByUid(ctx context.Context, uid uuid.UUID) (entity.Delivery, bool, error)
}

type AddressesService interface {
	GetCities(ctx context.Context) ([]entity.City, error)
	GetAddresses(ctx context.Context, qFilters entity.QueryFilters) ([]entity.Address, error)
}

type PaymentsService interface {
	CreatePayment(ctx context.Context, payment entity.Payment) (uuid.UUID, error)
	GetPayment(ctx context.Context, paymentUid uuid.UUID) (entity.Payment, bool, error)
	GetOrderPayment(ctx context.Context, orderUid uuid.UUID) (entity.Payment, bool, error)
	GetUserPayments(ctx context.Context, userUid uuid.UUID) ([]entity.Payment, error)
	GetCardPayments(ctx context.Context, cardUid uuid.UUID) ([]entity.Payment, error)

	AddUserPaymentCard(ctx context.Context, userFullCard entity.UserFullPaymentCard) (uuid.UUID, error)
	GetUserPaymentCardByUid(ctx context.Context, cardUid uuid.UUID) (entity.UserPaymentCard, bool, error)
	GetUserPaymentCards(ctx context.Context, userUid uuid.UUID) ([]entity.UserPaymentCard, error)
	DeleteUserPaymentCard(ctx context.Context, cardUid uuid.UUID) error
	DeleteUserPaymentCards(ctx context.Context, userUid uuid.UUID) error
}

type SupportService interface {
	CreateTicketsTopic(ctx context.Context, topic entity.SupportTicketsTopic) (uuid.UUID, error)
	UpdateTicketsTopic(ctx context.Context, topic entity.SupportTicketsTopic) error
	DeleteTicketsTopic(ctx context.Context, uid uuid.UUID) error
	GetTicketsTopicByUid(ctx context.Context, uid uuid.UUID) (entity.SupportTicketsTopic, bool, error)
	GetAllTicketsTopics(ctx context.Context) ([]entity.SupportTicketsTopic, error)

	CreateTicket(ctx context.Context, ticket entity.SupportTicket) (uuid.UUID, error)
	EditTicket(ctx context.Context, ticket entity.SupportTicket) error
	GetTicketByUid(ctx context.Context, uid uuid.UUID) (entity.SupportTicket, bool, error)
	GetTickets(ctx context.Context, qFilters entity.QueryFilters) ([]entity.SupportTicket, error)
	TakeTicket(ctx context.Context, ticketUid, solverUid uuid.UUID) error

	AddTicketMessage(ctx context.Context, message entity.SupportTicketCommentMessage) (uuid.UUID, error)
	EditTicketMessage(ctx context.Context, message entity.SupportTicketCommentMessage) error
	GetTicketMessages(ctx context.Context, ticketUid uuid.UUID, qFilters entity.QueryFilters) ([]entity.SupportTicketCommentMessage, error)

	CreateTicketSolution(ctx context.Context, solution entity.SupportTicketSolution) error
	GetTicketSolution(ctx context.Context, ticketUid uuid.UUID) (entity.SupportTicketSolution, bool, error)
	GetTicketsSolutionsByTopic(ctx context.Context, topicUid uuid.UUID, qFilters entity.QueryFilters) ([]entity.SupportTicketSolution, error)
}
