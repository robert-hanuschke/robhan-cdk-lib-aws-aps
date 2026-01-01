package robhancdklibawsaps


// The configuration for the anomaly detection algorithm.
type AnomalyDetectorConfiguration struct {
	// The Random Cut Forest algorithm configuration for anomaly detection.
	RandomCutForest *RandomCutForestConfiguration `field:"required" json:"randomCutForest" yaml:"randomCutForest"`
}

