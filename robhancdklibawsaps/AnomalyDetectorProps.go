package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsaps"
)

// Properties for an Amazon Managed Service for Prometheus anomaly detector.
type AnomalyDetectorProps struct {
	// A user-friendly name for the detector.
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The algorithm configuration.
	Configuration *AnomalyDetectorConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The workspace whose metrics are evaluated.
	Workspace interfacesawsaps.IWorkspaceRef `field:"required" json:"workspace" yaml:"workspace"`
	// The interval, in seconds, between evaluations.
	//
	// Literal values must be integers from 30 through 86,400. Unresolved values are validated by
	// the service after deployment.
	// Default: 60.
	//
	EvaluationIntervalInSeconds *float64 `field:"optional" json:"evaluationIntervalInSeconds" yaml:"evaluationIntervalInSeconds"`
	// Metric labels associated with the detector.
	// Default: - no labels.
	//
	Labels *[]*AnomalyDetectorLabel `field:"optional" json:"labels" yaml:"labels"`
	// The action to take when evaluation data is missing.
	// Default: - service default.
	//
	MissingDataAction MissingDataAction `field:"optional" json:"missingDataAction" yaml:"missingDataAction"`
	// Tags to apply to the anomaly detector.
	// Default: - no tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

