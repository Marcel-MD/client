package domain

import (
	"time"

	"github.com/rs/zerolog/log"
)

type DistributionResponse struct {
	OrderId        int64           `json:"order_id"`
	IsReady        bool            `json:"is_ready"`
	EstimatedWait  float64         `json:"estimated_waiting_time"`
	Priority       int             `json:"priority"`
	MaxWait        float64         `json:"max_wait"`
	CreatedTime    int64           `json:"created_time"`
	RegisteredTime int64           `json:"registered_time"`
	PreparedTime   int64           `json:"prepared_time"`
	CookingTime    int             `json:"cooking_time"`
	CookingDetails []CookingDetail `json:"cooking_details"`
}

type CookingDetail struct {
	FoodId int `json:"food_id"`
	CookId int `json:"cook_id"`
}

func (dr DistributionResponse) CalculateRating() int {
	orderTime := float64((time.Now().UnixMilli() - dr.RegisteredTime) / int64(cfg.TimeUnit))
	maxWaitTime := dr.MaxWait

	log.Debug().Int64("order_id", dr.OrderId).Float64("order_time", orderTime).Float64("max_wait", maxWaitTime).Msg("Calculating rating")

	if orderTime < maxWaitTime {
		return 5
	}

	if orderTime < maxWaitTime*1.1 {
		return 4
	}

	if orderTime < maxWaitTime*1.2 {
		return 3
	}

	if orderTime < maxWaitTime*1.3 {
		return 2
	}

	if orderTime < maxWaitTime*1.4 {
		return 1
	}

	return 0
}
