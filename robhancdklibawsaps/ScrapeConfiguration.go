package robhancdklibawsaps


// A scrape configuration for a scraper, base 64 encoded.
type ScrapeConfiguration struct {
	// The base 64 encoded scrape configuration file.
	ConfigurationBlob *string `field:"required" json:"configurationBlob" yaml:"configurationBlob"`
}

