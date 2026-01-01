package robhancdklibawsaps


// Properties for creating an Amazon Managed Service for Prometheus Anomaly Detector.
type AnomalyDetectorProps struct {
	// The user-friendly name of the anomaly detector.
	//
	// 1 to 128 characters length.
	Alias *string `field:"required" json:"alias" yaml:"alias"`
	// The algorithm configuration of the anomaly detector.
	Configuration *AnomalyDetectorConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// An Amazon Managed Service for Prometheus workspace is a logical and isolated Prometheus server dedicated to ingesting, storing, and querying your Prometheus-compatible metrics.
	Workspace IWorkspace `field:"required" json:"workspace" yaml:"workspace"`
	// The frequency, in seconds, at which the anomaly detector evaluates metrics.
	//
	// Minimum value of 30. Maximum value of 86400.
	EvaluationIntervalInSeconds *float64 `field:"optional" json:"evaluationIntervalInSeconds" yaml:"evaluationIntervalInSeconds"`
	// The Amazon Managed Service for Prometheus metric labels associated with the anomaly detector.
	//
	// Map Entries: Minimum number of 0 items. Maximum number of 140 items.
	// Key Length Constraints: Minimum length of 1. Maximum length of 7168.
	// Key Pattern: (?!__)[a-zA-Z_][a-zA-Z0-9_]*
	// Value Length Constraints: Minimum length of 1. Maximum length of 7168.
	Labels *[]*Label `field:"optional" json:"labels" yaml:"labels"`
	// The action taken when data is missing during evaluation.
	MissingDataAction *MissingDataAction `field:"optional" json:"missingDataAction" yaml:"missingDataAction"`
}

