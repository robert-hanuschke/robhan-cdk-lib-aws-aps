package robhancdklibawsaps


// Properties for creating a rule groups namespace in an Amazon Managed Service for Prometheus workspace.
type RuleGroupsNamespaceProps struct {
	// The rules file used in the namespace.
	Data *string `field:"required" json:"data" yaml:"data"`
	// The name of the rule groups namespace.
	//
	// Between 1 and 64 characters.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The workspace to add the rule groups namespace.
	Workspace IWorkspace `field:"required" json:"workspace" yaml:"workspace"`
}

