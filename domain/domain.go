package domain

// Domain business
type Domain struct {
	store Store
}

func New(store Store) *Domain {
	return &Domain{
		store: store,
	}
}
