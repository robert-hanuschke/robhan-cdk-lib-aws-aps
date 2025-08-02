package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
)

// Contains information about the rules and alerting logging configuration for the workspace.
type LoggingConfiguration struct {
	// The CloudWatch log group to which the vended log data will be published.
	LogGroup awslogs.ILogGroup `field:"optional" json:"logGroup" yaml:"logGroup"`
}

