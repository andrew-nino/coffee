package repository

type Authorisation interface {
}

type CoffeeList interface {
}

type CoffeeItem interface {
}

type Repository struct {
	Authorisation
	CoffeeList
	CoffeeItem
}

func NewRepository() *Repository {

	return &Repository{}
}
