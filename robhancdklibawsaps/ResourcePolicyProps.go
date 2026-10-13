package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsaps"
)

// Properties for an Amazon Managed Service for Prometheus workspace resource policy.
type ResourcePolicyProps struct {
	// The IAM policy document that grants access to the workspace.
	//
	// Every statement must include a principal. When the workspace uses a customer-managed KMS key,
	// principals also need permission to use that key.
	PolicyDocument awsiam.PolicyDocument `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The workspace protected by the policy.
	Workspace interfacesawsaps.IWorkspaceRef `field:"required" json:"workspace" yaml:"workspace"`
}

