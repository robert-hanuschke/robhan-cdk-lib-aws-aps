//go:build no_runtime_type_checking

package robhancdklibawsaps

// Building without runtime type checking enabled, so all the below just return nil

func validateScraperSource_EksParameters(options *EksScraperSourceOptions) error {
	return nil
}

func validateScraperSource_VpcParameters(options *VpcScraperSourceOptions) error {
	return nil
}

