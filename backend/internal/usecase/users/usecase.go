package useCaseUsers

type UseCaseUsers struct {
	usersRepository UsersRepository
}

func New(
	usersRepository UsersRepository,
) *UseCaseUsers {
	return &UseCaseUsers{
		usersRepository: usersRepository,
	}
}
