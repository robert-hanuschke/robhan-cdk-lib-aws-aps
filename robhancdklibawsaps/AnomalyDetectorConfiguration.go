package robhancdklibawsaps


// Algorithm configuration for an anomaly detector.
type AnomalyDetectorConfiguration struct {
	// Random Cut Forest algorithm settings.
	RandomCutForest *RandomCutForestConfiguration `field:"required" json:"randomCutForest" yaml:"randomCutForest"`
}

