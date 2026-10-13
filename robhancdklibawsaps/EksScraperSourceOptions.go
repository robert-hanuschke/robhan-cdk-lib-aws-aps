package robhancdklibawsaps

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawseks"
)

// Options for collecting metrics from an Amazon EKS cluster.
type EksScraperSourceOptions struct {
	// The EKS cluster from which metrics are collected.
	//
	// The cluster must have private endpoint access and VPC DNS enabled. Automatic scraper access
	// entries require EKS authentication mode `API` or `API_AND_CONFIG_MAP`; otherwise configure the
	// required Kubernetes RBAC permissions manually.
	Cluster interfacesawseks.IClusterRef `field:"required" json:"cluster" yaml:"cluster"`
	// Subnets used to reach scrape targets.
	//
	// Each subnet must have an available IP address.
	Subnets *[]interfacesawsec2.ISubnetRef `field:"required" json:"subnets" yaml:"subnets"`
	// Security groups used to reach scrape targets.
	//
	// `undefined` omits `securityGroupIds` and leaves selection to the service. An explicit empty
	// array is preserved in the synthesized resource. Supply only groups whose rules allow the
	// intended private scrape traffic.
	// Default: - no explicit security groups.
	//
	SecurityGroups *[]interfacesawsec2.ISecurityGroupRef `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

