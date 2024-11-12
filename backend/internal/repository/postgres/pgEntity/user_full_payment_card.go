package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const UserFullPaymentCardsTable = "full_cards"

type UserFullPaymentCardRow struct {
	Uid     pgtype.UUID
	UserUid pgtype.UUID
	Number  string
	Expired string
	CVV     string
}

func NewUserFullPaymentCardRow() *UserFullPaymentCardRow {
	return &UserFullPaymentCardRow{}
}

var UserFullPaymentCardsTableColumns = []string{
	"uid", "user_uid", "number", "expired", "cvv",
}

func (ufr *UserFullPaymentCardRow) FromEntity(card entity.UserFullPaymentCard) *UserFullPaymentCardRow {
	ufr.Uid = pgUidFromUUID(card.Uid)
	ufr.UserUid = pgUidFromUUID(card.UserUid)
	ufr.Number = card.Number
	ufr.Expired = card.Expired
	ufr.CVV = card.CVV
	return ufr
}

func (ufr *UserFullPaymentCardRow) ToEntity() entity.UserFullPaymentCard {
	return entity.UserFullPaymentCard{
		Uid:     ufr.Uid.Bytes,
		UserUid: ufr.UserUid.Bytes,
		Number:  ufr.Number,
		Expired: ufr.Expired,
		CVV:     ufr.CVV,
	}
}

func (ufr *UserFullPaymentCardRow) Values() []interface{} {
	return []interface{}{
		ufr.Uid, ufr.UserUid, ufr.Number, ufr.Expired, ufr.CVV,
	}
}

func (ufr *UserFullPaymentCardRow) Columns() []string {
	return UserFullPaymentCardsTableColumns
}

func (ufr *UserFullPaymentCardRow) Table() string {
	return UserFullPaymentCardsTable
}

func (ufr *UserFullPaymentCardRow) Scan(row pgx.Row) error {
	return row.Scan(&ufr.Uid, &ufr.UserUid, &ufr.Number, &ufr.Expired, &ufr.CVV)
}

func (ufr *UserFullPaymentCardRow) ColumnsForUpdate() []string {
	return nil
}

func (ufr *UserFullPaymentCardRow) ValuesForUpdate() []interface{} {
	return nil
}

func (ufr *UserFullPaymentCardRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{
		"uid": ufr.Uid,
	}
}

func (ufr *UserFullPaymentCardRow) ConditionUserUidEqual() sq.Eq {
	return sq.Eq{
		"user_uid": ufr.UserUid,
	}
}

func NewUserFullPaymentCardRows() *Rows[*UserFullPaymentCardRow, entity.UserFullPaymentCard] {
	return &Rows[*UserFullPaymentCardRow, entity.UserFullPaymentCard]{}
}
