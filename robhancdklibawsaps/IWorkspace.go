package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
	"github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/internal"
)

type IWorkspace interface {
	awscdk.IResource
	// The alert manager definition, a YAML configuration for the alert manager in your Amazon Managed Service for Prometheus workspace.
	AlertManagerDefinition() *string
	// The alias that is assigned to this workspace to help identify it.
	//
	// It does not need to be
	// unique.
	Alias() *string
	// The customer managed AWS KMS key to use for encrypting data within your workspace.
	KmsKey() awskms.IKey
	// Contains information about the current rules and alerting logging configuration for the workspace.
	//
	// Note: These logging configurations are only for rules and alerting logs.
	LoggingConfiguration() *LoggingConfiguration
	// The definition of logging configuration in an Amazon Managed Service for Prometheus workspace.
	QueryLoggingConfiguration() *QueryLoggingConfiguration
	// The ARN of the workspace.
	WorkspaceArn() *string
	// Use this structure to define label sets and the ingestion limits for time series that match label sets, and to specify the retention period of the workspace.
	WorkspaceConfiguration() *WorkspaceConfiguration
	// The unique ID for the workspace.
	WorkspaceId() *string
}

// The jsii proxy for IWorkspace
type jsiiProxy_IWorkspace struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IWorkspace) AlertManagerDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alertManagerDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) KmsKey() awskms.IKey {
	var returns awskms.IKey
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) LoggingConfiguration() *LoggingConfiguration {
	var returns *LoggingConfiguration
	_jsii_.Get(
		j,
		"loggingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkspace) QueryLoggingConfiguration() *QueryLoggingConfiguration {
	var returns *QueryLoggingConfiguration
	_jsii_.Get(
		j,
		"queryLoggingConfiguration",
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

func (j *jsiiProxy_IWorkspace) WorkspaceConfiguration() *WorkspaceConfiguration {
	var returns *WorkspaceConfiguration
	_jsii_.Get(
		j,
		"workspaceConfiguration",
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

