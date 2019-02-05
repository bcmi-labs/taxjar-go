package taxjar

import "encoding/json"

type Address struct {
	Street  string `json:"street,omitempty" bson:"street,omitempty"`
	City    string `json:"city,omitempty" bson:"city,omitempty"`
	State   string `json:"state,omitempty" bson:"state,omitempty"`
	Zip     string `json:"zip,omitempty" bson:"zip,omitempty"`
	Country string `json:"country,omitempty" bson:"country,omitempty"`
}

type LineItem struct {
	ID                int         `json:"id,omitempty" bson:"id,omitempty"`
	Quantity          int64       `json:"quantity,omitempty" bson:"quantity,omitempty"`
	ProductTaxCode    string      `json:"product_tax_code,omitempty" bson:"product_tax_code,omitempty"`
	ProductIdentifier string      `json:"product_identifier,omitempty" bson:"product_identifier,omitempty"`
	Description       string      `json:"description,omitempty" bson:"description,omitempty"`
	UnitPrice         json.Number `json:"unit_price,omitempty,Number" bson:"unit_price,omitempty"`
	Discount          json.Number `json:"discount,omitempty,Number" bson:"discount,omitempty"`
	SalesTax          json.Number `json:"sales_tax,omitempty,Number" bson:"sales_tax,omitempty"`
}

type Shipping struct {
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

type Breakdown struct {
	Shipping  Shipping      `json:"shipping" bson:"shipping"`
	LineItems []TaxLineItem `json:"line_items" bson:"line_items"`

	TaxCollectable float64 `json:"tax_collectable" bson:"tax_collectable"`
	TaxableAmount  float64 `json:"taxable_amount" bson:"taxable_amount"`

	// For US transactions
	StateTaxableAmount    float64 `json:"state_taxable_amount" bson:"state_taxable_amount"`
	StateTaxRate          float64 `json:"state_tax_rate" bson:"state_tax_rate"`
	StateTaxCollectable   float64 `json:"state_tax_collectable" bson:"state_tax_collectable"`
	CountyTaxableAmount   float64 `json:"county_taxable_amount" bson:"county_taxable_amount"`
	CountyTaxRate         float64 `json:"county_tax_rate" bson:"county_tax_rate"`
	CountyTaxCollectable  float64 `json:"county_tax_collectable" bson:"county_tax_collectable"`
	CityTaxableAmount     float64 `json:"city_taxable_amount" bson:"city_taxable_amount"`
	CityTaxRate           float64 `json:"city_tax_rate" bson:"city_tax_rate"`
	CityTaxCollectable    float64 `json:"city_tax_collectable" bson:"city_tax_collectable"`
	SpecialTaxableAmount  float64 `json:"special_district_taxable_amount" bson:"special_district_taxable_amount"`
	SpecialTaxRate        float64 `json:"special_tax_rate" bson:"special_tax_rate"`
	SpecialTaxCollectable float64 `json:"special_district_tax_collectable" bson:"special_district_tax_collectable"`

	// For CA transactions
	GstTaxableAmount  float64 `json:"gst_taxable_amount" bson:"gst_taxable_amount"`
	GstTaxCollectable float64 `json:"gst" bson:"gst"`
	GstTaxRate        float64 `json:"gst_tax_rate" bson:"gst_tax_rate"`
	PstTaxableAmount  float64 `json:"pst_taxable_amount" bson:"pst_taxable_amount"`
	PstTaxCollectable float64 `json:"pst" bson:"pst"`
	PstTaxRate        float64 `json:"pst_tax_rate" bson:"pst_tax_rate"`
	QstTaxableAmount  float64 `json:"qst_taxable_amount" bson:"qst_taxable_amount"`
	QstTaxCollectable float64 `json:"qst" bson:"qst"`
	QstTaxRate        float64 `json:"qst_tax_rate" bson:"qst_tax_rate"`
}
