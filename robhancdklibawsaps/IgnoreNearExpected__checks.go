//go:build !no_runtime_type_checking

package robhancdklibawsaps

import (
	"fmt"
)

func validateIgnoreNearExpected_AmountParameters(amount *float64) error {
	if amount == nil {
		return fmt.Errorf("parameter amount is required, but nil was provided")
	}

	return nil
}

func validateIgnoreNearExpected_RatioParameters(ratio *float64) error {
	if ratio == nil {
		return fmt.Errorf("parameter ratio is required, but nil was provided")
	}

	return nil
}

