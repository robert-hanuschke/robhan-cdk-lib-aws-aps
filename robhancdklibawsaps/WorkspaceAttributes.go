package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for importing an Amazon Managed Service for Prometheus Workspace from attributes.
type WorkspaceAttributes struct {
	// The arn of this workspace.
	WorkspaceArn *string `field:"required" json:"workspaceArn" yaml:"workspaceArn"`
	// The alert manager definition, a YAML configuration for the alert manager in your Amazon Managed Service for Prometheus workspace.
	AlertManagerDefinition *string `field:"optional" json:"alertManagerDefinition" yaml:"alertManagerDefinition"`
	// The alias that is assigned to this workspace to help identify it.
	//
	// It does not need to be
	// unique.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The customer managed AWS KMS key to use for encrypting data within your workspace.
	KmsKey awskms.IKey `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Contains information about the current rules and alerting logging configuration for the workspace.
	//
	// Note: These logging configurations are only for rules and alerting logs.
	LoggingConfiguration *LoggingConfiguration `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// The definition of logging configuration in an Amazon Managed Service for Prometheus workspace.
	QueryLoggingConfiguration *QueryLoggingConfiguration `field:"optional" json:"queryLoggingConfiguration" yaml:"queryLoggingConfiguration"`
	// Use this structure to define label sets and the ingestion limits for time series that match label sets, and to specify the retention period of the workspace.
	WorkspaceConfiguration *WorkspaceConfiguration `field:"optional" json:"workspaceConfiguration" yaml:"workspaceConfiguration"`
}

