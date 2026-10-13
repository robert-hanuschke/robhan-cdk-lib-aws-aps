package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslogs"
)

// Diagnostic logging configuration for a scraper.
type ScraperLoggingConfiguration struct {
	// The scraper components for which logging is enabled.
	Components *[]*ScraperComponent `field:"required" json:"components" yaml:"components"`
	// The CloudWatch log group that receives scraper logs.
	LogGroup interfacesawslogs.ILogGroupRef `field:"required" json:"logGroup" yaml:"logGroup"`
}

