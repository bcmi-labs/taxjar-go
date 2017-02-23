package taxjar

import (
	"encoding/json"
)

// OrderRepository defines the interface for working with Order through the API.
type OrderRepository interface {
	get(orderParams) (Order, error)
}

// OrderApi implements OrderRepository
type OrderApi struct {
	client *Client
}

func (api OrderApi) get(params orderParams) (Order, error) {
	orderList := OrderList{}
	data, err := api.client.Post("/transactions/orders", params)
	if err != nil {
		return orderList.Order, err
	}
	err = json.Unmarshal(data, &orderList)
	return orderList.Order, err
}
