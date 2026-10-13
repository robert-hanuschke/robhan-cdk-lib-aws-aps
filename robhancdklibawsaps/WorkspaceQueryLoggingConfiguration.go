package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// The query log destination and filtering criteria for a workspace.
type WorkspaceQueryLoggingConfiguration struct {
	// The CloudWatch log group to which query logs are published.
	LogGroup interfacesawslogs.ILogGroupRef `field:"required" json:"logGroup" yaml:"logGroup"`
	// The Query Samples Processed threshold above which queries are logged.
	//
	// When query logging is enabled, the default `0` makes every query that processes one or more
	// samples eligible for logging. Consider log volume, cost, and query-content sensitivity when
	// choosing this value.
	// Default: 0.
	//
	QspThreshold *float64 `field:"optional" json:"qspThreshold" yaml:"qspThreshold"`
}

