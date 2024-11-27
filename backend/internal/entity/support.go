package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type SupportTicketsTopic struct {
	Uid         uuid.UUID
	Name        string
	Description string
}

type SupportTicketStatus string

const (
	SupportTicketStatusCreated   SupportTicketStatus = "created"
	SupportTicketStatusInProcess SupportTicketStatus = "in_process"
	SupportTicketStatusSolved    SupportTicketStatus = "solved"
	SupportTicketStatusCantSolve SupportTicketStatus = "cant_solve"
)

type SupportTicket struct {
	Uid         uuid.UUID
	UserUid     uuid.UUID
	TopicUid    uuid.UUID
	SolverUid   uuid.UUID
	FromEmail   string
	FromPhone   string
	Title       string
	Description string
	Status      SupportTicketStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SupportTicketSolution struct {
	TicketUid   uuid.UUID
	Description string
	EmailText   string
	IsSuccess   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type SupportTicketCommentMessage struct {
	Uid                 uuid.UUID
	TicketUid           uuid.UUID
	SenderUid           uuid.UUID
	Content             string
	IsImportedFromEmail bool
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
