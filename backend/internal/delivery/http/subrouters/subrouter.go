package subrouters

type SubrouterDeps struct {
	AuthService     AuthService
	ProductsService ProductsService
	UsersService    UsersService
	RecipesService  RecipesService
}
