package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const supportTicketsTableName = "support_tickets"

type SupportTicketRow struct {
	Uid         pgtype.UUID
	UserUid     pgtype.UUID
	TopicUid    pgtype.UUID
	SolverUid   pgtype.UUID
	FromEmail   string
	FromPhone   string
	Title       string
	Description string
	Status      string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func NewSupportTicketRow() *SupportTicketRow {
	return &SupportTicketRow{}
}

func (st *SupportTicketRow) New() *SupportTicketRow {
	return &SupportTicketRow{}
}

func (st *SupportTicketRow) FromEntity(ticket entity.SupportTicket) *SupportTicketRow {
	st.Uid = pgUidFromUUID(ticket.Uid)
	st.UserUid = pgUidFromUUID(ticket.UserUid)
	st.TopicUid = pgUidFromUUID(ticket.TopicUid)
	st.SolverUid = pgUidFromUUID(ticket.SolverUid)
	st.FromEmail = ticket.FromEmail
	st.FromPhone = ticket.FromPhone
	st.Title = ticket.Title
	st.Description = ticket.Description
	st.Status = string(ticket.Status)
	st.CreatedAt = pgtype.Timestamp{
		Time:   ticket.CreatedAt,
		Status: pgStatusFromTime(ticket.CreatedAt),
	}
	st.UpdatedAt = pgtype.Timestamp{
		Time:   ticket.UpdatedAt,
		Status: pgStatusFromTime(ticket.UpdatedAt),
	}
	return st
}

func (st *SupportTicketRow) ToEntity() entity.SupportTicket {
	return entity.SupportTicket{
		Uid:         st.Uid.Bytes,
		UserUid:     st.UserUid.Bytes,
		TopicUid:    st.TopicUid.Bytes,
		SolverUid:   st.SolverUid.Bytes,
		FromEmail:   st.FromEmail,
		FromPhone:   st.FromPhone,
		Title:       st.Title,
		Description: st.Description,
		Status:      entity.SupportTicketStatus(st.Status),
		CreatedAt:   st.CreatedAt.Time,
		UpdatedAt:   st.UpdatedAt.Time,
	}
}

var supportTicketsTableColumns = []string{
	"uid", "user_uid", "topic_uid", "solver_uid",
	"from_email", "from_phone", "title", "description",
	"status", "created_at", "updated_at",
}

func (st *SupportTicketRow) Values() []interface{} {
	return []interface{}{
		st.Uid, st.UserUid, st.TopicUid, st.SolverUid,
		st.FromEmail, st.FromPhone, st.Title, st.Description,
		st.Status, st.CreatedAt, st.UpdatedAt,
	}
}

func (st *SupportTicketRow) Columns() []string {
	return supportTicketsTableColumns
}

func (st *SupportTicketRow) Table() string {
	return supportTicketsTableName
}

func (st *SupportTicketRow) Scan(row pgx.Row) error {
	return row.Scan(
		&st.Uid, &st.UserUid, &st.TopicUid, &st.SolverUid,
		&st.FromEmail, &st.FromPhone, &st.Title, &st.Description,
		&st.Status, &st.CreatedAt, &st.UpdatedAt,
	)
}

func (st *SupportTicketRow) ColumnsForUpdate() []string {
	return []string{
		"user_uid", "topic_uid", "solver_uid",
		"from_email", "from_phone", "title", "description",
		"status", "updated_at",
	}
}

func (st *SupportTicketRow) ValuesForUpdate() []interface{} {
	return []interface{}{
		st.UserUid, st.TopicUid, st.SolverUid,
		st.FromEmail, st.FromPhone, st.Title, st.Description,
		st.Status, st.UpdatedAt,
	}
}

func (st *SupportTicketRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{"uid": st.Uid}
}

func (st *SupportTicketRow) ConditionUserUidEqual() sq.Eq {
	return sq.Eq{"user_uid": st.UserUid}
}

func (st *SupportTicketRow) ConditionTopicUidEqual() sq.Eq {
	return sq.Eq{"topic_uid": st.TopicUid}
}

func (st *SupportTicketRow) ConditionSolverUidEqual() sq.Eq {
	return sq.Eq{"solver_uid": st.SolverUid}
}

func (st *SupportTicketRow) ConditionStatusEqual() sq.Eq {
	return sq.Eq{"status": st.Status}
}

func (st *SupportTicketRow) ConditionFromEmailEqual() sq.Eq {
	return sq.Eq{"from_email": st.FromEmail}
}

func NewSupportTicketRows() *Rows[*SupportTicketRow, entity.SupportTicket] {
	return &Rows[*SupportTicketRow, entity.SupportTicket]{}
}
