package robhancdklibawsaps


// A label set and the ingestion limits that apply to matching time series.
type WorkspaceLabelSet struct {
	// The labels that identify matching time series.
	Labels *[]*WorkspaceLabel `field:"required" json:"labels" yaml:"labels"`
	// The limits for time series matching this label set.
	Limits *WorkspaceLabelSetLimits `field:"required" json:"limits" yaml:"limits"`
}

