package robhancdklibawsaps


// Properties for importing an Amazon Managed Service for Prometheus Scraper from attributes.
type ScraperAttributes struct {
	// The Amazon Managed Service for Prometheus workspace the scraper sends metrics to.
	Destination *Destination `field:"required" json:"destination" yaml:"destination"`
	// The configuration in use by the scraper.
	ScrapeConfiguration *ScrapeConfiguration `field:"required" json:"scrapeConfiguration" yaml:"scrapeConfiguration"`
	// The ARN of the scraper.
	ScraperArn *string `field:"required" json:"scraperArn" yaml:"scraperArn"`
	// The Amazon EKS cluster from which the scraper collects metrics.
	Source *Source `field:"required" json:"source" yaml:"source"`
	// An optional user-assigned scraper alias.
	//
	// 1-100 characters.
	//
	// Pattern: ^[0-9A-Za-z][-.0-9A-Z_a-z]*$
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// The role configuration in an Amazon Managed Service for Prometheus scraper.
	RoleConfiguration *RoleConfiguration `field:"optional" json:"roleConfiguration" yaml:"roleConfiguration"`
}

