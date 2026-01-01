package robhancdklibawsaps


// Properties for creating an Amazon Managed Service for Prometheus Resource Policy.
type ResourcePolicyProps struct {
	// The JSON to use as the Resource-based Policy.
	PolicyDocument *string `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The workspace to attach the policy to.
	Workspace IWorkspace `field:"required" json:"workspace" yaml:"workspace"`
}

