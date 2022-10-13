package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog/log"
)

type Client struct {
	Id            int64 `json:"id"`
	Menu          Menu
	OrderResponse OrderResponse
}

var NrOfClients int64

var clientId int64

func NewClient() *Client {
	c := &Client{}

	r, err := http.Get(cfg.FoodOrderingUrl + "/menu")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get menu")
	}

	var menu Menu
	err = json.NewDecoder(r.Body).Decode(&menu)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to decode menu")
	}

	c.Id = atomic.AddInt64(&clientId, 1)
	c.Menu = menu

	log.Debug().Int64("client_id", c.Id).Msg("Client received menu")

	order := c.generateOrder()

	log.Debug().Int64("client_id", c.Id).Msg("Client generated order")

	jsonBody, err := json.Marshal(order)
	if err != nil {
		log.Fatal().Err(err).Msg("Error marshalling order")
	}
	contentType := "application/json"

	r, err = http.Post(cfg.FoodOrderingUrl+"/order", contentType, bytes.NewReader(jsonBody))
	if err != nil {
		log.Fatal().Err(err).Msg("Error sending order to restaurant")
	}

	var orderResponse OrderResponse
	err = json.NewDecoder(r.Body).Decode(&orderResponse)
	if err != nil {
		log.Fatal().Err(err).Msg("Error decoding order response")
	}

	c.OrderResponse = orderResponse

	log.Debug().Int64("client_id", c.Id).Msg("Client sent order to food ordering")

	return c
}

func (c *Client) Run() {
	log.Info().Int64("client_id", c.Id).Msg("Client waits for order")

	var wg sync.WaitGroup

	for _, order := range c.OrderResponse.Orders {
		wg.Add(1)
		go c.waitForOrder(order, &wg)
	}

	wg.Wait()

	log.Info().Int64("client_id", c.Id).Msg("Client picked up all orders")
	atomic.AddInt64(&NrOfClients, -1)
}

func (c *Client) waitForOrder(order OrderResponseData, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(order.EstimatedWait*float64(cfg.TimeUnit)) * time.Millisecond)

	isOrderPickedUp := false

	for !isOrderPickedUp {
		r, err := http.Get(order.RestaurantAddress + "/v2/order/" + fmt.Sprintf("%d", order.OrderId))
		if err != nil {
			log.Fatal().Err(err).Msg("Error getting distribution from restaurant")
		}

		var distribution DistributionResponse
		err = json.NewDecoder(r.Body).Decode(&distribution)
		if err != nil {
			log.Fatal().Err(err).Msg("Error decoding distribution")
		}

		if distribution.IsReady {
			log.Info().Int64("client_id", c.Id).Int("order_id", order.OrderId).Msg("Client picked up order")
			isOrderPickedUp = true
		} else {
			log.Debug().Int64("client_id", c.Id).Int("order_id", order.OrderId).Msg("Order not ready yet")
			time.Sleep(time.Duration(15*cfg.TimeUnit) * time.Millisecond)
		}
	}

	wg.Done()
}

func (c *Client) generateOrder() Order {
	order := Order{
		ClientId: c.Id,
		Orders:   make([]OrderData, 0),
	}

	for _, restaurantData := range c.Menu.RestaurantsData {

		foodCount := rand.Intn(cfg.MaxOrderItemsCount) + 1

		orderData := OrderData{
			RestaurantId: restaurantData.RestaurantId,
			Items:        make([]int, foodCount),
		}

		orderData.Priority = (cfg.MaxOrderItemsCount - foodCount) / (cfg.MaxOrderItemsCount / 5)

		maxTime := 0
		for i := 0; i < foodCount; i++ {
			orderData.Items[i] = rand.Intn(restaurantData.MenuItems) + 1
			prepTime := restaurantData.Menu[i].PreparationTime
			if prepTime > maxTime {
				maxTime = prepTime
			}
		}

		orderData.MaxWait = float64(maxTime) * cfg.MaxWaitTimeCoefficient
		orderData.CreatedTime = time.Now().UnixMilli()

		order.Orders = append(order.Orders, orderData)
	}

	return order
}
