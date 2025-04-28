package merch

type Repository interface {
	GetByName(name string) (*Item, error)
}
