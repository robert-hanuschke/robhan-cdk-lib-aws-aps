package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsaps"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/internal"
)

// An Amazon Managed Service for Prometheus workspace resource policy.
type IResourcePolicy interface {
	awscdk.IResource
	interfacesawsaps.IResourcePolicyRef
	// The ARN of the workspace protected by the policy.
	WorkspaceArn() *string
}

// The jsii proxy for IResourcePolicy
type jsiiProxy_IResourcePolicy struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsapsIResourcePolicyRef
}

func (i *jsiiProxy_IResourcePolicy) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IResourcePolicy) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IResourcePolicy) WorkspaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourcePolicy) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourcePolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourcePolicy) ResourcePolicyRef() *interfacesawsaps.ResourcePolicyReference {
	var returns *interfacesawsaps.ResourcePolicyReference
	_jsii_.Get(
		j,
		"resourcePolicyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResourcePolicy) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

