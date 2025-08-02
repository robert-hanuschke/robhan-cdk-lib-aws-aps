package robhancdklibawsaps


// The query logging configuration in an Amazon Managed Service for Prometheus workspace.
type QueryLoggingConfiguration struct {
	// Defines a destination and its associated filtering criteria for query logging.
	//
	// Minimum 1 and maximum 1 item in array.
	Destinations *[]*LoggingDestination `field:"required" json:"destinations" yaml:"destinations"`
}

