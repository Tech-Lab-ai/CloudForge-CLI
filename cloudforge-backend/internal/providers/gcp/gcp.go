package gcp

import "fmt"

// GCP is the GCP provider.
type GCP struct{}

// Apply applies the configuration to GCP.
func (g *GCP) Apply() {
	fmt.Println("Applying configuration to GCP...")
}

// Destroy destroys the infrastructure on GCP.
func (g *GCP) Destroy() {
	fmt.Println("Destroying infrastructure on GCP...")
}

// Plan plans the changes to be made on GCP.
func (g *GCP) Plan() (string, error) {
	return "Planning changes on GCP...", nil
}

// Deploy deploys the application to GCP.
func (g *GCP) Deploy(path string) {
	fmt.Printf("Deploying application from %s to GCP...\n", path)
}
