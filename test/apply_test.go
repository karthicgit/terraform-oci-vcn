package test

import (
	"testing"
	"os"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"fmt"
)

func TerratestOCI(t *testing.T) {
	t.Parallel()

	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	compartment_id := os.Getenv("TF_VAR_compartment_id")
	tenancy_ocid := os.Getenv("TF_VAR_tenancy_ocid")
	fingerprint := os.Getenv("TF_VAR_fingerprint")
	private_key := os.Getenv("TF_VAR_private_key")
	user_ocid := os.Getenv("TF_VAR_user_ocid")
	
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"region": "us-ashburn-1",
			"internet_gateway_enabled": false,
			"vcn_cidr": "10.0.0.0/16",
			"vcn_dns_label": "vcn",
			"vcn_name": "testvcn",
			"compartment_id": compartment_id,
			"tenancy_ocid": tenancy_ocid,
			"fingerprint": fingerprint,
			"private_key" : private_key,
			"user_ocid" : user_ocid,
		},
		
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created.
	//defer terraform.Destroy(t, terraformOptions)
	fmt.Println(os.Environ())
	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndPlan(t, terraformOptions)

	// Run `terraform output` to get the IP of the instance
	//bastionIp := terraform.Output(t, terraformOptions, "bastion_public_ip")

	// Make  get back a 200 OK with the body "Hello, World!"
	
// 	shell.RunCommandAndGetStdOut(t *testing.T,command )
 }
