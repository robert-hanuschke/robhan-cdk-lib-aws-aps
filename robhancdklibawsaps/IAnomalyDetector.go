package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsaps"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/internal"
)

// An Amazon Managed Service for Prometheus anomaly detector.
type IAnomalyDetector interface {
	interfacesawsaps.IAnomalyDetectorRef
	awscdk.IResource
	// The ARN of the anomaly detector.
	AnomalyDetectorArn() *string
}

// The jsii proxy for IAnomalyDetector
type jsiiProxy_IAnomalyDetector struct {
	internal.Type__interfacesawsapsIAnomalyDetectorRef
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_IAnomalyDetector) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IAnomalyDetector) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAnomalyDetector) AnomalyDetectorArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"anomalyDetectorArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyDetector) AnomalyDetectorRef() *interfacesawsaps.AnomalyDetectorReference {
	var returns *interfacesawsaps.AnomalyDetectorReference
	_jsii_.Get(
		j,
		"anomalyDetectorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyDetector) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyDetector) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAnomalyDetector) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

