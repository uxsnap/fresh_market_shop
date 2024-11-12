package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const userPaymentCardsTable = "users_cards"

type UserPaymentCardRow struct {
	Uid         pgtype.UUID
	UserUid     pgtype.UUID
	ExternalUid pgtype.UUID
	Number      string
	Expired     string
}

func NewUserPaymentCardRow() *UserPaymentCardRow {
	return &UserPaymentCardRow{}
}

func (ur *UserPaymentCardRow) FromEntity(card entity.UserPaymentCard) *UserPaymentCardRow {
	ur.Uid = pgUidFromUUID(card.Uid)
	ur.UserUid = pgUidFromUUID(card.UserUid)
	ur.ExternalUid = pgUidFromUUID(card.ExternalUid)
	ur.Number = card.Number
	ur.Expired = card.Expired
	return ur
}

func (ur *UserPaymentCardRow) ToEntity() entity.UserPaymentCard {
	return entity.UserPaymentCard{
		Uid:         ur.Uid.Bytes,
		UserUid:     ur.UserUid.Bytes,
		ExternalUid: ur.ExternalUid.Bytes,
		Number:      ur.Number,
		Expired:     ur.Expired,
	}
}

var userPaymentCardsTableColumns = []string{
	"uid", "user_uid", "external_uid",
	"number", "expired",
}

func (ur *UserPaymentCardRow) Values() []interface{} {
	return []interface{}{
		ur.Uid, ur.UserUid, ur.ExternalUid, ur.Number, ur.Expired,
	}
}

func (ur *UserPaymentCardRow) Columns() []string {
	return userPaymentCardsTableColumns
}

func (ur *UserPaymentCardRow) Table() string {
	return userPaymentCardsTable
}

func (ur *UserPaymentCardRow) Scan(row pgx.Row) error {
	return row.Scan(&ur.Uid, &ur.UserUid, &ur.ExternalUid, &ur.Number, &ur.Expired)
}

func (ur *UserPaymentCardRow) ColumnsForUpdate() []string {
	return nil
}

func (ur *UserPaymentCardRow) ValuesForUpdate() []interface{} {
	return nil
}

func (ur *UserPaymentCardRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{
		"uid": ur.Uid,
	}
}

func (ur *UserPaymentCardRow) ConditionUserUidEqual() sq.Eq {
	return sq.Eq{
		"user_uid": ur.UserUid,
	}
}

func NewUserPaymentCardRows() *Rows[*UserPaymentCardRow, entity.UserPaymentCard] {
	return &Rows[*UserPaymentCardRow, entity.UserPaymentCard]{}
}
