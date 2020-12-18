package test

import (
	"testing"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraformAwsHelloWorldExample(t *testing.T) {
	t.Parallel()

	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"region": "us-ashburn-1",
			"internet_gateway_enabled": false,
			"vcn_cidr": "10.0.0.0/16",
			"vcn_dns_label": "vcn",
			"vcn_name": "testvcn",
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created.
	//defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndPlan(t, terraformOptions)

	// Run `terraform output` to get the IP of the instance
	//bastionIp := terraform.Output(t, terraformOptions, "bastion_public_ip")

	// Make   sure we get back a 200 OK with the body "Hello, World!"
	//fmt.Println(bastionIp)
// 	shell.RunCommandAndGetStdOut(t *testing.T,command )
 }
