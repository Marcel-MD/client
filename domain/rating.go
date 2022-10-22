package domain

type Rating struct {
	ClientId int64         `json:"client_id"`
	OrderId  int64         `json:"order_id"`
	Orders   []OrderRating `json:"orders"`
}

type OrderRating struct {
	RestaurantId      int `json:"restaurant_id"`
	OrderId           int `json:"order_id"`
	Rating            int `json:"rating"`
	EstimatedWaitTime int `json:"estimated_waiting_time"`
	WaitTime          int `json:"waiting_time"`
}
