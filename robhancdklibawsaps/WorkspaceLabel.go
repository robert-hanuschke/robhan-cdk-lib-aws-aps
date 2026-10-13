package robhancdklibawsaps


// A label used to identify a set of ingested time series.
type WorkspaceLabel struct {
	// The label name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The label value.
	Value *string `field:"required" json:"value" yaml:"value"`
}

