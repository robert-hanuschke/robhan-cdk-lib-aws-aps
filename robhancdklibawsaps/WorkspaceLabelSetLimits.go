package robhancdklibawsaps


// The limits that apply to time series matching a workspace label set.
type WorkspaceLabelSetLimits struct {
	// The maximum number of active series that can be ingested for the label set.
	//
	// A value of `0` disables the limit while still publishing label-set metrics.
	// Default: 0.
	//
	MaxSeries *float64 `field:"optional" json:"maxSeries" yaml:"maxSeries"`
}

