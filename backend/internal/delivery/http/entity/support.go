package httpEntity

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

type SupportTicketsTopic struct {
	Uid         uuid.UUID `json:"uid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func ConvertSupportTicketsTopicFromEntity(topic entity.SupportTicketsTopic) SupportTicketsTopic {
	return SupportTicketsTopic{
		Uid:         topic.Uid,
		Name:        topic.Name,
		Description: topic.Description,
	}
}

func ConvertSupportTicketsTopicToEntity(topic SupportTicketsTopic) entity.SupportTicketsTopic {
	return entity.SupportTicketsTopic{
		Uid:         topic.Uid,
		Name:        topic.Name,
		Description: topic.Description,
	}
}

type SupportTicket struct {
	Uid         uuid.UUID `json:"uid"`
	UserUid     uuid.UUID `json:"userUid"`
	TopicUid    uuid.UUID `json:"topicUid"`
	SolverUid   uuid.UUID `json:"solverUid"`
	FromEmail   string    `json:"fromEmail"`
	FromPhone   string    `json:"fromPhone"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ConvertSupportTicketFromEntity(ticket entity.SupportTicket) SupportTicket {
	return SupportTicket{
		Uid:         ticket.Uid,
		UserUid:     ticket.UserUid,
		TopicUid:    ticket.TopicUid,
		SolverUid:   ticket.SolverUid,
		FromEmail:   ticket.FromEmail,
		FromPhone:   ticket.FromPhone,
		Title:       ticket.Title,
		Description: ticket.Description,
		Status:      string(ticket.Status),
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
	}
}

func ConvertSupportTicketToEntity(ticket SupportTicket) entity.SupportTicket {
	return entity.SupportTicket{
		Uid:         ticket.Uid,
		UserUid:     ticket.UserUid,
		TopicUid:    ticket.TopicUid,
		SolverUid:   ticket.SolverUid,
		FromEmail:   ticket.FromEmail,
		FromPhone:   ticket.FromPhone,
		Title:       ticket.Title,
		Description: ticket.Description,
		Status:      entity.SupportTicketStatus(ticket.Status),
		CreatedAt:   ticket.CreatedAt,
		UpdatedAt:   ticket.UpdatedAt,
	}
}

type SupportTicketSolution struct {
	TicketUid   uuid.UUID `json:"ticketUid"`
	Description string    `json:"description"`
	EmailText   string    `json:"emailText"`
	IsSuccess   bool      `json:"isSuccess"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func ConvertSupportTicketSolutionFromEntity(solution entity.SupportTicketSolution) SupportTicketSolution {
	return SupportTicketSolution{
		TicketUid:   solution.TicketUid,
		Description: solution.Description,
		EmailText:   solution.EmailText,
		IsSuccess:   solution.IsSuccess,
		CreatedAt:   solution.CreatedAt,
		UpdatedAt:   solution.UpdatedAt,
	}
}

func ConvertSupportTicketSolutionToEntity(solution SupportTicketSolution) entity.SupportTicketSolution {
	return entity.SupportTicketSolution{
		TicketUid:   solution.TicketUid,
		Description: solution.Description,
		EmailText:   solution.EmailText,
		IsSuccess:   solution.IsSuccess,
		CreatedAt:   solution.CreatedAt,
		UpdatedAt:   solution.UpdatedAt,
	}
}

type SupportTicketCommentMessage struct {
	Uid                 uuid.UUID `json:"uid"`
	TicketUid           uuid.UUID `json:"ticketUid"`
	SenderUid           uuid.UUID `json:"senderUid"`
	Content             string    `json:"content"`
	IsImportedFromEmail bool      `json:"isImportedFromEmail"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
}

func ConvertSupportTicketCommentMessageFromEntity(message entity.SupportTicketCommentMessage) SupportTicketCommentMessage {
	return SupportTicketCommentMessage{
		Uid:                 message.Uid,
		TicketUid:           message.TicketUid,
		SenderUid:           message.SenderUid,
		Content:             message.Content,
		IsImportedFromEmail: message.IsImportedFromEmail,
		CreatedAt:           message.CreatedAt,
		UpdatedAt:           message.UpdatedAt,
	}
}

func ConvertSupportTicketCommentMessageToEntity(message SupportTicketCommentMessage) entity.SupportTicketCommentMessage {
	return entity.SupportTicketCommentMessage{
		Uid:                 message.Uid,
		TicketUid:           message.TicketUid,
		SenderUid:           message.SenderUid,
		Content:             message.Content,
		IsImportedFromEmail: message.IsImportedFromEmail,
		CreatedAt:           message.CreatedAt,
		UpdatedAt:           message.UpdatedAt,
	}
}
