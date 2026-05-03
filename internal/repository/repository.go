package repository

type Repository[T any] interface {
	GetAll() ([]T, error)
	GetByID(id string) (*T, error)
	Create(v T) (*T, error)
	Delete(id string) (int, error)
	Count() (int, error)
	Clear() error
}
