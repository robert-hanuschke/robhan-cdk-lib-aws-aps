package robhancdklibawsaps


// Random Cut Forest configuration for anomaly detection.
type RandomCutForestConfiguration struct {
	// The PromQL query evaluated by the detector.
	//
	// The service requires the query to be wrapped by one of these aggregation operators: `avg`,
	// `count`, `group`, `max`, `min`, `quantile`, `stddev`, `stdvar`, or `sum`. Query syntax and
	// aggregation semantics are validated by Amazon Managed Service for Prometheus.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Values above the expected value that should not be treated as anomalies.
	// Default: - no upper tolerance.
	//
	IgnoreNearExpectedFromAbove IgnoreNearExpected `field:"optional" json:"ignoreNearExpectedFromAbove" yaml:"ignoreNearExpectedFromAbove"`
	// Values below the expected value that should not be treated as anomalies.
	// Default: - no lower tolerance.
	//
	IgnoreNearExpectedFromBelow IgnoreNearExpected `field:"optional" json:"ignoreNearExpectedFromBelow" yaml:"ignoreNearExpectedFromBelow"`
	// The number of data points sampled by the algorithm.
	// Default: 256.
	//
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
	// The number of consecutive data points in each shingle.
	// Default: 8.
	//
	ShingleSize *float64 `field:"optional" json:"shingleSize" yaml:"shingleSize"`
}

