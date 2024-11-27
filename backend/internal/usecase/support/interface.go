package useCaseSupport

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type RepositorySupport interface {
	CreateSupportTicketsTopic(ctx context.Context, topic entity.SupportTicketsTopic) error
	UpdateSupportTicketsTopic(ctx context.Context, topic entity.SupportTicketsTopic) error
	DeleteSupportTicketsTopic(ctx context.Context, uid uuid.UUID) error
	GetSupportTicketsTopicByUid(ctx context.Context, uid uuid.UUID) (entity.SupportTicketsTopic, bool, error)
	GetAllSupportTicketsTopics(ctx context.Context) ([]entity.SupportTicketsTopic, error)

	CreateSupportTicket(ctx context.Context, ticket entity.SupportTicket) error
	UpdateSupportTicket(ctx context.Context, ticket entity.SupportTicket) error
	DeleteSupportTicket(ctx context.Context, uid uuid.UUID) error
	GetSupportTicketByUid(ctx context.Context, uid uuid.UUID) (entity.SupportTicket, bool, error)
	GetSupportTickets(ctx context.Context, qFilters entity.QueryFilters) ([]entity.SupportTicket, error)

	CreateSupportTicketCommentMessage(ctx context.Context, message entity.SupportTicketCommentMessage) error
	UpdateSupportTicketCommentMessage(ctx context.Context, message entity.SupportTicketCommentMessage) error
	GetSupportTicketCommentMessage(ctx context.Context, uid uuid.UUID) (entity.SupportTicketCommentMessage, bool, error)
	GetSupportTicketCommentMessages(ctx context.Context, ticketUid uuid.UUID, qFilters entity.QueryFilters) ([]entity.SupportTicketCommentMessage, error)

	CreateSupportTicketSolution(ctx context.Context, solution entity.SupportTicketSolution) error
	UpdateSupportTicketSolution(ctx context.Context, solution entity.SupportTicketSolution) error
	GetSupportTicketSolution(ctx context.Context, ticketUid uuid.UUID) (entity.SupportTicketSolution, bool, error)
	GetSupportTicketSolutionsByTopic(ctx context.Context, topicUid uuid.UUID, qFilters entity.QueryFilters) ([]entity.SupportTicketSolution, error)
	DeleteSupportTicketSolution(ctx context.Context, ticketUid uuid.UUID) error
}
