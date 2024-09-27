package clientAuthService

import (
	"context"
	"log"

	"github.com/uxsnap/fresh_market_shop/backend/internal/entity"
)

func (c *AuthClient) HealthCheck(ctx context.Context) entity.AuthServiceHealthCheck {
	log.Printf("authClient.HealthCheck")

	resp, err := c.client.HealthCheck(ctx, nil)
	if err != nil {
		return entity.AuthServiceHealthCheck{
			Status:  entity.HealthStatusNotAvailable,
			Message: err.Error(),
		}
	}
	return entity.AuthServiceHealthCheck{
		Status:    resp.GetStatus(),
		GitTag:    resp.GetGitTag(),
		GitBranch: resp.GetGitBranch(),
		UpTime:    resp.GetUpTime().AsTime(),
	}
}
