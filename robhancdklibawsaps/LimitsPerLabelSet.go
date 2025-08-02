package robhancdklibawsaps


// This defines a label set for the workspace, and defines the ingestion limit for active time series that match that label set.
//
// Each label name in a label set must be unique.
type LimitsPerLabelSet struct {
	// This defines one label set that will have an enforced ingestion limit.
	//
	// You can set ingestion
	// limits on time series that match defined label sets, to help prevent a workspace from being
	// overwhelmed with unexpected spikes in time series ingestion.
	//
	// Label values accept all UTF-8 characters with one exception. If the label name is metric
	// name label __name__, then the metric part of the name must conform to the following pattern:
	// [a-zA-Z_:][a-zA-Z0-9_:]*
	//
	// Minimum 0.
	LabelSet *[]*Label `field:"required" json:"labelSet" yaml:"labelSet"`
	// This structure contains the information about the limits that apply to time series that match this label set.
	Limits *LimitsPerLabelSetEntry `field:"required" json:"limits" yaml:"limits"`
}

