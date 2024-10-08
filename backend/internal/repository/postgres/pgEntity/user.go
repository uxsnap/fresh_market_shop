package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const usersTableName = "users"

type UserRow struct {
	Uid       pgtype.UUID
	Username  string
	Email     string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

func NewUserRow() *UserRow {
	return &UserRow{}
}

func (ur *UserRow) FromEntity(user entity.User) *UserRow {
	ur.Uid = pgtype.UUID{
		Bytes:  user.Uid,
		Status: pgtype.Present,
	}
	ur.Username = user.Username
	ur.Email = user.Email
	ur.CreatedAt = pgtype.Timestamp{
		Time:   user.CreatedAt,
		Status: pgStatusFromTime(user.CreatedAt),
	}
	ur.UpdatedAt = pgtype.Timestamp{
		Time:   user.UpdatedAt,
		Status: pgStatusFromTime(user.UpdatedAt),
	}
	return ur
}

func (ur *UserRow) ToEntity() entity.User {
	return entity.User{
		Uid:       ur.Uid.Bytes,
		Username:  ur.Username,
		Email:     ur.Email,
		CreatedAt: ur.CreatedAt.Time,
		UpdatedAt: ur.UpdatedAt.Time,
	}
}

var usersTableColumns = []string{
	"uid", "username", "email", "created_at", "updated_at",
}

func (ur *UserRow) Values() []interface{} {
	return []interface{}{ur.Uid, ur.Username, ur.Email, ur.CreatedAt, ur.UpdatedAt}
}

func (ur *UserRow) Columns() []string {
	return usersTableColumns
}

func (ur *UserRow) Table() string {
	return usersTableName
}

func (ur *UserRow) Scan(row pgx.Row) error {
	return row.Scan(&ur.Uid, &ur.Username, &ur.Email, &ur.CreatedAt, &ur.UpdatedAt)
}

func (ur *UserRow) ColumnsForUpdate() []string {
	return []string{"username", "email", "updated_at"}
}

func (ur *UserRow) ValuesForUpdate() []interface{} {
	return []interface{}{ur.Username, ur.Email, ur.UpdatedAt}
}

func (ur *UserRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{"uid": ur.Uid}
}
