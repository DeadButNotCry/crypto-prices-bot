package user

import "context"

type Repository interface {
	Create(ctx context.Context, user AppUser) error
	Update(ctx context.Context, user AppUser) error
	FindAll(ctx context.Context) ([]AppUser, error)
	Find(ctx context.Context, id int64)
	Delete(ctx context.Context, id string)
}
