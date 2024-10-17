package goshopify

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

const recurringApplicationChargesBasePath = "recurring_application_charges"

// RecurringApplicationChargeService is an interface for interacting with the
// RecurringApplicationCharge endpoints of the Shopify API.
// See https://help.shopify.com/api/reference/billing/recurringapplicationcharge
type RecurringApplicationChargeService interface {
	Create(context.Context, RecurringApplicationCharge) (*RecurringApplicationCharge, error)
	Get(context.Context, uint64, interface{}) (*RecurringApplicationCharge, error)
	List(context.Context, interface{}) ([]RecurringApplicationCharge, error)
	Activate(context.Context, RecurringApplicationCharge) (*RecurringApplicationCharge, error)
	Delete(context.Context, uint64) error
	Update(context.Context, uint64, uint64) (*RecurringApplicationCharge, error)
}

// RecurringApplicationChargeServiceOp handles communication with the
// RecurringApplicationCharge related methods of the Shopify API.
type RecurringApplicationChargeServiceOp struct {
	client *Client
}

// RecurringApplicationCharge represents a Shopify RecurringApplicationCharge.
type RecurringApplicationCharge struct {
	APIClientId           uint64     `json:"api_client_id"`
	ActivatedOn           *time.Time `json:"activated_on"`
	BalanceRemaining      *Decimal   `json:"balance_remaining"`
	BalanceUsed           *Decimal   `json:"balance_used"`
	BillingOn             *time.Time `json:"billing_on"`
	CancelledOn           *time.Time `json:"cancelled_on"`
	CappedAmount          *Decimal   `json:"capped_amount"`
	ConfirmationURL       string     `json:"confirmation_url"`
	CreatedAt             *time.Time `json:"created_at"`
	DecoratedReturnURL    string     `json:"decorated_return_url"`
	Id                    uint64     `json:"id"`
	Name                  string     `json:"name"`
	Price                 *Decimal   `json:"price"`
	ReturnURL             string     `json:"return_url"`
	RiskLevel             *Decimal   `json:"risk_level"`
	Status                string     `json:"status"`
	Terms                 string     `json:"terms"`
	Test                  *bool      `json:"test"`
	TrialDays             int        `json:"trial_days"`
	TrialEndsOn           *time.Time `json:"trial_ends_on"`
	UpdateCappedAmountURL string     `json:"update_capped_amount_url"`
	UpdatedAt             *time.Time `json:"updated_at"`
	Currency              string     `json:"currency"`
}

func parse(dest **time.Time, data *string) error {
	if data == nil {
		return nil
	}
	// This is what API doc says: "2013-06-27T08:48:27-04:00"
	format := time.RFC3339
	if len(*data) == 10 {
		// This is how the date looks.
		format = "2006-01-02"
	}
	t, err := time.Parse(format, *data)
	if err != nil {
		return err
	}
	*dest = &t
	return nil
}

func (r *RecurringApplicationCharge) UnmarshalJSON(data []byte) error {
	// This is a workaround for the API returning incomplete results:
	// https://ecommerce.shopify.com/c/shopify-apis-and-technology/t/-523203
	// For a longer explanation of the hack check:
	// http://choly.ca/post/go-json-marshalling/
	type alias RecurringApplicationCharge
	aux := &struct {
		ActivatedOn *string `json:"activated_on"`
		BillingOn   *string `json:"billing_on"`
		CancelledOn *string `json:"cancelled_on"`
		CreatedAt   *string `json:"created_at"`
		TrialEndsOn *string `json:"trial_ends_on"`
		UpdatedAt   *string `json:"updated_at"`
		*alias
	}{alias: (*alias)(r)}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if err := parse(&r.ActivatedOn, aux.ActivatedOn); err != nil {
		return err
	}
	if err := parse(&r.BillingOn, aux.BillingOn); err != nil {
		return err
	}
	if err := parse(&r.CancelledOn, aux.CancelledOn); err != nil {
		return err
	}
	if err := parse(&r.CreatedAt, aux.CreatedAt); err != nil {
		return err
	}
	if err := parse(&r.TrialEndsOn, aux.TrialEndsOn); err != nil {
		return err
	}
	if err := parse(&r.UpdatedAt, aux.UpdatedAt); err != nil {
		return err
	}
	return nil
}

// RecurringApplicationChargeResource represents the result from the
// admin/recurring_application_charges{/X{/activate.json}.json}.json endpoints.
type RecurringApplicationChargeResource struct {
	Charge *RecurringApplicationCharge `json:"recurring_application_charge"`
}

// RecurringApplicationChargesResource represents the result from the
// admin/recurring_application_charges.json endpoint.
type RecurringApplicationChargesResource struct {
	Charges []RecurringApplicationCharge `json:"recurring_application_charges"`
}

// Create creates new recurring application charge.
func (r *RecurringApplicationChargeServiceOp) Create(ctx context.Context, charge RecurringApplicationCharge) (
	*RecurringApplicationCharge, error,
) {
	path := fmt.Sprintf("%s.json", recurringApplicationChargesBasePath)
	wrappedData := RecurringApplicationChargeResource{Charge: &charge}
	resource := &RecurringApplicationChargeResource{}
	err := r.client.Post(ctx, path, wrappedData, resource)
	return resource.Charge, err
}

// Get gets individual recurring application charge.
func (r *RecurringApplicationChargeServiceOp) Get(ctx context.Context, chargeId uint64, options interface{}) (
	*RecurringApplicationCharge, error,
) {
	path := fmt.Sprintf("%s/%d.json", recurringApplicationChargesBasePath, chargeId)
	resource := &RecurringApplicationChargeResource{}
	err := r.client.Get(ctx, path, resource, options)
	return resource.Charge, err
}

// List gets all recurring application charges.
func (r *RecurringApplicationChargeServiceOp) List(ctx context.Context, options interface{}) (
	[]RecurringApplicationCharge, error,
) {
	path := fmt.Sprintf("%s.json", recurringApplicationChargesBasePath)
	resource := &RecurringApplicationChargesResource{}
	err := r.client.Get(ctx, path, resource, options)
	return resource.Charges, err
}

// Activate activates recurring application charge.
func (r *RecurringApplicationChargeServiceOp) Activate(ctx context.Context, charge RecurringApplicationCharge) (
	*RecurringApplicationCharge, error,
) {
	path := fmt.Sprintf("%s/%d/activate.json", recurringApplicationChargesBasePath, charge.Id)
	wrappedData := RecurringApplicationChargeResource{Charge: &charge}
	resource := &RecurringApplicationChargeResource{}
	err := r.client.Post(ctx, path, wrappedData, resource)
	return resource.Charge, err
}

// Delete deletes recurring application charge.
func (r *RecurringApplicationChargeServiceOp) Delete(ctx context.Context, chargeId uint64) error {
	return r.client.Delete(ctx, fmt.Sprintf("%s/%d.json", recurringApplicationChargesBasePath, chargeId))
}

// Update updates recurring application charge.
func (r *RecurringApplicationChargeServiceOp) Update(ctx context.Context, chargeId, newCappedAmount uint64) (
	*RecurringApplicationCharge, error,
) {
	path := fmt.Sprintf("%s/%d/customize.json?recurring_application_charge[capped_amount]=%d",
		recurringApplicationChargesBasePath, chargeId, newCappedAmount)
	resource := &RecurringApplicationChargeResource{}
	err := r.client.Put(ctx, path, nil, resource)
	return resource.Charge, err
}
