package domain

type Client interface {
	UserRepository() UserRepository
}
