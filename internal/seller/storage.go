package seller

import "context"

type Repository interface {
	GetSellers(ctx context.Context) ([]Seller, error)
}
