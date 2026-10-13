package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsaps"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/internal"
)

// An Amazon Managed Service for Prometheus scraper.
type IScraper interface {
	awscdk.IResource
	interfacesawsaps.IScraperRef
	// The ARN of the scraper.
	ScraperArn() *string
	// The unique ID of the scraper.
	ScraperId() *string
}

// The jsii proxy for IScraper
type jsiiProxy_IScraper struct {
	internal.Type__awscdkIResource
	internal.Type__interfacesawsapsIScraperRef
}

func (i *jsiiProxy_IScraper) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_IScraper) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IScraper) ScraperArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scraperArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScraper) ScraperId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scraperId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScraper) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScraper) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScraper) ScraperRef() *interfacesawsaps.ScraperReference {
	var returns *interfacesawsaps.ScraperReference
	_jsii_.Get(
		j,
		"scraperRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScraper) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

