package robhancdklibawsaps


// Workspace-level retention, query, and ingestion settings.
type WorkspaceConfiguration struct {
	// Label sets with individual active-series limits.
	// Default: - no label-set-specific limits.
	//
	LabelSets *[]*WorkspaceLabelSet `field:"optional" json:"labelSets" yaml:"labelSets"`
	// The time window, in seconds, during which out-of-order samples are accepted.
	// Default: - the service default.
	//
	OutOfOrderTimeWindowInSeconds *float64 `field:"optional" json:"outOfOrderTimeWindowInSeconds" yaml:"outOfOrderTimeWindowInSeconds"`
	// The number of days to retain metrics.
	// Default: - the service default.
	//
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
	// The duration, in seconds, by which rule evaluation queries are offset into the past.
	// Default: - the service default.
	//
	RuleQueryOffsetInSeconds *float64 `field:"optional" json:"ruleQueryOffsetInSeconds" yaml:"ruleQueryOffsetInSeconds"`
}

