package robhancdklibawsaps


// Properties for an Amazon Managed Service for Prometheus scraper.
type ScraperProps struct {
	// The destination for collected metrics.
	Destination ScraperDestination `field:"required" json:"destination" yaml:"destination"`
	// The Prometheus scrape configuration.
	ScrapeConfiguration ScrapeConfiguration `field:"required" json:"scrapeConfiguration" yaml:"scrapeConfiguration"`
	// The source from which metrics are collected.
	Source ScraperSource `field:"required" json:"source" yaml:"source"`
	// A user-assigned alias for the scraper.
	// Default: - no alias.
	//
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Diagnostic logging configuration.
	// Default: - scraper diagnostic logging is disabled.
	//
	LoggingConfiguration *ScraperLoggingConfiguration `field:"optional" json:"loggingConfiguration" yaml:"loggingConfiguration"`
	// IAM roles used to access the source and destination.
	//
	// Required for cross-account collection and for workspaces encrypted with a customer-managed
	// KMS key. The caller is responsible for the roles' trust and permissions.
	// Default: - service-created roles.
	//
	RoleConfiguration *ScraperRoleConfiguration `field:"optional" json:"roleConfiguration" yaml:"roleConfiguration"`
	// Tags to apply to the scraper.
	// Default: - no tags.
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

