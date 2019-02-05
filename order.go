package taxjar

import (
	"encoding/json"
	"strconv"
)

type Order struct {
	orderParams
}

type OrderList struct {
	Order Order `json:"order" bson:"order"`
}

type OrderDetails struct {
	ID        string     `json:"transaction_id" bson:"transaction_id"`
	Date      string     `json:"transaction_date" bson:"transaction_date"`
	Amount    float64    `json:"amount" bson:"amount"`
	Shipping  float64    `json:"shipping" bson:"shipping"`
	SalesTax  float64    `json:"sales_tax" bson:"sales_tax"`
	LineItems []LineItem `json:"line_items,omitempty" bson:"line_items,omitempty"`
}

type orderParams struct {
	ID          string      `json:"transaction_id" bson:"transaction_id"`
	Date        string      `json:"transaction_date" bson:"transaction_date"`
	Amount      json.Number `json:"amount,Number" bson:"amount"`
	Shipping    json.Number `json:"shipping,Number" bson:"shipping"`
	SalesTax    json.Number `json:"sales_tax,Number" bson:"sales_tax"`
	FromCountry string      `json:"from_country" bson:"from_country"`
	FromZip     string      `json:"from_zip" bson:"from_zip"`
	FromState   string      `json:"from_state,omitempty" bson:"from_state,omitempty"`
	FromCity    string      `json:"from_city,omitempty" bson:"from_city,omitempty"`
	FromStreet  string      `json:"from_street,omitempty" bson:"from_street,omitempty"`
	ToCountry   string      `json:"to_country,omitempty" bson:"to_country,omitempty"`
	ToZip       string      `json:"to_zip" bson:"to_zip"`
	ToState     string      `json:"to_state" bson:"to_state"`
	ToStreet    string      `json:"to_street,omitempty" bson:"to_street,omitempty"`
	ToCity      string      `json:"to_city,omitempty" bson:"to_city,omitempty"`
	LineItems   []LineItem  `json:"line_items,omitempty" bson:"line_items,omitempty"`
}

type OrderService struct {
	Repository OrderRepository
}

func (s *OrderService) Create(details OrderDetails, from, to Address) (Order, error) {
	return s.Repository.get(orderParams{
		ID:          details.ID,
		Date:        details.Date,
		Amount:      json.Number(strconv.FormatFloat(details.Amount, 'f', -1, 64)),
		Shipping:    json.Number(strconv.FormatFloat(details.Shipping, 'f', -1, 64)),
		SalesTax:    json.Number(strconv.FormatFloat(details.SalesTax, 'f', -1, 64)),
		LineItems:   details.LineItems,
		FromStreet:  from.Street,
		FromCity:    from.City,
		FromState:   from.State,
		FromZip:     from.Zip,
		FromCountry: from.Country,
		ToStreet:    to.Street,
		ToCity:      to.City,
		ToState:     to.State,
		ToZip:       to.Zip,
		ToCountry:   to.Country,
	})
}
