//go:build !no_runtime_type_checking

package robhancdklibawsaps

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsaps"
)

func validateScraperDestination_AmpParameters(workspace interfacesawsaps.IWorkspaceRef) error {
	if workspace == nil {
		return fmt.Errorf("parameter workspace is required, but nil was provided")
	}

	return nil
}

func validateScraperDestination_CloudWatchParameters(datasetArn *string) error {
	if datasetArn == nil {
		return fmt.Errorf("parameter datasetArn is required, but nil was provided")
	}

	return nil
}

