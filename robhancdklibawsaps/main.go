// AWS CDK Construct Library for Amazon Managed Service for Prometheus
package robhancdklibawsaps

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.AmpConfiguration",
		reflect.TypeOf((*AmpConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.AnomalyDetector",
		reflect.TypeOf((*AnomalyDetector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "anomalyDetectorArn", GoGetter: "AnomalyDetectorArn"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationIntervalInSeconds", GoGetter: "EvaluationIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "missingDataAction", GoGetter: "MissingDataAction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspace", GoGetter: "Workspace"},
		},
		func() interface{} {
			j := jsiiProxy_AnomalyDetector{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_AnomalyDetectorBase)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.AnomalyDetectorBase",
		reflect.TypeOf((*AnomalyDetectorBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "anomalyDetectorArn", GoGetter: "AnomalyDetectorArn"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationIntervalInSeconds", GoGetter: "EvaluationIntervalInSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "missingDataAction", GoGetter: "MissingDataAction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspace", GoGetter: "Workspace"},
		},
		func() interface{} {
			j := jsiiProxy_AnomalyDetectorBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IAnomalyDetector)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.AnomalyDetectorConfiguration",
		reflect.TypeOf((*AnomalyDetectorConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.AnomalyDetectorProps",
		reflect.TypeOf((*AnomalyDetectorProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.CloudWatchLogDestination",
		reflect.TypeOf((*CloudWatchLogDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.Destination",
		reflect.TypeOf((*Destination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.EksConfiguration",
		reflect.TypeOf((*EksConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@robhan-cdk-lib/aws_aps.IAnomalyDetector",
		reflect.TypeOf((*IAnomalyDetector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "anomalyDetectorArn", GoGetter: "AnomalyDetectorArn"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "configuration", GoGetter: "Configuration"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "evaluationIntervalInSeconds", GoGetter: "EvaluationIntervalInSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "missingDataAction", GoGetter: "MissingDataAction"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "workspace", GoGetter: "Workspace"},
		},
		func() interface{} {
			j := jsiiProxy_IAnomalyDetector{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@robhan-cdk-lib/aws_aps.IResourcePolicy",
		reflect.TypeOf((*IResourcePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "workspace", GoGetter: "Workspace"},
		},
		func() interface{} {
			j := jsiiProxy_IResourcePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@robhan-cdk-lib/aws_aps.IRuleGroupsNamespace",
		reflect.TypeOf((*IRuleGroupsNamespace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "data", GoGetter: "Data"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ruleGroupsNamespaceArn", GoGetter: "RuleGroupsNamespaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "workspace", GoGetter: "Workspace"},
		},
		func() interface{} {
			j := jsiiProxy_IRuleGroupsNamespace{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@robhan-cdk-lib/aws_aps.IScraper",
		reflect.TypeOf((*IScraper)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "roleConfiguration", GoGetter: "RoleConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "scrapeConfiguration", GoGetter: "ScrapeConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "scraperArn", GoGetter: "ScraperArn"},
			_jsii_.MemberProperty{JsiiProperty: "scraperId", GoGetter: "ScraperId"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
		},
		func() interface{} {
			j := jsiiProxy_IScraper{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@robhan-cdk-lib/aws_aps.IWorkspace",
		reflect.TypeOf((*IWorkspace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alertManagerDefinition", GoGetter: "AlertManagerDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfiguration", GoGetter: "LoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "queryLoggingConfiguration", GoGetter: "QueryLoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceArn", GoGetter: "WorkspaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceConfiguration", GoGetter: "WorkspaceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
		},
		func() interface{} {
			j := jsiiProxy_IWorkspace{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.IgnoreNearExpected",
		reflect.TypeOf((*IgnoreNearExpected)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.Label",
		reflect.TypeOf((*Label)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.LimitsPerLabelSet",
		reflect.TypeOf((*LimitsPerLabelSet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.LimitsPerLabelSetEntry",
		reflect.TypeOf((*LimitsPerLabelSetEntry)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.LoggingConfiguration",
		reflect.TypeOf((*LoggingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.LoggingDestination",
		reflect.TypeOf((*LoggingDestination)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.LoggingFilter",
		reflect.TypeOf((*LoggingFilter)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.MissingDataAction",
		reflect.TypeOf((*MissingDataAction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.QueryLoggingConfiguration",
		reflect.TypeOf((*QueryLoggingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.RandomCutForestConfiguration",
		reflect.TypeOf((*RandomCutForestConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.ResourcePolicy",
		reflect.TypeOf((*ResourcePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspace", GoGetter: "Workspace"},
		},
		func() interface{} {
			j := jsiiProxy_ResourcePolicy{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ResourcePolicyBase)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.ResourcePolicyBase",
		reflect.TypeOf((*ResourcePolicyBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspace", GoGetter: "Workspace"},
		},
		func() interface{} {
			j := jsiiProxy_ResourcePolicyBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResourcePolicy)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.ResourcePolicyProps",
		reflect.TypeOf((*ResourcePolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.RoleConfiguration",
		reflect.TypeOf((*RoleConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.RuleGroupsNamespace",
		reflect.TypeOf((*RuleGroupsNamespace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "data", GoGetter: "Data"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleGroupsNamespaceArn", GoGetter: "RuleGroupsNamespaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspace", GoGetter: "Workspace"},
		},
		func() interface{} {
			j := jsiiProxy_RuleGroupsNamespace{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_RuleGroupsNamespaceBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.RuleGroupsNamespaceAttributes",
		reflect.TypeOf((*RuleGroupsNamespaceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.RuleGroupsNamespaceBase",
		reflect.TypeOf((*RuleGroupsNamespaceBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "data", GoGetter: "Data"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleGroupsNamespaceArn", GoGetter: "RuleGroupsNamespaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspace", GoGetter: "Workspace"},
		},
		func() interface{} {
			j := jsiiProxy_RuleGroupsNamespaceBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRuleGroupsNamespace)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.RuleGroupsNamespaceProps",
		reflect.TypeOf((*RuleGroupsNamespaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.ScrapeConfiguration",
		reflect.TypeOf((*ScrapeConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.Scraper",
		reflect.TypeOf((*Scraper)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getScraperId", GoMethod: "GetScraperId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "roleConfiguration", GoGetter: "RoleConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "scrapeConfiguration", GoGetter: "ScrapeConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "scraperArn", GoGetter: "ScraperArn"},
			_jsii_.MemberProperty{JsiiProperty: "scraperId", GoGetter: "ScraperId"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_Scraper{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ScraperBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.ScraperAttributes",
		reflect.TypeOf((*ScraperAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.ScraperBase",
		reflect.TypeOf((*ScraperBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getScraperId", GoMethod: "GetScraperId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "roleConfiguration", GoGetter: "RoleConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "scrapeConfiguration", GoGetter: "ScrapeConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "scraperArn", GoGetter: "ScraperArn"},
			_jsii_.MemberProperty{JsiiProperty: "scraperId", GoGetter: "ScraperId"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_ScraperBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IScraper)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.ScraperProps",
		reflect.TypeOf((*ScraperProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.Source",
		reflect.TypeOf((*Source)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.Workspace",
		reflect.TypeOf((*Workspace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alertManagerDefinition", GoGetter: "AlertManagerDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getWorkspaceId", GoMethod: "GetWorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfiguration", GoGetter: "LoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "prometheusEndpoint", GoGetter: "PrometheusEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "queryLoggingConfiguration", GoGetter: "QueryLoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceArn", GoGetter: "WorkspaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceConfiguration", GoGetter: "WorkspaceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
		},
		func() interface{} {
			j := jsiiProxy_Workspace{}
			_jsii_.InitJsiiProxy(&j.jsiiProxy_WorkspaceBase)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.WorkspaceAttributes",
		reflect.TypeOf((*WorkspaceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.WorkspaceBase",
		reflect.TypeOf((*WorkspaceBase)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "alertManagerDefinition", GoGetter: "AlertManagerDefinition"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getWorkspaceId", GoMethod: "GetWorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfiguration", GoGetter: "LoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "queryLoggingConfiguration", GoGetter: "QueryLoggingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceArn", GoGetter: "WorkspaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceConfiguration", GoGetter: "WorkspaceConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
		},
		func() interface{} {
			j := jsiiProxy_WorkspaceBase{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWorkspace)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.WorkspaceConfiguration",
		reflect.TypeOf((*WorkspaceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.WorkspaceProps",
		reflect.TypeOf((*WorkspaceProps)(nil)).Elem(),
	)
}
