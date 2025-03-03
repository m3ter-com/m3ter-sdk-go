// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/m3ter-com/m3ter-sdk-go/internal/apijson"
	"github.com/m3ter-com/m3ter-sdk-go/internal/param"
)

// An array of currency conversion rates from Bill currency to Organization
// currency. For example, if Account is billed in GBP and Organization is set to
// USD, Bill line items are calculated in GBP and then converted to USD using the
// defined rate.
type CurrencyConversion struct {
	// Currency to convert from. For example: GBP.
	From string `json:"from,required"`
	// Currency to convert to. For example: USD.
	To string `json:"to,required"`
	// Conversion rate between currencies.
	Multiplier float64                `json:"multiplier"`
	JSON       currencyConversionJSON `json:"-"`
}

// currencyConversionJSON contains the JSON metadata for the struct
// [CurrencyConversion]
type currencyConversionJSON struct {
	From        apijson.Field
	To          apijson.Field
	Multiplier  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CurrencyConversion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r currencyConversionJSON) RawJSON() string {
	return r.raw
}

// An array of currency conversion rates from Bill currency to Organization
// currency. For example, if Account is billed in GBP and Organization is set to
// USD, Bill line items are calculated in GBP and then converted to USD using the
// defined rate.
type CurrencyConversionParam struct {
	// Currency to convert from. For example: GBP.
	From param.Field[string] `json:"from,required"`
	// Currency to convert to. For example: USD.
	To param.Field[string] `json:"to,required"`
	// Conversion rate between currencies.
	Multiplier param.Field[float64] `json:"multiplier"`
}

func (r CurrencyConversionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PricingBand struct {
	// Fixed price charged for the Pricing band.
	FixedPrice float64 `json:"fixedPrice,required"`
	// Lower limit for the Pricing band.
	LowerLimit float64 `json:"lowerLimit,required"`
	// Unit price charged for the Pricing band.
	UnitPrice float64 `json:"unitPrice,required"`
	// The ID for the Pricing band.
	ID string `json:"id"`
	// **OBSOLETE - this is deprecated and no longer used.**
	CreditTypeID string          `json:"creditTypeId"`
	JSON         pricingBandJSON `json:"-"`
}

// pricingBandJSON contains the JSON metadata for the struct [PricingBand]
type pricingBandJSON struct {
	FixedPrice   apijson.Field
	LowerLimit   apijson.Field
	UnitPrice    apijson.Field
	ID           apijson.Field
	CreditTypeID apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PricingBand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pricingBandJSON) RawJSON() string {
	return r.raw
}

type PricingBandParam struct {
	// Fixed price charged for the Pricing band.
	FixedPrice param.Field[float64] `json:"fixedPrice,required"`
	// Lower limit for the Pricing band.
	LowerLimit param.Field[float64] `json:"lowerLimit,required"`
	// Unit price charged for the Pricing band.
	UnitPrice param.Field[float64] `json:"unitPrice,required"`
	// The ID for the Pricing band.
	ID param.Field[string] `json:"id"`
	// **OBSOLETE - this is deprecated and no longer used.**
	CreditTypeID param.Field[string] `json:"creditTypeId"`
}

func (r PricingBandParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SetString struct {
	Empty bool          `json:"empty"`
	JSON  setStringJSON `json:"-"`
}

// setStringJSON contains the JSON metadata for the struct [SetString]
type setStringJSON struct {
	Empty       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setStringJSON) RawJSON() string {
	return r.raw
}
