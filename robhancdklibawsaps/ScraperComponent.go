package robhancdklibawsaps


// A scraper component that can emit diagnostic logs.
type ScraperComponent struct {
	// The component type.
	Type ScraperComponentType `field:"required" json:"type" yaml:"type"`
	// Component-specific logging options.
	// Default: - no component-specific options.
	//
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
}

