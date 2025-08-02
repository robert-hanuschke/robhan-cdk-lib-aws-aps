package robhancdklibawsaps


// Where to send the metrics from a scraper.
type Destination struct {
	// The Amazon Managed Service for Prometheus workspace to send metrics to.
	AmpConfiguration *AmpConfiguration `field:"required" json:"ampConfiguration" yaml:"ampConfiguration"`
}

