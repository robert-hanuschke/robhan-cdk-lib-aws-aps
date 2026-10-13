//go:build no_runtime_type_checking

package robhancdklibawsaps

// Building without runtime type checking enabled, so all the below just return nil

func validateScraperDestination_AmpParameters(workspace interfacesawsaps.IWorkspaceRef) error {
	return nil
}

func validateScraperDestination_CloudWatchParameters(datasetArn *string) error {
	return nil
}

