package pgEntity

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

const supportTicketsSolutionsTableName = "support_tickets_solutions"

type SupportTicketSolutionRow struct {
	TicketUid   pgtype.UUID
	Description string
	EmailText   string
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func NewSupportTicketSolutionRow() *SupportTicketSolutionRow {
	return &SupportTicketSolutionRow{}
}

func (sts *SupportTicketSolutionRow) New() *SupportTicketSolutionRow {
	return &SupportTicketSolutionRow{}
}

func (sts *SupportTicketSolutionRow) FromEntity(solution entity.SupportTicketSolution) *SupportTicketSolutionRow {
	sts.TicketUid = pgUidFromUUID(solution.TicketUid)
	sts.Description = solution.Description
	sts.EmailText = solution.EmailText
	sts.CreatedAt = pgtype.Timestamp{
		Time:   solution.CreatedAt,
		Status: pgStatusFromTime(solution.CreatedAt),
	}
	sts.UpdatedAt = pgtype.Timestamp{
		Time:   solution.UpdatedAt,
		Status: pgStatusFromTime(solution.UpdatedAt),
	}
	return sts
}

func (sts *SupportTicketSolutionRow) ToEntity() entity.SupportTicketSolution {
	return entity.SupportTicketSolution{
		TicketUid:   sts.TicketUid.Bytes,
		Description: sts.Description,
		EmailText:   sts.EmailText,
		CreatedAt:   sts.CreatedAt.Time,
		UpdatedAt:   sts.UpdatedAt.Time,
	}
}

var supportTicketsSolutionsTableColumns = []string{
	"ticket_uid", "description", "email_text", "created_at", "updated_at",
}

func (sts *SupportTicketSolutionRow) Values() []interface{} {
	return []interface{}{
		sts.TicketUid, sts.Description, sts.EmailText, sts.CreatedAt, sts.UpdatedAt,
	}
}

func (sts *SupportTicketSolutionRow) Columns() []string {
	return supportTicketsSolutionsTableColumns
}

func (sts *SupportTicketSolutionRow) Table() string {
	return supportTicketsSolutionsTableName
}

func (sts *SupportTicketSolutionRow) Scan(row pgx.Row) error {
	return row.Scan(&sts.TicketUid, &sts.Description, &sts.EmailText, &sts.CreatedAt, &sts.UpdatedAt)
}

func (sts *SupportTicketSolutionRow) ColumnsForUpdate() []string {
	return []string{"description", "email_text", "updated_at"}
}

func (sts *SupportTicketSolutionRow) ValuesForUpdate() []interface{} {
	return []interface{}{
		sts.Description, sts.EmailText, sts.CreatedAt, sts.UpdatedAt,
	}
}

func (sts *SupportTicketSolutionRow) ConditionTicketUidEqual() sq.Eq {
	return sq.Eq{"ticket_uid": sts.TicketUid}
}

func NewSupportTicketSolutionRows() *Rows[*SupportTicketSolutionRow, entity.SupportTicketSolution] {
	return &Rows[*SupportTicketSolutionRow, entity.SupportTicketSolution]{}
}
