package supportSubrouterTopics

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	httpEntity "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/entity"
	httpUtils "github.com/uxsnap/fresh_market_shop/backend/internal/delivery/http/utils"
)

func (h SupportSubrouterTopics) GetTicketsTopicByUid(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	topicUid, err := uuid.FromString(chi.URLParam(r, "uid"))
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusBadRequest, nil)
		return
	}

	topic, isFound, err := h.SupportService.GetTicketsTopicByUid(ctx, topicUid)
	if err != nil {
		httpUtils.WriteErrorResponse(w, http.StatusInternalServerError, err)
		return
	}
	if !isFound {
		httpUtils.WriteErrorResponse(w, http.StatusNotFound, nil)
		return
	}

	httpUtils.WriteResponseJson(w, httpEntity.ConvertSupportTicketsTopicFromEntity(topic))
}
