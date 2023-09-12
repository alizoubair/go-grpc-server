package user

import (
	"context"
	"fmt"

	"github.com/alizoubair/go-grpc-server/api/api_struct/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.UserModel) (*model.UserModel, error)
	GetUser(ctx context.Context, id, selectField string) (*model.UserModel, error)
	GetUsers(ctx context.Context, filter map[string]interface{}, where, orderBy, selectField string)([]*model.UserModel, error)
	UpdateUser(ctx context.Context, user *model.UserModel) (*model.UserModel, error)
	DeleteUser(ctx context.Context, id string) error
}

type userRepository struct {
	DBRead  *sqlx.DB
	DBWrite *sqlx.DB
}

var _ UserRepository = (*userRepository)(nil)

func NewUserRepository(DBRead *sqlx.DB, DBWrite *sqlx.DB) UserRepository {
	return &userRepository{
		DBRead:  DBRead,
		DBWrite: DBWrite,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.UserModel) (*model.UserModel, error) {
	_, err := r.DBWrite.NamedExec(`INSERT INTO users (id, name, email, address, phone, created_at, updated_at) VALUES (:id, :name, :email, :address, :phone, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`, user)
	return user, err
}

func (r *userRepository) GetUser(ctx context.Context, id, selectField string) (*model.UserModel, error) {
	where := "WHERE deleted_at IS NULL"
	filter := map[string]interface{}{}

	where += " AND id = :id"
	filter["id"] = id

	user := &model.UserModel{}
	if len(selectField) == 0 {
		selectField = model.UserSelectField
	}
	query := fmt.Sprintf("SELECT %s FROM users %s", selectField, where)
	namedQuery, args, _ := r.DBRead.BindNamed(query, filter)
	err := r.DBRead.Get(user, namedQuery, args...)

	return user, err
}

func (r *userRepository) GetUsers(ctx context.Context, filter map[string]interface{}, where, orderBy, selectField string) ([]*model.UserModel, error) {
	users := []*model.UserModel{}
	if len(selectField) == 0 {
		selectField = model.UserSelectField
	}
	query := fmt.Sprintf("SELECT %s FROM users %s ORDER BY %s LIMIT :limit OFFSET :offset", selectField, where, orderBy)
	namedQuery, args, _ := r.DBRead.BindNamed(query, filter)
	err := r.DBRead.Select(&users, namedQuery, args...)
	return users, err
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.UserModel) (*model.UserModel, error) {
	_, err := r.DBWrite.NamedExec(`UPDATE users SET name = :name, email = :email, address = :address, phone = :phone, updated_at = CURRENT_TIMESTAMP WHERE id = :id`, user)
	return user, err
}

func (r *userRepository) DeleteUser(ctx context.Context, id string) error {
	user := &model.UserModel{
		ID: id,
	}
	_, err := r.DBWrite.NamedExec(`UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id = :id`, user)
	return err
}
