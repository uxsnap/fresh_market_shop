package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const supportTicketsCommentMessagesTableName = "support_tickets_comment_messages"

type SupportTicketCommentMessageRow struct {
	Uid                 pgtype.UUID
	TicketUid           pgtype.UUID
	SenderUid           pgtype.UUID
	Content             string
	IsImportedFromEmail bool
	CreatedAt           pgtype.Timestamp
	UpdatedAt           pgtype.Timestamp
}

func NewSupportTicketCommentMessageRow() *SupportTicketCommentMessageRow {
	return &SupportTicketCommentMessageRow{}
}

func (sm *SupportTicketCommentMessageRow) New() *SupportTicketCommentMessageRow {
	return &SupportTicketCommentMessageRow{}
}

func (sm *SupportTicketCommentMessageRow) FromEntity(message entity.SupportTicketCommentMessage) *SupportTicketCommentMessageRow {
	sm.Uid = pgUidFromUUID(message.Uid)
	sm.TicketUid = pgUidFromUUID(message.TicketUid)
	sm.SenderUid = pgUidFromUUID(message.SenderUid)
	sm.Content = message.Content
	sm.IsImportedFromEmail = message.IsImportedFromEmail
	sm.CreatedAt = pgtype.Timestamp{
		Time:   message.CreatedAt,
		Status: pgStatusFromTime(message.CreatedAt),
	}
	sm.UpdatedAt = pgtype.Timestamp{
		Time:   message.UpdatedAt,
		Status: pgStatusFromTime(message.UpdatedAt),
	}
	return sm
}

func (sm *SupportTicketCommentMessageRow) ToEntity() entity.SupportTicketCommentMessage {
	return entity.SupportTicketCommentMessage{
		Uid:                 sm.Uid.Bytes,
		TicketUid:           sm.TicketUid.Bytes,
		SenderUid:           sm.SenderUid.Bytes,
		Content:             sm.Content,
		IsImportedFromEmail: sm.IsImportedFromEmail,
		CreatedAt:           sm.CreatedAt.Time,
		UpdatedAt:           sm.UpdatedAt.Time,
	}
}

var supportTicketsCommentMessagesTableColumns = []string{
	"uid", "ticket_uid", "sender_uid", "content",
	"is_imported_from_email", "created_at", "updated_at",
}

func (sm *SupportTicketCommentMessageRow) Values() []interface{} {
	return []interface{}{
		sm.Uid, sm.TicketUid, sm.SenderUid, sm.Content,
		sm.IsImportedFromEmail, sm.CreatedAt, sm.UpdatedAt,
	}
}

func (sm *SupportTicketCommentMessageRow) Columns() []string {
	return supportTicketsCommentMessagesTableColumns
}

func (sm *SupportTicketCommentMessageRow) Table() string {
	return supportTicketsCommentMessagesTableName
}

func (sm *SupportTicketCommentMessageRow) Scan(row pgx.Row) error {
	return row.Scan(
		&sm.Uid, &sm.TicketUid, &sm.SenderUid, &sm.Content,
		&sm.IsImportedFromEmail, &sm.CreatedAt, &sm.UpdatedAt,
	)
}

func (sm *SupportTicketCommentMessageRow) ColumnsForUpdate() []string {
	return []string{
		"content", "is_imported_from_email", "updated_at",
	}
}

func (sm *SupportTicketCommentMessageRow) ValuesForUpdate() []interface{} {
	return []interface{}{
		sm.Content, sm.IsImportedFromEmail, sm.UpdatedAt,
	}
}

func (sm *SupportTicketCommentMessageRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{
		"uid": sm.Uid,
	}
}

func (sm *SupportTicketCommentMessageRow) ConditionTicketUidEqual() sq.Eq {
	return sq.Eq{
		"ticket_uid": sm.TicketUid,
	}
}

func NewSupportTicketCommentMessageRows() *Rows[*SupportTicketCommentMessageRow, entity.SupportTicketCommentMessage] {
	return &Rows[*SupportTicketCommentMessageRow, entity.SupportTicketCommentMessage]{}
}
