package robhancdklibawsaps


// This structure contains the limits that apply to time series that match one label set.
type LimitsPerLabelSetEntry struct {
	// The maximum number of active series that can be ingested that match this label set.
	//
	// Setting this to 0 causes no label set limit to be enforced, but it does cause Amazon Managed
	// Service for Prometheus to vend label set metrics to CloudWatch Logs.
	//
	// Minimum 0.
	MaxSeries *float64 `field:"optional" json:"maxSeries" yaml:"maxSeries"`
}

