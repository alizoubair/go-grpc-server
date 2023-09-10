package user

import (
	"context"

	"github.com/alizoubair/go-grpc-server/api/api_struct/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(ctx context.Context, userReq *model.UserModel) (*model.UserModel, error)
}

type userRepository struct {
	DBRead  *sqlx.DB
	DBWrite *sqlx.DB
}

func NewUserRepository(DBRead *sqlx.DB, DBWrite *sqlx.DB) UserRepository {
	return &userRepository{
		DBRead: DBRead,
		DBWrite: DBWrite,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.UserModel) (*model.UserModel, error) {
	_, err := r.DBWrite.NamedExec(`INSERT INTO users (id, name, email, address, phone, created_at, updated_at) VALUES (:id, :name, :email, :address, :phone, :created_at, :updated_at)`, user)
	return user, err
}
