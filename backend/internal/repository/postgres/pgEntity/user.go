package pgEntity

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const usersTableName = "users"

type UserRow struct {
	Uid       pgtype.UUID
	FirstName string
	LastName  string
	Email     string
	Birthday  pgtype.Timestamp
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

	ur.FirstName = user.FirstName
	ur.LastName = user.LastName
	ur.Email = user.Email

	if user.Birthday.Unix() <= 0 {
		ur.Birthday = pgtype.Timestamp{
			Status: pgtype.Null,
		}
	} else {
		ur.Birthday = pgtype.Timestamp{
			Time:   user.Birthday,
			Status: pgtype.Present,
		}
	}

	if user.CreatedAt.Unix() <= 0 {
		ur.CreatedAt = pgtype.Timestamp{
			Time:   time.Now().UTC(),
			Status: pgtype.Present,
		}
	} else {
		ur.CreatedAt = pgtype.Timestamp{
			Time:   user.CreatedAt,
			Status: pgtype.Present,
		}
	}

	if user.UpdatedAt.Unix() <= 0 {
		ur.UpdatedAt = pgtype.Timestamp{
			Time:   time.Now().UTC(),
			Status: pgtype.Present,
		}
	} else {
		ur.UpdatedAt = pgtype.Timestamp{
			Time:   user.UpdatedAt,
			Status: pgtype.Present,
		}
	}
	return ur
}

func (ur *UserRow) ToEntity() entity.User {
	return entity.User{
		Uid:       ur.Uid.Bytes,
		FirstName: ur.FirstName,
		LastName:  ur.LastName,
		Birthday:  ur.Birthday.Time,
		Email:     ur.Email,
		CreatedAt: ur.CreatedAt.Time,
		UpdatedAt: ur.UpdatedAt.Time,
	}
}

var usersTableColumns = []string{
	"uid", "first_name", "last_name", "email", "birthday", "created_at", "updated_at",
}

func (ur *UserRow) Values() []interface{} {
	return []interface{}{ur.Uid, ur.FirstName, ur.LastName, ur.Email, ur.Birthday, ur.CreatedAt, ur.UpdatedAt}
}

func (ur *UserRow) Columns() []string {
	return usersTableColumns
}

func (ur *UserRow) Table() string {
	return usersTableName
}

func (ur *UserRow) Scan(row pgx.Row) error {
	return row.Scan(&ur.Uid, &ur.FirstName, &ur.LastName, &ur.Email, &ur.Birthday, &ur.CreatedAt, &ur.UpdatedAt)
}

func (ur *UserRow) ColumnsForUpdate() []string {
	return []string{"first_name", "last_name", "birthday", "email", "updated_at"}
}

func (ur *UserRow) ValuesForUpdate() []interface{} {
	return []interface{}{ur.FirstName, ur.LastName, ur.Birthday.Time, ur.Email, ur.UpdatedAt.Time}
}

func (ur *UserRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{"uid": ur.Uid}
}
