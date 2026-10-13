package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
)

// Options for collecting metrics from resources reachable in a VPC.
type VpcScraperSourceOptions struct {
	// Security groups used to reach scrape targets.
	SecurityGroups *[]interfacesawsec2.ISecurityGroupRef `field:"required" json:"securityGroups" yaml:"securityGroups"`
	// Subnets used to reach scrape targets.
	//
	// Prefer private subnets and narrowly scoped security
	// group rules for the intended Prometheus endpoints.
	Subnets *[]interfacesawsec2.ISubnetRef `field:"required" json:"subnets" yaml:"subnets"`
}

