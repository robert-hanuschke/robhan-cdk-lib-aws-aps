package robhancdklibawsaps


// Components of a scraper that support diagnostic logging.
type ScraperComponentType string

const (
	// Service discovery.
	ScraperComponentType_SERVICE_DISCOVERY ScraperComponentType = "SERVICE_DISCOVERY"
	// Metric collection.
	ScraperComponentType_COLLECTOR ScraperComponentType = "COLLECTOR"
	// Metric export.
	ScraperComponentType_EXPORTER ScraperComponentType = "EXPORTER"
)

