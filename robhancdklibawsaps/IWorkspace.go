package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsaps"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/internal"
)

// An Amazon Managed Service for Prometheus workspace.
type IWorkspace interface {
	awscdk.IResource
	interfacesawsaps.IWorkspaceRef
	// The Prometheus-compatible query endpoint for the workspace.
	PrometheusEndpoint() *string
	// The ARN of the workspace.
	WorkspaceArn() *string
	// The unique ID of the workspace.
	WorkspaceId() *string
}

// The jsii proxy for IWorkspace
type jsiiProxy_IWorkspace struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsapsIWorkspaceRef
}

func (i *jsiiProxy_IWorkspace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IWorkspace) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IWorkspace) PrometheusEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prometheusEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) WorkspaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) WorkspaceRef() *interfacesawsaps.WorkspaceReference {
	var returns *interfacesawsaps.WorkspaceReference
	_jsii_.Get(
		j,
		"workspaceRef",
		&returns,
	)
	return returns
}

