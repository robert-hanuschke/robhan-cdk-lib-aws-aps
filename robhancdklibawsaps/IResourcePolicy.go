package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/internal"
)

type IResourcePolicy interface {
	awscdk.IResource
	// The JSON to use as the Resource-based Policy.
	PolicyDocument() *string
	// The workspace to attach the policy to.
	Workspace() IWorkspace
}

// The jsii proxy for IResourcePolicy
type jsiiProxy_IResourcePolicy struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IResourcePolicy) PolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourcePolicy) Workspace() IWorkspace {
	var returns IWorkspace
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}

