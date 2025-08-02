package robhancdklibawsaps


// The logging destination in an Amazon Managed Service for Prometheus workspace.
type LoggingDestination struct {
	// Configuration details for logging to CloudWatch Logs.
	CloudWatchLogs *CloudWatchLogDestination `field:"required" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
	// Filtering criteria that determine which queries are logged.
	Filters *LoggingFilter `field:"required" json:"filters" yaml:"filters"`
}

