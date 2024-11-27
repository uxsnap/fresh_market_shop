package useCaseSupport

import (
	"errors"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func validateTicketsTopic(topic entity.SupportTicketsTopic) error {
	if len(topic.Name) == 0 {
		return errors.New("не указано название темы обращения")
	}
	if len(topic.Description) == 0 {
		return errors.New("не указано описание темы обращения")
	}
	return nil
}
