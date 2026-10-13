package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/jsii"
)

// Threshold settings for values close to the expected value.
type IgnoreNearExpected interface {
}

// The jsii proxy struct for IgnoreNearExpected
type jsiiProxy_IgnoreNearExpected struct {
	_ byte // padding
}

func NewIgnoreNearExpected_Override(i IgnoreNearExpected) {
	_init_.Initialize()

	_jsii_.Create(
		"@robhan-cdk-lib/aws_aps.IgnoreNearExpected",
		nil, // no parameters
		i,
	)
}

// Ignore values within a non-negative absolute amount of the expected value.
func IgnoreNearExpected_Amount(amount *float64) IgnoreNearExpected {
	_init_.Initialize()

	if err := validateIgnoreNearExpected_AmountParameters(amount); err != nil {
		panic(err)
	}
	var returns IgnoreNearExpected

	_jsii_.StaticInvoke(
		"@robhan-cdk-lib/aws_aps.IgnoreNearExpected",
		"amount",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

// Ignore values within a non-negative ratio of the expected value.
func IgnoreNearExpected_Ratio(ratio *float64) IgnoreNearExpected {
	_init_.Initialize()

	if err := validateIgnoreNearExpected_RatioParameters(ratio); err != nil {
		panic(err)
	}
	var returns IgnoreNearExpected

	_jsii_.StaticInvoke(
		"@robhan-cdk-lib/aws_aps.IgnoreNearExpected",
		"ratio",
		[]interface{}{ratio},
		&returns,
	)

	return returns
}

