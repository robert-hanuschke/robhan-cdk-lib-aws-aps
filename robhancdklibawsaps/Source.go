package robhancdklibawsaps


// The source of collected metrics for a scraper.
type Source struct {
	// The Amazon EKS cluster from which a scraper collects metrics.
	EksConfiguration *EksConfiguration `field:"required" json:"eksConfiguration" yaml:"eksConfiguration"`
}

