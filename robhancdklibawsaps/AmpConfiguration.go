package robhancdklibawsaps


// The AmpConfiguration structure defines the Amazon Managed Service for Prometheus instance a scraper should send metrics to.
type AmpConfiguration struct {
	// The Amazon Managed Service for Prometheus workspace.
	Workspace IWorkspace `field:"required" json:"workspace" yaml:"workspace"`
}

