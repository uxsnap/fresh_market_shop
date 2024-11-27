package useCaseSupport

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func validateTicket(ticket entity.SupportTicket) error {
	if uuid.Equal(ticket.TopicUid, uuid.UUID{}) {
		return errors.New("пустой uid темы обращения")
	}
	if len(ticket.FromEmail) == 0 {
		return errors.New("пустой email обращающегося")
	}
	if len(ticket.Title) == 0 {
		return errors.New("пустой заголовок обращения")
	}
	if len(ticket.Description) == 0 {
		return errors.New("пустое описание проблемы обращения")
	}
	return nil
}
