package robhancdklibawsaps


// A label is a name:value pair used to add context to ingested metrics.
//
// This structure defines the
// name and value for one label that is used in a label set. You can set ingestion limits on time
// series that match defined label sets, to help prevent a workspace from being overwhelmed with
// unexpected spikes in time series ingestion.
type Label struct {
	// The name for this label.
	//
	// Pattern: ^[a-zA-Z_][a-zA-Z0-9_]*$
	//
	// At least one character.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The value for this label.
	//
	// At least one character.
	Value *string `field:"required" json:"value" yaml:"value"`
}

