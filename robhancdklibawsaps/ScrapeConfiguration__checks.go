//go:build !no_runtime_type_checking

package robhancdklibawsaps

import (
	"fmt"
)

func validateScrapeConfiguration_FromStringParameters(configuration *string) error {
	if configuration == nil {
		return fmt.Errorf("parameter configuration is required, but nil was provided")
	}

	return nil
}

