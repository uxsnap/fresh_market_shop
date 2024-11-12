package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const PaymentsTableName = "payments"

type PaymentRow struct {
	Uid         pgtype.UUID
	UserUid     pgtype.UUID
	OrderUid    pgtype.UUID
	CardUid     pgtype.UUID
	Sum         int64
	Currency    string
	PaymentTime pgtype.Timestamp
}

func NewPaymentRow() *PaymentRow {
	return &PaymentRow{}
}

func (pr *PaymentRow) FromEntity(payment entity.Payment) *PaymentRow {
	pr.Uid = pgUidFromUUID(payment.Uid)
	pr.UserUid = pgUidFromUUID(payment.UserUid)
	pr.OrderUid = pgUidFromUUID(payment.OrderUid)
	pr.CardUid = pgUidFromUUID(payment.CardUid)
	pr.Sum = payment.Sum
	pr.Currency = payment.Currency
	pr.PaymentTime = pgtype.Timestamp{
		Time:   payment.PaymentTime,
		Status: pgStatusFromTime(payment.PaymentTime),
	}
	return pr
}

func (pr *PaymentRow) ToEntity() entity.Payment {
	return entity.Payment{
		Uid:         pr.Uid.Bytes,
		UserUid:     pr.UserUid.Bytes,
		OrderUid:    pr.OrderUid.Bytes,
		CardUid:     pr.CardUid.Bytes,
		Sum:         pr.Sum,
		Currency:    pr.Currency,
		PaymentTime: pr.PaymentTime.Time,
	}
}

var paymentsTableColumns = []string{
	"uid", "user_uid", "order_uid", "card_uid", "sum", "currency", "time",
}

func (pr *PaymentRow) Values() []interface{} {
	return []interface{}{
		pr.Uid, pr.UserUid, pr.OrderUid, pr.CardUid, pr.Sum, pr.Currency, pr.PaymentTime,
	}
}

func (pr *PaymentRow) Columns() []string {
	return paymentsTableColumns
}

func (pr *PaymentRow) Table() string {
	return PaymentsTableName
}

func (pr *PaymentRow) Scan(row pgx.Row) error {
	return row.Scan(&pr.Uid, &pr.UserUid, &pr.OrderUid, &pr.CardUid, &pr.Sum, &pr.Currency, &pr.PaymentTime)
}

func (pr *PaymentRow) ColumnsForUpdate() []string {
	return nil
}

func (pr *PaymentRow) ValuesForUpdate() []interface{} {
	return nil
}

func (pr *PaymentRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{
		"uid": pr.Uid,
	}
}

func (pr *PaymentRow) ConditionUserUidEqual() sq.Eq {
	return sq.Eq{
		"user_uid": pr.UserUid,
	}
}

func (pr *PaymentRow) ConditionOrderUidEqual() sq.Eq {
	return sq.Eq{
		"order_uid": pr.Uid,
	}
}

func (pr *PaymentRow) ConditionCardUidEqual() sq.Eq {
	return sq.Eq{
		"order_uid": pr.Uid,
	}
}

func NewPaymentRows() *Rows[*PaymentRow, entity.Payment] {
	return &Rows[*PaymentRow, entity.Payment]{}
}
