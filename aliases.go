// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package m3ter

import (
	"github.com/m3ter-com/m3ter-sdk-go/internal/apierror"
	"github.com/m3ter-com/m3ter-sdk-go/shared"
)

type Error = apierror.Error

// An array of currency conversion rates from Bill currency to Organization
// currency. For example, if Account is billed in GBP and Organization is set to
// USD, Bill line items are calculated in GBP and then converted to USD using the
// defined rate.
//
// This is an alias to an internal type.
type CurrencyConversion = shared.CurrencyConversion

// An array of currency conversion rates from Bill currency to Organization
// currency. For example, if Account is billed in GBP and Organization is set to
// USD, Bill line items are calculated in GBP and then converted to USD using the
// defined rate.
//
// This is an alias to an internal type.
type CurrencyConversionParam = shared.CurrencyConversionParam

// This is an alias to an internal type.
type PricingBand = shared.PricingBand

// This is an alias to an internal type.
type PricingBandParam = shared.PricingBandParam

// This is an alias to an internal type.
type SetString = shared.SetString
