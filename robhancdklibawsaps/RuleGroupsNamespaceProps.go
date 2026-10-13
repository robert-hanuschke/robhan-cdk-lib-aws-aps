package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsaps"
)

// Properties for an Amazon Managed Service for Prometheus rule groups namespace.
type RuleGroupsNamespaceProps struct {
	// The Prometheus rules file as YAML.
	Data *string `field:"required" json:"data" yaml:"data"`
	// The name of the rule groups namespace.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The workspace that contains the rule groups namespace.
	Workspace interfacesawsaps.IWorkspaceRef `field:"required" json:"workspace" yaml:"workspace"`
	// Tags to apply to the rule groups namespace.
	// Default: - no tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

