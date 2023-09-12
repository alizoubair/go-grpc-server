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