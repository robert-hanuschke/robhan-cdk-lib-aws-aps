package robhancdklibawsaps


// A metric label associated with an anomaly detector.
type AnomalyDetectorLabel struct {
	// The label key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The label value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

