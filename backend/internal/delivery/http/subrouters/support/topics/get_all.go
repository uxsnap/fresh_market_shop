package supportSubrouterTopics

import (
	"context"
	"net/http"

	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h SupportSubrouterTopics) GetAllTicketsTopics(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	topics, err := h.SupportService.GetAllTicketsTopics(ctx)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	resp := make([]httpEntity.SupportTicketsTopic, len(topics))
	for i := 0; i < len(topics); i++ {
		resp[i] = httpEntity.ConvertSupportTicketsTopicFromEntity(topics[i])
	}

	httpUtils.WriteResponseJson(w, resp)
}
