package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawskms"
)

// Properties for an Amazon Managed Service for Prometheus workspace.
type WorkspaceProps struct {
	// The alert manager definition as YAML.
	// Default: - no alert manager definition.
	//
	AlertManagerDefinition *string `field:"optional" json:"alertManagerDefinition" yaml:"alertManagerDefinition"`
	// A user-assigned alias for the workspace.
	// Default: - no alias.
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The customer-managed KMS key used to encrypt workspace data.
	//
	// A managed scraper writing to a workspace encrypted with a customer-managed key requires an
	// explicit scraper role configuration, even when the scraper and workspace share an account.
	// Resource-policy principals also require access to this key.
	// Default: - an AWS-owned key.
	//
	KmsKey interfacesawskms.IKeyRef `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Rules and alerting logging configuration.
	// Default: - rules and alerting logs are not published.
	//
	LoggingConfiguration *WorkspaceLoggingConfiguration `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// Query logging configuration.
	// Default: - queries are not logged.
	//
	QueryLoggingConfiguration *WorkspaceQueryLoggingConfiguration `field:"optional" json:"queryLoggingConfiguration" yaml:"queryLoggingConfiguration"`
	// Tags to apply to the workspace.
	// Default: - no tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Workspace retention and ingestion settings.
	// Default: - service defaults.
	//
	WorkspaceConfiguration *WorkspaceConfiguration `field:"optional" json:"workspaceConfiguration" yaml:"workspaceConfiguration"`
}

