package robhancdklibawsaps


// Configuration for threshold settings that determine when values near expected values should be ignored during anomaly detection.
type IgnoreNearExpected struct {
	// The absolute amount by which values can differ from expected values before being considered anomalous.
	Amount *float64 `field:"optional" json:"amount" yaml:"amount"`
	// The ratio by which values can differ from expected values before being considered anomalous.
	Ratio *float64 `field:"optional" json:"ratio" yaml:"ratio"`
}

