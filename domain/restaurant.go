package domain

var Restaurants map[int]Restaurant = make(map[int]Restaurant)

type Menu struct {
	Restaurants     int          `json:"restaurants"`
	RestaurantsData []Restaurant `json:"restaurants_data"`
}

type Restaurant struct {
	RestaurantId int     `json:"restaurant_id"`
	Name         string  `json:"name"`
	Address      string  `json:"address"`
	MenuItems    int     `json:"menu_items"`
	Menu         []Food  `json:"menu"`
	Rating       float64 `json:"rating"`
}

type Food struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	PreparationTime  int    `json:"preparation_time"`
	Complexity       int    `json:"complexity"`
	CookingApparatus string `json:"cooking_apparatus"`
}
