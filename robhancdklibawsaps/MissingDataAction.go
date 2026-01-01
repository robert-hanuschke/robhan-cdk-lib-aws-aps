package robhancdklibawsaps


// Specifies the action to take when data is missing during anomaly detection evaluation.
type MissingDataAction struct {
	// Marks missing data points as anomalies.
	MarkAsAnomaly *bool `field:"optional" json:"markAsAnomaly" yaml:"markAsAnomaly"`
	// Skips evaluation when data is missing.
	Skip *bool `field:"optional" json:"skip" yaml:"skip"`
}

