package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseks"
)

// The EksConfiguration structure describes the connection to the Amazon EKS cluster from which a scraper collects metrics.
type EksConfiguration struct {
	// The Amazon EKS cluster.
	Cluster awseks.ICluster `field:"required" json:"cluster" yaml:"cluster"`
	// A list of subnets for the Amazon EKS cluster VPC configuration.
	//
	// Min 1, max 5.
	Subnets *[]awsec2.ISubnet `field:"required" json:"subnets" yaml:"subnets"`
	// A list of the security group IDs for the Amazon EKS cluster VPC configuration.
	//
	// Min 1, max 5.
	SecurityGroups *[]awsec2.ISecurityGroup `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

