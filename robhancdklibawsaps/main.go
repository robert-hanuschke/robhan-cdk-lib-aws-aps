// AWS CDK Construct Library for Amazon Managed Service for Prometheus
package robhancdklibawsaps

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.AnomalyDetector",
		reflect.TypeOf((*AnomalyDetector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "anomalyDetectorArn", GoGetter: "AnomalyDetectorArn"},
			_jsii_.MemberProperty{JsiiProperty: "anomalyDetectorRef", GoGetter: "AnomalyDetectorRef"},
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_AnomalyDetector{}
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
		"@robhan-cdk-lib/aws_aps.AnomalyDetectorLabel",
		reflect.TypeOf((*AnomalyDetectorLabel)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.AnomalyDetectorProps",
		reflect.TypeOf((*AnomalyDetectorProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.EksScraperSourceOptions",
		reflect.TypeOf((*EksScraperSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"@robhan-cdk-lib/aws_aps.IAnomalyDetector",
		reflect.TypeOf((*IAnomalyDetector)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "anomalyDetectorArn", GoGetter: "AnomalyDetectorArn"},
			_jsii_.MemberProperty{JsiiProperty: "anomalyDetectorRef", GoGetter: "AnomalyDetectorRef"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IAnomalyDetector{}
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsapsIAnomalyDetectorRef)
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
			_jsii_.MemberProperty{JsiiProperty: "resourcePolicyRef", GoGetter: "ResourcePolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceArn", GoGetter: "WorkspaceArn"},
		},
		func() interface{} {
			j := jsiiProxy_IResourcePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsapsIResourcePolicyRef)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@robhan-cdk-lib/aws_aps.IRuleGroupsNamespace",
		reflect.TypeOf((*IRuleGroupsNamespace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "ruleGroupsNamespaceArn", GoGetter: "RuleGroupsNamespaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "ruleGroupsNamespaceRef", GoGetter: "RuleGroupsNamespaceRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IRuleGroupsNamespace{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsapsIRuleGroupsNamespaceRef)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@robhan-cdk-lib/aws_aps.IScraper",
		reflect.TypeOf((*IScraper)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "scraperArn", GoGetter: "ScraperArn"},
			_jsii_.MemberProperty{JsiiProperty: "scraperId", GoGetter: "ScraperId"},
			_jsii_.MemberProperty{JsiiProperty: "scraperRef", GoGetter: "ScraperRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_IScraper{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsapsIScraperRef)
			return &j
		},
	)
	_jsii_.RegisterInterface(
		"@robhan-cdk-lib/aws_aps.IWorkspace",
		reflect.TypeOf((*IWorkspace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "prometheusEndpoint", GoGetter: "PrometheusEndpoint"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceArn", GoGetter: "WorkspaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceRef", GoGetter: "WorkspaceRef"},
		},
		func() interface{} {
			j := jsiiProxy_IWorkspace{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkIResource)
			_jsii_.InitJsiiProxy(&j.Type__interfacesawsapsIWorkspaceRef)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.IgnoreNearExpected",
		reflect.TypeOf((*IgnoreNearExpected)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IgnoreNearExpected{}
		},
	)
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_aps.MissingDataAction",
		reflect.TypeOf((*MissingDataAction)(nil)).Elem(),
		map[string]interface{}{
			"MARK_AS_ANOMALY": MissingDataAction_MARK_AS_ANOMALY,
			"SKIP": MissingDataAction_SKIP,
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.RandomCutForestConfiguration",
		reflect.TypeOf((*RandomCutForestConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.ResourcePolicy",
		reflect.TypeOf((*ResourcePolicy)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "policyDocument", GoGetter: "PolicyDocument"},
			_jsii_.MemberProperty{JsiiProperty: "resourcePolicyRef", GoGetter: "ResourcePolicyRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceArn", GoGetter: "WorkspaceArn"},
		},
		func() interface{} {
			j := jsiiProxy_ResourcePolicy{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IResourcePolicy)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.ResourcePolicyProps",
		reflect.TypeOf((*ResourcePolicyProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.RuleGroupsNamespace",
		reflect.TypeOf((*RuleGroupsNamespace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "ruleGroupsNamespaceArn", GoGetter: "RuleGroupsNamespaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "ruleGroupsNamespaceRef", GoGetter: "RuleGroupsNamespaceRef"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_RuleGroupsNamespace{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IRuleGroupsNamespace)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.RuleGroupsNamespaceProps",
		reflect.TypeOf((*RuleGroupsNamespaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.ScrapeConfiguration",
		reflect.TypeOf((*ScrapeConfiguration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ScrapeConfiguration{}
		},
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.Scraper",
		reflect.TypeOf((*Scraper)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "scraperArn", GoGetter: "ScraperArn"},
			_jsii_.MemberProperty{JsiiProperty: "scraperId", GoGetter: "ScraperId"},
			_jsii_.MemberMethod{JsiiMethod: "scraperIdFromArn", GoMethod: "ScraperIdFromArn"},
			_jsii_.MemberProperty{JsiiProperty: "scraperRef", GoGetter: "ScraperRef"},
			_jsii_.MemberProperty{JsiiProperty: "scraperRoleArn", GoGetter: "ScraperRoleArn"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_Scraper{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IScraper)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.ScraperComponent",
		reflect.TypeOf((*ScraperComponent)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@robhan-cdk-lib/aws_aps.ScraperComponentType",
		reflect.TypeOf((*ScraperComponentType)(nil)).Elem(),
		map[string]interface{}{
			"SERVICE_DISCOVERY": ScraperComponentType_SERVICE_DISCOVERY,
			"COLLECTOR": ScraperComponentType_COLLECTOR,
			"EXPORTER": ScraperComponentType_EXPORTER,
		},
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.ScraperDestination",
		reflect.TypeOf((*ScraperDestination)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ScraperDestination{}
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.ScraperLoggingConfiguration",
		reflect.TypeOf((*ScraperLoggingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.ScraperProps",
		reflect.TypeOf((*ScraperProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.ScraperRoleConfiguration",
		reflect.TypeOf((*ScraperRoleConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.ScraperSource",
		reflect.TypeOf((*ScraperSource)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ScraperSource{}
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.VpcScraperSourceOptions",
		reflect.TypeOf((*VpcScraperSourceOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@robhan-cdk-lib/aws_aps.Workspace",
		reflect.TypeOf((*Workspace)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyCrossStackReferenceStrength", GoMethod: "ApplyCrossStackReferenceStrength"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "generatePhysicalName", GoMethod: "GeneratePhysicalName"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceArnAttribute", GoMethod: "GetResourceArnAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getResourceNameAttribute", GoMethod: "GetResourceNameAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "physicalName", GoGetter: "PhysicalName"},
			_jsii_.MemberProperty{JsiiProperty: "prometheusEndpoint", GoGetter: "PrometheusEndpoint"},
			_jsii_.MemberMethod{JsiiMethod: "prometheusEndpointFromArn", GoMethod: "PrometheusEndpointFromArn"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceArn", GoGetter: "WorkspaceArn"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceId", GoGetter: "WorkspaceId"},
			_jsii_.MemberMethod{JsiiMethod: "workspaceIdFromArn", GoMethod: "WorkspaceIdFromArn"},
			_jsii_.MemberProperty{JsiiProperty: "workspaceRef", GoGetter: "WorkspaceRef"},
		},
		func() interface{} {
			j := jsiiProxy_Workspace{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_IWorkspace)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.WorkspaceAttributes",
		reflect.TypeOf((*WorkspaceAttributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.WorkspaceConfiguration",
		reflect.TypeOf((*WorkspaceConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.WorkspaceLabel",
		reflect.TypeOf((*WorkspaceLabel)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.WorkspaceLabelSet",
		reflect.TypeOf((*WorkspaceLabelSet)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.WorkspaceLabelSetLimits",
		reflect.TypeOf((*WorkspaceLabelSetLimits)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.WorkspaceLoggingConfiguration",
		reflect.TypeOf((*WorkspaceLoggingConfiguration)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.WorkspaceProps",
		reflect.TypeOf((*WorkspaceProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@robhan-cdk-lib/aws_aps.WorkspaceQueryLoggingConfiguration",
		reflect.TypeOf((*WorkspaceQueryLoggingConfiguration)(nil)).Elem(),
	)
}
