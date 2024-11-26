package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const supportTicketsTopicsTableName = "support_tickets_topics"

type SupportTicketsTopicRow struct {
	Uid         pgtype.UUID
	Name        string
	Description string
}

func NewSupportTicketsTopicRow() *SupportTicketsTopicRow {
	return &SupportTicketsTopicRow{}
}

func (s *SupportTicketsTopicRow) New() *SupportTicketsTopicRow {
	return &SupportTicketsTopicRow{}
}

func (s *SupportTicketsTopicRow) FromEntity(topic entity.SupportTicketsTopic) *SupportTicketsTopicRow {
	s.Uid = pgUidFromUUID(topic.Uid)
	s.Name = topic.Name
	s.Description = topic.Description
	return s
}

func (s *SupportTicketsTopicRow) ToEntity() entity.SupportTicketsTopic {
	return entity.SupportTicketsTopic{
		Uid:         s.Uid.Bytes,
		Name:        s.Name,
		Description: s.Description,
	}
}

var supportTicketsTopicsTableColumns = []string{"uid", "name", "description"}

func (s *SupportTicketsTopicRow) Values() []interface{} {
	return []interface{}{s.Uid, s.Name, s.Description}
}

func (s *SupportTicketsTopicRow) Columns() []string {
	return supportTicketsTopicsTableColumns
}

func (s *SupportTicketsTopicRow) Table() string {
	return supportTicketsTopicsTableName
}

func (s *SupportTicketsTopicRow) Scan(row pgx.Row) error {
	return row.Scan(&s.Uid, &s.Name, &s.Description)
}

func (s *SupportTicketsTopicRow) ColumnsForUpdate() []string {
	return []string{"name", "description"}
}

func (s *SupportTicketsTopicRow) ValuesForUpdate() []interface{} {
	return []interface{}{s.Name, s.Description}
}

func (s *SupportTicketsTopicRow) ConditionUidEqual() sq.Eq {
	return sq.Eq{"uid": s.Uid}
}

func NewSupportTicketsTopicRows() *Rows[*SupportTicketsTopicRow, entity.SupportTicketsTopic] {
	return &Rows[*SupportTicketsTopicRow, entity.SupportTicketsTopic]{}
}
