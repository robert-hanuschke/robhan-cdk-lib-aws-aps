package robhancdklibawsaps


// Configuration for the Random Cut Forest algorithm used for anomaly detection in time-series data.
type RandomCutForestConfiguration struct {
	// The Prometheus query used to retrieve the time-series data for anomaly detection.
	//
	// 1 to 8192 characters length.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Configuration for ignoring values that are near expected values from above during anomaly detection.
	IgnoreNearExpectedFromAbove *IgnoreNearExpected `field:"optional" json:"ignoreNearExpectedFromAbove" yaml:"ignoreNearExpectedFromAbove"`
	// Configuration for ignoring values that are near expected values from below during anomaly detection.
	IgnoreNearExpectedFromBelow *IgnoreNearExpected `field:"optional" json:"ignoreNearExpectedFromBelow" yaml:"ignoreNearExpectedFromBelow"`
	// The number of data points sampled from the input stream for the Random Cut Forest algorithm.
	//
	// The default number is 256 consecutive data points.
	//
	// Minimum: 256
	// Maximum: 1024.
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
	// The number of consecutive data points used to create a shingle for the Random Cut Forest algorithm.
	//
	// The default number is 8 consecutive data points.
	//
	// Minimum: 2
	// Maximum: 1024.
	ShingleSize *float64 `field:"optional" json:"shingleSize" yaml:"shingleSize"`
}

