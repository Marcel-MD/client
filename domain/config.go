package domain

type Config struct {
	TimeUnit    int `json:"time_unit"`
	NrOfClients int `json:"nr_of_clients"`

	MaxOrderItemsCount     int     `json:"max_order_items_count"`
	MaxWaitTimeCoefficient float64 `json:"max_wait_time_coefficient"`

	FoodOrderingUrl string `json:"food_ordering_url"`
	ClientPort      string `json:"client_port"`
}

var cfg Config = Config{
	TimeUnit:    250,
	NrOfClients: 5,

	MaxOrderItemsCount:     5,
	MaxWaitTimeCoefficient: 1.3,
	FoodOrderingUrl:        "http://food-ordering:8090",
	ClientPort:             "8091",
}

func SetConfig(c Config) {
	cfg = c
}
