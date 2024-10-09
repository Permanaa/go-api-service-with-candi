package domain

import (
	shareddomain "go-api-service-with-candi/pkg/shared/domain"
	"time"

	"github.com/golangid/candi/candishared"
)

// ResponseAuthList model
type ResponseAuthList struct {
	Meta candishared.Meta `json:"meta"`
	Data []ResponseAuth   `json:"data"`
}

// ResponseAuth model
type ResponseAuth struct {
	ID        int    `json:"id"`
	Field     string `json:"field"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseAuth) Serialize(source *shareddomain.Auth) {
	r.ID = source.ID
	r.Field = source.Field
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}

type LoginResponse struct {
	AccessToken  string `json:"acccessToken"`
	RefreshToken string `json:"refreshToken"`
}
