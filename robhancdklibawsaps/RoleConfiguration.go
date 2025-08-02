package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
)

// The role configuration in an Amazon Managed Service for Prometheus scraper.
type RoleConfiguration struct {
	// The source role.
	SourceRole awsiam.IRole `field:"optional" json:"sourceRole" yaml:"sourceRole"`
	// The target role.
	TargetRole awsiam.IRole `field:"optional" json:"targetRole" yaml:"targetRole"`
}

