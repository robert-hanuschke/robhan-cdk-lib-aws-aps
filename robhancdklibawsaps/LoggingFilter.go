package robhancdklibawsaps


// Filtering criteria that determine which queries are logged.
type LoggingFilter struct {
	// Integer.
	//
	// Minimum 0.
	QspThreshold *float64 `field:"required" json:"qspThreshold" yaml:"qspThreshold"`
}

