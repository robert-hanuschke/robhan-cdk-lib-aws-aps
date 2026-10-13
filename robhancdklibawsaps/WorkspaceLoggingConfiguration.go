package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// The rules and alerting log destination for an Amazon Managed Service for Prometheus workspace.
type WorkspaceLoggingConfiguration struct {
	// The CloudWatch log group to which rules and alerting logs are published.
	LogGroup interfacesawslogs.ILogGroupRef `field:"required" json:"logGroup" yaml:"logGroup"`
}

