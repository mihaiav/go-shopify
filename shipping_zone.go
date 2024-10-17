package goshopify

import (
	"context"
)

// ShippingZoneService is an interface for interfacing with the shipping zones endpoint
// of the Shopify API.
// See: https://help.shopify.com/api/reference/store-properties/shippingzone
type ShippingZoneService interface {
	List(context.Context) ([]ShippingZone, error)
}

// ShippingZoneServiceOp handles communication with the shipping zone related methods
// of the Shopify API.
type ShippingZoneServiceOp struct {
	client *Client
}

// ShippingZone represents a Shopify shipping zone
type ShippingZone struct {
	Id                           uint64                        `json:"id,omitempty"`
	Name                         string                        `json:"name,omitempty"`
	ProfileId                    string                        `json:"profile_id,omitempty"`
	LocationGroupId              string                        `json:"location_group_id,omitempty"`
	AdminGraphqlApiId            string                        `json:"admin_graphql_api_id,omitempty"`
	Countries                    []ShippingCountry             `json:"countries,omitempty"`
	WeightBasedShippingRates     []WeightBasedShippingRate     `json:"weight_based_shipping_rates,omitempty"`
	PriceBasedShippingRates      []PriceBasedShippingRate      `json:"price_based_shipping_rates,omitempty"`
	CarrierShippingRateProviders []CarrierShippingRateProvider `json:"carrier_shipping_rate_providers,omitempty"`
}

// ShippingCountry represents a Shopify shipping country
type ShippingCountry struct {
	Id             uint64             `json:"id,omitempty"`
	ShippingZoneId uint64             `json:"shipping_zone_id,omitempty"`
	Name           string             `json:"name,omitempty"`
	Tax            *Decimal           `json:"tax,omitempty"`
	Code           string             `json:"code,omitempty"`
	TaxName        string             `json:"tax_name,omitempty"`
	Provinces      []ShippingProvince `json:"provinces,omitempty"`
}

// ShippingProvince represents a Shopify shipping province
type ShippingProvince struct {
	Id             uint64   `json:"id,omitempty"`
	CountryId      uint64   `json:"country_id,omitempty"`
	ShippingZoneId uint64   `json:"shipping_zone_id,omitempty"`
	Name           string   `json:"name,omitempty"`
	Code           string   `json:"code,omitempty"`
	Tax            *Decimal `json:"tax,omitempty"`
	TaxName        string   `json:"tax_name,omitempty"`
	TaxType        string   `json:"tax_type,omitempty"`
	TaxPercentage  *Decimal `json:"tax_percentage,omitempty"`
}

// WeightBasedShippingRate represents a Shopify weight-constrained shipping rate
type WeightBasedShippingRate struct {
	Id             uint64   `json:"id,omitempty"`
	ShippingZoneId uint64   `json:"shipping_zone_id,omitempty"`
	Name           string   `json:"name,omitempty"`
	Price          *Decimal `json:"price,omitempty"`
	WeightLow      *Decimal `json:"weight_low,omitempty"`
	WeightHigh     *Decimal `json:"weight_high,omitempty"`
}

// PriceBasedShippingRate represents a Shopify subtotal-constrained shipping rate
type PriceBasedShippingRate struct {
	Id               uint64   `json:"id,omitempty"`
	ShippingZoneId   uint64   `json:"shipping_zone_id,omitempty"`
	Name             string   `json:"name,omitempty"`
	Price            *Decimal `json:"price,omitempty"`
	MinOrderSubtotal *Decimal `json:"min_order_subtotal,omitempty"`
	MaxOrderSubtotal *Decimal `json:"max_order_subtotal,omitempty"`
}

// CarrierShippingRateProvider represents a Shopify carrier-constrained shipping rate
type CarrierShippingRateProvider struct {
	Id               uint64            `json:"id,omitempty"`
	CarrierServiceId uint64            `json:"carrier_service_id,omitempty"`
	ShippingZoneId   uint64            `json:"shipping_zone_id,omitempty"`
	FlatModifier     *Decimal          `json:"flat_modifier,omitempty"`
	PercentModifier  *Decimal          `json:"percent_modifier,omitempty"`
	ServiceFilter    map[string]string `json:"service_filter,omitempty"`
}

// Represents the result from the shipping_zones.json endpoint
type ShippingZonesResource struct {
	ShippingZones []ShippingZone `json:"shipping_zones"`
}

// List shipping zones
func (s *ShippingZoneServiceOp) List(ctx context.Context) ([]ShippingZone, error) {
	resource := new(ShippingZonesResource)
	err := s.client.Get(ctx, "shipping_zones.json", resource, nil)
	return resource.ShippingZones, err
}
