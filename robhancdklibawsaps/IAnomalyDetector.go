package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/internal"
)

type IAnomalyDetector interface {
	awscdk.IResource
	// The user-friendly name of the anomaly detector. 1 to 128 characters length.
	//
	// Minimum length of 1. Maximum length of 64.
	// Pattern: [0-9A-Za-z][-.0-9A-Z_a-z]*
	Alias() *string
	// The Amazon Resource Name (ARN) of the anomaly detector.
	AnomalyDetectorArn() *string
	// The algorithm configuration of the anomaly detector.
	Configuration() *AnomalyDetectorConfiguration
	// The frequency, in seconds, at which the anomaly detector evaluates metrics.
	//
	// Minimum value of 30. Maximum value of 86400.
	EvaluationIntervalInSeconds() *float64
	// The Amazon Managed Service for Prometheus metric labels associated with the anomaly detector.
	//
	// Map Entries: Minimum number of 0 items. Maximum number of 140 items.
	// Key Length Constraints: Minimum length of 1. Maximum length of 7168.
	// Key Pattern: (?!__)[a-zA-Z_][a-zA-Z0-9_]*
	// Value Length Constraints: Minimum length of 1. Maximum length of 7168.
	Labels() *[]*Label
	// The action taken when data is missing during evaluation.
	MissingDataAction() *MissingDataAction
	// An Amazon Managed Service for Prometheus workspace is a logical and isolated Prometheus server dedicated to ingesting, storing, and querying your Prometheus-compatible metrics.
	Workspace() IWorkspace
}

// The jsii proxy for IAnomalyDetector
type jsiiProxy_IAnomalyDetector struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IAnomalyDetector) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyDetector) AnomalyDetectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anomalyDetectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyDetector) Configuration() *AnomalyDetectorConfiguration {
	var returns *AnomalyDetectorConfiguration
	_jsii_.Get(
		j,
		"configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyDetector) EvaluationIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"evaluationIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyDetector) Labels() *[]*Label {
	var returns *[]*Label
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyDetector) MissingDataAction() *MissingDataAction {
	var returns *MissingDataAction
	_jsii_.Get(
		j,
		"missingDataAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyDetector) Workspace() IWorkspace {
	var returns IWorkspace
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}

