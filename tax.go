package taxjar

type Tax struct {
	Breakdown        Breakdown `json:"breakdown" bson:"breakdown"`
	OrderTotalAmount float64   `json:"order_total_amount" bson:"order_total_amount"`
	Shipping         float64   `json:"shipping" bson:"shipping"`
	TaxableAmount    float64   `json:"taxable_amount" bson:"taxable_amount"`
	Rate             float64   `json:"rate" bson:"rate"`
	AmountToCollect  float64   `json:"amount_to_collect" bson:"amount_to_collect"`
	HasNexus         bool      `json:"has_nexus" bson:"has_nexus"`
	FreightTaxable   bool      `json:"freight_taxable" bson:"freight_taxable"`
	TaxSource        string    `json:"tax_source" bson:"tax_source"`
}

type TaxList struct {
	Tax Tax `json:"tax" bson:"tax"`
}

type taxParams struct {
	FromCountry    string     `json:"from_country" bson:"from_country"`
	FromZip        string     `json:"from_zip" bson:"from_zip"`
	FromState      string     `json:"from_state,omitempty" bson:"from_state,omitempty"`
	FromCity       string     `json:"from_city,omitempty" bson:"from_city,omitempty"`
	FromStreet     string     `json:"from_street,omitempty" bson:"from_street,omitempty"`
	ToCountry      string     `json:"to_country,omitempty" bson:"to_country,omitempty"`
	ToZip          string     `json:"to_zip" bson:"to_zip"`
	ToState        string     `json:"to_state" bson:"to_state"`
	ToStreet       string     `json:"to_street,omitempty" bson:"to_street,omitempty"`
	ToCity         string     `json:"to_city,omitempty" bson:"to_city,omitempty"`
	Shipping       float64    `json:"shipping" bson:"shipping"`
	Amount         float64    `json:"amount,omitempty" bson:"amount,omitempty"`
	LineItems      []LineItem `json:"line_items,omitempty" bson:"line_items,omitempty"`
	NexusAddresses []Address  `json:"nexus_addresses,omitempty" bson:"nexus_addresses,omitempty"`
}

type TaxService struct {
	Repository TaxRepository
}

// Calculate sales Tax for a given order
func (s *TaxService) Calculate(from, to Address, shipping, amount float64) (Tax, error) {
	return s.Repository.get(taxParams{
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
		Shipping:    shipping,
		Amount:      amount,
	})
}

func (s *TaxService) CalculateItems(from, to Address, nexuses []Address, shipping float64, items []LineItem) (Tax, error) {
	return s.Repository.get(taxParams{
		FromStreet:     from.Street,
		FromCity:       from.City,
		FromState:      from.State,
		FromZip:        from.Zip,
		FromCountry:    from.Country,
		ToStreet:       to.Street,
		ToCity:         to.City,
		ToState:        to.State,
		ToZip:          to.Zip,
		ToCountry:      to.Country,
		Shipping:       shipping,
		LineItems:      items,
		NexusAddresses: nexuses,
	})
}

type TaxLineItem struct {
	Id string `json:"id" bson:"id"`

	// For US transactions
	StateTaxableAmount   float64 `json:"state_taxable_amount" bson:"state_taxable_amount"`
	StateSalesTaxRate    float64 `json:"state_sales_tax_rate" bson:"state_sales_tax_rate"`
	StateAmount          float64 `json:"state_amount" bson:"state_amount"`
	CountyTaxableAmount  float64 `json:"county_taxable_amount" bson:"county_taxable_amount"`
	CountyTaxRate        float64 `json:"county_tax_rate" bson:"county_tax_rate"`
	CountyAmount         float64 `json:"county_amount" bson:"county_amount"`
	CityTaxableAmount    float64 `json:"city_taxable_amount" bson:"city_taxable_amount"`
	CityTaxRate          float64 `json:"city_tax_rate" bson:"city_tax_rate"`
	CityAmount           float64 `json:"city_amount" bson:"city_amount"`
	SpecialTaxableAmount float64 `json:"special_district_taxable_amount" bson:"special_district_taxable_amount"`
	SpecialTaxRate       float64 `json:"special_tax_rate" bson:"special_tax_rate"`
	SpecialAmount        float64 `json:"special_district_amount" bson:"special_district_amount"`

	// For CA transactions
	GstTaxableAmount float64 `json:"gst_taxable_amount" bson:"gst_taxable_amount"`
	GstTaxRate       float64 `json:"gst_tax_rate" bson:"gst_tax_rate"`
	GstAmount        float64 `json:"gst" bson:"gst"`
	PstTaxableAmount float64 `json:"pst_taxable_amount" bson:"pst_taxable_amount"`
	PstTaxRate       float64 `json:"pst_tax_rate" bson:"pst_tax_rate"`
	PstAmount        float64 `json:"pst" bson:"pst"`
	QstTaxableAmount float64 `json:"qst_taxable_amount" bson:"qst_taxable_amount"`
	QstTaxRate       float64 `json:"qst_tax_rate" bson:"qst_tax_rate"`
	QstAmount        float64 `json:"qst" bson:"qst"`
}
