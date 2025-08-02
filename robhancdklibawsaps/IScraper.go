package robhancdklibawsaps

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/robert-hanuschke/robhan-cdk-lib-aws-aps/robhancdklibawsaps/internal"
)

type IScraper interface {
	awscdk.IResource
	// An optional user-assigned scraper alias.
	//
	// 1-100 characters.
	//
	// Pattern: ^[0-9A-Za-z][-.0-9A-Z_a-z]*$
	Alias() *string
	// The Amazon Managed Service for Prometheus workspace the scraper sends metrics to.
	Destination() *Destination
	// The role configuration in an Amazon Managed Service for Prometheus scraper.
	RoleConfiguration() *RoleConfiguration
	// The configuration in use by the scraper.
	ScrapeConfiguration() *ScrapeConfiguration
	// The ARN of the scraper.
	ScraperArn() *string
	// The ID of the scraper.
	ScraperId() *string
	// The Amazon EKS cluster from which the scraper collects metrics.
	Source() *Source
}

// The jsii proxy for IScraper
type jsiiProxy_IScraper struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_IScraper) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScraper) Destination() *Destination {
	var returns *Destination
	_jsii_.Get(
		j,
		"destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScraper) RoleConfiguration() *RoleConfiguration {
	var returns *RoleConfiguration
	_jsii_.Get(
		j,
		"roleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScraper) ScrapeConfiguration() *ScrapeConfiguration {
	var returns *ScrapeConfiguration
	_jsii_.Get(
		j,
		"scrapeConfiguration",
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

func (j *jsiiProxy_IScraper) Source() *Source {
	var returns *Source
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

