package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/internal"
)

type IRuleGroupsNamespace interface {
	awscdk.IResource
	// The rules file used in the namespace.
	Data() *string
	// The name of the rule groups namespace.
	Name() *string
	// The ARN of the rule groups namespace.
	RuleGroupsNamespaceArn() *string
	// The workspace to add the rule groups namespace.
	Workspace() IWorkspace
}

// The jsii proxy for IRuleGroupsNamespace
type jsiiProxy_IRuleGroupsNamespace struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IRuleGroupsNamespace) Data() *string {
	var returns *string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleGroupsNamespace) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleGroupsNamespace) RuleGroupsNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleGroupsNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleGroupsNamespace) Workspace() IWorkspace {
	var returns IWorkspace
	_jsii_.Get(
		j,
		"workspace",
		&returns,
	)
	return returns
}

