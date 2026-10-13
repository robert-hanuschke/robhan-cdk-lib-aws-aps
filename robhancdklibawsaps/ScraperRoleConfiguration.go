package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiam"
)

// IAM roles used by a scraper to access its source and destination.
//
// Supply both roles for cross-account collection. The caller owns the roles' trust and permissions.
// Both roles are also required when writing to a workspace encrypted with a customer-managed KMS
// key, including same-account collection.
type ScraperRoleConfiguration struct {
	// The role used to access the metric source.
	// Default: - a service role is created by Amazon Managed Service for Prometheus.
	//
	SourceRole interfacesawsiam.IRoleRef `field:"optional" json:"sourceRole" yaml:"sourceRole"`
	// The role used to write to the destination.
	// Default: - a service role is created by Amazon Managed Service for Prometheus.
	//
	TargetRole interfacesawsiam.IRoleRef `field:"optional" json:"targetRole" yaml:"targetRole"`
}

