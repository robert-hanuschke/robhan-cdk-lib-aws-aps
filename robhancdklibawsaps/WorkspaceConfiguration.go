package robhancdklibawsaps


// Use this structure to define label sets and the ingestion limits for time series that match label sets, and to specify the retention period of the workspace.
type WorkspaceConfiguration struct {
	// This is an array of structures, where each structure defines a label set for the workspace, and defines the ingestion limit for active time series for each of those label sets.
	//
	// Each
	// label name in a label set must be unique.
	//
	// Minimum 0.
	LimitsPerLabelSets *[]*LimitsPerLabelSet `field:"optional" json:"limitsPerLabelSets" yaml:"limitsPerLabelSets"`
	// Specifies how many days that metrics will be retained in the workspace.
	//
	// Minimum 1.
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
}

