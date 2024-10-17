package goshopify

import (
	"encoding/json"

	"github.com/shopspring/decimal"
)

type Decimal struct {
	Decimal *decimal.Decimal
}

func (d *Decimal) UnmarshalJSON(p []byte) error {
	// we wrap decimal.Decimal to handle cases where shopify
	// provides an empty string instead for empty price/decimals
	if string(p) == `""` {
		return nil
	}
	dc := &decimal.Decimal{}
	if err := dc.UnmarshalJSON(p); err != nil {
		return err
	}
	if dc.IsZero() {
		return nil
	}
	d.Decimal = dc
	return nil
}

func (d *Decimal) MarshalJSON() ([]byte, error) {
	if d == nil || d.Decimal == nil || d.Decimal.IsZero() {
		return []byte("null"), nil
	}
	return json.Marshal(d.Decimal)
}
