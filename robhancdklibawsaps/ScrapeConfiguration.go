package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/jsii"
)

// A Prometheus scrape configuration.
type ScrapeConfiguration interface {
}

// The jsii proxy struct for ScrapeConfiguration
type jsiiProxy_ScrapeConfiguration struct {
	_ byte // padding
}

func NewScrapeConfiguration_Override(s ScrapeConfiguration) {
	_init_.Initialize()

	_jsii_.Create(
		"@robhan-cdk-lib/aws_aps.ScrapeConfiguration",
		nil, // no parameters
		s,
	)
}

// Create a scrape configuration from Prometheus YAML.
//
// The CloudFormation resource provider accepts raw YAML and serializes it to the service API's
// blob shape. Do not base64-encode this value.
func ScrapeConfiguration_FromString(configuration *string) ScrapeConfiguration {
	_init_.Initialize()

	if err := validateScrapeConfiguration_FromStringParameters(configuration); err != nil {
		panic(err)
	}
	var returns ScrapeConfiguration

	_jsii_.StaticInvoke(
		"@robhan-cdk-lib/aws_aps.ScrapeConfiguration",
		"fromString",
		[]interface{}{configuration},
		&returns,
	)

	return returns
}

