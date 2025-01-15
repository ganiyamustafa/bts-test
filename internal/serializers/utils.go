package serializers

import (
	"math"
	"time"
)

type MetaResponse struct {
	Total       int64 `json:"total"`
	PerPage     int   `json:"perPage"`
	CurrentPage int   `json:"currentPage"`
	LastPage    int   `json:"lastPage"`
}

func (t *MetaResponse) GeneratePaginateData(limit, page int) {
	t.CurrentPage = page
	t.PerPage = limit
	t.LastPage = int(math.Ceil(float64(t.Total) / float64(limit)))
}

type BaseResponse struct {
	Status     string        `json:"status"`
	StatusCode int           `json:"statusCode"`
	Message    string        `json:"message"`
	Timestamp  time.Time     `json:"timestamp"`
	Data       any           `json:"data"`
	Meta       *MetaResponse `json:"meta"`
}
