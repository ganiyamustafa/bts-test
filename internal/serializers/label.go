package serializers

import (
	"time"

	"github.com/ganiyamustafa/bts/internal/models"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type GetLabelResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u GetLabelResponse) FromModel(m *models.Label) *GetLabelResponse {
	var res GetLabelResponse
	copier.CopyWithOption(&res, &m, copier.Option{IgnoreEmpty: true})
	return &res
}
