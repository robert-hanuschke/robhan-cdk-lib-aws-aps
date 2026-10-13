package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsaps"
)

// A destination for metrics collected by a scraper.
type ScraperDestination interface {
}

// The jsii proxy struct for ScraperDestination
type jsiiProxy_ScraperDestination struct {
	_ byte // padding
}

func NewScraperDestination_Override(s ScraperDestination) {
	_init_.Initialize()

	_jsii_.Create(
		"@robhan-cdk-lib/aws_aps.ScraperDestination",
		nil, // no parameters
		s,
	)
}

// Send collected metrics from an Amazon EKS source to an Amazon Managed Service for Prometheus workspace.
//
// VPC sources must use {@link cloudWatch} instead.
func ScraperDestination_Amp(workspace interfacesawsaps.IWorkspaceRef) ScraperDestination {
	_init_.Initialize()

	if err := validateScraperDestination_AmpParameters(workspace); err != nil {
		panic(err)
	}
	var returns ScraperDestination

	_jsii_.StaticInvoke(
		"@robhan-cdk-lib/aws_aps.ScraperDestination",
		"amp",
		[]interface{}{workspace},
		&returns,
	)

	return returns
}

// Send collected metrics from an EKS or VPC source to a CloudWatch dataset.
func ScraperDestination_CloudWatch(datasetArn *string) ScraperDestination {
	_init_.Initialize()

	if err := validateScraperDestination_CloudWatchParameters(datasetArn); err != nil {
		panic(err)
	}
	var returns ScraperDestination

	_jsii_.StaticInvoke(
		"@robhan-cdk-lib/aws_aps.ScraperDestination",
		"cloudWatch",
		[]interface{}{datasetArn},
		&returns,
	)

	return returns
}

