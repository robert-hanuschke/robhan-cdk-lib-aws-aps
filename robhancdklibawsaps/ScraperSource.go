package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/jsii"
)

// A source from which a scraper collects metrics.
type ScraperSource interface {
}

// The jsii proxy struct for ScraperSource
type jsiiProxy_ScraperSource struct {
	_ byte // padding
}

func NewScraperSource_Override(s ScraperSource) {
	_init_.Initialize()

	_jsii_.Create(
		"@robhan-cdk-lib/aws_aps.ScraperSource",
		nil, // no parameters
		s,
	)
}

// Collect metrics from an Amazon EKS cluster.
func ScraperSource_Eks(options *EksScraperSourceOptions) ScraperSource {
	_init_.Initialize()

	if err := validateScraperSource_EksParameters(options); err != nil {
		panic(err)
	}
	var returns ScraperSource

	_jsii_.StaticInvoke(
		"@robhan-cdk-lib/aws_aps.ScraperSource",
		"eks",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Collect metrics from resources reachable in a VPC.
//
// This source must be paired with {@link ScraperDestination.cloudWatch}.
func ScraperSource_Vpc(options *VpcScraperSourceOptions) ScraperSource {
	_init_.Initialize()

	if err := validateScraperSource_VpcParameters(options); err != nil {
		panic(err)
	}
	var returns ScraperSource

	_jsii_.StaticInvoke(
		"@robhan-cdk-lib/aws_aps.ScraperSource",
		"vpc",
		[]interface{}{options},
		&returns,
	)

	return returns
}

