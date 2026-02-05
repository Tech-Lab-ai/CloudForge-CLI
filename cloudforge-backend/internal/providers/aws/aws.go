package aws

import "fmt"

// AWS is the AWS provider.
type AWS struct{}

// Apply applies the configuration to AWS.
func (a *AWS) Apply() {
	fmt.Println("Applying configuration to AWS...")
}

// Destroy destroys the infrastructure on AWS.
func (a *AWS) Destroy() {
	fmt.Println("Destroying infrastructure on AWS...")
}

// Plan plans the changes to be made on AWS.
func (a *AWS) Plan() {
	fmt.Println("Planning changes on AWS...")
}
