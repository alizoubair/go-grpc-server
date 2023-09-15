package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alizoubair/go-grpc-server/api/api_struct/model"
	"github.com/alizoubair/go-grpc-server/api/user"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when openning a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "address", "created_at", "updated_at", "deleted_at"}).
		AddRow(xid.New().String(), "John", "John@email.com", "(012)-345-6789", "Japan", time.Now().UTC(), time.Now().UTC(), nil)

	query := "SELECT " + model.UserSelectField + " FROM users WHERE deleted_at IS NULL ORDER BY id ASC LIMIT \\? OFFSET \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	u := user.NewUserRepository(sqlxDB, sqlxDB)

	orderBy := "id ASC"
	where := "WHERE deleted_at IS NULL"
	filter := map[string]interface{}{}
	filter["limit"] = "10"
	filter["offset"] = "10"

	users, err := u.GetUsers(context.Background(), filter, where, orderBy, "")

	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	req := &model.UserModel{
		ID:        xid.New().String(),
		Name:      "John",
		Email:     "john@email.com",
		Address:   "Japan",
		Phone:     "(012)-345-6789",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "email", "address", "phone", "created_at", "updated_at"}).
		AddRow(xid.New().String(), "John", "john@email.com", "Japan", "(012)-345-6789", time.Now().UTC(), time.Now().UTC())

	query := "SELECT " + model.UserSelectField + " FROM users WHERE deleted_at IS NULL AND id = \\?"

	mock.ExpectQuery(query).WithArgs(req.ID).WillReturnRows(rows)
	u := user.NewUserRepository(sqlxDB, sqlxDB)

	userRow, err := u.GetUser(context.Background(), req.ID, "")

	assert.NoError(t, err)
	assert.NotNil(t, userRow)
}

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"id", "name", "email", "address", "phone", "created_at", "updated_at", "deleted_at"}).
		AddRow(xid.New().String(), "John", "john@email.com", "Japan", "(012)-345-6789", time.Now().UTC(), time.Now().UTC(), nil)

	query := "SELECT " + model.UserSelectField + " FROM users WHERE deleted_at IS NULL ORDER BY id ASC LIMIT \\? OFFSET \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	u := user.NewUserRepository(sqlxDB, sqlxDB)

	orderBy := "id ASC"
	where := "WHERE deleted_at IS NULL"
	filter := map[string]interface{}{}
	filter["limit"] = "10"
	filter["offset"] = "0"

	users, err := u.GetUsers(context.Background(), filter, where, orderBy, "")

	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestCount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a new stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"COUNT(id)"}).AddRow(1)

	query := "SELECT COUNT\\(id\\) FROM users WHERE deleted_at IS NULL"

	mock.ExpectQuery(query).WillReturnRows(rows)
	u := user.NewUserRepository(sqlxDB, sqlxDB)

	where := "WHERE deleted_at IS NULL"
	filter := map[string]interface{}{}

	count, err := u.CountUsers(context.Background(), filter, where)

	assert.NoError(t, err)
	assert.Equal(t, count, 1)
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	req := &model.UserModel{
		ID:      xid.New().String(),
		Name:    "John",
		Email:   "john@email.com",
		Address: "Japan",
	}

	query := "UPDATE users SET name = \\?, email = \\?, address = \\?, phone = \\?, updated_at = CURRENT_TIMESTAMP WHERE id = \\?"

	mock.ExpectExec(query).WithArgs(req.Name, req.Email, req.Address, req.Phone, req.ID).WillReturnResult(sqlmock.NewResult(0, 1))
	u := user.NewUserRepository(sqlxDB, sqlxDB)

	userRow, err := u.UpdateUser(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, req.ID, userRow.ID)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	id := xid.New().String()

	query := "UPDATE users SET deleted_at = CURRENT_STAMP WHERE id = \\?"

	mock.ExpectExec(query).WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1))
	u := user.NewUserRepository(sqlxDB, sqlxDB)

	u.DeleteUser(context.Background(), id)

	assert.NoError(t, err)
}
