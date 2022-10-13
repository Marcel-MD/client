package domain

type Order struct {
	OrderId  int64       `json:"order_id"`
	ClientId int64       `json:"client_id"`
	Orders   []OrderData `json:"orders"`
}

type OrderData struct {
	RestaurantId int     `json:"restaurant_id"`
	Items        []int   `json:"items"`
	Priority     int     `json:"priority"`
	MaxWait      float64 `json:"max_wait"`
	CreatedTime  int64   `json:"created_time"`
}

type OrderResponse struct {
	OrderId int64               `json:"order_id"`
	Orders  []OrderResponseData `json:"orders"`
}

type OrderResponseData struct {
	OrderId           int     `json:"order_id"`
	RestaurantId      int     `json:"restaurant_id"`
	RestaurantAddress string  `json:"restaurant_address"`
	EstimatedWait     float64 `json:"estimated_waiting_time"`
	CreatedTime       int64   `json:"created_time"`
	RegisteredTime    int64   `json:"registered_time"`
}
