package postgresql

import (
	"context"

	"github.com/deadbutnotcry/crypto-prices-bot/internal/database/postgresql"
	"github.com/deadbutnotcry/crypto-prices-bot/internal/user"
)

type repository struct {
	client postgresql.Client
}

func (r *repository) Create(ctx context.Context, u user.AppUser) error {
	q := `INSERT INTO app_user (id, is_admin, is_banned, state, last_message_date, register_date) 
			VALUES ($1,$2,$3,$4,$5,$6)`
	_, err := r.client.Exec(ctx, q, u.Id, u.IsAdmin, u.IsBanned, u.State, u.LastMessageDate, u.RegisterDate)
	return err
}

func (r *repository) Update(ctx context.Context, u user.AppUser) error {
	q := `UPDATE app_user
    SET state = $1,
        is_banned = $2,
        last_message_date = $3
    WHERE id = $4`
	_, err := r.client.Exec(ctx, q, u.State, u.IsBanned, u.LastMessageDate, u.Id)
	return err
}

func (r *repository) FindAll(ctx context.Context) ([]user.AppUser, error) {
	q := `SELECT id,is_admin,is_banned,state,last_message_date,register_date FROM app_user`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	users := make([]user.AppUser, 0)

	for rows.Next() {
		var u user.AppUser
		err = rows.Scan(u.Id, u.IsAdmin, u.IsBanned, u.State, u.LastMessageDate, u.RegisterDate)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *repository) Find(ctx context.Context, id int64) (user.AppUser, error) {
	var u user.AppUser
	q := `SELECT id,is_admin,is_banned,state,last_message_date,register_date FROM app_user WHERE id = $1`
	err := r.client.QueryRow(ctx, q, id).Scan(&u.Id, &u.IsAdmin, &u.IsBanned, &u.State, &u.LastMessageDate, &u.RegisterDate)
	if err != nil {
		return user.AppUser{}, err
	}
	return u, nil
}

func NewRepository(client postgresql.Client) (user.Repository, error) {
	return &repository{
		client: client,
	}, nil
}
