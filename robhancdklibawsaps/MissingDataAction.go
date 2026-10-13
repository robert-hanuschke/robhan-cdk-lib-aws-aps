package robhancdklibawsaps


// The action to take when evaluation data is missing.
type MissingDataAction string

const (
	// Treat the missing data point as anomalous.
	MissingDataAction_MARK_AS_ANOMALY MissingDataAction = "MARK_AS_ANOMALY"
	// Skip evaluation for the missing data point.
	MissingDataAction_SKIP MissingDataAction = "SKIP"
)

