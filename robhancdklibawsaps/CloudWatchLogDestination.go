package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Configuration details for logging to CloudWatch Logs.
type CloudWatchLogDestination struct {
	// The CloudWatch log group.
	LogGroup awslogs.ILogGroup `field:"required" json:"logGroup" yaml:"logGroup"`
}

