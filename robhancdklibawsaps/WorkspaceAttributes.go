package robhancdklibawsaps


// Attributes for referencing an existing workspace.
type WorkspaceAttributes struct {
	// The ARN of the workspace.
	WorkspaceArn *string `field:"required" json:"workspaceArn" yaml:"workspaceArn"`
	// The Prometheus-compatible query endpoint.
	// Default: - derived from the workspace ARN.
	//
	PrometheusEndpoint *string `field:"optional" json:"prometheusEndpoint" yaml:"prometheusEndpoint"`
}

