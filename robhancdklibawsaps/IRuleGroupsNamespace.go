package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsaps"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/internal"
)

// An Amazon Managed Service for Prometheus rule groups namespace.
type IRuleGroupsNamespace interface {
	awscdk.IResource
	interfacesawsaps.IRuleGroupsNamespaceRef
	// The ARN of the rule groups namespace.
	RuleGroupsNamespaceArn() *string
}

// The jsii proxy for IRuleGroupsNamespace
type jsiiProxy_IRuleGroupsNamespace struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsapsIRuleGroupsNamespaceRef
}

func (i *jsiiProxy_IRuleGroupsNamespace) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IRuleGroupsNamespace) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRuleGroupsNamespace) RuleGroupsNamespaceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ruleGroupsNamespaceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleGroupsNamespace) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleGroupsNamespace) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleGroupsNamespace) RuleGroupsNamespaceRef() *interfacesawsaps.RuleGroupsNamespaceReference {
	var returns *interfacesawsaps.RuleGroupsNamespaceReference
	_jsii_.Get(
		j,
		"ruleGroupsNamespaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuleGroupsNamespace) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

